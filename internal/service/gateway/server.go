package gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/rss3-gateway/internal/config"
	"github.com/naturalselectionlabs/rss3-gateway/internal/database"
	"github.com/naturalselectionlabs/rss3-gateway/internal/service"
	"github.com/naturalselectionlabs/rss3-gateway/internal/service/gateway/gen/oapi"
	"github.com/naturalselectionlabs/rss3-gateway/internal/service/gateway/handlers"
	"github.com/naturalselectionlabs/rss3-gateway/internal/service/gateway/jwt"
	"github.com/naturalselectionlabs/rss3-gateway/internal/service/gateway/middlewares"
	"github.com/naturalselectionlabs/rss3-gateway/internal/service/gateway/processors"
	"github.com/naturalselectionlabs/rss3-gateway/internal/service/gateway/siwe"
	"github.com/naturalselectionlabs/rss3-gateway/internal/service/gateway/swagger"
	"github.com/redis/go-redis/v9"
	"github.com/rss3-network/gateway-common/accesslog"
	"github.com/rss3-network/gateway-common/control"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	config         config.Gateway
	redis          *redis.Client
	databaseClient database.Client
	controlClient  *control.StateClientWriter
}

func (s *Server) Run(ctx context.Context) error {
	errorPool := pool.New().WithContext(ctx).WithCancelOnError().WithFirstError()

	// Run echo server.
	errorPool.Go(func(ctx context.Context) error {
		// Prepare JWT
		jwtClient, err := jwt.New(s.config.API.JWTKey)
		if err != nil {
			return fmt.Errorf("prepare JWT: %w", err)
		}

		// Prepare SIWE
		siweClient, err := siwe.New(s.config.API.SIWEDomain, s.redis)
		if err != nil {
			return fmt.Errorf("prepare SIWE: %w", err)
		}

		// Connect to kafka for access logs
		kafkaClient, err := accesslog.NewConsumer(
			s.config.Kafka.Brokers,
			s.config.Kafka.Topic,
			"gateway",
		)
		if err != nil {
			return fmt.Errorf("prepare kafka: %w", err)
		}

		// Prepare processors
		processorApp, err := processors.NewApp(s.controlClient, s.databaseClient.Raw())
		if err != nil {
			return fmt.Errorf("prepare processors: %w", err)
		}

		err = kafkaClient.Start(processorApp.ProcessAccessLog)
		if err != nil {
			return fmt.Errorf("start kafka: %w", err)
		}

		// Prepare handler
		e := echo.New()
		handlerApp, err := handlers.NewApp(
			s.controlClient,
			s.redis,
			s.databaseClient.Raw(),
			jwtClient,
			siweClient,
		)
		if err != nil {
			return fmt.Errorf("start handler: %w", err)
		}

		// Configure middlewares
		configureMiddlewares(e, handlerApp, jwtClient, s.databaseClient.Raw(), s.controlClient)

		// Start echo API server
		return e.Start(fmt.Sprintf("%s:%d", s.config.API.Listen.Host, s.config.API.Listen.Port))
	})

	errorChan := make(chan error)
	go func() { errorChan <- errorPool.Wait() }()

	select {
	case err := <-errorChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func New(databaseClient database.Client, redis *redis.Client, controlClient *control.StateClientWriter, config config.Gateway) (service.Server, error) {
	instance := Server{
		config:         config,
		redis:          redis,
		databaseClient: databaseClient,
		controlClient:  controlClient,
	}

	return &instance, nil
}

func configureMiddlewares(e *echo.Echo, app *handlers.App, jwtClient *jwt.JWT, databaseClient *gorm.DB, controlClient *control.StateClientWriter) {
	oapi.RegisterHandlers(e, app)

	// Add api docs
	if os.Getenv(config.Environment) == config.EnvironmentDevelopment {
		swg, err := oapi.GetSwagger()

		if err != nil {
			// Log but ignore
			zap.L().Error("get swagger doc", zap.Error(err))
		}

		swgJSON, err := swg.MarshalJSON()

		if err != nil {
			// Log but ignore
			zap.L().Error("marshal swagger doc", zap.Error(err))
		}

		e.Pre(swagger.SwaggerDoc("/", swgJSON))
	}

	// Check user authentication
	e.Use(middlewares.UserAuthenticationMiddleware(databaseClient, controlClient, jwtClient))

	e.HTTPErrorHandler = customHTTPErrorHandler
}

func customHTTPErrorHandler(err error, c echo.Context) {
	// ignore user cancelled error
	switch {
	case errors.Is(err, context.Canceled):
		_ = c.NoContent(0)
	case errors.Is(err, gorm.ErrRecordNotFound):
		_ = JSONResponseMsg(c, err.Error(), http.StatusNotFound)
	case errors.Is(err, gorm.ErrInvalidField):
		_ = JSONResponseMsg(c, err.Error(), http.StatusBadRequest)
	case errors.Is(err, errors.New(http.StatusText(http.StatusUnauthorized))) && err.Error() == http.StatusText(http.StatusUnauthorized):
		_ = JSONResponseMsg(c, "Your credentials have expired.", http.StatusUnauthorized)
	case strings.Contains(err.Error(), "Path was not found"):
		_ = JSONResponseMsg(c, err.Error(), http.StatusNotFound)
	}

	c.Echo().DefaultHTTPErrorHandler(err, c)
}

func JSONResponseMsg(ctx echo.Context, msg string, code int) error {
	return ctx.JSON(code, map[string]interface{}{
		"msg":    msg,
		"errors": struct{}{},
	})
}

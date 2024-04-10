package indexer

import (
	"context"
	"github.com/redis/go-redis/v9"

	"github.com/rss3-network/gateway-common/control"
	"github.com/rss3-network/payment-processor/internal/config"
	"github.com/rss3-network/payment-processor/internal/database"
	"github.com/rss3-network/payment-processor/internal/service/indexer/l2"
	"github.com/sourcegraph/conc/pool"
)

type Server struct {
	chainConfig    config.RSS3Chain
	billingConfig  config.Billing
	databaseClient database.Client
	controlClient  *control.StateClientWriter
	redisClient    *redis.Client
}

func (s *Server) Run(ctx context.Context) error {
	errorPool := pool.New().WithContext(ctx).WithCancelOnError().WithFirstError()

	// Run L2 indexer.
	errorPool.Go(func(ctx context.Context) error {
		l2Config := l2.Config{
			Endpoint: s.chainConfig.EndpointL2,
		}

		serverL2, err := l2.NewServer(ctx, s.databaseClient, s.controlClient, s.redisClient, l2Config, s.billingConfig)
		if err != nil {
			return err
		}

		return serverL2.Run(ctx)
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

func New(databaseClient database.Client, controlClient *control.StateClientWriter, redisClient *redis.Client, chainConfig config.RSS3Chain, billingConfig config.Billing) (*Server, error) {
	instance := Server{
		chainConfig:    chainConfig,
		billingConfig:  billingConfig,
		databaseClient: databaseClient,
		controlClient:  controlClient,
		redisClient:    redisClient,
	}

	return &instance, nil
}

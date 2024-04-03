package handlers

import (
	"github.com/redis/go-redis/v9"
	"github.com/rss3-network/gateway-common/control"
	"github.com/rss3-network/payment-processor/internal/service/hub/gen/oapi"
	"github.com/rss3-network/payment-processor/internal/service/hub/jwt"
	"github.com/rss3-network/payment-processor/internal/service/hub/siwe"
	"gorm.io/gorm"
)

var _ oapi.ServerInterface = (*App)(nil)

type App struct {
	controlClient  *control.StateClientWriter
	redisClient    *redis.Client
	databaseClient *gorm.DB
	jwtClient      *jwt.JWT
	siweClient     *siwe.SIWE
}

func NewApp(controlClient *control.StateClientWriter, redisClient *redis.Client, databaseClient *gorm.DB, jwtClient *jwt.JWT, siweClient *siwe.SIWE) (*App, error) {
	return &App{
		controlClient:  controlClient,
		redisClient:    redisClient,
		databaseClient: databaseClient,
		jwtClient:      jwtClient,
		siweClient:     siweClient,
	}, nil
}

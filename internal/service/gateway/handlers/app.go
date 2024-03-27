package handlers

import (
	"github.com/redis/go-redis/v9"
	"github.com/rss3-network/gateway-common/control"
	"github.com/rss3-network/payment-processor/internal/service/gateway/gen/oapi"
	"github.com/rss3-network/payment-processor/internal/service/gateway/jwt"
	"github.com/rss3-network/payment-processor/internal/service/gateway/siwe"
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

func NewApp(controlClient *control.StateClientWriter, redis *redis.Client, databaseClient *gorm.DB, jwtClient *jwt.JWT, siweClient *siwe.SIWE) (*App, error) {
	return &App{
		controlClient:  controlClient,
		redisClient:    redis,
		databaseClient: databaseClient,
		jwtClient:      jwtClient,
		siweClient:     siweClient,
	}, nil
}

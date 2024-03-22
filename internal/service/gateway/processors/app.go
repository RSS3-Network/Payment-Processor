package processors

import (
	"github.com/rss3-network/gateway-common/control"
	"gorm.io/gorm"
)

type App struct {
	controlClient  *control.StateClientWriter
	databaseClient *gorm.DB
}

func NewApp(controlClient *control.StateClientWriter, databaseClient *gorm.DB) (*App, error) {
	return &App{
		controlClient:  controlClient,
		databaseClient: databaseClient,
	}, nil
}

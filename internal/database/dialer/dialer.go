package dialer

import (
	"context"
	"fmt"

	"github.com/rss3-network/payment-processor/internal/config"
	"github.com/rss3-network/payment-processor/internal/database"
	"github.com/rss3-network/payment-processor/internal/database/dialer/postgresql"
)

func Dial(ctx context.Context, config *config.Database) (database.Client, error) {
	switch config.Driver {
	case database.DriverPostgreSQL:
		return postgresql.Dial(ctx, config.URI)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.Driver)
	}
}

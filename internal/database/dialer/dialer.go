package dialer

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-gateway/internal/config"
	"github.com/naturalselectionlabs/rss3-gateway/internal/database"
	"github.com/naturalselectionlabs/rss3-gateway/internal/database/dialer/cockroachdb"
)

func Dial(ctx context.Context, config *config.Database) (database.Client, error) {
	switch config.Driver {
	case database.DriverCockroachDB:
		return cockroachdb.Dial(ctx, config.URI)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.Driver)
	}
}

package indexer

import (
	"context"

	"github.com/naturalselectionlabs/rss3-gateway/internal/config"
	"github.com/naturalselectionlabs/rss3-gateway/internal/database"
	"github.com/naturalselectionlabs/rss3-gateway/internal/service/indexer/l2"
	"github.com/rss3-network/gateway-common/control"
	"github.com/sourcegraph/conc/pool"
)

type Server struct {
	config         config.RSS3Chain
	databaseClient database.Client
	controlClient  *control.StateClientWriter
	ruPerToken     int64
}

func (s *Server) Run(ctx context.Context) error {
	errorPool := pool.New().WithContext(ctx).WithCancelOnError().WithFirstError()

	// Run L2 indexer.
	errorPool.Go(func(ctx context.Context) error {
		l2Config := l2.Config{
			Endpoint: s.config.EndpointL2,
		}

		serverL2, err := l2.NewServer(ctx, s.databaseClient, s.controlClient, s.ruPerToken, l2Config)
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

func New(databaseClient database.Client, controlClient *control.StateClientWriter, ruPerToken int64, config config.RSS3Chain) (*Server, error) {
	instance := Server{
		config:         config,
		databaseClient: databaseClient,
		controlClient:  controlClient,
		ruPerToken:     ruPerToken,
	}

	return &instance, nil
}

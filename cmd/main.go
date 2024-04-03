package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/redis/go-redis/v9"
	"github.com/rss3-network/gateway-common/control"
	"github.com/rss3-network/payment-processor/internal/config"
	"github.com/rss3-network/payment-processor/internal/config/flag"
	"github.com/rss3-network/payment-processor/internal/database/dialer"
	"github.com/rss3-network/payment-processor/internal/service/hub"
	"github.com/rss3-network/payment-processor/internal/service/indexer"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

var flags *pflag.FlagSet

var indexCommand = &cobra.Command{
	Use: "index",
	RunE: func(cmd *cobra.Command, args []string) error {
		flags = cmd.PersistentFlags()

		cfg, err := config.Setup(lo.Must(flags.GetString(flag.KeyConfig)))
		if err != nil {
			return fmt.Errorf("setup config file: %w", err)
		}

		isDevEnv := cfg.Environment == config.EnvironmentDevelopment

		initializeLogger(isDevEnv)

		// Initialize database client
		databaseClient, err := dialer.Dial(cmd.Context(), cfg.Database)
		if err != nil {
			return fmt.Errorf("dial database: %w", err)
		}

		if err := databaseClient.Migrate(cmd.Context()); err != nil {
			return fmt.Errorf("migrate database: %w", err)
		}

		// Initialize control configurations
		controlClient, err := control.NewWriter(cfg.Gateway.Etcd.Endpoints, cfg.Gateway.Etcd.Username, cfg.Gateway.Etcd.Password)
		if err != nil {
			return fmt.Errorf("prepare control service: %w", err)
		}

		// Initialize Redis client
		options, err := redis.ParseURL(cfg.Redis.URI)
		if err != nil {
			return fmt.Errorf("parse redis uri: %w", err)
		}

		redisClient := redis.NewClient(options)

		// Start
		instance, err := indexer.New(databaseClient, controlClient, redisClient, cfg.Billing.RuPerToken, *cfg.RSS3Chain)
		if err != nil {
			return fmt.Errorf("create indexer: %w", err)
		}

		return instance.Run(cmd.Context())
	},
}

var command = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		flags = cmd.PersistentFlags()

		cfg, err := config.Setup(lo.Must(flags.GetString(flag.KeyConfig)))
		if err != nil {
			return fmt.Errorf("setup config file: %w", err)
		}

		isDevEnv := cfg.Environment == config.EnvironmentDevelopment

		initializeLogger(isDevEnv)

		// Initialize database client
		databaseClient, err := dialer.Dial(cmd.Context(), cfg.Database)
		if err != nil {
			return fmt.Errorf("dial database: %w", err)
		}

		if err := databaseClient.Migrate(cmd.Context()); err != nil {
			return fmt.Errorf("migrate database: %w", err)
		}

		// Initialize redis client
		options, err := redis.ParseURL(cfg.Redis.URI)
		if err != nil {
			return fmt.Errorf("parse redis uri: %w", err)
		}

		redisClient := redis.NewClient(options)

		// Initialize control configurations
		controlClient, err := control.NewWriter(cfg.Gateway.Etcd.Endpoints, cfg.Gateway.Etcd.Username, cfg.Gateway.Etcd.Password)
		if err != nil {
			return fmt.Errorf("prepare control service: %w", err)
		}

		// Start
		instance, err := hub.New(isDevEnv, databaseClient, redisClient, controlClient, *cfg.Gateway)
		if err != nil {
			return fmt.Errorf("create hub: %w", err)
		}

		return instance.Run(cmd.Context())
	},
}

func initializeLogger(isDevEnv bool) {
	if isDevEnv {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	} else {
		zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
	}
}

func init() {
	command.AddCommand(indexCommand)

	command.PersistentFlags().String(flag.KeyConfig, "./deploy/config.yaml", "config file path")
	indexCommand.PersistentFlags().String(flag.KeyConfig, "./deploy/config.yaml", "config file path")
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := command.ExecuteContext(ctx); err != nil {
		zap.L().Fatal("execute command", zap.Error(err))
	}
}

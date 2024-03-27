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
	"github.com/rss3-network/payment-processor/internal/service/epoch"
	"github.com/rss3-network/payment-processor/internal/service/gateway"
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

		config, err := config.Setup(lo.Must(flags.GetString(flag.KeyConfig)))
		if err != nil {
			return fmt.Errorf("setup config file: %w", err)
		}

		databaseClient, err := dialer.Dial(cmd.Context(), config.Database)
		if err != nil {
			return err
		}

		if err := databaseClient.Migrate(cmd.Context()); err != nil {
			return fmt.Errorf("migrate database: %w", err)
		}

		// Initialize control configurations
		controlClient, err := control.NewWriter(config.Gateway.Etcd.Endpoints)
		if err != nil {
			return fmt.Errorf("prepare control service: %w", err)
		}

		instance, err := indexer.New(databaseClient, controlClient, config.Billing.RuPerToken, *config.RSS3Chain)
		if err != nil {
			return err
		}

		return instance.Run(cmd.Context())
	},
}

var epochCommand = &cobra.Command{
	Use: "epoch",
	RunE: func(cmd *cobra.Command, args []string) error {
		flags = cmd.PersistentFlags()

		config, err := config.Setup(lo.Must(flags.GetString(flag.KeyConfig)))
		if err != nil {
			return fmt.Errorf("setup config file: %w", err)
		}

		databaseClient, err := dialer.Dial(cmd.Context(), config.Database)
		if err != nil {
			return err
		}

		if err := databaseClient.Migrate(cmd.Context()); err != nil {
			return fmt.Errorf("migrate database: %w", err)
		}

		options, err := redis.ParseURL(config.Redis.URI)
		if err != nil {
			return fmt.Errorf("parse redis uri: %w", err)
		}

		redisClient := redis.NewClient(options)

		instance, err := epoch.New(cmd.Context(), databaseClient, redisClient, *config)
		if err != nil {
			return err
		}

		return instance.Run(cmd.Context())
	},
}

var command = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		flags = cmd.PersistentFlags()

		config, err := config.Setup(lo.Must(flags.GetString(flag.KeyConfig)))
		if err != nil {
			return fmt.Errorf("setup config file: %w", err)
		}

		databaseClient, err := dialer.Dial(cmd.Context(), config.Database)
		if err != nil {
			return err
		}

		if err := databaseClient.Migrate(cmd.Context()); err != nil {
			return fmt.Errorf("migrate database: %w", err)
		}

		options, err := redis.ParseURL(config.Redis.URI)
		if err != nil {
			return fmt.Errorf("parse redis uri: %w", err)
		}

		redisClient := redis.NewClient(options)

		// Initialize control configurations
		controlClient, err := control.NewWriter(config.Gateway.Etcd.Endpoints)
		if err != nil {
			return fmt.Errorf("prepare control service: %w", err)
		}

		instance, err := gateway.New(databaseClient, redisClient, controlClient, *config.Gateway)
		if err != nil {
			return err
		}

		return instance.Run(cmd.Context())
	},
}

func initializeLogger() {
	if os.Getenv(config.Environment) == config.EnvironmentDevelopment {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	} else {
		zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
	}
}

func init() {
	initializeLogger()

	command.AddCommand(indexCommand)
	command.AddCommand(epochCommand)

	command.PersistentFlags().String(flag.KeyConfig, "./deploy/config.yaml", "config file path")
	indexCommand.PersistentFlags().String(flag.KeyConfig, "./deploy/config.yaml", "config file path")
	epochCommand.PersistentFlags().String(flag.KeyConfig, "./deploy/config.yaml", "config file path")
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := command.ExecuteContext(ctx); err != nil {
		zap.L().Fatal("execute command", zap.Error(err))
	}
}

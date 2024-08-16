package config

import (
	"fmt"
	"os"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/rss3-network/payment-processor/internal/database"
	"gopkg.in/yaml.v3"
)

const (
	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
)

type File struct {
	Environment string     `yaml:"environment" validate:"required" default:"development"`
	Database    *Database  `yaml:"database"`
	Redis       *Redis     `yaml:"redis"`
	RSS3Chain   *RSS3Chain `yaml:"rss3_chain"`
	Settler     *Settler   `yaml:"settler"`
	Gateway     *Gateway   `yaml:"gateway"`
	Billing     *Billing   `yaml:"billing"`
}

type Database struct {
	Driver database.Driver `mapstructure:"driver" validate:"required" default:"postgresql"`
	URI    string          `mapstructure:"uri" validate:"required" default:"postgres://postgres:password@localhost:5432/postgres"`
}

type Redis struct {
	URI string `mapstructure:"uri" validate:"required" default:"redis://localhost:6379/0"`
}

type RSS3Chain struct {
	EndpointL2 string `yaml:"endpoint_l2" validate:"required"`
}

type Gateway struct {
	API struct {
		Listen struct {
			Host     string `yaml:"host" default:"0.0.0.0"`
			Port     uint64 `yaml:"port" default:"5555"`
			PromPort uint64 `yaml:"prom_port" default:"9000"`
		} `yaml:"listen"`
		JWTKey     string `yaml:"jwt_key" validate:"required"`
		SIWEDomain string `yaml:"siwe_domain" validate:"required"`
	} `yaml:"api" validate:"required"`
	Kafka struct {
		Brokers []string `yaml:"brokers" validate:"required"`
		Topic   string   `yaml:"topic" validate:"required"`
	} `yaml:"kafka" validate:"required"`
	Etcd struct {
		Endpoints []string `yaml:"endpoints" validate:"required"`
		Username  *string  `yaml:"username"`
		Password  *string  `yaml:"password"`
	} `yaml:"etcd" validate:"required"`
}

type Billing struct {
	RuPerToken        int64 `yaml:"ru_per_token" default:"1000"`
	SlackNotification struct {
		BotToken       string `yaml:"bot_token"`
		Channel        string `yaml:"channel"`
		BlockchainScan string `yaml:"blockchain_scan" validate:"required"`
	} `yaml:"slack_notification"`
}

type Settler struct {
	PrivateKey     string `yaml:"private_key"`
	WalletAddress  string `yaml:"wallet_address"`
	SignerEndpoint string `yaml:"signer_endpoint"`
	// EpochIntervalInHours
	EpochIntervalInHours int    `yaml:"epoch_interval_in_hours" default:"18"`
	GasLimit             uint64 `yaml:"gas_limit" default:"2500000"`
	// BatchSize is the number of Nodes to process in each batch.
	// This is to prevent the contract call from running out of gas.
	BatchSize int `yaml:"batch_size" default:"200"`
}

func Setup(configFilePath string) (*File, error) {
	// Read config file.
	config, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	// Unmarshal config file.
	var configFile File
	if err := yaml.Unmarshal(config, &configFile); err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	// Set default values.yaml.
	if err := defaults.Set(&configFile); err != nil {
		return nil, fmt.Errorf("set default values.yaml: %w", err)
	}

	// Validate config values.yaml.
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&configFile); err != nil {
		return nil, fmt.Errorf("validate config file: %w", err)
	}

	return &configFile, nil
}

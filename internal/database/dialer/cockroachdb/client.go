package cockroachdb

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"github.com/naturalselectionlabs/rss3-gateway/internal/database"
	"github.com/naturalselectionlabs/rss3-gateway/internal/database/dialer/cockroachdb/table"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var _ database.Client = (*client)(nil)

//go:embed migration/*.sql
var migrationFS embed.FS

type client struct {
	database *gorm.DB
}

func (c *client) Migrate(ctx context.Context) error {
	goose.SetBaseFS(migrationFS)
	goose.SetTableName("versions")
	goose.SetLogger(&database.SugaredLogger{Logger: zap.L().Sugar()})

	if err := goose.SetDialect(new(postgres.Dialector).Name()); err != nil {
		return fmt.Errorf("set migration dialect: %w", err)
	}

	connector, err := c.database.DB()
	if err != nil {
		return fmt.Errorf("get database connector: %w", err)
	}

	return goose.UpContext(ctx, connector, "migration")
}

func (c *client) WithTransaction(ctx context.Context, transactionFunction func(ctx context.Context, client database.Client) error, transactionOptions ...*sql.TxOptions) error {
	transaction, err := c.Begin(ctx, transactionOptions...)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	if err := transactionFunction(ctx, transaction); err != nil {
		_ = transaction.Rollback()

		return fmt.Errorf("execute transaction: %w", err)
	}

	if err := transaction.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (c *client) Begin(ctx context.Context, transactionOptions ...*sql.TxOptions) (database.Client, error) {
	transaction := c.database.WithContext(ctx).Begin(transactionOptions...)
	if err := transaction.Error; err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}

	return &client{database: transaction}, nil
}

func (c *client) Rollback() error {
	return c.database.Rollback().Error
}

func (c *client) Commit() error {
	return c.database.Commit().Error
}

func (c *client) RollbackBlock(ctx context.Context, chainID, blockNUmber uint64) error {
	databaseClient := c.database.WithContext(ctx)

	// Delete billing records
	if err := databaseClient.
		Where(`"chain_id" = ? AND "block_number" >= ?`, chainID, blockNUmber).
		Delete(&table.BillingRecordDeposited{}).
		Error; err != nil {
		return fmt.Errorf("delete billing record deposited: %w", err)
	}

	if err := databaseClient.
		Where(`"chain_id" = ? AND "block_number" >= ?`, chainID, blockNUmber).
		Delete(&table.BillingRecordWithdrawal{}).
		Error; err != nil {
		return fmt.Errorf("delete billing record withdrawal: %w", err)
	}

	if err := databaseClient.
		Where(`"chain_id" = ? AND "block_number" >= ?`, chainID, blockNUmber).
		Delete(&table.BillingRecordCollected{}).
		Error; err != nil {
		return fmt.Errorf("delete billing record collected: %w", err)
	}

	return nil
}

// Raw : get the raw db client
func (c *client) Raw() *gorm.DB {
	return c.database
}

// Dial dials a database.
func Dial(_ context.Context, dataSourceName string) (database.Client, error) {
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()

	config := gorm.Config{
		Logger: logger,
	}

	databaseClient, err := gorm.Open(postgres.Open(dataSourceName), &config)
	if err != nil {
		return nil, fmt.Errorf("dial database: %w", err)
	}

	return &client{
		database: databaseClient,
	}, nil
}

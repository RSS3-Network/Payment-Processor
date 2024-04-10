package database

import (
	"context"
	"database/sql"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pressly/goose/v3"
	"github.com/rss3-network/payment-processor/schema"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ErrorRowNotFound = errors.New("row not found")
)

type Client interface {
	Session
	Transaction

	RollbackBlock(ctx context.Context, chainID, blockNUmber uint64) error

	FindCheckpoint(ctx context.Context, chainID uint64) (*schema.Checkpoint, error)
	SaveCheckpoint(ctx context.Context, checkpoint *schema.Checkpoint) error

	SaveBillingRecordDeposited(ctx context.Context, billingRecord *schema.BillingRecordDeposited) error
	SaveBillingRecordWithdrawal(ctx context.Context, billingRecord *schema.BillingRecordWithdrawal) error
	SaveBillingRecordCollected(ctx context.Context, billingRecord *schema.BillingRecordCollected) error

	PrepareBillingCollectTokens(ctx context.Context, nowTime time.Time) (*map[common.Address]schema.BillingCollectDataPerAddress, error)
	PrepareBillingWithdrawTokens(ctx context.Context) (*map[common.Address]float64, error)
	UpdateBillingRuLimit(ctx context.Context, succeededUsersWithRu map[common.Address]int64) error

	GatewayDeposit(ctx context.Context, address common.Address, ruIncrease int64) (bool, error)

	FindNodeRequestRewardsByEpoch(ctx context.Context, epoch *big.Int) ([]schema.NodeRequestRecord, error)
	SaveNodeRequestCount(ctx context.Context, record *schema.NodeRequestRecord) error
	SetNodeRequestRewards(ctx context.Context, nodeAddr common.Address, reward *big.Int) error

	Raw() *gorm.DB
}

type Session interface {
	Migrate(ctx context.Context) error
	WithTransaction(ctx context.Context, transactionFunction func(ctx context.Context, client Client) error, transactionOptions ...*sql.TxOptions) error
	Begin(ctx context.Context, transactionOptions ...*sql.TxOptions) (Client, error)
}

type Transaction interface {
	Rollback() error
	Commit() error
}

var _ goose.Logger = (*SugaredLogger)(nil)

type SugaredLogger struct {
	Logger *zap.SugaredLogger
}

func (s SugaredLogger) Fatalf(format string, v ...interface{}) {
	s.Logger.Fatalf(format, v...)
}

func (s SugaredLogger) Printf(format string, v ...interface{}) {
	s.Logger.Infof(format, v...)
}

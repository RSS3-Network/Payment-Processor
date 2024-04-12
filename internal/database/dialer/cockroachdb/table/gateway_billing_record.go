package table

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/payment-processor/schema"
	"github.com/shopspring/decimal"
	gormSchema "gorm.io/gorm/schema"
)

var (
	_ gormSchema.Tabler = (*BillingRecordDeposited)(nil)
	_ gormSchema.Tabler = (*BillingRecordWithdrawal)(nil)
	_ gormSchema.Tabler = (*BillingRecordCollected)(nil)
)

type BillingRecordBase struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	TxHash         common.Hash `gorm:"primaryKey;type:bytea;column:tx_hash"`
	Index          uint        `gorm:"column:index"`
	ChainID        uint64      `gorm:"column:chain_id"`
	BlockTimestamp time.Time   `gorm:"index;column:block_timestamp"`
	BlockNumber    uint64      `gorm:"column:block_number"`

	User   common.Address  `gorm:"type:bytea;column:user"`
	Amount decimal.Decimal `gorm:"column:amount"`
}

type BillingRecordDeposited struct {
	BillingRecordBase
}

type BillingRecordWithdrawal struct {
	BillingRecordBase

	Fee decimal.Decimal
}

type BillingRecordCollected struct {
	BillingRecordBase
}

func (r *BillingRecordDeposited) TableName() string {
	return "gateway.br_deposited"
}

func (r *BillingRecordWithdrawal) TableName() string {
	return "gateway.br_withdrawn"
}

func (r *BillingRecordCollected) TableName() string {
	return "gateway.br_collected"
}

func (r *BillingRecordDeposited) Import(billingRecord schema.BillingRecordDeposited) error {
	r.TxHash = billingRecord.TxHash
	r.Index = billingRecord.Index
	r.ChainID = billingRecord.ChainID
	r.BlockTimestamp = billingRecord.BlockTimestamp
	r.BlockNumber = billingRecord.BlockNumber.Uint64()
	r.User = billingRecord.User
	r.Amount = decimal.NewFromBigInt(billingRecord.Amount, 0)

	return nil
}

func (r *BillingRecordDeposited) Export() (*schema.BillingRecordDeposited, error) {
	billingRecord := schema.BillingRecordDeposited{
		BillingRecordBase: schema.BillingRecordBase{
			TxHash:         r.TxHash,
			Index:          r.Index,
			ChainID:        r.ChainID,
			BlockTimestamp: r.BlockTimestamp,
			BlockNumber:    new(big.Int).SetUint64(r.BlockNumber),
			User:           r.User,
			Amount:         r.Amount.BigInt(),
		},
	}

	return &billingRecord, nil
}

func (r *BillingRecordWithdrawal) Import(billingRecord schema.BillingRecordWithdrawal) error {
	r.TxHash = billingRecord.TxHash
	r.Index = billingRecord.Index
	r.ChainID = billingRecord.ChainID
	r.BlockTimestamp = billingRecord.BlockTimestamp
	r.BlockNumber = billingRecord.BlockNumber.Uint64()
	r.User = billingRecord.User
	r.Amount = decimal.NewFromBigInt(billingRecord.Amount, 0)
	r.Fee = decimal.NewFromBigInt(billingRecord.Fee, 0)

	return nil
}

func (r *BillingRecordWithdrawal) Export() (*schema.BillingRecordWithdrawal, error) {
	billingRecord := schema.BillingRecordWithdrawal{
		BillingRecordBase: schema.BillingRecordBase{
			TxHash:         r.TxHash,
			Index:          r.Index,
			ChainID:        r.ChainID,
			BlockTimestamp: r.BlockTimestamp,
			BlockNumber:    new(big.Int).SetUint64(r.BlockNumber),
			User:           r.User,
			Amount:         r.Amount.BigInt(),
		},
		Fee: r.Fee.BigInt(),
	}

	return &billingRecord, nil
}

func (r *BillingRecordCollected) Import(billingRecord schema.BillingRecordCollected) error {
	r.TxHash = billingRecord.TxHash
	r.Index = billingRecord.Index
	r.ChainID = billingRecord.ChainID
	r.BlockTimestamp = billingRecord.BlockTimestamp
	r.BlockNumber = billingRecord.BlockNumber.Uint64()
	r.User = billingRecord.User
	r.Amount = decimal.NewFromBigInt(billingRecord.Amount, 0)

	return nil
}

func (r *BillingRecordCollected) Export() (*schema.BillingRecordCollected, error) {
	billingRecord := schema.BillingRecordCollected{
		BillingRecordBase: schema.BillingRecordBase{
			TxHash:         r.TxHash,
			Index:          r.Index,
			ChainID:        r.ChainID,
			BlockTimestamp: r.BlockTimestamp,
			BlockNumber:    new(big.Int).SetUint64(r.BlockNumber),
			User:           r.User,
			Amount:         r.Amount.BigInt(),
		},
	}

	return &billingRecord, nil
}

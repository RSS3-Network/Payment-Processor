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
	_ gormSchema.Tabler = (*BillingRecordWithdrawal)(nil)
)

type BillingRecordWithdrawal struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	TxHash         common.Hash `gorm:"primaryKey;type:bytea;column:tx_hash"`
	Index          uint        `gorm:"primaryKey;column:index"`
	ChainID        uint64      `gorm:"index;column:chain_id"`
	BlockTimestamp time.Time   `gorm:"index;column:block_timestamp"`
	BlockNumber    uint64      `gorm:"index;column:block_number"`

	User   common.Address  `gorm:"type:bytea;column:user"`
	Amount decimal.Decimal `gorm:"column:amount"`

	Fee decimal.Decimal `gorm:"column:fee"`
}

func (r *BillingRecordWithdrawal) TableName() string {
	return "br_withdrawn"
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

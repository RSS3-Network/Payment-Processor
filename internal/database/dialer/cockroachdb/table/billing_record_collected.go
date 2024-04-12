package table

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/payment-processor/schema"
	"github.com/shopspring/decimal"
	gormSchema "gorm.io/gorm/schema"
	"math/big"
	"time"
)

var (
	_ gormSchema.Tabler = (*BillingRecordCollected)(nil)
)

type BillingRecordCollected struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	TxHash         common.Hash `gorm:"primaryKey;type:bytea;column:tx_hash"`
	Index          uint        `gorm:"column:index"`
	ChainID        uint64      `gorm:"index;column:chain_id"`
	BlockTimestamp time.Time   `gorm:"index;column:block_timestamp"`
	BlockNumber    uint64      `gorm:"index;column:block_number"`

	User   common.Address  `gorm:"type:bytea;column:user"`
	Amount decimal.Decimal `gorm:"column:amount"`
}

func (r *BillingRecordCollected) TableName() string {
	return "br_collected"
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

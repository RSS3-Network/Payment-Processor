package table

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/payment-processor/schema"
	"github.com/shopspring/decimal"
	gormSchema "gorm.io/gorm/schema"
	"math/big"
)

var (
	_ gormSchema.Tabler = (*NodeRequestRecord)(nil)
)

type NodeRequestRecord struct {
	// Composite Primary Key (epoch-node_address)
	Epoch       uint64         `gorm:"primaryKey;autoIncrement:false;column:epoch"`
	NodeAddress common.Address `gorm:"primaryKey;type:bytea;column:node_address"`

	// Event data (ignore `operationRewards` `stakingRewards` `taxAmounts` because they are not required here)
	RequestCount uint64 `gorm:"column:request_counts"` // From GI on-chain requests

	RequestReward decimal.Decimal `gorm:"column:request_rewards"` // Update after request rewards distribute
}

func (r *NodeRequestRecord) TableName() string {
	return "node_request_record"
}

func (r *NodeRequestRecord) Import(nodeRequestRewards schema.NodeRequestRecord) error {
	r.NodeAddress = nodeRequestRewards.NodeAddress
	r.Epoch = nodeRequestRewards.Epoch.Uint64()
	r.RequestCount = nodeRequestRewards.RequestCount.Uint64()
	r.RequestReward = decimal.NewFromBigInt(nodeRequestRewards.RequestReward, 0)

	return nil
}

func (r *NodeRequestRecord) Export() (*schema.NodeRequestRecord, error) {
	nodeRequestRewards := schema.NodeRequestRecord{
		NodeAddress:   r.NodeAddress,
		Epoch:         new(big.Int).SetUint64(r.Epoch),
		RequestCount:  new(big.Int).SetUint64(r.RequestCount),
		RequestReward: r.RequestReward.BigInt(),
	}

	return &nodeRequestRewards, nil
}

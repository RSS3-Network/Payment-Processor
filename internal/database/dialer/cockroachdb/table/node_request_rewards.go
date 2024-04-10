package table

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/payment-processor/schema"
	"github.com/shopspring/decimal"
	gormSchema "gorm.io/gorm/schema"
	"math/big"
)

var (
	_ gormSchema.Tabler = (*NodeRequestRewards)(nil)
)

type NodeRequestRewards struct {
	// Composite Primary Key (epoch-node_address)
	NodeAddress   common.Address `gorm:"primaryKey;type:bytea;column:node_address"`
	Epoch         uint64         `gorm:"primaryKey;autoIncrement:false;column:epoch"`
	RequestCounts uint64         `gorm:"column:request_counts"` // From GI on-chain requests

	RequestRewards decimal.Decimal `gorm:"column:request_rewards"` // Update after request rewards distribute
}

func (r *NodeRequestRewards) TableName() string {
	return "node_request_rewards"
}

func (r *NodeRequestRewards) Import(nodeRequestRewards schema.NodeRequestRewards) error {
	r.NodeAddress = nodeRequestRewards.NodeAddress
	r.Epoch = nodeRequestRewards.Epoch.Uint64()
	r.RequestCounts = nodeRequestRewards.RequestCounts.Uint64()
	r.RequestRewards = decimal.NewFromBigInt(nodeRequestRewards.RequestRewards, 0)

	return nil
}

func (r *NodeRequestRewards) Export() (*schema.NodeRequestRewards, error) {
	nodeRequestRewards := schema.NodeRequestRewards{
		NodeAddress:    r.NodeAddress,
		Epoch:          new(big.Int).SetUint64(r.Epoch),
		RequestCounts:  new(big.Int).SetUint64(r.RequestCounts),
		RequestRewards: r.RequestRewards.BigInt(),
	}

	return &nodeRequestRewards, nil
}

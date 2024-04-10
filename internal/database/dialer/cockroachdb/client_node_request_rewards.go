package cockroachdb

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/payment-processor/internal/database/dialer/cockroachdb/table"
	"github.com/rss3-network/payment-processor/schema"
	"github.com/shopspring/decimal"
	"math/big"
)

func (c *client) FindNodeRequestRewardsByEpoch(ctx context.Context, epoch *big.Int) ([]*schema.NodeRequestRewards, error) {
	var rewardsRecord []table.NodeRequestRewards

	if err := c.database.
		WithContext(ctx).
		Find(&rewardsRecord, table.NodeRequestRewards{Epoch: epoch.Uint64()}).Error; err != nil {
		return nil, err
	}

	rewardsSchema := make([]*schema.NodeRequestRewards, len(rewardsRecord))

	var err error

	for i, reward := range rewardsRecord {
		rewardsSchema[i], err = reward.Export()
		if err != nil {
			return nil, err
		}
	}

	return rewardsSchema, nil
}

func (c *client) SaveNodeRequestRewards(ctx context.Context, rewards *schema.NodeRequestRewards) error {
	var value table.NodeRequestRewards
	if err := value.Import(*rewards); err != nil {
		return fmt.Errorf("import node request rewards: %w", err)
	}

	return c.database.WithContext(ctx).Create(&value).Error
}

func (c *client) SetNodeRequestRewards(ctx context.Context, epoch *big.Int, nodeAddr common.Address, rewards *big.Int) error {
	return c.database.WithContext(ctx).
		Updates(table.NodeRequestRewards{
			RequestRewards: decimal.NewFromBigInt(rewards, 0),
		}).
		Where(table.NodeRequestRewards{
			Epoch:       epoch.Uint64(),
			NodeAddress: nodeAddr,
		}).
		Error
}

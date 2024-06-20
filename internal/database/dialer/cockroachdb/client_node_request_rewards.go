package cockroachdb

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/payment-processor/internal/database/dialer/cockroachdb/table"
	"github.com/rss3-network/payment-processor/schema"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func (c *client) FindNodeRequestRewardsByEpoch(ctx context.Context, epoch *big.Int) ([]*schema.NodeRequestRecord, error) {
	var rewardsRecord []table.NodeRequestRecord

	zap.L().Debug("FindNodeRequestRewardsByEpoch: before query")

	if err := c.database.
		WithContext(ctx).
		Find(&rewardsRecord, table.NodeRequestRecord{Epoch: epoch.Uint64()}).Error; err != nil {
		return nil, err
	}

	zap.L().Debug("FindNodeRequestRewardsByEpoch: after query")

	rewardsSchema := make([]*schema.NodeRequestRecord, len(rewardsRecord))

	var err error

	zap.L().Debug("FindNodeRequestRewardsByEpoch: before export")

	for i, reward := range rewardsRecord {
		rewardsSchema[i], err = reward.Export()
		if err != nil {
			return nil, err
		}
	}

	zap.L().Debug("FindNodeRequestRewardsByEpoch: after export")

	return rewardsSchema, nil
}

func (c *client) SaveNodeRequestCount(ctx context.Context, record *schema.NodeRequestRecord) error {
	var value table.NodeRequestRecord

	// Record count means it doesn't include RequestReward data.
	// So to ensure this, let's have an additional check
	processedRecord := *record
	if processedRecord.RequestReward == nil {
		processedRecord.RequestReward = big.NewInt(0)
	}

	if err := value.Import(processedRecord); err != nil {
		return fmt.Errorf("import node request rewards: %w", err)
	}

	return c.database.WithContext(ctx).Create(&value).Error
}

func (c *client) SetNodeRequestRewards(ctx context.Context, epoch *big.Int, nodeAddr common.Address, reward *big.Int) error {
	return c.database.WithContext(ctx).
		Where(table.NodeRequestRecord{
			Epoch:       epoch.Uint64(),
			NodeAddress: nodeAddr,
		}).
		Updates(table.NodeRequestRecord{
			RequestReward: decimal.NewFromBigInt(reward, 0),
		}).
		Error
}

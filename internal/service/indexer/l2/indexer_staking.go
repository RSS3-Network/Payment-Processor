package l2

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rss3-network/payment-processor/contract/l2"
	"github.com/rss3-network/payment-processor/internal/database"
	"github.com/rss3-network/payment-processor/schema"
	"go.uber.org/zap"
)

func (s *server) indexStakingLog(ctx context.Context, header *types.Header, _ *types.Transaction, _ *types.Receipt, log *types.Log, _ int, databaseTransaction database.Client) (*big.Int, error) {
	switch eventHash := log.Topics[0]; eventHash {
	case l2.EventHashStakingRewardDistributed:
		return s.indexStakingDistributeRewardsLog(ctx, header, log, databaseTransaction)
	default:
		return nil, nil
	}
}

func (s *server) indexStakingDistributeRewardsLog(ctx context.Context, _ *types.Header, log *types.Log, databaseTransaction database.Client) (*big.Int, error) {
	stakingDistributeRewardsEvent, err := s.contractStaking.ParseRewardDistributed(*log)
	if err != nil {
		return nil, fmt.Errorf("parse RewardDistributed event: %w", err)
	}

	// The workflow here is:
	// 1. collect all data (nodeAddrs and requestCounts)
	// 2. if this is the latest batch of this epoch (isFinal = true):
	//    1. lock state with redis (SET paymentprocessor.lock.settlement locked NX)
	//    2. billing: collect tokens (count how much collected)
	//    3. billing: withdraw tokens
	//    4. calc requestCounts percentage of each node collected in step 1
	//    5. billing: distribute request rewards

	// Step 1: collect all data
	bigZero := big.NewInt(0)
	for i, nodeAddr := range stakingDistributeRewardsEvent.NodeAddrs {
		if stakingDistributeRewardsEvent.RequestCounts[i].Cmp(bigZero) != 1 {
			// No contribution in this epoch, skip
			zap.L().Debug("node has no contribution in this epoch, skip",
				zap.Uint64("epoch", stakingDistributeRewardsEvent.Epoch.Uint64()),
				zap.String("nodeAddr", nodeAddr.String()),
			)

			continue
		}

		// Save all node request count
		//   Since closeEpoch depends on this,
		//   here we use databaseClient directly rather than insert transaction
		err = databaseTransaction.SaveNodeRequestCount(ctx, &schema.NodeRequestRecord{
			NodeAddress:  nodeAddr,
			Epoch:        stakingDistributeRewardsEvent.Epoch,
			RequestCount: stakingDistributeRewardsEvent.RequestCounts[i],
		})
		if err != nil {
			// Log error and abort
			zap.L().Error("save node request count",
				zap.Uint64("epoch", stakingDistributeRewardsEvent.Epoch.Uint64()),
				zap.String("nodeAddr", nodeAddr.Hex()),
				zap.Int64("requestCount", stakingDistributeRewardsEvent.RequestCounts[i].Int64()),
				zap.Error(err),
			)

			return nil, fmt.Errorf("save node %s request count %d: %w", nodeAddr.Hex(), stakingDistributeRewardsEvent.RequestCounts[i].Int64(), err)
		}
	}

	return stakingDistributeRewardsEvent.Epoch, nil // No error
}

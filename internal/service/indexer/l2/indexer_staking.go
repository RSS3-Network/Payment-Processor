package l2

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rss3-network/payment-processor/contract/l2"
	"github.com/rss3-network/payment-processor/internal/database"
	"github.com/rss3-network/payment-processor/internal/service/indexer/constants"
	"github.com/rss3-network/payment-processor/schema"
	"go.uber.org/zap"
	"math/big"
)

func (s *server) indexStakingLog(ctx context.Context, header *types.Header, transaction *types.Transaction, receipt *types.Receipt, log *types.Log, logIndex int, databaseTransaction database.Client) error {
	switch eventHash := log.Topics[0]; eventHash {
	case l2.EventHashStakingRewardDistributed:
		return s.indexStakingDistributeRewardsLog(ctx, header, transaction, receipt, log, logIndex, databaseTransaction)
	default:
		return nil
	}
}

func (s *server) indexStakingDistributeRewardsLog(ctx context.Context, header *types.Header, transaction *types.Transaction, receipt *types.Receipt, log *types.Log, _ int, databaseTransaction database.Client) error {
	stakingDistributeRewardsEvent, err := s.contractStaking.ParseRewardDistributed(*log)
	if err != nil {
		return fmt.Errorf("parse RewardDistributed event: %w", err)
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
	for i, nodeAddr := range stakingDistributeRewardsEvent.NodeAddrs {
		err = s.databaseClient.SaveNodeRequestCount(ctx, &schema.NodeRequestRecord{
			NodeAddress:  nodeAddr,
			Epoch:        stakingDistributeRewardsEvent.Epoch,
			RequestCount: stakingDistributeRewardsEvent.RequestCounts[i],
		})
		if err != nil {
			// Error, but no need to abort
			zap.L().Error("save node request count", zap.Any("index", i), zap.Any("event", stakingDistributeRewardsEvent), zap.Error(err))
		}
	}

	// Step 2: check if is last batch of this epoch
	isStillProceeding, err := s.contractStaking.IsSettlementPhase(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: header.Number,
	})
	if err != nil {
		return fmt.Errorf("failed to check is settlement phase: %w", err)
	}
	if isStillProceeding {
		// Not last batch, return
		return nil
	} // else is last batch, start proceed

	// 2.1. Set mutex lock
	isMutexLockSuccessful, err := s.redisClient.SetNX(ctx, constants.EpochMutexLockKey, 1, constants.EpochMutexExpiration).Result()
	if err != nil {
		return fmt.Errorf("failed to set mutex lock with redis: %w", err)
	}
	if !isMutexLockSuccessful {
		// A process already running, skip
		return nil
	}

	// Defer release mutex lock
	defer s.redisClient.Del(ctx, constants.EpochMutexLockKey)

	// 2.2-3. billing
	totalCollected, err := s.billingFlow(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute billing flow: %w", err)
	}

	// 2.4. calc request percentage
	allNodes, err := s.databaseClient.FindNodeRequestRewardsByEpoch(ctx, stakingDistributeRewardsEvent.Epoch)
	if err != nil {
		return fmt.Errorf("failed to find node requests record: %w", err)
	}

	// Sum all requests count
	totalRequestCount := big.NewInt(0)
	for _, node := range allNodes {
		totalRequestCount.Add(totalRequestCount, node.RequestCount)
	}

	// Calculate reward per request
	rewardPerRequest := new(big.Int).Quo(totalCollected, totalRequestCount)
	zap.L().Info(
		"epoch reward per request",
		zap.Uint64("epoch", stakingDistributeRewardsEvent.Epoch.Uint64()),
		zap.String("totalRewards", totalCollected.String()),
		zap.String("totalRequests", totalRequestCount.String()),
		zap.String("rewardPerRequest", rewardPerRequest.String()),
	)

	// Calculate reward for nodes
	rewardNodesAddress := make([]common.Address, len(allNodes))
	rewardNodesAmount := make([]*big.Int, len(allNodes))

	for i, node := range allNodes {
		// Calculate reward per node
		rewardNodesAddress[i] = node.NodeAddress
		rewardNodesAmount[i] = new(big.Int).Mul(rewardPerRequest, node.RequestCount)

		// Save into database
		err = s.databaseClient.SetNodeRequestRewards(ctx, rewardNodesAddress[i], rewardNodesAmount[i])
		if err != nil {
			// Error, but no need to abort
			zap.L().Error("update node request rewards", zap.String("address", rewardNodesAddress[i].String()), zap.String("amount", rewardNodesAmount[i].String()), zap.Any("node", node), zap.Error(err))
		}
	}

	// 2.5. billing: distribute request rewards
	err = s.distributeRequestRewards(ctx, rewardNodesAddress, rewardNodesAmount)
	if err != nil {
		return fmt.Errorf("trigger distribute request rewards: %w", err)
	}

	return nil

}

package l2

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/payment-processor/internal/service/indexer/constants"
	"go.uber.org/zap"
)

func (s *server) closeEpoch(ctx context.Context, blockNumber *big.Int, epoch *big.Int) error {
	isStillProceeding, err := s.contractStaking.IsSettlementPhase(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: blockNumber,
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

	s.closeEpochExec(ctx, epoch) // This cannot retry when error happens, so just report errors to slack rather than retry it

	return nil
}

func (s *server) closeEpochExec(ctx context.Context, epoch *big.Int) {
	// 2.2-3. billing
	zap.L().Debug("closeEpochExec: 2.2-3. billing")

	totalCollected, err := s.billingFlow(ctx, epoch)
	if err != nil {
		zap.L().Error("failed to execute closeEpochExec: 2.2-3. billing", zap.Error(err))
		s.ReportFailedTransactionToSlack(err, nil, "closeEpochExec: 2.2-3. billing", nil, nil)

		return
	}

	zap.L().Debug("billing flow total collect", zap.String("token", totalCollected.String()))

	if totalCollected.Cmp(big.NewInt(0)) == 0 {
		// No request fees collect in this epoch, skip
		zap.L().Info("no request fees collect in this epoch, skip")

		return
	}

	// 2.4. calc request percentage
	zap.L().Debug("closeEpochExec: 2.4. calc request percentage")

	allNodes, err := s.databaseClient.FindNodeRequestRewardsByEpoch(ctx, epoch)
	if err != nil {
		zap.L().Error("failed to execute closeEpochExec: 2.4. calc request percentage", zap.Error(err))
		s.ReportFailedTransactionToSlack(err, nil, "closeEpochExec: 2.4. calc request percentage", nil, nil)

		return
	}

	if len(allNodes) == 0 {
		zap.L().Debug("No active nodes in current epoch, skip")
		return
	}

	zap.L().Debug("All nodes found, start contribution calc", zap.Uint64("epoch", epoch.Uint64()), zap.Any("nodes", allNodes))

	// Sum all requests count
	totalRequestCount := big.NewInt(0)

	for _, node := range allNodes {
		totalRequestCount.Add(totalRequestCount, node.RequestCount)
	}

	if totalRequestCount.Cmp(big.NewInt(0)) == 0 {
		// No requests happened in this epoch, skip
		zap.L().Info("no requests happened in this epoch, skip")

		return
	}

	// Calculate reward per request
	rewardPerRequest := new(big.Int).Quo(totalCollected, totalRequestCount)
	zap.L().Info(
		"epoch reward per request",
		zap.Uint64("epoch", epoch.Uint64()),
		zap.String("totalRewards", totalCollected.String()),
		zap.String("totalRequests", totalRequestCount.String()),
		zap.String("rewardPerRequest", rewardPerRequest.String()),
	)

	// Calculate reward for nodes
	rewardNodesAddress := []common.Address{}
	rewardNodesAmount := []*big.Int{}

	for _, node := range allNodes {
		// Calculate reward per node
		reward := new(big.Int).Mul(rewardPerRequest, node.RequestCount)

		// Save into database
		err = s.databaseClient.SetNodeRequestRewards(ctx, epoch, node.NodeAddress, reward)
		if err != nil {
			// Error, but no need to abort
			zap.L().Error("update node request rewards", zap.String("address", node.NodeAddress.String()), zap.String("amount", reward.String()), zap.Any("node", node), zap.Error(err))
		}

		rewardNodesAddress = append(rewardNodesAddress, node.NodeAddress)
		rewardNodesAmount = append(rewardNodesAmount, reward)
	}

	// 2.5. billing: distribute request rewards
	zap.L().Debug("closeEpochExec: 2.5. billing: distribute request rewards")

	s.distributeRequestRewards(ctx, rewardNodesAddress, rewardNodesAmount)
}

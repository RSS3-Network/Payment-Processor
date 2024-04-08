package l2

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rss3-network/payment-processor/contract/l2"
	"github.com/rss3-network/payment-processor/internal/database"
	"github.com/rss3-network/payment-processor/internal/service/indexer/constants"
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

	// The workflow here is:
	// 1. collect all data (nodeAddrs and requestCounts)
	// 2. if this is the latest batch of this epoch (isFinal = true):
	//    1. lock state with redis (SET paymentprocessor.lock.settlement locked NX)
	//    2. billing: collect tokens (count how much collected)
	//    3. billing: withdraw tokens
	//    4. calc requestCounts percentage of each node collected in step 1
	//    5. billing: distribute request rewards

	// TODO:
	//   1. Done ~~prepare settlement contract (seems like there's no DistributeRewards event)~~
	//   2. prepare database tables required with new orm technology (requires main branch update)

	// Step 1: collect all data
	// TODO

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
	// TODO

	// 2.5. billing: distribute request rewards
	// TODO

	return nil

}

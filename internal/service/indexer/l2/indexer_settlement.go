package l2

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rss3-network/payment-processor/contract/l2"
	"github.com/rss3-network/payment-processor/internal/database"
)

func (s *server) indexSettlementLog(ctx context.Context, header *types.Header, transaction *types.Transaction, receipt *types.Receipt, log *types.Log, logIndex int, databaseTransaction database.Client) error {
	switch eventHash := log.Topics[0]; eventHash {
	case l2.EventHashSettlementDistributeRewards:
		return s.indexSettlementDistributeRewardsLog(ctx, header, transaction, receipt, log, logIndex, databaseTransaction)
	default:
		return nil
	}
}

func (s *server) indexSettlementDistributeRewardsLog(ctx context.Context, header *types.Header, transaction *types.Transaction, receipt *types.Receipt, log *types.Log, _ int, databaseTransaction database.Client) error {
	settlementDistributeRewardsEvent, err := s.contractSettlement.ParseDistributeRewards(*log)

	// The workflow here is:
	// 1. collect all data (nodeAddrs and requestCounts)
	// 2. if this is the latest batch of this epoch (isFinal = true):
	//    1. lock state with redis (SET paymentprocessor.lock.settlement locked NX)
	//    2. billing: collect tokens (count how much collected)
	//    3. billing: withdraw tokens
	//    4. calc requestCounts percentage of each node collected in step 1
	//    5. billing: distribute request rewards

	// TODO:
	//   1. prepare settlement contract (seems like there's no DistributeRewards event)
	//   2. prepare database tables required with new orm technology (requires main branch update)

}

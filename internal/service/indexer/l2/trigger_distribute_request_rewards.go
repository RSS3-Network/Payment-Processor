package l2

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/payment-processor/contract/l2"
	"go.uber.org/zap"
)

func (s *server) distributeRequestRewards(ctx context.Context, nodeAddress []common.Address, amounts []*big.Int) {
	for len(nodeAddress) > 0 {
		limit := len(nodeAddress)
		if limit > s.settlerConfig.BatchSize {
			limit = s.settlerConfig.BatchSize
		}

		err := s.triggerDistributeRequestRewards(ctx, nodeAddress[:limit], amounts[:limit])
		if err != nil {
			zap.L().Error("trigger distribute request rewards", zap.Any("nodeAddress", nodeAddress[:limit]), zap.Any("amount", amounts[:limit]), zap.Error(err))
		}

		nodeAddress = nodeAddress[limit:]
		amounts = amounts[limit:]
	}
}

func (s *server) triggerDistributeRequestRewards(ctx context.Context, nodeAddress []common.Address, amounts []*big.Int) error {
	// Trigger distribute requests contract.
	input, err := s.encodeInput(l2.BillingMetaData.ABI, l2.MethodDistributeRewards, nodeAddress, amounts)
	if err != nil {
		return fmt.Errorf("encode input: %w", err)
	}

	receipt, err := s.sendTransaction(ctx, input)
	if err != nil {
		s.ReportFailedTransactionToSlack(err, "", l2.MethodDistributeRewards, nodeAddress, amounts)
		return fmt.Errorf("send transaction receipt: %w", err)
	}

	zap.L().Info("distribute requests successfully", zap.String("tx", receipt.TxHash.String()))

	return nil
}

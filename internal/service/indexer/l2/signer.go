package l2

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rss3-network/payment-processor/common/txmgr"
	"github.com/rss3-network/payment-processor/contract/l2"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// sendTransaction sends the transaction and returns the receipt if successful
func (s *server) sendTransaction(ctx context.Context, input []byte) (*types.Receipt, error) {
	txCandidate := txmgr.TxCandidate{
		TxData:   input,
		To:       lo.ToPtr(l2.ContractMap[s.chainID.Uint64()].AddressBillingProxy),
		GasLimit: s.settlerConfig.GasLimit,
		Value:    big.NewInt(0),
	}

	receipt, err := s.txManager.Send(ctx, txCandidate)
	if err != nil {
		return nil, fmt.Errorf("failed to send tx: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		zap.L().Error("received an invalid transaction receipt", zap.String("tx", receipt.TxHash.String()))

		// select {} purposely block the process as it is a critical error and meaningless to continue
		// if panic() is called, the process will be restarted by the supervisor
		// we do not want that as it will be stuck in the same state
		// select {} // Move this process blocker after message report to notification system
		return receipt, fmt.Errorf("received an invalid transaction receipt")
	}

	// return the receipt if the transaction is successful
	return receipt, nil
}

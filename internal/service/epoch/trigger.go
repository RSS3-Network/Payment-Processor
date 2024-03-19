package epoch

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"strings"
)

const BatchSize = 200

func (s *Server) trigger(ctx context.Context, epoch uint64) error {
	if err := s.mutex.Lock(); err != nil {
		zap.L().Error("lock error", zap.String("key", s.mutex.Name()), zap.Error(err))

		return nil
	}

	defer func() {
		if _, err := s.mutex.Unlock(); err != nil {
			zap.L().Error("release lock error", zap.String("key", s.mutex.Name()), zap.Error(err))
		}
	}()

	// TODO: Collect tokens

	// TODO: Withdraw

	zap.L().Info("Reward distribution completed")

	return nil
}

func (s *Server) encodeInput(contractABI, methodName string, args ...interface{}) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}

	encodedArgs, err := parsedABI.Pack(methodName, args...)
	if err != nil {
		return nil, err
	}

	return encodedArgs, nil
}

func (s *Server) transactionReceipt(ctx context.Context, txHash common.Hash) error {
	for {
		receipt, err := s.ethereumClient.TransactionReceipt(ctx, txHash)
		if err != nil {
			zap.L().Warn("wait for transaction", zap.Error(err), zap.String("tx", txHash.String()))

			continue
		}

		if receipt.Status == types.ReceiptStatusSuccessful {
			return nil
		}

		if receipt.Status == types.ReceiptStatusFailed {
			return fmt.Errorf("transaction failed: %s", receipt.TxHash.String())
		}
	}
}

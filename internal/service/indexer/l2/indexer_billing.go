package l2

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rss3-network/payment-processor/common/ethereum"
	"github.com/rss3-network/payment-processor/contract/l2"
	"github.com/rss3-network/payment-processor/internal/database"
	"github.com/rss3-network/payment-processor/schema"
	"go.uber.org/zap"
)

func (s *server) indexBillingLog(ctx context.Context, header *types.Header, transaction *types.Transaction, receipt *types.Receipt, log *types.Log, logIndex int, databaseTransaction database.Client) error {
	switch eventHash := log.Topics[0]; eventHash {
	case l2.EventHashBillingTokensDeposited:
		return s.indexBillingTokensDepositedLog(ctx, header, transaction, receipt, log, logIndex, databaseTransaction)
	case l2.EventHashBillingTokensWithdrawn:
		return s.indexBillingTokensWithdrawnLog(ctx, header, transaction, receipt, log, logIndex, databaseTransaction)
	case l2.EventHashBillingTokensCollected:
		return s.indexBillingTokensCollectedLog(ctx, header, transaction, receipt, log, logIndex, databaseTransaction)
	default:
		return nil
	}
}

func (s *server) indexBillingTokensDepositedLog(ctx context.Context, header *types.Header, transaction *types.Transaction, _ *types.Receipt, log *types.Log, _ int, databaseTransaction database.Client) error {
	billingTokensDepositedEvent, err := s.contractBilling.ParseTokensDeposited(*log)
	if err != nil {
		return fmt.Errorf("parse TokensDeposited event: %w", err)
	}

	zap.L().Debug("indexing TokensDeposited event for Billing", zap.Stringer("transaction.hash", transaction.Hash()), zap.Any("event", billingTokensDepositedEvent))

	billingRecord := schema.BillingRecordDeposited{
		BillingRecordBase: schema.BillingRecordParseBase(s.chainID.Uint64(), header, transaction, log.Index, billingTokensDepositedEvent.User, billingTokensDepositedEvent.Amount),
	}

	if err := databaseTransaction.SaveBillingRecordDeposited(ctx, &billingRecord); err != nil {
		return fmt.Errorf("save billing record: %w", err)
	}

	parsedRu, _ := new(big.Float).Mul(new(big.Float).Quo(
		new(big.Float).SetInt(billingTokensDepositedEvent.Amount),
		new(big.Float).SetInt(big.NewInt(ethereum.BillingTokenDecimals)),
	), big.NewFloat(float64(s.billingConfig.RuPerToken))).Int64()

	isResumed, err := databaseTransaction.GatewayDeposit(ctx, billingTokensDepositedEvent.User, parsedRu)

	if isResumed {
		// Try to resume anyway
		err = s.controlClient.ResumeAccount(ctx, billingTokensDepositedEvent.User.Hex())
		if err != nil {
			return fmt.Errorf("resume account in control: %w", err)
		}
	}

	if err != nil {
		return fmt.Errorf("resume account in database: %w", err)
	}

	return nil
}

func (s *server) indexBillingTokensWithdrawnLog(ctx context.Context, header *types.Header, transaction *types.Transaction, _ *types.Receipt, log *types.Log, _ int, databaseTransaction database.Client) error {
	billingTokensWithdrawnEvent, err := s.contractBilling.ParseTokensWithdrawn(*log)
	if err != nil {
		return fmt.Errorf("parse TokensWithdrawn event: %w", err)
	}

	zap.L().Debug("indexing TokensWithdrawn event for Billing", zap.Stringer("transaction.hash", transaction.Hash()), zap.Any("event", billingTokensWithdrawnEvent))

	billingRecord := schema.BillingRecordWithdrawal{
		BillingRecordBase: schema.BillingRecordParseBase(s.chainID.Uint64(), header, transaction, log.Index, billingTokensWithdrawnEvent.User, billingTokensWithdrawnEvent.Amount),
		Fee:               big.NewInt(0),
	}

	if err := databaseTransaction.SaveBillingRecordWithdrawal(ctx, &billingRecord); err != nil {
		return fmt.Errorf("save billing record: %w", err)
	}

	return nil
}

func (s *server) indexBillingTokensCollectedLog(ctx context.Context, header *types.Header, transaction *types.Transaction, _ *types.Receipt, log *types.Log, _ int, databaseTransaction database.Client) error {
	billingTokensCollected, err := s.contractBilling.ParseTokensCollected(*log)
	if err != nil {
		return fmt.Errorf("parse TokensCollected event: %w", err)
	}

	zap.L().Debug("indexing TokensCollected event for Billing", zap.Stringer("transaction.hash", transaction.Hash()), zap.Any("event", billingTokensCollected))

	billingRecord := schema.BillingRecordCollected{
		BillingRecordBase: schema.BillingRecordParseBase(s.chainID.Uint64(), header, transaction, log.Index, billingTokensCollected.User, billingTokensCollected.Amount),
	}

	if err := databaseTransaction.SaveBillingRecordCollected(ctx, &billingRecord); err != nil {
		return fmt.Errorf("save billing record: %w", err)
	}

	return nil
}

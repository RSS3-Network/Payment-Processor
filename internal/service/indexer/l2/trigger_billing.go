package l2

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/payment-processor/common/ethereum"
	"github.com/rss3-network/payment-processor/common/txmgr"
	"github.com/rss3-network/payment-processor/contract/l2"
	"go.uber.org/zap"
)

func (s *server) billingFlow(ctx context.Context, epoch *big.Int) (*big.Int, error) {
	// billing
	var usersRequireRuLimitRefresh []common.Address

	// Step 1: Collect
	succeededUsers, totalCollected, err := s.billingCollect(ctx, epoch)

	if err != nil {
		return nil, fmt.Errorf("billing collect failed with %w", err)
	}

	if succeededUsers != nil {
		usersRequireRuLimitRefresh = append(usersRequireRuLimitRefresh, succeededUsers...)
	}

	// Step 2: Withdraw
	succeededUsers, err = s.billingWithdraw(ctx)

	if err != nil {
		return nil, fmt.Errorf("billing withdraw failed with %w", err)
	}

	if succeededUsers != nil {
		usersRequireRuLimitRefresh = append(usersRequireRuLimitRefresh, succeededUsers...)
	}

	// Step 3: Merge succeed lists and refresh their RU limit
	err = s.billingUpdateRuLimit(ctx, usersRequireRuLimitRefresh)

	if err != nil {
		return nil, fmt.Errorf("billing update ru limit failed with %w", err)
	}

	return totalCollected, nil
}

func (s *server) billingCollect(ctx context.Context, epoch *big.Int) ([]common.Address, *big.Int, error) {
	// billing collect tokens
	nowTime := time.Now() // Epoch round identifier for billing

	users, amounts, err := s.buildBillingCollectTokens(ctx, nowTime)

	if err != nil {
		zap.L().Error("build billing collect tokens", zap.Error(err))
		return nil, nil, fmt.Errorf("build billing collect tokens: %w", err)
	}

	if users == nil || amounts == nil {
		// Nothing to do
		return nil, big.NewInt(0), nil
	}

	// else Need collect
	var succeededUsers []common.Address

	totalCollected := big.NewInt(0)

	// call contract slice by slice
	for len(users) > 0 {
		limit := len(users)

		if limit > s.settlerConfig.BatchSize {
			limit = s.settlerConfig.BatchSize
		}

		err = s.triggerBillingCollectTokens(ctx, epoch, users[:limit], amounts[:limit])

		if err != nil {
			zap.L().Error("trigger billing collect tokens", zap.Error(err), zap.Int64("epoch", epoch.Int64()), zap.Any("users", users[:limit]), zap.Any("amounts", amounts[:limit]))
		} else {
			succeededUsers = append(succeededUsers, users[:limit]...)

			for _, amount := range amounts[limit:] {
				totalCollected.Add(totalCollected, amount)
			}
		}

		users = users[limit:]
		amounts = amounts[limit:]
	}

	zap.L().Info("Collect tokens ran successfully")

	return succeededUsers, totalCollected, nil
}

func (s *server) billingWithdraw(ctx context.Context) ([]common.Address, error) {
	// billing withdraw
	users, amounts, err := s.buildBillingWithdrawTokens(ctx)

	if err != nil {
		zap.L().Error("build billing withdraw", zap.Error(err))
		return nil, fmt.Errorf("build billing withdraw: %w", err)
	}

	if users == nil || amounts == nil {
		// Nothing to do
		return nil, nil
	}

	// else Need withdraw
	var succeededUsers []common.Address

	// call contract slice by slice
	for len(users) > 0 {
		limit := len(users)

		if limit > s.settlerConfig.BatchSize {
			limit = s.settlerConfig.BatchSize
		}

		err = s.triggerBillingWithdrawTokens(ctx, users[:limit], amounts[:limit])

		if err == nil {
			succeededUsers = append(succeededUsers, users[:limit]...)
		}

		users = users[limit:]
		amounts = amounts[limit:]
	}

	return succeededUsers, nil
}

func (s *server) billingUpdateRuLimit(ctx context.Context, usersRequireRuLimitRefresh []common.Address) error {
	// update ru limit
	currentBalance := s.getCurrentRuBalance(ctx, usersRequireRuLimitRefresh)

	if currentBalance != nil {
		err := s.databaseClient.UpdateBillingRuLimit(ctx, currentBalance)

		if err != nil {
			zap.L().Error("update ru limit", zap.Any("usersRequireRuLimitRefresh", usersRequireRuLimitRefresh), zap.Error(err))
			return fmt.Errorf("update ru limit: %w", err)
		}
	}

	return nil
}

func (s *server) buildBillingCollectTokens(ctx context.Context, nowTime time.Time) ([]common.Address, []*big.Int, error) {
	zap.L().Debug("Build billing collect tokens")

	collectTokensData, err := s.databaseClient.PrepareBillingCollectTokens(ctx, nowTime)

	if err != nil {
		zap.L().Error("prepare billing data", zap.Error(err))
		return nil, nil, fmt.Errorf("prepare billing data: %w", err)
	}

	if collectTokensData == nil {
		// Nothing to do
		return nil, nil, nil
	}

	// Prepare result storage arrays
	var (
		users   []common.Address
		amounts []*big.Int
	)

	// Calculate consumed token (w/ billing rate) per address
	for addr, ruC := range *collectTokensData {
		consumedTokenRaw := new(big.Int).Quo(
			new(big.Int).Mul(big.NewInt(ruC.Ru), big.NewInt(ethereum.BillingTokenDecimals)),
			big.NewInt(s.billingConfig.RuPerToken),
		)

		consumedToken, _ := new(big.Float).Mul(
			new(big.Float).SetInt(consumedTokenRaw),
			big.NewFloat(ruC.BillingRate),
		).Int(nil)

		zap.L().Debug("calculate consumed token (real)", zap.String("addr", addr.Hex()), zap.String("token (raw)", consumedToken.String()))

		// Check address balance, prevent from exceeding
		balance, err := s.contractBilling.BalanceOf(&bind.CallOpts{
			Context: ctx,
		}, addr)

		if err != nil {
			zap.L().Error("check balance of address", zap.String("addr", addr.Hex()), zap.Error(err))
		} else if consumedToken.Cmp(balance) == 1 {
			// Balance not enough, only get balance to prevent calculation exceeds
			consumedToken = balance
		}

		zap.L().Debug("calculate consumed token (collect)", zap.String("addr", addr.Hex()), zap.String("token (raw)", consumedToken.String()))

		if consumedToken.Cmp(big.NewInt(0)) == 1 {
			// consumedToken > 0
			users = append(users, addr)
			amounts = append(amounts, consumedToken)
		}
	}

	return users, amounts, nil
}

func (s *server) buildBillingWithdrawTokens(ctx context.Context) ([]common.Address, []*big.Int, error) {
	withdrawData, err := s.databaseClient.PrepareBillingWithdrawTokens(ctx)

	if err != nil {
		zap.L().Error("prepare billing data", zap.Error(err))
		return nil, nil, fmt.Errorf("prepare billing data: %w", err)
	}

	if withdrawData == nil {
		// Nothing to do
		return nil, nil, nil
	}

	// Prepare result storage arrays
	var (
		users   []common.Address
		amounts []*big.Int
	)

	for addr, withdrawAmount := range *withdrawData {
		amount, _ := new(big.Float).Mul(big.NewFloat(withdrawAmount), big.NewFloat(ethereum.BillingTokenDecimals)).Int(nil)

		if amount == nil {
			zap.L().Error("parse withdraw amount", zap.String("address", addr.Hex()), zap.Float64("amount", withdrawAmount))
		} else {
			users = append(users, addr)
			amounts = append(amounts, amount)
		}
	}

	return users, amounts, nil
}

func (s *server) triggerBillingCollectTokens(ctx context.Context, epoch *big.Int, users []common.Address, amounts []*big.Int) error {
	// Trigger collectTokens contract.
	input, err := txmgr.EncodeInput(l2.BillingMetaData.ABI, l2.MethodCollectTokens, epoch, users, amounts)

	if err != nil {
		return fmt.Errorf("encode input: %w", err)
	}

	receipt, err := s.sendTransaction(ctx, input)

	if err != nil {
		s.ReportFailedTransactionToSlack(err, receipt, l2.MethodCollectTokens, users, amounts)
		return fmt.Errorf("send transaction receipt: %w", err)
	}

	zap.L().Info("collect tokens successfully", zap.String("tx", receipt.TxHash.String()))

	return nil
}

func (s *server) triggerBillingWithdrawTokens(ctx context.Context, users []common.Address, amounts []*big.Int) error {
	// Trigger collectTokens contract.
	input, err := txmgr.EncodeInput(l2.BillingMetaData.ABI, l2.MethodWithdrawTokens, users, amounts)

	if err != nil {
		return fmt.Errorf("encode input: %w", err)
	}

	receipt, err := s.sendTransaction(ctx, input)

	if err != nil {
		s.ReportFailedTransactionToSlack(err, receipt, l2.MethodWithdrawTokens, users, amounts)
		return fmt.Errorf("send transaction receipt: %w", err)
	}

	zap.L().Info("collect tokens successfully", zap.String("tx", receipt.TxHash.String()))

	return nil
}

func (s *server) getCurrentRuBalance(ctx context.Context, users []common.Address) map[common.Address]int64 {
	if len(users) == 0 {
		return nil
	}

	latestRuLimit := make(map[common.Address]int64)

	for _, address := range users {
		// Get from chain
		balance, err := s.contractBilling.BalanceOf(&bind.CallOpts{
			Context: ctx,
		}, address)

		if err != nil {
			// Something is wrong
			zap.L().Error("get current balance", zap.Error(err), zap.String("address", address.Hex()))
			continue
		}

		// Parse balance to RU
		parsedRu, _ := new(big.Float).Mul(new(big.Float).Quo(
			new(big.Float).SetInt(balance),
			new(big.Float).SetInt(big.NewInt(ethereum.BillingTokenDecimals)),
		), big.NewFloat(float64(s.billingConfig.RuPerToken))).Int64()

		latestRuLimit[address] = parsedRu
	}

	return latestRuLimit
}

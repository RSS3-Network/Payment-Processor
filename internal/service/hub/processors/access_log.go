package processors

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rss3-network/gateway-common/accesslog"
	"github.com/rss3-network/payment-processor/internal/database/dialer/cockroachdb/table"
	"github.com/rss3-network/payment-processor/internal/service/hub/model"
	"go.uber.org/zap"
)

func (app *App) ProcessAccessLog(l *accesslog.Log) {
	ctx := context.Background()

	zap.L().Debug("new access log arrive")

	// Check billing eligibility
	if l.KeyID == nil {
		return
	}

	// Find account and key
	key, account, err := app.findAccount(ctx, *l.KeyID)
	if err != nil {
		zap.L().Error("find account", zap.Error(err))
		return
	}

	// Update consumption
	if err = app.updateConsumption(ctx, key, l.Status == http.StatusOK && !key.Account.IsPaused); err != nil {
		zap.L().Error("update consumption", zap.Error(err))
	}

	// Pause account
	if err = app.pauseAccount(ctx, account); err != nil {
		zap.L().Error("pause account", zap.Error(err))
	}
}

func (app *App) findAccount(ctx context.Context, keyID string) (*model.Key, *model.Account, error) {
	keyIDParsed, err := strconv.ParseUint(keyID, 10, 64)
	if err != nil {
		return nil, nil, fmt.Errorf("recover key id (%s): %w", keyID, err)
	}

	key, _, err := model.KeyGetByID(ctx, keyIDParsed, false, app.databaseClient, app.controlClient) // Deleted key could also be used for pending bills
	if err != nil {
		return nil, nil, fmt.Errorf("get key by id (%s): %w", keyID, err)
	}

	user, err := key.GetAccount(ctx)
	if err != nil {
		// Failed to get account
		return nil, nil, fmt.Errorf("get account (%s): %w", keyID, err)
	}

	return key, user, nil
}

func (app *App) updateConsumption(ctx context.Context, key *model.Key, isRUConsumption bool) error {
	// Request failed or is in free tier, only increase API call count
	if err := key.ConsumeRu(ctx, 0); err != nil {
		// Failed to consumer RU
		return fmt.Errorf("key (%v) increase API call count: %w", key, err)
	}

	if isRUConsumption {
		ru := int64(1) // Default // TODO

		if err := key.ConsumeRu(ctx, ru); err != nil {
			// Failed to consume RU
			return fmt.Errorf("key (%v) consume RU: %w", key, err)
		}
	}

	return nil
}

func (app *App) pauseAccount(ctx context.Context, account *model.Account) error {
	ruRemain, err := account.GetBalance(ctx)
	if err != nil {
		// Failed to get remain RU
		return fmt.Errorf("get account (%v) remain RU: %w", account, err)
	}

	if ruRemain < 0 {
		zap.L().Info("Insufficient remain RU, pause account", zap.Any("account", account))
		// Pause account
		if !account.IsPaused {
			if err = app.controlClient.PauseAccount(ctx, account.Address.Hex()); err != nil {
				return fmt.Errorf("pause account (%v) in control: %w", account, err)
			} else if err = app.databaseClient.WithContext(ctx).
				Model(&table.GatewayAccount{}).
				Where("address = ?", account.Address).
				Update("is_paused", true).
				Error; err != nil {
				return fmt.Errorf("save paused account (%v) in into db: %w", account, err)
			}
		}
	}

	return nil
}

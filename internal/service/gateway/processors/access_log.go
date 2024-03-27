package processors

import (
	"context"
	"net/http"
	"strconv"

	"github.com/rss3-network/gateway-common/accesslog"
	"github.com/rss3-network/payment-processor/internal/database/dialer/cockroachdb/table"
	"github.com/rss3-network/payment-processor/internal/service/gateway/model"
	"go.uber.org/zap"
)

func (app *App) ProcessAccessLog(l *accesslog.Log) {
	rctx := context.Background()

	zap.L().Debug("new access log arrive")

	// Check billing eligibility
	if l.KeyID == nil {
		return
	}

	// Find user
	keyIDParsed, err := strconv.ParseUint(*l.KeyID, 10, 64)

	if err != nil {
		zap.L().Error("recover key id", zap.String("keyID string", *l.KeyID), zap.Error(err))
		return
	}

	key, _, err := model.KeyGetByID(rctx, keyIDParsed, false, app.databaseClient, app.controlClient) // Deleted key could also be used for pending bills

	if err != nil {
		zap.L().Error("get key by id", zap.Uint64("keyID", keyIDParsed), zap.Error(err))
		return
	}

	user, err := key.GetAccount(rctx)

	if err != nil {
		// Failed to get account
		zap.L().Error("get account", zap.Error(err))
		return
	}

	if l.Status != http.StatusOK || key.Account.IsPaused {
		// Request failed or is in free tier, only increase API call count
		if err = key.ConsumeRu(rctx, 0); err != nil {
			// Failed to consumer RU
			zap.L().Error("increase API call count", zap.Any("account", user), zap.Any("key", key), zap.Error(err))
		}

		return
	}

	// Consumer RU
	ru := int64(1) // Default // TODO

	if err = key.ConsumeRu(rctx, ru); err != nil {
		// Failed to consume RU
		zap.L().Error("consume RU", zap.Any("account", user), zap.Any("key", key), zap.Error(err))

		return
	}

	ruRemain, err := user.GetBalance(rctx)

	if err != nil {
		// Failed to get remain RU
		zap.L().Error("get account remain RU", zap.Any("account", user), zap.Error(err))

		return
	}

	if ruRemain < 0 {
		zap.L().Info("Insufficient remain RU, pause account", zap.Any("account", user))
		// Pause user account
		if !user.IsPaused {
			if err = app.controlClient.PauseAccount(rctx, user.Address.Hex()); err != nil {
				zap.L().Error("pause account in control", zap.Any("account", user), zap.Error(err))
			} else if err = app.databaseClient.WithContext(rctx).
				Model(&table.GatewayAccount{}).
				Where("address = ?", user.Address).
				Update("is_paused", true).
				Error; err != nil {
				zap.L().Error("save paused account into db", zap.Any("account", user), zap.Error(err))
			}
		}
	}
}

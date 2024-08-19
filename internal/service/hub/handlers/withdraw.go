package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/payment-processor/internal/database/dialer/postgresql/table"
	"github.com/rss3-network/payment-processor/internal/service/hub/gen/oapi"
	"github.com/rss3-network/payment-processor/internal/service/hub/model"
	"github.com/rss3-network/payment-processor/internal/service/hub/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (app *App) GetPendingRequestWithdraw(ctx echo.Context) error {
	user := ctx.Get("user").(*model.Account)

	amount := float32(0)

	// Check if there's any pending withdraw requests
	var pendingWithdrawRequest table.GatewayPendingWithdrawRequest
	err := app.databaseClient.WithContext(ctx.Request().Context()).
		Model(&pendingWithdrawRequest).
		Where("account_address = ?", user.Address).
		First(&pendingWithdrawRequest).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil // Not real error
		}
	} else {
		amount = float32(pendingWithdrawRequest.Amount)
	}

	if err != nil {
		zap.L().Error("GetPendingRequestWithdraw", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, oapi.GetRequestWithdrawResponse{Amount: &amount})
}

func (app *App) SetPendingRequestWithdraw(ctx echo.Context, params oapi.SetPendingRequestWithdrawParams) error {
	user := ctx.Get("user").(*model.Account)

	// Check if there's any pending withdraw requests
	var pendingWithdrawRequest table.GatewayPendingWithdrawRequest
	err := app.databaseClient.WithContext(ctx.Request().Context()).
		Model(&pendingWithdrawRequest).
		Where("account_address = ?", user.Address).
		First(&pendingWithdrawRequest).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil // Not real error
			if params.Amount > 0 {
				// Create
				err = app.databaseClient.WithContext(ctx.Request().Context()).
					Model(&pendingWithdrawRequest).
					Create(&table.GatewayPendingWithdrawRequest{
						AccountAddress: user.Address,
						Amount:         float64(params.Amount),
					}).
					Error
			}
		}
	} else {
		// Found record with no error
		if params.Amount > 0 {
			// Update
			err = app.databaseClient.WithContext(ctx.Request().Context()).
				Model(&pendingWithdrawRequest).
				Where("account_address = ?", user.Address).
				Update("amount", float64(params.Amount)).
				Error
		} else {
			// Delete
			err = app.databaseClient.WithContext(ctx.Request().Context()).
				Model(&pendingWithdrawRequest).
				Where("account_address = ?", user.Address).
				Delete(&pendingWithdrawRequest).
				Error
		}
	}

	if err != nil {
		zap.L().Error("SetPendingRequestWithdraw", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusOK)
}

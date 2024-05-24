package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/payment-processor/internal/service/hub/gen/oapi"
	"github.com/rss3-network/payment-processor/internal/service/hub/model"
	"github.com/rss3-network/payment-processor/internal/service/hub/utils"
	"go.uber.org/zap"
)

func (app *App) GetRUStatus(ctx echo.Context) error {
	user := ctx.Get("user").(*model.Account)

	ruUsedTotal, ruUsedCurrent, apiCallsTotal, apiCallsCurrent, err := user.GetUsage(ctx.Request().Context())
	if err != nil {
		zap.L().Error("GetRUStatus", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	resp := oapi.RUStatus{
		RuLimit:         &user.RuLimit,
		RuUsedTotal:     &ruUsedTotal,
		RuUsedCurrent:   &ruUsedCurrent,
		ApiCallsTotal:   &apiCallsTotal,
		ApiCallsCurrent: &apiCallsCurrent,
	}

	return ctx.JSON(http.StatusOK, resp)
}

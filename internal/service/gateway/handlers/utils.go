package handlers

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/api-gateway/app/model"
	"github.com/naturalselectionlabs/api-gateway/app/oapi/constants"
	"github.com/naturalselectionlabs/api-gateway/app/oapi/middlewares"
	"github.com/naturalselectionlabs/api-gateway/app/oapi/types"
	"github.com/naturalselectionlabs/rss3-global-indexer/internal/service/gateway/gen/oapi"
)

func getCtx(ctx echo.Context) (context.Context, *types.UserContext) {
	uctx := &types.UserContext{Context: ctx, User: middlewares.ParseUserWithToken(ctx)}
	rctx := ctx.Request().Context()
	return rctx, uctx
}

func getKey(ctx echo.Context, key int) (*model.Key, error) {
	user := ctx.Get("user").(*model.Account)

	k, err := user.GetKey(ctx.Request().Context(), key)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

func parseDates(since *oapi.Since, until *oapi.Until) (time.Time, time.Time) {
	var startFrom, untilTo time.Time
	nowTime := time.Now()

	if since != nil {
		startFrom = time.UnixMilli(*since)
	} else {
		startFrom = nowTime.Add(-constants.DEFAULT_HISTORY_SINCE)
	}

	if until != nil {
		untilTo = time.UnixMilli(*until)
	} else {
		untilTo = nowTime
	}

	if untilTo.Before(startFrom) {
		// Swap
		startFrom, untilTo = untilTo, startFrom
	}
	return startFrom, untilTo
}

func parseLimitPage(limit *oapi.Limit, page *oapi.Page) (int, int) {
	var (
		l = constants.DEFAULT_PAGINATION_LIMIT
		p = 1
	)

	if limit != nil {
		l = int(*limit)
	}

	if page != nil && *page >= 1 {
		p = int(*page)
	}

	return l, p
}

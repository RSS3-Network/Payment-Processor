package handlers

import (
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/payment-processor/internal/database/dialer/postgresql/table"
	"github.com/rss3-network/payment-processor/internal/service/hub/gen/oapi"
	"github.com/rss3-network/payment-processor/internal/service/hub/model"
	"github.com/rss3-network/payment-processor/internal/service/hub/utils"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (app *App) GetDepositHistory(ctx echo.Context, params oapi.GetDepositHistoryParams) error {
	user := ctx.Get("user").(*model.Account)

	// Parse date
	limit, page := parseLimitPage(params.Limit, params.Page)

	// Query from database
	query := app.databaseClient.WithContext(ctx.Request().Context()).
		Model(&table.BillingRecordDeposited{}).
		Where(&table.BillingRecordDeposited{
			User: user.Address,
		})

	var totalCount int64
	err := query.Count(&totalCount).Error

	if err != nil {
		zap.L().Error("GetDepositHistory count", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	var records []table.BillingRecordDeposited
	err = query.Order("block_timestamp DESC").Offset(limit * (page - 1)).Limit(limit).Find(&records).Error

	if err != nil {
		zap.L().Error("GetDepositHistory query", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	list := make([]oapi.DepositRecord, len(records))

	for i, record := range records {
		blockTS := record.BlockTimestamp.UnixMilli()
		amount, _ := utils.ParseAmount(record.Amount.BigInt()).Float32()
		list[i] = oapi.DepositRecord{
			TxHash:         lo.ToPtr(record.TxHash.Hex()),
			BlockTimestamp: &blockTS,
			Index:          lo.ToPtr(int(record.Index)),
			Amount:         &amount,
		}
	}

	count := uint(len(list))
	pageCurrent := int64(page)
	pageMax := int64(math.Ceil(float64(totalCount) / float64(limit)))

	return ctx.JSON(http.StatusOK, oapi.DepositHistoryResponse{
		Count:       &count,
		PageCurrent: &pageCurrent,
		PageMax:     &pageMax,
		List:        &list,
	})
}

func (app *App) GetWithdrawalHistory(ctx echo.Context, params oapi.GetWithdrawalHistoryParams) error {
	user := ctx.Get("user").(*model.Account)

	// Parse date
	limit, page := parseLimitPage(params.Limit, params.Page)

	// Query from database
	query := app.databaseClient.WithContext(ctx.Request().Context()).
		Model(&table.BillingRecordWithdrawal{}).
		Where(&table.BillingRecordWithdrawal{
			User: user.Address,
		})

	var totalCount int64
	err := query.Count(&totalCount).Error

	if err != nil {
		zap.L().Error("GetWithdrawalHistory count", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	var records []table.BillingRecordWithdrawal
	err = query.Order("block_timestamp DESC").Offset(limit * (page - 1)).Limit(limit).Find(&records).Error

	if err != nil {
		zap.L().Error("GetWithdrawalHistory query", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	list := make([]oapi.WithdrawalRecord, len(records))

	for i, record := range records {
		blockTS := record.BlockTimestamp.UnixMilli()
		amount, _ := utils.ParseAmount(record.Amount.BigInt()).Float32()
		fee, _ := utils.ParseAmount(record.Fee.BigInt()).Float32()
		list[i] = oapi.WithdrawalRecord{
			TxHash:         lo.ToPtr(record.TxHash.Hex()),
			BlockTimestamp: &blockTS,
			Index:          lo.ToPtr(int(record.Index)),
			User:           lo.ToPtr(record.User.Hex()),
			Amount:         &amount,
			Fee:            &fee,
		}
	}

	count := uint(len(list))
	pageCurrent := int64(page)
	pageMax := int64(math.Ceil(float64(totalCount) / float64(limit)))

	return ctx.JSON(http.StatusOK, oapi.WithdrawalHistoryResponse{
		Count:       &count,
		PageCurrent: &pageCurrent,
		PageMax:     &pageMax,
		List:        &list,
	})
}

func (app *App) GetCollectionHistory(ctx echo.Context, params oapi.GetCollectionHistoryParams) error {
	user := ctx.Get("user").(*model.Account)

	limit, page := parseLimitPage(params.Limit, params.Page)

	// Query from database
	query := app.databaseClient.WithContext(ctx.Request().Context()).
		Model(&table.BillingRecordCollected{}).
		Where(&table.BillingRecordCollected{
			User: user.Address,
		})

	var totalCount int64
	err := query.Count(&totalCount).Error

	if err != nil {
		zap.L().Error("GetCollectionHistory count", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	var records []table.BillingRecordCollected
	err = query.Order("block_timestamp DESC").Offset(limit * (page - 1)).Limit(limit).Find(&records).Error

	if err != nil {
		zap.L().Error("GetCollectionHistory query", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	list := make([]oapi.CollectionRecord, len(records))

	for i, record := range records {
		blockTS := record.BlockTimestamp.UnixMilli()
		amount, _ := utils.ParseAmount(record.Amount.BigInt()).Float32()
		list[i] = oapi.CollectionRecord{
			TxHash:         lo.ToPtr(record.TxHash.Hex()),
			BlockTimestamp: &blockTS,
			Index:          lo.ToPtr(int(record.Index)),
			Amount:         &amount,
		}
	}

	count := uint(len(list))
	pageCurrent := int64(page)
	pageMax := int64(math.Ceil(float64(totalCount) / float64(limit)))

	return ctx.JSON(http.StatusOK, oapi.CollectionHistoryResponse{
		Count:       &count,
		PageCurrent: &pageCurrent,
		PageMax:     &pageMax,
		List:        &list,
	})
}

func (app *App) GetConsumptionHistoryByKey(ctx echo.Context, keyID string, params oapi.GetConsumptionHistoryByKeyParams) error {
	since, until := parseDates(params.Since, params.Until)

	// Query from database
	k, exist, err := app.getKey(ctx, keyID)
	if err != nil {
		zap.L().Error("GetConsumptionHistoryByKey getKey", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	} else if !exist {
		return utils.SendJSONError(ctx, http.StatusNotFound)
	}

	var logs []table.GatewayConsumptionLog
	err = app.databaseClient.WithContext(ctx.Request().Context()).
		Model(&table.GatewayConsumptionLog{}).
		Where("key_id = ? AND consumption_date >= ? AND consumption_date <= ?", k.ID, since, until).
		Order("consumption_date DESC").
		Find(&logs).
		Error

	if err != nil {
		zap.L().Error("GetConsumptionHistoryByKey query", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	sinceInt64 := since.UnixMilli()
	untilInt64 := until.UnixMilli()
	resp := &oapi.ConsumptionLogResponse{
		Since:   &sinceInt64,
		Until:   &untilInt64,
		History: &[]oapi.ConsumptionLogByKey{},
	}

	if params.Merge != nil && *params.Merge {
		var (
			apiCalls int64
			ruUsed   int64
		)

		for _, log := range logs {
			apiCalls += log.APICalls
			ruUsed += log.RuUsed
		}

		*resp.History = append(*resp.History, oapi.ConsumptionLogByKey{
			KeyName:  &k.Name,
			ApiCalls: &apiCalls,
			RuUsed:   &ruUsed,
		})
	} else {
		for _, log := range logs {
			consumptionDate := log.ConsumptionDate.UnixMilli()
			*resp.History = append(*resp.History, oapi.ConsumptionLogByKey{
				KeyName:         lo.ToPtr(k.Name),
				ConsumptionDate: lo.ToPtr(consumptionDate),
				ApiCalls:        lo.ToPtr(log.APICalls),
				RuUsed:          lo.ToPtr(log.RuUsed),
			})
		}
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (app *App) GetConsumptionHistoryByAccount(ctx echo.Context, params oapi.GetConsumptionHistoryByAccountParams) error {
	user := ctx.Get("user").(*model.Account)

	since, until := parseDates(params.Since, params.Until)

	// Query from database
	logs, err := user.GetUsageByDate(ctx.Request().Context(), since, until)

	if err != nil {
		zap.L().Error("GetUsageByDate", zap.Error(err))
		return utils.SendJSONError(ctx, http.StatusInternalServerError)
	}

	sinceInt64 := since.UnixMilli()
	untilInt64 := until.UnixMilli()
	resp := &oapi.ConsumptionLogResponse{
		Since:   lo.ToPtr(sinceInt64),
		Until:   lo.ToPtr(untilInt64),
		History: lo.ToPtr([]oapi.ConsumptionLogByKey{}),
	}

	if params.Merge != nil && *params.Merge {
		var (
			apiCalls int64
			ruUsed   int64
		)

		for _, log := range *logs {
			apiCalls += log.APICalls
			ruUsed += log.RuUsed
		}

		*resp.History = append(*resp.History, oapi.ConsumptionLogByKey{
			ApiCalls: &apiCalls,
			RuUsed:   &ruUsed,
		})
	} else {
		for _, log := range *logs {
			*resp.History = append(*resp.History, oapi.ConsumptionLogByKey{
				ConsumptionDate: lo.ToPtr(log.ConsumptionTimestamp),
				ApiCalls:        lo.ToPtr(log.APICalls),
				RuUsed:          lo.ToPtr(log.RuUsed),
			})
		}
	}

	return ctx.JSON(http.StatusOK, resp)
}

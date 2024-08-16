package model

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/gateway-common/control"
	"github.com/rss3-network/payment-processor/internal/database/dialer/postgresql/table"
	"github.com/rss3-network/payment-processor/schema"
	"gorm.io/gorm"
)

type Account struct {
	table.GatewayAccount

	databaseClient *gorm.DB
	controlClient  *control.StateClientWriter
}

func AccountCreate(ctx context.Context, address common.Address, databaseClient *gorm.DB, controlClient *control.StateClientWriter) (*Account, error) {
	acc := table.GatewayAccount{
		Address:     address,
		BillingRate: 1,
	}
	err := databaseClient.WithContext(ctx).
		Save(&acc).
		Error

	if err != nil {
		return nil, err
	}

	return &Account{acc, databaseClient, controlClient}, nil
}

func AccountGetByAddress(ctx context.Context, address common.Address, databaseClient *gorm.DB, controlClient *control.StateClientWriter) (*Account, bool, error) {
	var acc table.GatewayAccount

	err := databaseClient.WithContext(ctx).
		Model(&table.GatewayAccount{}).
		Where("address = ?", address).
		First(&acc).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}

		return nil, false, err
	}

	return &Account{acc, databaseClient, controlClient}, true, nil
}

func AccountGetOrCreate(ctx context.Context, address common.Address, databaseClient *gorm.DB, controlClient *control.StateClientWriter) (*Account, error) {
	acc, exist, err := AccountGetByAddress(ctx, address, databaseClient, controlClient)

	if err != nil {
		return nil, err
	} else if !exist {
		return AccountCreate(ctx, address, databaseClient, controlClient)
	}

	return acc, nil
}

func (acc *Account) ListKeys(ctx context.Context) ([]*Key, error) {
	var keys []table.GatewayKey

	err := acc.databaseClient.WithContext(ctx).
		Model(&table.GatewayKey{}).
		Where("account_address = ?", acc.Address).
		Find(&keys).
		Error

	if err != nil {
		return nil, err
	}

	wrappedKeys := make([]*Key, len(keys))
	for i, k := range keys {
		wrappedKeys[i] = &Key{k, acc.databaseClient, acc.controlClient}
	}

	return wrappedKeys, nil
}

func (acc *Account) GetUsage(ctx context.Context) (int64, int64, int64, int64, error) {
	var status struct {
		RuUsedTotal     int64 `gorm:"column:ru_used_total"`
		RuUsedCurrent   int64 `gorm:"column:ru_used_current"`
		APICallsTotal   int64 `gorm:"column:api_calls_total"`
		APICallsCurrent int64 `gorm:"column:api_calls_current"`
	}

	err := acc.databaseClient.WithContext(ctx).
		Model(&table.GatewayKey{}).
		Unscoped().
		Select("SUM(ru_used_total) AS ru_used_total, SUM(ru_used_current) AS ru_used_current, SUM(api_calls_total) AS api_calls_total, SUM(api_calls_current) AS api_calls_current").
		Where("account_address = ?", acc.Address).
		Find(&status).
		Error

	return status.RuUsedTotal, status.RuUsedCurrent, status.APICallsTotal, status.APICallsCurrent, err
}

func (acc *Account) GetUsageByDate(ctx context.Context, since time.Time, until time.Time) (*[]schema.UsageByDate, error) {
	var logs []schema.UsageByDate

	err := acc.databaseClient.WithContext(ctx).
		Model(&table.GatewayConsumptionLog{}).
		Joins("LEFT JOIN key ON consumption_log.key_id = key.id").
		Where("account_address = ? AND consumption_date BETWEEN ? AND ?", acc.Address, since, until).
		Select("SUM(ru_used) AS ru_used, SUM(api_calls) AS api_calls, (EXTRACT(EPOCH FROM consumption_date)*1000)::BIGINT AS consumption_timestamp").
		Group("consumption_timestamp").
		Order("consumption_timestamp DESC").
		Find(&logs).
		Error

	if err != nil {
		return nil, err
	}

	return &logs, nil
}

func (acc *Account) GetBalance(ctx context.Context) (int64, error) {
	_, ruUsed, _, _, err := acc.GetUsage(ctx)
	if err != nil {
		return 0, err
	}

	return acc.RuLimit - ruUsed, nil
}

func (acc *Account) GetKey(ctx context.Context, keyID uint64) (*Key, bool, error) {
	var k table.GatewayKey

	err := acc.databaseClient.WithContext(ctx).
		Model(&table.GatewayKey{}).
		Where("account_address = ? AND id = ?", acc.Address, keyID).
		First(&k).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}

		return nil, false, err
	}

	return &Key{k, acc.databaseClient, acc.controlClient}, true, nil
}

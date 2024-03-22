package model

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/naturalselectionlabs/rss3-gateway/internal/database/dialer/cockroachdb/table"
	"github.com/rss3-network/gateway-common/control"
	"gorm.io/gorm"
)

type Key struct {
	table.GatewayKey

	databaseClient *gorm.DB
	controlClient  *control.StateClientWriter
}

func KeyCreate(ctx context.Context, accountAddress common.Address, keyName string, databaseClient *gorm.DB, controlClient *control.StateClientWriter) (*Key, error) {
	keyUUID := uuid.New()
	k := table.GatewayKey{
		Key:            keyUUID,
		Name:           keyName,
		AccountAddress: accountAddress,
	}

	err := databaseClient.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// DB
		err := tx.
			Create(&k).
			Error
		if err != nil {
			return err
		}
		// APISix
		err = controlClient.CreateKey(ctx, accountAddress.Hex(), strconv.FormatUint(k.ID, 10), keyUUID.String())
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &Key{k, databaseClient, controlClient}, nil
}

func KeyGetByID(ctx context.Context, KeyID uint64, activeOnly bool, databaseClient *gorm.DB, controlClient *control.StateClientWriter) (*Key, bool, error) {
	queryBase := databaseClient.WithContext(ctx).Model(&table.GatewayKey{})

	if activeOnly {
		queryBase = queryBase.Unscoped()
	}

	var k table.GatewayKey

	err := queryBase.Where("id = ?", KeyID).First(&k).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}

		return nil, false, err
	}

	return &Key{k, databaseClient, controlClient}, true, nil
}

func (k *Key) ConsumeRu(ctx context.Context, ru int64) error {
	// Add API calls count
	err := k.databaseClient.WithContext(ctx).
		Model(&table.GatewayKey{}).
		Unscoped().
		Where("id = ?", k.ID).
		Updates(map[string]interface{}{
			"api_calls_current": gorm.Expr("api_calls_current + ?", 1),
			"ru_used_current":   gorm.Expr("ru_used_current + ?", ru),
		}).
		Error

	if err != nil {
		// Failed to consumer RU
		log.Printf("Faield to increase API call count with error: %v", err)
		return err
	}

	return nil
}

func (k *Key) GetAccount(_ context.Context) (*Account, error) {
	return &Account{k.Account, k.databaseClient, k.controlClient}, nil
}

func (k *Key) Delete(ctx context.Context) error {
	err := k.controlClient.DeleteKey(ctx, k.Key.String())

	if err != nil {
		return err
	}

	err = k.databaseClient.WithContext(ctx).
		Delete(&k).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (k *Key) UpdateInfo(ctx context.Context, name string) error {
	k.Name = name
	err := k.databaseClient.WithContext(ctx).
		Model(&table.GatewayKey{}).
		Where("id = ?", k.ID).
		Update("name", name).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (k *Key) Rotate(ctx context.Context) error {
	err := k.controlClient.DeleteKey(ctx, k.Key.String())

	if err != nil {
		return err
	}

	k.Key = uuid.New()

	err = k.databaseClient.WithContext(ctx).
		Model(&table.GatewayKey{}).
		Where("id = ?", k.ID).
		Update("key", k.Key).
		Error
	if err != nil {
		return err
	}

	// Update consumer
	err = k.controlClient.CreateKey(ctx, k.Account.Address.Hex(), strconv.FormatUint(k.ID, 10), k.Key.String())
	if err != nil {
		return err
	}

	return nil
}

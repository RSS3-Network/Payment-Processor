package table

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	gormSchema "gorm.io/gorm/schema"
)

var (
	_ gormSchema.Tabler = (*GatewayPendingWithdrawRequest)(nil)
)

type GatewayPendingWithdrawRequest struct {
	gorm.Model
	CreatedAt time.Time
	UpdatedAt time.Time

	Amount float64 `gorm:"column:amount"`

	AccountAddress common.Address `gorm:"primarykey;type:bytea;column:account_address"` // Foreign key of GatewayAccount
	Account        GatewayAccount `gorm:"foreignKey:AccountAddress"`                    // Belongs to GatewayAccount
}

func (r *GatewayPendingWithdrawRequest) TableName() string {
	return "pending_withdraw_request"
}

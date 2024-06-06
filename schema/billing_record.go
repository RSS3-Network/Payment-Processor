package schema

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BillingRecordBase struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	TxHash         common.Hash
	Index          uint
	ChainID        uint64
	BlockTimestamp time.Time
	BlockNumber    *big.Int

	User   common.Address
	Amount *big.Int
}

type BillingRecordDeposited struct {
	BillingRecordBase
}

type BillingRecordWithdrawal struct {
	BillingRecordBase

	Fee *big.Int
}

type BillingRecordCollected struct {
	BillingRecordBase
}

func BillingRecordParseBase(chainID uint64, header *types.Header, transaction *types.Transaction, logIndex uint, user common.Address, amount *big.Int) BillingRecordBase {
	return BillingRecordBase{
		TxHash:         transaction.Hash(),
		Index:          logIndex,
		ChainID:        chainID,
		BlockTimestamp: time.Unix(int64(header.Time), 0),
		BlockNumber:    header.Number,

		User:   user,
		Amount: amount,
	}
}

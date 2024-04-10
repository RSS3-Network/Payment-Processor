package schema

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type NodeRequestRecord struct {
	NodeAddress  common.Address
	Epoch        *big.Int
	RequestCount *big.Int

	RequestReward *big.Int
}

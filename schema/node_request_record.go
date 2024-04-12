package schema

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NodeRequestRecord struct {
	NodeAddress  common.Address
	Epoch        *big.Int
	RequestCount *big.Int

	RequestReward *big.Int
}

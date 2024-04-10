package schema

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type NodeRequestRewards struct {
	NodeAddress   common.Address
	Epoch         *big.Int
	RequestCounts *big.Int

	RequestRewards *big.Int
}
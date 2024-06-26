package utils

import (
	"math/big"

	"github.com/rss3-network/payment-processor/common/ethereum"
)

func ParseAmount(rawAmount *big.Int) *big.Float {
	return new(big.Float).Quo(
		new(big.Float).SetInt(rawAmount),
		new(big.Float).SetInt(big.NewInt(ethereum.BillingTokenDecimals)),
	)
}

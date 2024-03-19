package l2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/Billing.abi --pkg l2 --type Billing --out contract_billing.go

var ContractMap = map[uint64]*struct {
	AddressBillingProxy common.Address
}{
	2331: {
		AddressBillingProxy: common.HexToAddress("0x4630b9ad9b149ebf13d185ab7b96cb4afe95e6c4"), // https://scan.testnet.rss3.io/address/0x4630b9ad9b149ebf13d185ab7b96cb4afe95e6c4
	},
}

var (
	EventHashBillingTokensDeposited = crypto.Keccak256Hash([]byte("TokensDeposited(address,uint256)"))
	EventHashBillingTokensWithdrawn = crypto.Keccak256Hash([]byte("TokensWithdrawn(address,uint256,uint256)"))
	EventHashBillingTokensCollected = crypto.Keccak256Hash([]byte("TokensCollected(address,uint256)"))
)

var (
	MethodCollectTokens  = "collectTokens"
	MethodWithdrawTokens = "withdrawTokens"
)

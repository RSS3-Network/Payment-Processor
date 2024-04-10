package l2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/Billing.abi --pkg l2 --type Billing --out contract_billing.go
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/Staking.abi --pkg l2 --type Staking --out contract_staking.go

var ContractMap = map[uint64]*struct {
	AddressBillingProxy common.Address
	AddressStakingProxy common.Address
}{
	2331: {
		AddressBillingProxy: common.HexToAddress("0x2d2f2649f314100a17ead85c7a0b7099780e743e"), // https://scan.testnet.rss3.io/address/0x2d2f2649f314100a17ead85c7a0b7099780e743e
		AddressStakingProxy: common.HexToAddress("0xb1b209Ee24272C7EE8076764DAa27563c5add9FF"), // https://scan.testnet.rss3.io/address/0xb1b209Ee24272C7EE8076764DAa27563c5add9FF
	},
}

var (
	EventHashBillingTokensDeposited = crypto.Keccak256Hash([]byte("TokensDeposited(address,uint256)"))
	EventHashBillingTokensWithdrawn = crypto.Keccak256Hash([]byte("TokensWithdrawn(address,uint256,uint256)"))
	EventHashBillingTokensCollected = crypto.Keccak256Hash([]byte("TokensCollected(address,uint256)"))

	EventHashStakingRewardDistributed = crypto.Keccak256Hash([]byte("RewardDistributed(uint256,uint256,uint256,address[],uint256[],uint256[],uint256[],uint256[])"))
)

var (
	MethodCollectTokens  = "collectTokens"
	MethodWithdrawTokens = "withdrawTokens"
)

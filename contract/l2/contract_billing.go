// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BillingMetaData contains all meta data concerning the Billing contract.
var BillingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"COLLECTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collectTokens\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"users\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"totalCollected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collectedBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"distributeRewards\",\"inputs\":[{\"name\":\"nodeAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rewards\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSnapshot\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"totalCollected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDistributed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"staking\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakingContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawTokens\",\"inputs\":[{\"name\":\"users\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsDistributed\",\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isPublicGood\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensCollected\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensDeposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensWithdrawn\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressAlreadyCollected\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressAlreadyDistributed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidArrayLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpochNumber\",\"inputs\":[{\"name\":\"currentEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SubmissionIntervalNotElapsed\",\"inputs\":[]}]",
}

// BillingABI is the input ABI used to generate the binding from.
// Deprecated: Use BillingMetaData.ABI instead.
var BillingABI = BillingMetaData.ABI

// Billing is an auto generated Go binding around an Ethereum contract.
type Billing struct {
	BillingCaller     // Read-only binding to the contract
	BillingTransactor // Write-only binding to the contract
	BillingFilterer   // Log filterer for contract events
}

// BillingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BillingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BillingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BillingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BillingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BillingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BillingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BillingSession struct {
	Contract     *Billing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BillingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BillingCallerSession struct {
	Contract *BillingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BillingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BillingTransactorSession struct {
	Contract     *BillingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BillingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BillingRaw struct {
	Contract *Billing // Generic contract binding to access the raw methods on
}

// BillingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BillingCallerRaw struct {
	Contract *BillingCaller // Generic read-only contract binding to access the raw methods on
}

// BillingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BillingTransactorRaw struct {
	Contract *BillingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBilling creates a new instance of Billing, bound to a specific deployed contract.
func NewBilling(address common.Address, backend bind.ContractBackend) (*Billing, error) {
	contract, err := bindBilling(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Billing{BillingCaller: BillingCaller{contract: contract}, BillingTransactor: BillingTransactor{contract: contract}, BillingFilterer: BillingFilterer{contract: contract}}, nil
}

// NewBillingCaller creates a new read-only instance of Billing, bound to a specific deployed contract.
func NewBillingCaller(address common.Address, caller bind.ContractCaller) (*BillingCaller, error) {
	contract, err := bindBilling(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BillingCaller{contract: contract}, nil
}

// NewBillingTransactor creates a new write-only instance of Billing, bound to a specific deployed contract.
func NewBillingTransactor(address common.Address, transactor bind.ContractTransactor) (*BillingTransactor, error) {
	contract, err := bindBilling(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BillingTransactor{contract: contract}, nil
}

// NewBillingFilterer creates a new log filterer instance of Billing, bound to a specific deployed contract.
func NewBillingFilterer(address common.Address, filterer bind.ContractFilterer) (*BillingFilterer, error) {
	contract, err := bindBilling(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BillingFilterer{contract: contract}, nil
}

// bindBilling binds a generic wrapper to an already deployed contract.
func bindBilling(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BillingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Billing *BillingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Billing.Contract.BillingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Billing *BillingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Billing.Contract.BillingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Billing *BillingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Billing.Contract.BillingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Billing *BillingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Billing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Billing *BillingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Billing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Billing *BillingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Billing.Contract.contract.Transact(opts, method, params...)
}

// COLLECTORROLE is a free data retrieval call binding the contract method 0x9ca39ae9.
//
// Solidity: function COLLECTOR_ROLE() view returns(bytes32)
func (_Billing *BillingCaller) COLLECTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "COLLECTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// COLLECTORROLE is a free data retrieval call binding the contract method 0x9ca39ae9.
//
// Solidity: function COLLECTOR_ROLE() view returns(bytes32)
func (_Billing *BillingSession) COLLECTORROLE() ([32]byte, error) {
	return _Billing.Contract.COLLECTORROLE(&_Billing.CallOpts)
}

// COLLECTORROLE is a free data retrieval call binding the contract method 0x9ca39ae9.
//
// Solidity: function COLLECTOR_ROLE() view returns(bytes32)
func (_Billing *BillingCallerSession) COLLECTORROLE() ([32]byte, error) {
	return _Billing.Contract.COLLECTORROLE(&_Billing.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Billing *BillingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Billing *BillingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Billing.Contract.DEFAULTADMINROLE(&_Billing.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Billing *BillingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Billing.Contract.DEFAULTADMINROLE(&_Billing.CallOpts)
}

// EPOCHDURATION is a free data retrieval call binding the contract method 0xa70b9f0c.
//
// Solidity: function EPOCH_DURATION() view returns(uint256)
func (_Billing *BillingCaller) EPOCHDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "EPOCH_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EPOCHDURATION is a free data retrieval call binding the contract method 0xa70b9f0c.
//
// Solidity: function EPOCH_DURATION() view returns(uint256)
func (_Billing *BillingSession) EPOCHDURATION() (*big.Int, error) {
	return _Billing.Contract.EPOCHDURATION(&_Billing.CallOpts)
}

// EPOCHDURATION is a free data retrieval call binding the contract method 0xa70b9f0c.
//
// Solidity: function EPOCH_DURATION() view returns(uint256)
func (_Billing *BillingCallerSession) EPOCHDURATION() (*big.Int, error) {
	return _Billing.Contract.EPOCHDURATION(&_Billing.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_Billing *BillingCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_Billing *BillingSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _Billing.Contract.BalanceOf(&_Billing.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_Billing *BillingCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _Billing.Contract.BalanceOf(&_Billing.CallOpts, user)
}

// CollectedBalance is a free data retrieval call binding the contract method 0x8ddc3b19.
//
// Solidity: function collectedBalance() view returns(uint256)
func (_Billing *BillingCaller) CollectedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "collectedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectedBalance is a free data retrieval call binding the contract method 0x8ddc3b19.
//
// Solidity: function collectedBalance() view returns(uint256)
func (_Billing *BillingSession) CollectedBalance() (*big.Int, error) {
	return _Billing.Contract.CollectedBalance(&_Billing.CallOpts)
}

// CollectedBalance is a free data retrieval call binding the contract method 0x8ddc3b19.
//
// Solidity: function collectedBalance() view returns(uint256)
func (_Billing *BillingCallerSession) CollectedBalance() (*big.Int, error) {
	return _Billing.Contract.CollectedBalance(&_Billing.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Billing *BillingCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Billing *BillingSession) CurrentEpoch() (*big.Int, error) {
	return _Billing.Contract.CurrentEpoch(&_Billing.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Billing *BillingCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Billing.Contract.CurrentEpoch(&_Billing.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Billing *BillingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Billing *BillingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Billing.Contract.GetRoleAdmin(&_Billing.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Billing *BillingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Billing.Contract.GetRoleAdmin(&_Billing.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Billing *BillingCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Billing *BillingSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Billing.Contract.GetRoleMember(&_Billing.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Billing *BillingCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Billing.Contract.GetRoleMember(&_Billing.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Billing *BillingCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Billing *BillingSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Billing.Contract.GetRoleMemberCount(&_Billing.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Billing *BillingCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Billing.Contract.GetRoleMemberCount(&_Billing.CallOpts, role)
}

// GetSnapshot is a free data retrieval call binding the contract method 0x76f10ad0.
//
// Solidity: function getSnapshot(uint256 epoch) view returns(uint256 totalCollected, uint256 totalDistributed)
func (_Billing *BillingCaller) GetSnapshot(opts *bind.CallOpts, epoch *big.Int) (struct {
	TotalCollected   *big.Int
	TotalDistributed *big.Int
}, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "getSnapshot", epoch)

	outstruct := new(struct {
		TotalCollected   *big.Int
		TotalDistributed *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollected = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalDistributed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSnapshot is a free data retrieval call binding the contract method 0x76f10ad0.
//
// Solidity: function getSnapshot(uint256 epoch) view returns(uint256 totalCollected, uint256 totalDistributed)
func (_Billing *BillingSession) GetSnapshot(epoch *big.Int) (struct {
	TotalCollected   *big.Int
	TotalDistributed *big.Int
}, error) {
	return _Billing.Contract.GetSnapshot(&_Billing.CallOpts, epoch)
}

// GetSnapshot is a free data retrieval call binding the contract method 0x76f10ad0.
//
// Solidity: function getSnapshot(uint256 epoch) view returns(uint256 totalCollected, uint256 totalDistributed)
func (_Billing *BillingCallerSession) GetSnapshot(epoch *big.Int) (struct {
	TotalCollected   *big.Int
	TotalDistributed *big.Int
}, error) {
	return _Billing.Contract.GetSnapshot(&_Billing.CallOpts, epoch)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Billing *BillingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Billing *BillingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Billing.Contract.HasRole(&_Billing.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Billing *BillingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Billing.Contract.HasRole(&_Billing.CallOpts, role, account)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Billing *BillingCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Billing *BillingSession) StakingContract() (common.Address, error) {
	return _Billing.Contract.StakingContract(&_Billing.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Billing *BillingCallerSession) StakingContract() (common.Address, error) {
	return _Billing.Contract.StakingContract(&_Billing.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Billing *BillingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Billing.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Billing *BillingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Billing.Contract.SupportsInterface(&_Billing.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Billing *BillingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Billing.Contract.SupportsInterface(&_Billing.CallOpts, interfaceId)
}

// CollectTokens is a paid mutator transaction binding the contract method 0xc929f03d.
//
// Solidity: function collectTokens(uint256 epoch, address[] users, uint256[] amounts) returns(uint256 totalCollected)
func (_Billing *BillingTransactor) CollectTokens(opts *bind.TransactOpts, epoch *big.Int, users []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Billing.contract.Transact(opts, "collectTokens", epoch, users, amounts)
}

// CollectTokens is a paid mutator transaction binding the contract method 0xc929f03d.
//
// Solidity: function collectTokens(uint256 epoch, address[] users, uint256[] amounts) returns(uint256 totalCollected)
func (_Billing *BillingSession) CollectTokens(epoch *big.Int, users []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Billing.Contract.CollectTokens(&_Billing.TransactOpts, epoch, users, amounts)
}

// CollectTokens is a paid mutator transaction binding the contract method 0xc929f03d.
//
// Solidity: function collectTokens(uint256 epoch, address[] users, uint256[] amounts) returns(uint256 totalCollected)
func (_Billing *BillingTransactorSession) CollectTokens(epoch *big.Int, users []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Billing.Contract.CollectTokens(&_Billing.TransactOpts, epoch, users, amounts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Billing *BillingTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Billing.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Billing *BillingSession) Deposit() (*types.Transaction, error) {
	return _Billing.Contract.Deposit(&_Billing.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Billing *BillingTransactorSession) Deposit() (*types.Transaction, error) {
	return _Billing.Contract.Deposit(&_Billing.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x143ba4f3.
//
// Solidity: function distributeRewards(address[] nodeAddrs, uint256[] rewards) returns()
func (_Billing *BillingTransactor) DistributeRewards(opts *bind.TransactOpts, nodeAddrs []common.Address, rewards []*big.Int) (*types.Transaction, error) {
	return _Billing.contract.Transact(opts, "distributeRewards", nodeAddrs, rewards)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x143ba4f3.
//
// Solidity: function distributeRewards(address[] nodeAddrs, uint256[] rewards) returns()
func (_Billing *BillingSession) DistributeRewards(nodeAddrs []common.Address, rewards []*big.Int) (*types.Transaction, error) {
	return _Billing.Contract.DistributeRewards(&_Billing.TransactOpts, nodeAddrs, rewards)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x143ba4f3.
//
// Solidity: function distributeRewards(address[] nodeAddrs, uint256[] rewards) returns()
func (_Billing *BillingTransactorSession) DistributeRewards(nodeAddrs []common.Address, rewards []*big.Int) (*types.Transaction, error) {
	return _Billing.Contract.DistributeRewards(&_Billing.TransactOpts, nodeAddrs, rewards)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Billing *BillingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Billing.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Billing *BillingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Billing.Contract.GrantRole(&_Billing.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Billing *BillingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Billing.Contract.GrantRole(&_Billing.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address staking, address collector, uint256 startTime, uint256 startEpoch) returns()
func (_Billing *BillingTransactor) Initialize(opts *bind.TransactOpts, staking common.Address, collector common.Address, startTime *big.Int, startEpoch *big.Int) (*types.Transaction, error) {
	return _Billing.contract.Transact(opts, "initialize", staking, collector, startTime, startEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address staking, address collector, uint256 startTime, uint256 startEpoch) returns()
func (_Billing *BillingSession) Initialize(staking common.Address, collector common.Address, startTime *big.Int, startEpoch *big.Int) (*types.Transaction, error) {
	return _Billing.Contract.Initialize(&_Billing.TransactOpts, staking, collector, startTime, startEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address staking, address collector, uint256 startTime, uint256 startEpoch) returns()
func (_Billing *BillingTransactorSession) Initialize(staking common.Address, collector common.Address, startTime *big.Int, startEpoch *big.Int) (*types.Transaction, error) {
	return _Billing.Contract.Initialize(&_Billing.TransactOpts, staking, collector, startTime, startEpoch)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Billing *BillingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Billing.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Billing *BillingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Billing.Contract.RenounceRole(&_Billing.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Billing *BillingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Billing.Contract.RenounceRole(&_Billing.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Billing *BillingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Billing.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Billing *BillingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Billing.Contract.RevokeRole(&_Billing.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Billing *BillingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Billing.Contract.RevokeRole(&_Billing.TransactOpts, role, account)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0xbd13a803.
//
// Solidity: function withdrawTokens(address[] users, uint256[] amounts) returns()
func (_Billing *BillingTransactor) WithdrawTokens(opts *bind.TransactOpts, users []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Billing.contract.Transact(opts, "withdrawTokens", users, amounts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0xbd13a803.
//
// Solidity: function withdrawTokens(address[] users, uint256[] amounts) returns()
func (_Billing *BillingSession) WithdrawTokens(users []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Billing.Contract.WithdrawTokens(&_Billing.TransactOpts, users, amounts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0xbd13a803.
//
// Solidity: function withdrawTokens(address[] users, uint256[] amounts) returns()
func (_Billing *BillingTransactorSession) WithdrawTokens(users []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Billing.Contract.WithdrawTokens(&_Billing.TransactOpts, users, amounts)
}

// BillingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Billing contract.
type BillingInitializedIterator struct {
	Event *BillingInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BillingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BillingInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BillingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingInitialized represents a Initialized event raised by the Billing contract.
type BillingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Billing *BillingFilterer) FilterInitialized(opts *bind.FilterOpts) (*BillingInitializedIterator, error) {

	logs, sub, err := _Billing.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BillingInitializedIterator{contract: _Billing.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Billing *BillingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BillingInitialized) (event.Subscription, error) {

	logs, sub, err := _Billing.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingInitialized)
				if err := _Billing.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Billing *BillingFilterer) ParseInitialized(log types.Log) (*BillingInitialized, error) {
	event := new(BillingInitialized)
	if err := _Billing.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRewardsDistributedIterator is returned from FilterRewardsDistributed and is used to iterate over the raw logs and unpacked data for RewardsDistributed events raised by the Billing contract.
type BillingRewardsDistributedIterator struct {
	Event *BillingRewardsDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BillingRewardsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRewardsDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BillingRewardsDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BillingRewardsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRewardsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRewardsDistributed represents a RewardsDistributed event raised by the Billing contract.
type BillingRewardsDistributed struct {
	NodeAddr     common.Address
	Amount       *big.Int
	Receiver     common.Address
	IsPublicGood bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributed is a free log retrieval operation binding the contract event 0x6e84718920f8e2f7cd5f87a6d7578558898c604ebe2c631e81b50256f626e693.
//
// Solidity: event RewardsDistributed(address indexed nodeAddr, uint256 indexed amount, address indexed receiver, bool isPublicGood)
func (_Billing *BillingFilterer) FilterRewardsDistributed(opts *bind.FilterOpts, nodeAddr []common.Address, amount []*big.Int, receiver []common.Address) (*BillingRewardsDistributedIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Billing.contract.FilterLogs(opts, "RewardsDistributed", nodeAddrRule, amountRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &BillingRewardsDistributedIterator{contract: _Billing.contract, event: "RewardsDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributed is a free log subscription operation binding the contract event 0x6e84718920f8e2f7cd5f87a6d7578558898c604ebe2c631e81b50256f626e693.
//
// Solidity: event RewardsDistributed(address indexed nodeAddr, uint256 indexed amount, address indexed receiver, bool isPublicGood)
func (_Billing *BillingFilterer) WatchRewardsDistributed(opts *bind.WatchOpts, sink chan<- *BillingRewardsDistributed, nodeAddr []common.Address, amount []*big.Int, receiver []common.Address) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Billing.contract.WatchLogs(opts, "RewardsDistributed", nodeAddrRule, amountRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRewardsDistributed)
				if err := _Billing.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsDistributed is a log parse operation binding the contract event 0x6e84718920f8e2f7cd5f87a6d7578558898c604ebe2c631e81b50256f626e693.
//
// Solidity: event RewardsDistributed(address indexed nodeAddr, uint256 indexed amount, address indexed receiver, bool isPublicGood)
func (_Billing *BillingFilterer) ParseRewardsDistributed(log types.Log) (*BillingRewardsDistributed, error) {
	event := new(BillingRewardsDistributed)
	if err := _Billing.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Billing contract.
type BillingRoleAdminChangedIterator struct {
	Event *BillingRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BillingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BillingRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BillingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRoleAdminChanged represents a RoleAdminChanged event raised by the Billing contract.
type BillingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Billing *BillingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BillingRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Billing.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BillingRoleAdminChangedIterator{contract: _Billing.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Billing *BillingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BillingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Billing.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRoleAdminChanged)
				if err := _Billing.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Billing *BillingFilterer) ParseRoleAdminChanged(log types.Log) (*BillingRoleAdminChanged, error) {
	event := new(BillingRoleAdminChanged)
	if err := _Billing.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Billing contract.
type BillingRoleGrantedIterator struct {
	Event *BillingRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BillingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BillingRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BillingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRoleGranted represents a RoleGranted event raised by the Billing contract.
type BillingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Billing *BillingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BillingRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Billing.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BillingRoleGrantedIterator{contract: _Billing.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Billing *BillingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BillingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Billing.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRoleGranted)
				if err := _Billing.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Billing *BillingFilterer) ParseRoleGranted(log types.Log) (*BillingRoleGranted, error) {
	event := new(BillingRoleGranted)
	if err := _Billing.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Billing contract.
type BillingRoleRevokedIterator struct {
	Event *BillingRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BillingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BillingRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BillingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRoleRevoked represents a RoleRevoked event raised by the Billing contract.
type BillingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Billing *BillingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BillingRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Billing.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BillingRoleRevokedIterator{contract: _Billing.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Billing *BillingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BillingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Billing.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRoleRevoked)
				if err := _Billing.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Billing *BillingFilterer) ParseRoleRevoked(log types.Log) (*BillingRoleRevoked, error) {
	event := new(BillingRoleRevoked)
	if err := _Billing.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingTokensCollectedIterator is returned from FilterTokensCollected and is used to iterate over the raw logs and unpacked data for TokensCollected events raised by the Billing contract.
type BillingTokensCollectedIterator struct {
	Event *BillingTokensCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BillingTokensCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingTokensCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BillingTokensCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BillingTokensCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingTokensCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingTokensCollected represents a TokensCollected event raised by the Billing contract.
type BillingTokensCollected struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensCollected is a free log retrieval operation binding the contract event 0x9381e53ffdc9733a6783a6f8665be3f89c231bb81a6771996ed553b4e75c0fe3.
//
// Solidity: event TokensCollected(address indexed user, uint256 indexed amount)
func (_Billing *BillingFilterer) FilterTokensCollected(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*BillingTokensCollectedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Billing.contract.FilterLogs(opts, "TokensCollected", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &BillingTokensCollectedIterator{contract: _Billing.contract, event: "TokensCollected", logs: logs, sub: sub}, nil
}

// WatchTokensCollected is a free log subscription operation binding the contract event 0x9381e53ffdc9733a6783a6f8665be3f89c231bb81a6771996ed553b4e75c0fe3.
//
// Solidity: event TokensCollected(address indexed user, uint256 indexed amount)
func (_Billing *BillingFilterer) WatchTokensCollected(opts *bind.WatchOpts, sink chan<- *BillingTokensCollected, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Billing.contract.WatchLogs(opts, "TokensCollected", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingTokensCollected)
				if err := _Billing.contract.UnpackLog(event, "TokensCollected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensCollected is a log parse operation binding the contract event 0x9381e53ffdc9733a6783a6f8665be3f89c231bb81a6771996ed553b4e75c0fe3.
//
// Solidity: event TokensCollected(address indexed user, uint256 indexed amount)
func (_Billing *BillingFilterer) ParseTokensCollected(log types.Log) (*BillingTokensCollected, error) {
	event := new(BillingTokensCollected)
	if err := _Billing.contract.UnpackLog(event, "TokensCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingTokensDepositedIterator is returned from FilterTokensDeposited and is used to iterate over the raw logs and unpacked data for TokensDeposited events raised by the Billing contract.
type BillingTokensDepositedIterator struct {
	Event *BillingTokensDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BillingTokensDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingTokensDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BillingTokensDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BillingTokensDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingTokensDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingTokensDeposited represents a TokensDeposited event raised by the Billing contract.
type BillingTokensDeposited struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensDeposited is a free log retrieval operation binding the contract event 0x59062170a285eb80e8c6b8ced60428442a51910635005233fc4ce084a475845e.
//
// Solidity: event TokensDeposited(address indexed user, uint256 indexed amount)
func (_Billing *BillingFilterer) FilterTokensDeposited(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*BillingTokensDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Billing.contract.FilterLogs(opts, "TokensDeposited", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &BillingTokensDepositedIterator{contract: _Billing.contract, event: "TokensDeposited", logs: logs, sub: sub}, nil
}

// WatchTokensDeposited is a free log subscription operation binding the contract event 0x59062170a285eb80e8c6b8ced60428442a51910635005233fc4ce084a475845e.
//
// Solidity: event TokensDeposited(address indexed user, uint256 indexed amount)
func (_Billing *BillingFilterer) WatchTokensDeposited(opts *bind.WatchOpts, sink chan<- *BillingTokensDeposited, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Billing.contract.WatchLogs(opts, "TokensDeposited", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingTokensDeposited)
				if err := _Billing.contract.UnpackLog(event, "TokensDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensDeposited is a log parse operation binding the contract event 0x59062170a285eb80e8c6b8ced60428442a51910635005233fc4ce084a475845e.
//
// Solidity: event TokensDeposited(address indexed user, uint256 indexed amount)
func (_Billing *BillingFilterer) ParseTokensDeposited(log types.Log) (*BillingTokensDeposited, error) {
	event := new(BillingTokensDeposited)
	if err := _Billing.contract.UnpackLog(event, "TokensDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the Billing contract.
type BillingTokensWithdrawnIterator struct {
	Event *BillingTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BillingTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BillingTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BillingTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingTokensWithdrawn represents a TokensWithdrawn event raised by the Billing contract.
type BillingTokensWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed user, uint256 indexed amount)
func (_Billing *BillingFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*BillingTokensWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Billing.contract.FilterLogs(opts, "TokensWithdrawn", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &BillingTokensWithdrawnIterator{contract: _Billing.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed user, uint256 indexed amount)
func (_Billing *BillingFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *BillingTokensWithdrawn, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Billing.contract.WatchLogs(opts, "TokensWithdrawn", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingTokensWithdrawn)
				if err := _Billing.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed user, uint256 indexed amount)
func (_Billing *BillingFilterer) ParseTokensWithdrawn(log types.Log) (*BillingTokensWithdrawn, error) {
	event := new(BillingTokensWithdrawn)
	if err := _Billing.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Enrockvault

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
)

// EnrockvaultMetaData contains all meta data concerning the Enrockvault contract.
var EnrockvaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonateDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"h\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dot\",\"type\":\"uint256\"}],\"name\":\"GameVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdatedBank\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficier\",\"type\":\"address\"}],\"name\":\"UpdatedWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UsrDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UsrWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalGov\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"contractPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"value\",\"type\":\"uint8\"}],\"name\":\"traillock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bank_vault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"uint256[6]\",\"name\":\"chr\",\"type\":\"uint256[6]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit_erc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"donate_erc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game_round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"total_users\",\"type\":\"uint16\"}],\"name\":\"oracle_post_result\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pool_fee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pool_fee_discount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pool_vault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelistAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"round_vault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dot\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"users\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"perc\",\"type\":\"uint16\"}],\"name\":\"set_discount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"x\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"percent_1000\",\"type\":\"uint16\"}],\"name\":\"set_fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n_round\",\"type\":\"uint256\"}],\"name\":\"set_round\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"}],\"name\":\"signature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"take_usr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseSc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"coins\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"name\":\"up_bank_vault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"up_single\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coins\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"up_usr_batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"k\",\"type\":\"address\"}],\"name\":\"updateProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistAdmins\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EnrockvaultABI is the input ABI used to generate the binding from.
// Deprecated: Use EnrockvaultMetaData.ABI instead.
var EnrockvaultABI = EnrockvaultMetaData.ABI

// Enrockvault is an auto generated Go binding around an Ethereum contract.
type Enrockvault struct {
	EnrockvaultCaller     // Read-only binding to the contract
	EnrockvaultTransactor // Write-only binding to the contract
	EnrockvaultFilterer   // Log filterer for contract events
}

// EnrockvaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnrockvaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnrockvaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnrockvaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnrockvaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnrockvaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnrockvaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnrockvaultSession struct {
	Contract     *Enrockvault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnrockvaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnrockvaultCallerSession struct {
	Contract *EnrockvaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EnrockvaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnrockvaultTransactorSession struct {
	Contract     *EnrockvaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EnrockvaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnrockvaultRaw struct {
	Contract *Enrockvault // Generic contract binding to access the raw methods on
}

// EnrockvaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnrockvaultCallerRaw struct {
	Contract *EnrockvaultCaller // Generic read-only contract binding to access the raw methods on
}

// EnrockvaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnrockvaultTransactorRaw struct {
	Contract *EnrockvaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnrockvault creates a new instance of Enrockvault, bound to a specific deployed contract.
func NewEnrockvault(address common.Address, backend bind.ContractBackend) (*Enrockvault, error) {
	contract, err := bindEnrockvault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Enrockvault{EnrockvaultCaller: EnrockvaultCaller{contract: contract}, EnrockvaultTransactor: EnrockvaultTransactor{contract: contract}, EnrockvaultFilterer: EnrockvaultFilterer{contract: contract}}, nil
}

// NewEnrockvaultCaller creates a new read-only instance of Enrockvault, bound to a specific deployed contract.
func NewEnrockvaultCaller(address common.Address, caller bind.ContractCaller) (*EnrockvaultCaller, error) {
	contract, err := bindEnrockvault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnrockvaultCaller{contract: contract}, nil
}

// NewEnrockvaultTransactor creates a new write-only instance of Enrockvault, bound to a specific deployed contract.
func NewEnrockvaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EnrockvaultTransactor, error) {
	contract, err := bindEnrockvault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnrockvaultTransactor{contract: contract}, nil
}

// NewEnrockvaultFilterer creates a new log filterer instance of Enrockvault, bound to a specific deployed contract.
func NewEnrockvaultFilterer(address common.Address, filterer bind.ContractFilterer) (*EnrockvaultFilterer, error) {
	contract, err := bindEnrockvault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnrockvaultFilterer{contract: contract}, nil
}

// bindEnrockvault binds a generic wrapper to an already deployed contract.
func bindEnrockvault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnrockvaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enrockvault *EnrockvaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Enrockvault.Contract.EnrockvaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enrockvault *EnrockvaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enrockvault.Contract.EnrockvaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enrockvault *EnrockvaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enrockvault.Contract.EnrockvaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enrockvault *EnrockvaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Enrockvault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enrockvault *EnrockvaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enrockvault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enrockvault *EnrockvaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enrockvault.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Enrockvault *EnrockvaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Enrockvault *EnrockvaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Enrockvault.Contract.DEFAULTADMINROLE(&_Enrockvault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Enrockvault *EnrockvaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Enrockvault.Contract.DEFAULTADMINROLE(&_Enrockvault.CallOpts)
}

// BankVault is a free data retrieval call binding the contract method 0xbad43ca9.
//
// Solidity: function bank_vault(address ) view returns(uint256)
func (_Enrockvault *EnrockvaultCaller) BankVault(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "bank_vault", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BankVault is a free data retrieval call binding the contract method 0xbad43ca9.
//
// Solidity: function bank_vault(address ) view returns(uint256)
func (_Enrockvault *EnrockvaultSession) BankVault(arg0 common.Address) (*big.Int, error) {
	return _Enrockvault.Contract.BankVault(&_Enrockvault.CallOpts, arg0)
}

// BankVault is a free data retrieval call binding the contract method 0xbad43ca9.
//
// Solidity: function bank_vault(address ) view returns(uint256)
func (_Enrockvault *EnrockvaultCallerSession) BankVault(arg0 common.Address) (*big.Int, error) {
	return _Enrockvault.Contract.BankVault(&_Enrockvault.CallOpts, arg0)
}

// Check is a free data retrieval call binding the contract method 0xb3154db0.
//
// Solidity: function check(address coin, address user) view returns(uint256[6] chr)
func (_Enrockvault *EnrockvaultCaller) Check(opts *bind.CallOpts, coin common.Address, user common.Address) ([6]*big.Int, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "check", coin, user)

	if err != nil {
		return *new([6]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([6]*big.Int)).(*[6]*big.Int)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0xb3154db0.
//
// Solidity: function check(address coin, address user) view returns(uint256[6] chr)
func (_Enrockvault *EnrockvaultSession) Check(coin common.Address, user common.Address) ([6]*big.Int, error) {
	return _Enrockvault.Contract.Check(&_Enrockvault.CallOpts, coin, user)
}

// Check is a free data retrieval call binding the contract method 0xb3154db0.
//
// Solidity: function check(address coin, address user) view returns(uint256[6] chr)
func (_Enrockvault *EnrockvaultCallerSession) Check(coin common.Address, user common.Address) ([6]*big.Int, error) {
	return _Enrockvault.Contract.Check(&_Enrockvault.CallOpts, coin, user)
}

// GameRound is a free data retrieval call binding the contract method 0x4b6b0cdf.
//
// Solidity: function game_round() view returns(uint256)
func (_Enrockvault *EnrockvaultCaller) GameRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "game_round")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GameRound is a free data retrieval call binding the contract method 0x4b6b0cdf.
//
// Solidity: function game_round() view returns(uint256)
func (_Enrockvault *EnrockvaultSession) GameRound() (*big.Int, error) {
	return _Enrockvault.Contract.GameRound(&_Enrockvault.CallOpts)
}

// GameRound is a free data retrieval call binding the contract method 0x4b6b0cdf.
//
// Solidity: function game_round() view returns(uint256)
func (_Enrockvault *EnrockvaultCallerSession) GameRound() (*big.Int, error) {
	return _Enrockvault.Contract.GameRound(&_Enrockvault.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Enrockvault *EnrockvaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Enrockvault *EnrockvaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Enrockvault.Contract.GetRoleAdmin(&_Enrockvault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Enrockvault *EnrockvaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Enrockvault.Contract.GetRoleAdmin(&_Enrockvault.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Enrockvault *EnrockvaultCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Enrockvault *EnrockvaultSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Enrockvault.Contract.GetRoleMember(&_Enrockvault.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Enrockvault *EnrockvaultCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Enrockvault.Contract.GetRoleMember(&_Enrockvault.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Enrockvault *EnrockvaultCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Enrockvault *EnrockvaultSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Enrockvault.Contract.GetRoleMemberCount(&_Enrockvault.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Enrockvault *EnrockvaultCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Enrockvault.Contract.GetRoleMemberCount(&_Enrockvault.CallOpts, role)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(bytes32)
func (_Enrockvault *EnrockvaultCaller) Governor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(bytes32)
func (_Enrockvault *EnrockvaultSession) Governor() ([32]byte, error) {
	return _Enrockvault.Contract.Governor(&_Enrockvault.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(bytes32)
func (_Enrockvault *EnrockvaultCallerSession) Governor() ([32]byte, error) {
	return _Enrockvault.Contract.Governor(&_Enrockvault.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Enrockvault *EnrockvaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Enrockvault *EnrockvaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Enrockvault.Contract.HasRole(&_Enrockvault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Enrockvault *EnrockvaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Enrockvault.Contract.HasRole(&_Enrockvault.CallOpts, role, account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address account) view returns(bool)
func (_Enrockvault *EnrockvaultCaller) IsGovernor(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "isGovernor", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address account) view returns(bool)
func (_Enrockvault *EnrockvaultSession) IsGovernor(account common.Address) (bool, error) {
	return _Enrockvault.Contract.IsGovernor(&_Enrockvault.CallOpts, account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address account) view returns(bool)
func (_Enrockvault *EnrockvaultCallerSession) IsGovernor(account common.Address) (bool, error) {
	return _Enrockvault.Contract.IsGovernor(&_Enrockvault.CallOpts, account)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_Enrockvault *EnrockvaultCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_Enrockvault *EnrockvaultSession) IsLocked() (bool, error) {
	return _Enrockvault.Contract.IsLocked(&_Enrockvault.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_Enrockvault *EnrockvaultCallerSession) IsLocked() (bool, error) {
	return _Enrockvault.Contract.IsLocked(&_Enrockvault.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Enrockvault *EnrockvaultCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Enrockvault *EnrockvaultSession) IsOwner() (bool, error) {
	return _Enrockvault.Contract.IsOwner(&_Enrockvault.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Enrockvault *EnrockvaultCallerSession) IsOwner() (bool, error) {
	return _Enrockvault.Contract.IsOwner(&_Enrockvault.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Enrockvault *EnrockvaultCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Enrockvault *EnrockvaultSession) IsPaused() (bool, error) {
	return _Enrockvault.Contract.IsPaused(&_Enrockvault.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Enrockvault *EnrockvaultCallerSession) IsPaused() (bool, error) {
	return _Enrockvault.Contract.IsPaused(&_Enrockvault.CallOpts)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_Enrockvault *EnrockvaultCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "isWhitelistAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_Enrockvault *EnrockvaultSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _Enrockvault.Contract.IsWhitelistAdmin(&_Enrockvault.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_Enrockvault *EnrockvaultCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _Enrockvault.Contract.IsWhitelistAdmin(&_Enrockvault.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Enrockvault *EnrockvaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Enrockvault *EnrockvaultSession) Owner() (common.Address, error) {
	return _Enrockvault.Contract.Owner(&_Enrockvault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Enrockvault *EnrockvaultCallerSession) Owner() (common.Address, error) {
	return _Enrockvault.Contract.Owner(&_Enrockvault.CallOpts)
}

// PoolFee is a free data retrieval call binding the contract method 0x61282ebe.
//
// Solidity: function pool_fee(address ) view returns(uint16)
func (_Enrockvault *EnrockvaultCaller) PoolFee(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "pool_fee", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolFee is a free data retrieval call binding the contract method 0x61282ebe.
//
// Solidity: function pool_fee(address ) view returns(uint16)
func (_Enrockvault *EnrockvaultSession) PoolFee(arg0 common.Address) (uint16, error) {
	return _Enrockvault.Contract.PoolFee(&_Enrockvault.CallOpts, arg0)
}

// PoolFee is a free data retrieval call binding the contract method 0x61282ebe.
//
// Solidity: function pool_fee(address ) view returns(uint16)
func (_Enrockvault *EnrockvaultCallerSession) PoolFee(arg0 common.Address) (uint16, error) {
	return _Enrockvault.Contract.PoolFee(&_Enrockvault.CallOpts, arg0)
}

// PoolFeeDiscount is a free data retrieval call binding the contract method 0x5d9d08b1.
//
// Solidity: function pool_fee_discount(address ) view returns(uint16)
func (_Enrockvault *EnrockvaultCaller) PoolFeeDiscount(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "pool_fee_discount", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolFeeDiscount is a free data retrieval call binding the contract method 0x5d9d08b1.
//
// Solidity: function pool_fee_discount(address ) view returns(uint16)
func (_Enrockvault *EnrockvaultSession) PoolFeeDiscount(arg0 common.Address) (uint16, error) {
	return _Enrockvault.Contract.PoolFeeDiscount(&_Enrockvault.CallOpts, arg0)
}

// PoolFeeDiscount is a free data retrieval call binding the contract method 0x5d9d08b1.
//
// Solidity: function pool_fee_discount(address ) view returns(uint16)
func (_Enrockvault *EnrockvaultCallerSession) PoolFeeDiscount(arg0 common.Address) (uint16, error) {
	return _Enrockvault.Contract.PoolFeeDiscount(&_Enrockvault.CallOpts, arg0)
}

// PoolVault is a free data retrieval call binding the contract method 0xf369acaa.
//
// Solidity: function pool_vault(address ) view returns(uint256)
func (_Enrockvault *EnrockvaultCaller) PoolVault(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "pool_vault", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolVault is a free data retrieval call binding the contract method 0xf369acaa.
//
// Solidity: function pool_vault(address ) view returns(uint256)
func (_Enrockvault *EnrockvaultSession) PoolVault(arg0 common.Address) (*big.Int, error) {
	return _Enrockvault.Contract.PoolVault(&_Enrockvault.CallOpts, arg0)
}

// PoolVault is a free data retrieval call binding the contract method 0xf369acaa.
//
// Solidity: function pool_vault(address ) view returns(uint256)
func (_Enrockvault *EnrockvaultCallerSession) PoolVault(arg0 common.Address) (*big.Int, error) {
	return _Enrockvault.Contract.PoolVault(&_Enrockvault.CallOpts, arg0)
}

// RoundVault is a free data retrieval call binding the contract method 0x38d4ee6e.
//
// Solidity: function round_vault(uint256 ) view returns(uint256 dot, uint16 users, string hash)
func (_Enrockvault *EnrockvaultCaller) RoundVault(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Dot   *big.Int
	Users uint16
	Hash  string
}, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "round_vault", arg0)

	outstruct := new(struct {
		Dot   *big.Int
		Users uint16
		Hash  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dot = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Users = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.Hash = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// RoundVault is a free data retrieval call binding the contract method 0x38d4ee6e.
//
// Solidity: function round_vault(uint256 ) view returns(uint256 dot, uint16 users, string hash)
func (_Enrockvault *EnrockvaultSession) RoundVault(arg0 *big.Int) (struct {
	Dot   *big.Int
	Users uint16
	Hash  string
}, error) {
	return _Enrockvault.Contract.RoundVault(&_Enrockvault.CallOpts, arg0)
}

// RoundVault is a free data retrieval call binding the contract method 0x38d4ee6e.
//
// Solidity: function round_vault(uint256 ) view returns(uint256 dot, uint16 users, string hash)
func (_Enrockvault *EnrockvaultCallerSession) RoundVault(arg0 *big.Int) (struct {
	Dot   *big.Int
	Users uint16
	Hash  string
}, error) {
	return _Enrockvault.Contract.RoundVault(&_Enrockvault.CallOpts, arg0)
}

// Signature is a free data retrieval call binding the contract method 0x5361d760.
//
// Solidity: function signature(address user, address coin) pure returns(bytes32)
func (_Enrockvault *EnrockvaultCaller) Signature(opts *bind.CallOpts, user common.Address, coin common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "signature", user, coin)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Signature is a free data retrieval call binding the contract method 0x5361d760.
//
// Solidity: function signature(address user, address coin) pure returns(bytes32)
func (_Enrockvault *EnrockvaultSession) Signature(user common.Address, coin common.Address) ([32]byte, error) {
	return _Enrockvault.Contract.Signature(&_Enrockvault.CallOpts, user, coin)
}

// Signature is a free data retrieval call binding the contract method 0x5361d760.
//
// Solidity: function signature(address user, address coin) pure returns(bytes32)
func (_Enrockvault *EnrockvaultCallerSession) Signature(user common.Address, coin common.Address) ([32]byte, error) {
	return _Enrockvault.Contract.Signature(&_Enrockvault.CallOpts, user, coin)
}

// WhitelistAdmins is a free data retrieval call binding the contract method 0x7ef581cc.
//
// Solidity: function whitelistAdmins() view returns(bytes32)
func (_Enrockvault *EnrockvaultCaller) WhitelistAdmins(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Enrockvault.contract.Call(opts, &out, "whitelistAdmins")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WhitelistAdmins is a free data retrieval call binding the contract method 0x7ef581cc.
//
// Solidity: function whitelistAdmins() view returns(bytes32)
func (_Enrockvault *EnrockvaultSession) WhitelistAdmins() ([32]byte, error) {
	return _Enrockvault.Contract.WhitelistAdmins(&_Enrockvault.CallOpts)
}

// WhitelistAdmins is a free data retrieval call binding the contract method 0x7ef581cc.
//
// Solidity: function whitelistAdmins() view returns(bytes32)
func (_Enrockvault *EnrockvaultCallerSession) WhitelistAdmins() ([32]byte, error) {
	return _Enrockvault.Contract.WhitelistAdmins(&_Enrockvault.CallOpts)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address account) returns()
func (_Enrockvault *EnrockvaultTransactor) AddGovernor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "addGovernor", account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address account) returns()
func (_Enrockvault *EnrockvaultSession) AddGovernor(account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.AddGovernor(&_Enrockvault.TransactOpts, account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address account) returns()
func (_Enrockvault *EnrockvaultTransactorSession) AddGovernor(account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.AddGovernor(&_Enrockvault.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_Enrockvault *EnrockvaultTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_Enrockvault *EnrockvaultSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.AddWhitelistAdmin(&_Enrockvault.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_Enrockvault *EnrockvaultTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.AddWhitelistAdmin(&_Enrockvault.TransactOpts, account)
}

// DepositErc20 is a paid mutator transaction binding the contract method 0xd16544f0.
//
// Solidity: function deposit_erc20(address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultTransactor) DepositErc20(opts *bind.TransactOpts, coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "deposit_erc20", coin, amount)
}

// DepositErc20 is a paid mutator transaction binding the contract method 0xd16544f0.
//
// Solidity: function deposit_erc20(address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultSession) DepositErc20(coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.DepositErc20(&_Enrockvault.TransactOpts, coin, amount)
}

// DepositErc20 is a paid mutator transaction binding the contract method 0xd16544f0.
//
// Solidity: function deposit_erc20(address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultTransactorSession) DepositErc20(coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.DepositErc20(&_Enrockvault.TransactOpts, coin, amount)
}

// DonateErc20 is a paid mutator transaction binding the contract method 0xd17af65b.
//
// Solidity: function donate_erc20(address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultTransactor) DonateErc20(opts *bind.TransactOpts, coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "donate_erc20", coin, amount)
}

// DonateErc20 is a paid mutator transaction binding the contract method 0xd17af65b.
//
// Solidity: function donate_erc20(address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultSession) DonateErc20(coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.DonateErc20(&_Enrockvault.TransactOpts, coin, amount)
}

// DonateErc20 is a paid mutator transaction binding the contract method 0xd17af65b.
//
// Solidity: function donate_erc20(address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultTransactorSession) DonateErc20(coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.DonateErc20(&_Enrockvault.TransactOpts, coin, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Enrockvault *EnrockvaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Enrockvault *EnrockvaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.GrantRole(&_Enrockvault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Enrockvault *EnrockvaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.GrantRole(&_Enrockvault.TransactOpts, role, account)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Enrockvault *EnrockvaultTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Enrockvault *EnrockvaultSession) Lock() (*types.Transaction, error) {
	return _Enrockvault.Contract.Lock(&_Enrockvault.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Enrockvault *EnrockvaultTransactorSession) Lock() (*types.Transaction, error) {
	return _Enrockvault.Contract.Lock(&_Enrockvault.TransactOpts)
}

// OraclePostResult is a paid mutator transaction binding the contract method 0x0db1b9df.
//
// Solidity: function oracle_post_result(uint256 round, uint256 dot, string hash, uint16 total_users) returns()
func (_Enrockvault *EnrockvaultTransactor) OraclePostResult(opts *bind.TransactOpts, round *big.Int, dot *big.Int, hash string, total_users uint16) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "oracle_post_result", round, dot, hash, total_users)
}

// OraclePostResult is a paid mutator transaction binding the contract method 0x0db1b9df.
//
// Solidity: function oracle_post_result(uint256 round, uint256 dot, string hash, uint16 total_users) returns()
func (_Enrockvault *EnrockvaultSession) OraclePostResult(round *big.Int, dot *big.Int, hash string, total_users uint16) (*types.Transaction, error) {
	return _Enrockvault.Contract.OraclePostResult(&_Enrockvault.TransactOpts, round, dot, hash, total_users)
}

// OraclePostResult is a paid mutator transaction binding the contract method 0x0db1b9df.
//
// Solidity: function oracle_post_result(uint256 round, uint256 dot, string hash, uint16 total_users) returns()
func (_Enrockvault *EnrockvaultTransactorSession) OraclePostResult(round *big.Int, dot *big.Int, hash string, total_users uint16) (*types.Transaction, error) {
	return _Enrockvault.Contract.OraclePostResult(&_Enrockvault.TransactOpts, round, dot, hash, total_users)
}

// PauseSc is a paid mutator transaction binding the contract method 0x1975ebaf.
//
// Solidity: function pauseSc() returns()
func (_Enrockvault *EnrockvaultTransactor) PauseSc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "pauseSc")
}

// PauseSc is a paid mutator transaction binding the contract method 0x1975ebaf.
//
// Solidity: function pauseSc() returns()
func (_Enrockvault *EnrockvaultSession) PauseSc() (*types.Transaction, error) {
	return _Enrockvault.Contract.PauseSc(&_Enrockvault.TransactOpts)
}

// PauseSc is a paid mutator transaction binding the contract method 0x1975ebaf.
//
// Solidity: function pauseSc() returns()
func (_Enrockvault *EnrockvaultTransactorSession) PauseSc() (*types.Transaction, error) {
	return _Enrockvault.Contract.PauseSc(&_Enrockvault.TransactOpts)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address account) returns()
func (_Enrockvault *EnrockvaultTransactor) RemoveGovernor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "removeGovernor", account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address account) returns()
func (_Enrockvault *EnrockvaultSession) RemoveGovernor(account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.RemoveGovernor(&_Enrockvault.TransactOpts, account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address account) returns()
func (_Enrockvault *EnrockvaultTransactorSession) RemoveGovernor(account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.RemoveGovernor(&_Enrockvault.TransactOpts, account)
}

// RemoveWhitelistAdmin is a paid mutator transaction binding the contract method 0x6897e974.
//
// Solidity: function removeWhitelistAdmin(address account) returns()
func (_Enrockvault *EnrockvaultTransactor) RemoveWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "removeWhitelistAdmin", account)
}

// RemoveWhitelistAdmin is a paid mutator transaction binding the contract method 0x6897e974.
//
// Solidity: function removeWhitelistAdmin(address account) returns()
func (_Enrockvault *EnrockvaultSession) RemoveWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.RemoveWhitelistAdmin(&_Enrockvault.TransactOpts, account)
}

// RemoveWhitelistAdmin is a paid mutator transaction binding the contract method 0x6897e974.
//
// Solidity: function removeWhitelistAdmin(address account) returns()
func (_Enrockvault *EnrockvaultTransactorSession) RemoveWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.RemoveWhitelistAdmin(&_Enrockvault.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Enrockvault *EnrockvaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Enrockvault *EnrockvaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _Enrockvault.Contract.RenounceOwnership(&_Enrockvault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Enrockvault *EnrockvaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Enrockvault.Contract.RenounceOwnership(&_Enrockvault.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Enrockvault *EnrockvaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Enrockvault *EnrockvaultSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.RenounceRole(&_Enrockvault.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Enrockvault *EnrockvaultTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.RenounceRole(&_Enrockvault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Enrockvault *EnrockvaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Enrockvault *EnrockvaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.RevokeRole(&_Enrockvault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Enrockvault *EnrockvaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.RevokeRole(&_Enrockvault.TransactOpts, role, account)
}

// SetDiscount is a paid mutator transaction binding the contract method 0x8bac6a64.
//
// Solidity: function set_discount(address user, uint16 perc) returns()
func (_Enrockvault *EnrockvaultTransactor) SetDiscount(opts *bind.TransactOpts, user common.Address, perc uint16) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "set_discount", user, perc)
}

// SetDiscount is a paid mutator transaction binding the contract method 0x8bac6a64.
//
// Solidity: function set_discount(address user, uint16 perc) returns()
func (_Enrockvault *EnrockvaultSession) SetDiscount(user common.Address, perc uint16) (*types.Transaction, error) {
	return _Enrockvault.Contract.SetDiscount(&_Enrockvault.TransactOpts, user, perc)
}

// SetDiscount is a paid mutator transaction binding the contract method 0x8bac6a64.
//
// Solidity: function set_discount(address user, uint16 perc) returns()
func (_Enrockvault *EnrockvaultTransactorSession) SetDiscount(user common.Address, perc uint16) (*types.Transaction, error) {
	return _Enrockvault.Contract.SetDiscount(&_Enrockvault.TransactOpts, user, perc)
}

// SetFee is a paid mutator transaction binding the contract method 0x8340d73f.
//
// Solidity: function set_fee(address x, uint16 percent_1000) returns()
func (_Enrockvault *EnrockvaultTransactor) SetFee(opts *bind.TransactOpts, x common.Address, percent_1000 uint16) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "set_fee", x, percent_1000)
}

// SetFee is a paid mutator transaction binding the contract method 0x8340d73f.
//
// Solidity: function set_fee(address x, uint16 percent_1000) returns()
func (_Enrockvault *EnrockvaultSession) SetFee(x common.Address, percent_1000 uint16) (*types.Transaction, error) {
	return _Enrockvault.Contract.SetFee(&_Enrockvault.TransactOpts, x, percent_1000)
}

// SetFee is a paid mutator transaction binding the contract method 0x8340d73f.
//
// Solidity: function set_fee(address x, uint16 percent_1000) returns()
func (_Enrockvault *EnrockvaultTransactorSession) SetFee(x common.Address, percent_1000 uint16) (*types.Transaction, error) {
	return _Enrockvault.Contract.SetFee(&_Enrockvault.TransactOpts, x, percent_1000)
}

// SetRound is a paid mutator transaction binding the contract method 0x8c91edae.
//
// Solidity: function set_round(uint256 n_round) returns()
func (_Enrockvault *EnrockvaultTransactor) SetRound(opts *bind.TransactOpts, n_round *big.Int) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "set_round", n_round)
}

// SetRound is a paid mutator transaction binding the contract method 0x8c91edae.
//
// Solidity: function set_round(uint256 n_round) returns()
func (_Enrockvault *EnrockvaultSession) SetRound(n_round *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.SetRound(&_Enrockvault.TransactOpts, n_round)
}

// SetRound is a paid mutator transaction binding the contract method 0x8c91edae.
//
// Solidity: function set_round(uint256 n_round) returns()
func (_Enrockvault *EnrockvaultTransactorSession) SetRound(n_round *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.SetRound(&_Enrockvault.TransactOpts, n_round)
}

// TakeUsr is a paid mutator transaction binding the contract method 0x7cb89783.
//
// Solidity: function take_usr(address user, address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultTransactor) TakeUsr(opts *bind.TransactOpts, user common.Address, coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "take_usr", user, coin, amount)
}

// TakeUsr is a paid mutator transaction binding the contract method 0x7cb89783.
//
// Solidity: function take_usr(address user, address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultSession) TakeUsr(user common.Address, coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.TakeUsr(&_Enrockvault.TransactOpts, user, coin, amount)
}

// TakeUsr is a paid mutator transaction binding the contract method 0x7cb89783.
//
// Solidity: function take_usr(address user, address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultTransactorSession) TakeUsr(user common.Address, coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.TakeUsr(&_Enrockvault.TransactOpts, user, coin, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Enrockvault *EnrockvaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Enrockvault *EnrockvaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.TransferOwnership(&_Enrockvault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Enrockvault *EnrockvaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.TransferOwnership(&_Enrockvault.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Enrockvault *EnrockvaultTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Enrockvault *EnrockvaultSession) Unlock() (*types.Transaction, error) {
	return _Enrockvault.Contract.Unlock(&_Enrockvault.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Enrockvault *EnrockvaultTransactorSession) Unlock() (*types.Transaction, error) {
	return _Enrockvault.Contract.Unlock(&_Enrockvault.TransactOpts)
}

// UnpauseSc is a paid mutator transaction binding the contract method 0xdfc877d3.
//
// Solidity: function unpauseSc() returns()
func (_Enrockvault *EnrockvaultTransactor) UnpauseSc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "unpauseSc")
}

// UnpauseSc is a paid mutator transaction binding the contract method 0xdfc877d3.
//
// Solidity: function unpauseSc() returns()
func (_Enrockvault *EnrockvaultSession) UnpauseSc() (*types.Transaction, error) {
	return _Enrockvault.Contract.UnpauseSc(&_Enrockvault.TransactOpts)
}

// UnpauseSc is a paid mutator transaction binding the contract method 0xdfc877d3.
//
// Solidity: function unpauseSc() returns()
func (_Enrockvault *EnrockvaultTransactorSession) UnpauseSc() (*types.Transaction, error) {
	return _Enrockvault.Contract.UnpauseSc(&_Enrockvault.TransactOpts)
}

// UpBankVault is a paid mutator transaction binding the contract method 0x4b689736.
//
// Solidity: function up_bank_vault(address[] coins, uint256[] balances) returns()
func (_Enrockvault *EnrockvaultTransactor) UpBankVault(opts *bind.TransactOpts, coins []common.Address, balances []*big.Int) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "up_bank_vault", coins, balances)
}

// UpBankVault is a paid mutator transaction binding the contract method 0x4b689736.
//
// Solidity: function up_bank_vault(address[] coins, uint256[] balances) returns()
func (_Enrockvault *EnrockvaultSession) UpBankVault(coins []common.Address, balances []*big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.UpBankVault(&_Enrockvault.TransactOpts, coins, balances)
}

// UpBankVault is a paid mutator transaction binding the contract method 0x4b689736.
//
// Solidity: function up_bank_vault(address[] coins, uint256[] balances) returns()
func (_Enrockvault *EnrockvaultTransactorSession) UpBankVault(coins []common.Address, balances []*big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.UpBankVault(&_Enrockvault.TransactOpts, coins, balances)
}

// UpSingle is a paid mutator transaction binding the contract method 0xf573870d.
//
// Solidity: function up_single(address user, address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultTransactor) UpSingle(opts *bind.TransactOpts, user common.Address, coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "up_single", user, coin, amount)
}

// UpSingle is a paid mutator transaction binding the contract method 0xf573870d.
//
// Solidity: function up_single(address user, address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultSession) UpSingle(user common.Address, coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.UpSingle(&_Enrockvault.TransactOpts, user, coin, amount)
}

// UpSingle is a paid mutator transaction binding the contract method 0xf573870d.
//
// Solidity: function up_single(address user, address coin, uint256 amount) returns()
func (_Enrockvault *EnrockvaultTransactorSession) UpSingle(user common.Address, coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.UpSingle(&_Enrockvault.TransactOpts, user, coin, amount)
}

// UpUsrBatch is a paid mutator transaction binding the contract method 0x701ba7c3.
//
// Solidity: function up_usr_batch(address[] users, address[] coins, uint256[] amounts) returns()
func (_Enrockvault *EnrockvaultTransactor) UpUsrBatch(opts *bind.TransactOpts, users []common.Address, coins []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "up_usr_batch", users, coins, amounts)
}

// UpUsrBatch is a paid mutator transaction binding the contract method 0x701ba7c3.
//
// Solidity: function up_usr_batch(address[] users, address[] coins, uint256[] amounts) returns()
func (_Enrockvault *EnrockvaultSession) UpUsrBatch(users []common.Address, coins []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.UpUsrBatch(&_Enrockvault.TransactOpts, users, coins, amounts)
}

// UpUsrBatch is a paid mutator transaction binding the contract method 0x701ba7c3.
//
// Solidity: function up_usr_batch(address[] users, address[] coins, uint256[] amounts) returns()
func (_Enrockvault *EnrockvaultTransactorSession) UpUsrBatch(users []common.Address, coins []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Enrockvault.Contract.UpUsrBatch(&_Enrockvault.TransactOpts, users, coins, amounts)
}

// UpdateProvider is a paid mutator transaction binding the contract method 0xef70768a.
//
// Solidity: function updateProvider(address k) returns()
func (_Enrockvault *EnrockvaultTransactor) UpdateProvider(opts *bind.TransactOpts, k common.Address) (*types.Transaction, error) {
	return _Enrockvault.contract.Transact(opts, "updateProvider", k)
}

// UpdateProvider is a paid mutator transaction binding the contract method 0xef70768a.
//
// Solidity: function updateProvider(address k) returns()
func (_Enrockvault *EnrockvaultSession) UpdateProvider(k common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.UpdateProvider(&_Enrockvault.TransactOpts, k)
}

// UpdateProvider is a paid mutator transaction binding the contract method 0xef70768a.
//
// Solidity: function updateProvider(address k) returns()
func (_Enrockvault *EnrockvaultTransactorSession) UpdateProvider(k common.Address) (*types.Transaction, error) {
	return _Enrockvault.Contract.UpdateProvider(&_Enrockvault.TransactOpts, k)
}

// EnrockvaultDonateDepositIterator is returned from FilterDonateDeposit and is used to iterate over the raw logs and unpacked data for DonateDeposit events raised by the Enrockvault contract.
type EnrockvaultDonateDepositIterator struct {
	Event *EnrockvaultDonateDeposit // Event containing the contract specifics and raw log

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
func (it *EnrockvaultDonateDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultDonateDeposit)
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
		it.Event = new(EnrockvaultDonateDeposit)
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
func (it *EnrockvaultDonateDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultDonateDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultDonateDeposit represents a DonateDeposit event raised by the Enrockvault contract.
type EnrockvaultDonateDeposit struct {
	Coin   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonateDeposit is a free log retrieval operation binding the contract event 0xb0e3012a8e81e2cdf7032b5a43ac70d4b5a9a79c4170d25a3343c8c361d2bcb2.
//
// Solidity: event DonateDeposit(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) FilterDonateDeposit(opts *bind.FilterOpts) (*EnrockvaultDonateDepositIterator, error) {

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "DonateDeposit")
	if err != nil {
		return nil, err
	}
	return &EnrockvaultDonateDepositIterator{contract: _Enrockvault.contract, event: "DonateDeposit", logs: logs, sub: sub}, nil
}

// WatchDonateDeposit is a free log subscription operation binding the contract event 0xb0e3012a8e81e2cdf7032b5a43ac70d4b5a9a79c4170d25a3343c8c361d2bcb2.
//
// Solidity: event DonateDeposit(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) WatchDonateDeposit(opts *bind.WatchOpts, sink chan<- *EnrockvaultDonateDeposit) (event.Subscription, error) {

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "DonateDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultDonateDeposit)
				if err := _Enrockvault.contract.UnpackLog(event, "DonateDeposit", log); err != nil {
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

// ParseDonateDeposit is a log parse operation binding the contract event 0xb0e3012a8e81e2cdf7032b5a43ac70d4b5a9a79c4170d25a3343c8c361d2bcb2.
//
// Solidity: event DonateDeposit(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) ParseDonateDeposit(log types.Log) (*EnrockvaultDonateDeposit, error) {
	event := new(EnrockvaultDonateDeposit)
	if err := _Enrockvault.contract.UnpackLog(event, "DonateDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultGameVaultIterator is returned from FilterGameVault and is used to iterate over the raw logs and unpacked data for GameVault events raised by the Enrockvault contract.
type EnrockvaultGameVaultIterator struct {
	Event *EnrockvaultGameVault // Event containing the contract specifics and raw log

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
func (it *EnrockvaultGameVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultGameVault)
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
		it.Event = new(EnrockvaultGameVault)
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
func (it *EnrockvaultGameVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultGameVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultGameVault represents a GameVault event raised by the Enrockvault contract.
type EnrockvaultGameVault struct {
	H   string
	Dot *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGameVault is a free log retrieval operation binding the contract event 0x911dff754ebcc7c039c915d0447cf27c9c977bed617ffb212e131da35cbd4762.
//
// Solidity: event GameVault(string h, uint256 dot)
func (_Enrockvault *EnrockvaultFilterer) FilterGameVault(opts *bind.FilterOpts) (*EnrockvaultGameVaultIterator, error) {

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "GameVault")
	if err != nil {
		return nil, err
	}
	return &EnrockvaultGameVaultIterator{contract: _Enrockvault.contract, event: "GameVault", logs: logs, sub: sub}, nil
}

// WatchGameVault is a free log subscription operation binding the contract event 0x911dff754ebcc7c039c915d0447cf27c9c977bed617ffb212e131da35cbd4762.
//
// Solidity: event GameVault(string h, uint256 dot)
func (_Enrockvault *EnrockvaultFilterer) WatchGameVault(opts *bind.WatchOpts, sink chan<- *EnrockvaultGameVault) (event.Subscription, error) {

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "GameVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultGameVault)
				if err := _Enrockvault.contract.UnpackLog(event, "GameVault", log); err != nil {
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

// ParseGameVault is a log parse operation binding the contract event 0x911dff754ebcc7c039c915d0447cf27c9c977bed617ffb212e131da35cbd4762.
//
// Solidity: event GameVault(string h, uint256 dot)
func (_Enrockvault *EnrockvaultFilterer) ParseGameVault(log types.Log) (*EnrockvaultGameVault, error) {
	event := new(EnrockvaultGameVault)
	if err := _Enrockvault.contract.UnpackLog(event, "GameVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Enrockvault contract.
type EnrockvaultOwnershipTransferredIterator struct {
	Event *EnrockvaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EnrockvaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultOwnershipTransferred)
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
		it.Event = new(EnrockvaultOwnershipTransferred)
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
func (it *EnrockvaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultOwnershipTransferred represents a OwnershipTransferred event raised by the Enrockvault contract.
type EnrockvaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Enrockvault *EnrockvaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EnrockvaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EnrockvaultOwnershipTransferredIterator{contract: _Enrockvault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Enrockvault *EnrockvaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EnrockvaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultOwnershipTransferred)
				if err := _Enrockvault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Enrockvault *EnrockvaultFilterer) ParseOwnershipTransferred(log types.Log) (*EnrockvaultOwnershipTransferred, error) {
	event := new(EnrockvaultOwnershipTransferred)
	if err := _Enrockvault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Enrockvault contract.
type EnrockvaultRoleGrantedIterator struct {
	Event *EnrockvaultRoleGranted // Event containing the contract specifics and raw log

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
func (it *EnrockvaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultRoleGranted)
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
		it.Event = new(EnrockvaultRoleGranted)
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
func (it *EnrockvaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultRoleGranted represents a RoleGranted event raised by the Enrockvault contract.
type EnrockvaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Enrockvault *EnrockvaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EnrockvaultRoleGrantedIterator, error) {

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

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EnrockvaultRoleGrantedIterator{contract: _Enrockvault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Enrockvault *EnrockvaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EnrockvaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultRoleGranted)
				if err := _Enrockvault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Enrockvault *EnrockvaultFilterer) ParseRoleGranted(log types.Log) (*EnrockvaultRoleGranted, error) {
	event := new(EnrockvaultRoleGranted)
	if err := _Enrockvault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Enrockvault contract.
type EnrockvaultRoleRevokedIterator struct {
	Event *EnrockvaultRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EnrockvaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultRoleRevoked)
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
		it.Event = new(EnrockvaultRoleRevoked)
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
func (it *EnrockvaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultRoleRevoked represents a RoleRevoked event raised by the Enrockvault contract.
type EnrockvaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Enrockvault *EnrockvaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EnrockvaultRoleRevokedIterator, error) {

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

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EnrockvaultRoleRevokedIterator{contract: _Enrockvault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Enrockvault *EnrockvaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EnrockvaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultRoleRevoked)
				if err := _Enrockvault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Enrockvault *EnrockvaultFilterer) ParseRoleRevoked(log types.Log) (*EnrockvaultRoleRevoked, error) {
	event := new(EnrockvaultRoleRevoked)
	if err := _Enrockvault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultUpdatedBankIterator is returned from FilterUpdatedBank and is used to iterate over the raw logs and unpacked data for UpdatedBank events raised by the Enrockvault contract.
type EnrockvaultUpdatedBankIterator struct {
	Event *EnrockvaultUpdatedBank // Event containing the contract specifics and raw log

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
func (it *EnrockvaultUpdatedBankIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultUpdatedBank)
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
		it.Event = new(EnrockvaultUpdatedBank)
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
func (it *EnrockvaultUpdatedBankIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultUpdatedBankIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultUpdatedBank represents a UpdatedBank event raised by the Enrockvault contract.
type EnrockvaultUpdatedBank struct {
	Coin   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBank is a free log retrieval operation binding the contract event 0x6766e7039d402e314ce5c5063014a084c485e79d0804bc3b35d96c528adf9040.
//
// Solidity: event UpdatedBank(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) FilterUpdatedBank(opts *bind.FilterOpts) (*EnrockvaultUpdatedBankIterator, error) {

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "UpdatedBank")
	if err != nil {
		return nil, err
	}
	return &EnrockvaultUpdatedBankIterator{contract: _Enrockvault.contract, event: "UpdatedBank", logs: logs, sub: sub}, nil
}

// WatchUpdatedBank is a free log subscription operation binding the contract event 0x6766e7039d402e314ce5c5063014a084c485e79d0804bc3b35d96c528adf9040.
//
// Solidity: event UpdatedBank(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) WatchUpdatedBank(opts *bind.WatchOpts, sink chan<- *EnrockvaultUpdatedBank) (event.Subscription, error) {

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "UpdatedBank")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultUpdatedBank)
				if err := _Enrockvault.contract.UnpackLog(event, "UpdatedBank", log); err != nil {
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

// ParseUpdatedBank is a log parse operation binding the contract event 0x6766e7039d402e314ce5c5063014a084c485e79d0804bc3b35d96c528adf9040.
//
// Solidity: event UpdatedBank(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) ParseUpdatedBank(log types.Log) (*EnrockvaultUpdatedBank, error) {
	event := new(EnrockvaultUpdatedBank)
	if err := _Enrockvault.contract.UnpackLog(event, "UpdatedBank", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultUpdatedWalletIterator is returned from FilterUpdatedWallet and is used to iterate over the raw logs and unpacked data for UpdatedWallet events raised by the Enrockvault contract.
type EnrockvaultUpdatedWalletIterator struct {
	Event *EnrockvaultUpdatedWallet // Event containing the contract specifics and raw log

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
func (it *EnrockvaultUpdatedWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultUpdatedWallet)
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
		it.Event = new(EnrockvaultUpdatedWallet)
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
func (it *EnrockvaultUpdatedWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultUpdatedWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultUpdatedWallet represents a UpdatedWallet event raised by the Enrockvault contract.
type EnrockvaultUpdatedWallet struct {
	Coin       common.Address
	Beneficier common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdatedWallet is a free log retrieval operation binding the contract event 0x27191f4a186003124963faf556d1ecdf18a0f8d3a2b59ad9e1ff7c381f47b7ac.
//
// Solidity: event UpdatedWallet(address coin, address beneficier)
func (_Enrockvault *EnrockvaultFilterer) FilterUpdatedWallet(opts *bind.FilterOpts) (*EnrockvaultUpdatedWalletIterator, error) {

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "UpdatedWallet")
	if err != nil {
		return nil, err
	}
	return &EnrockvaultUpdatedWalletIterator{contract: _Enrockvault.contract, event: "UpdatedWallet", logs: logs, sub: sub}, nil
}

// WatchUpdatedWallet is a free log subscription operation binding the contract event 0x27191f4a186003124963faf556d1ecdf18a0f8d3a2b59ad9e1ff7c381f47b7ac.
//
// Solidity: event UpdatedWallet(address coin, address beneficier)
func (_Enrockvault *EnrockvaultFilterer) WatchUpdatedWallet(opts *bind.WatchOpts, sink chan<- *EnrockvaultUpdatedWallet) (event.Subscription, error) {

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "UpdatedWallet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultUpdatedWallet)
				if err := _Enrockvault.contract.UnpackLog(event, "UpdatedWallet", log); err != nil {
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

// ParseUpdatedWallet is a log parse operation binding the contract event 0x27191f4a186003124963faf556d1ecdf18a0f8d3a2b59ad9e1ff7c381f47b7ac.
//
// Solidity: event UpdatedWallet(address coin, address beneficier)
func (_Enrockvault *EnrockvaultFilterer) ParseUpdatedWallet(log types.Log) (*EnrockvaultUpdatedWallet, error) {
	event := new(EnrockvaultUpdatedWallet)
	if err := _Enrockvault.contract.UnpackLog(event, "UpdatedWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultUsrDepositIterator is returned from FilterUsrDeposit and is used to iterate over the raw logs and unpacked data for UsrDeposit events raised by the Enrockvault contract.
type EnrockvaultUsrDepositIterator struct {
	Event *EnrockvaultUsrDeposit // Event containing the contract specifics and raw log

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
func (it *EnrockvaultUsrDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultUsrDeposit)
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
		it.Event = new(EnrockvaultUsrDeposit)
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
func (it *EnrockvaultUsrDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultUsrDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultUsrDeposit represents a UsrDeposit event raised by the Enrockvault contract.
type EnrockvaultUsrDeposit struct {
	Coin   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUsrDeposit is a free log retrieval operation binding the contract event 0x089f4869c98e4e8a6236b0a3ea67d4a99eed42aafc725bea496d8c1bd37c4d89.
//
// Solidity: event UsrDeposit(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) FilterUsrDeposit(opts *bind.FilterOpts) (*EnrockvaultUsrDepositIterator, error) {

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "UsrDeposit")
	if err != nil {
		return nil, err
	}
	return &EnrockvaultUsrDepositIterator{contract: _Enrockvault.contract, event: "UsrDeposit", logs: logs, sub: sub}, nil
}

// WatchUsrDeposit is a free log subscription operation binding the contract event 0x089f4869c98e4e8a6236b0a3ea67d4a99eed42aafc725bea496d8c1bd37c4d89.
//
// Solidity: event UsrDeposit(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) WatchUsrDeposit(opts *bind.WatchOpts, sink chan<- *EnrockvaultUsrDeposit) (event.Subscription, error) {

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "UsrDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultUsrDeposit)
				if err := _Enrockvault.contract.UnpackLog(event, "UsrDeposit", log); err != nil {
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

// ParseUsrDeposit is a log parse operation binding the contract event 0x089f4869c98e4e8a6236b0a3ea67d4a99eed42aafc725bea496d8c1bd37c4d89.
//
// Solidity: event UsrDeposit(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) ParseUsrDeposit(log types.Log) (*EnrockvaultUsrDeposit, error) {
	event := new(EnrockvaultUsrDeposit)
	if err := _Enrockvault.contract.UnpackLog(event, "UsrDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultUsrWithdrawIterator is returned from FilterUsrWithdraw and is used to iterate over the raw logs and unpacked data for UsrWithdraw events raised by the Enrockvault contract.
type EnrockvaultUsrWithdrawIterator struct {
	Event *EnrockvaultUsrWithdraw // Event containing the contract specifics and raw log

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
func (it *EnrockvaultUsrWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultUsrWithdraw)
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
		it.Event = new(EnrockvaultUsrWithdraw)
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
func (it *EnrockvaultUsrWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultUsrWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultUsrWithdraw represents a UsrWithdraw event raised by the Enrockvault contract.
type EnrockvaultUsrWithdraw struct {
	Coin   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUsrWithdraw is a free log retrieval operation binding the contract event 0x103cedaa5cbb0ad45069be12c8d6dd32ddad8cf11b20cf615fc7b0e245fca752.
//
// Solidity: event UsrWithdraw(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) FilterUsrWithdraw(opts *bind.FilterOpts) (*EnrockvaultUsrWithdrawIterator, error) {

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "UsrWithdraw")
	if err != nil {
		return nil, err
	}
	return &EnrockvaultUsrWithdrawIterator{contract: _Enrockvault.contract, event: "UsrWithdraw", logs: logs, sub: sub}, nil
}

// WatchUsrWithdraw is a free log subscription operation binding the contract event 0x103cedaa5cbb0ad45069be12c8d6dd32ddad8cf11b20cf615fc7b0e245fca752.
//
// Solidity: event UsrWithdraw(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) WatchUsrWithdraw(opts *bind.WatchOpts, sink chan<- *EnrockvaultUsrWithdraw) (event.Subscription, error) {

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "UsrWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultUsrWithdraw)
				if err := _Enrockvault.contract.UnpackLog(event, "UsrWithdraw", log); err != nil {
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

// ParseUsrWithdraw is a log parse operation binding the contract event 0x103cedaa5cbb0ad45069be12c8d6dd32ddad8cf11b20cf615fc7b0e245fca752.
//
// Solidity: event UsrWithdraw(address coin, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) ParseUsrWithdraw(log types.Log) (*EnrockvaultUsrWithdraw, error) {
	event := new(EnrockvaultUsrWithdraw)
	if err := _Enrockvault.contract.UnpackLog(event, "UsrWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultWithdrawalGovIterator is returned from FilterWithdrawalGov and is used to iterate over the raw logs and unpacked data for WithdrawalGov events raised by the Enrockvault contract.
type EnrockvaultWithdrawalGovIterator struct {
	Event *EnrockvaultWithdrawalGov // Event containing the contract specifics and raw log

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
func (it *EnrockvaultWithdrawalGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultWithdrawalGov)
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
		it.Event = new(EnrockvaultWithdrawalGov)
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
func (it *EnrockvaultWithdrawalGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultWithdrawalGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultWithdrawalGov represents a WithdrawalGov event raised by the Enrockvault contract.
type EnrockvaultWithdrawalGov struct {
	Coin       common.Address
	Beneficier common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalGov is a free log retrieval operation binding the contract event 0x89abea36d5c6738d65ee6b77714607188187e39939693f079ad2601e66ddbb0d.
//
// Solidity: event WithdrawalGov(address coin, address beneficier, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) FilterWithdrawalGov(opts *bind.FilterOpts) (*EnrockvaultWithdrawalGovIterator, error) {

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "WithdrawalGov")
	if err != nil {
		return nil, err
	}
	return &EnrockvaultWithdrawalGovIterator{contract: _Enrockvault.contract, event: "WithdrawalGov", logs: logs, sub: sub}, nil
}

// WatchWithdrawalGov is a free log subscription operation binding the contract event 0x89abea36d5c6738d65ee6b77714607188187e39939693f079ad2601e66ddbb0d.
//
// Solidity: event WithdrawalGov(address coin, address beneficier, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) WatchWithdrawalGov(opts *bind.WatchOpts, sink chan<- *EnrockvaultWithdrawalGov) (event.Subscription, error) {

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "WithdrawalGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultWithdrawalGov)
				if err := _Enrockvault.contract.UnpackLog(event, "WithdrawalGov", log); err != nil {
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

// ParseWithdrawalGov is a log parse operation binding the contract event 0x89abea36d5c6738d65ee6b77714607188187e39939693f079ad2601e66ddbb0d.
//
// Solidity: event WithdrawalGov(address coin, address beneficier, uint256 amount)
func (_Enrockvault *EnrockvaultFilterer) ParseWithdrawalGov(log types.Log) (*EnrockvaultWithdrawalGov, error) {
	event := new(EnrockvaultWithdrawalGov)
	if err := _Enrockvault.contract.UnpackLog(event, "WithdrawalGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultContractPausedIterator is returned from FilterContractPaused and is used to iterate over the raw logs and unpacked data for ContractPaused events raised by the Enrockvault contract.
type EnrockvaultContractPausedIterator struct {
	Event *EnrockvaultContractPaused // Event containing the contract specifics and raw log

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
func (it *EnrockvaultContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultContractPaused)
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
		it.Event = new(EnrockvaultContractPaused)
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
func (it *EnrockvaultContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultContractPaused represents a ContractPaused event raised by the Enrockvault contract.
type EnrockvaultContractPaused struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractPaused is a free log retrieval operation binding the contract event 0xf378156ac648519c9cd03e0977fbb5b5c0552e1c0cd584389461f5349ba18657.
//
// Solidity: event contractPaused(bool value)
func (_Enrockvault *EnrockvaultFilterer) FilterContractPaused(opts *bind.FilterOpts) (*EnrockvaultContractPausedIterator, error) {

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "contractPaused")
	if err != nil {
		return nil, err
	}
	return &EnrockvaultContractPausedIterator{contract: _Enrockvault.contract, event: "contractPaused", logs: logs, sub: sub}, nil
}

// WatchContractPaused is a free log subscription operation binding the contract event 0xf378156ac648519c9cd03e0977fbb5b5c0552e1c0cd584389461f5349ba18657.
//
// Solidity: event contractPaused(bool value)
func (_Enrockvault *EnrockvaultFilterer) WatchContractPaused(opts *bind.WatchOpts, sink chan<- *EnrockvaultContractPaused) (event.Subscription, error) {

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "contractPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultContractPaused)
				if err := _Enrockvault.contract.UnpackLog(event, "contractPaused", log); err != nil {
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

// ParseContractPaused is a log parse operation binding the contract event 0xf378156ac648519c9cd03e0977fbb5b5c0552e1c0cd584389461f5349ba18657.
//
// Solidity: event contractPaused(bool value)
func (_Enrockvault *EnrockvaultFilterer) ParseContractPaused(log types.Log) (*EnrockvaultContractPaused, error) {
	event := new(EnrockvaultContractPaused)
	if err := _Enrockvault.contract.UnpackLog(event, "contractPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnrockvaultTraillockIterator is returned from FilterTraillock and is used to iterate over the raw logs and unpacked data for Traillock events raised by the Enrockvault contract.
type EnrockvaultTraillockIterator struct {
	Event *EnrockvaultTraillock // Event containing the contract specifics and raw log

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
func (it *EnrockvaultTraillockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnrockvaultTraillock)
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
		it.Event = new(EnrockvaultTraillock)
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
func (it *EnrockvaultTraillockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnrockvaultTraillockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnrockvaultTraillock represents a Traillock event raised by the Enrockvault contract.
type EnrockvaultTraillock struct {
	Value uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTraillock is a free log retrieval operation binding the contract event 0xd02d7ece1b390124a8d8d40fc59fc1d6584a261d9a503b71166ec69b3eef00b3.
//
// Solidity: event traillock(uint8 value)
func (_Enrockvault *EnrockvaultFilterer) FilterTraillock(opts *bind.FilterOpts) (*EnrockvaultTraillockIterator, error) {

	logs, sub, err := _Enrockvault.contract.FilterLogs(opts, "traillock")
	if err != nil {
		return nil, err
	}
	return &EnrockvaultTraillockIterator{contract: _Enrockvault.contract, event: "traillock", logs: logs, sub: sub}, nil
}

// WatchTraillock is a free log subscription operation binding the contract event 0xd02d7ece1b390124a8d8d40fc59fc1d6584a261d9a503b71166ec69b3eef00b3.
//
// Solidity: event traillock(uint8 value)
func (_Enrockvault *EnrockvaultFilterer) WatchTraillock(opts *bind.WatchOpts, sink chan<- *EnrockvaultTraillock) (event.Subscription, error) {

	logs, sub, err := _Enrockvault.contract.WatchLogs(opts, "traillock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnrockvaultTraillock)
				if err := _Enrockvault.contract.UnpackLog(event, "traillock", log); err != nil {
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

// ParseTraillock is a log parse operation binding the contract event 0xd02d7ece1b390124a8d8d40fc59fc1d6584a261d9a503b71166ec69b3eef00b3.
//
// Solidity: event traillock(uint8 value)
func (_Enrockvault *EnrockvaultFilterer) ParseTraillock(log types.Log) (*EnrockvaultTraillock, error) {
	event := new(EnrockvaultTraillock)
	if err := _Enrockvault.contract.UnpackLog(event, "traillock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

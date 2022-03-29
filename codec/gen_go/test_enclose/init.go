// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test_enclose

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

// TestEncloseMetaData contains all meta data concerning the TestEnclose contract.
var TestEncloseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"UsrDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"contractPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"value\",\"type\":\"uint8\"}],\"name\":\"traillock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"coin_user_vault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"coin_user_vault2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit_erc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelistAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"}],\"name\":\"sig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseSc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistAdmins\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TestEncloseABI is the input ABI used to generate the binding from.
// Deprecated: Use TestEncloseMetaData.ABI instead.
var TestEncloseABI = TestEncloseMetaData.ABI

// TestEnclose is an auto generated Go binding around an Ethereum contract.
type TestEnclose struct {
	TestEncloseCaller     // Read-only binding to the contract
	TestEncloseTransactor // Write-only binding to the contract
	TestEncloseFilterer   // Log filterer for contract events
}

// TestEncloseCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestEncloseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEncloseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestEncloseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEncloseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestEncloseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEncloseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestEncloseSession struct {
	Contract     *TestEnclose      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestEncloseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestEncloseCallerSession struct {
	Contract *TestEncloseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestEncloseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestEncloseTransactorSession struct {
	Contract     *TestEncloseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestEncloseRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestEncloseRaw struct {
	Contract *TestEnclose // Generic contract binding to access the raw methods on
}

// TestEncloseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestEncloseCallerRaw struct {
	Contract *TestEncloseCaller // Generic read-only contract binding to access the raw methods on
}

// TestEncloseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestEncloseTransactorRaw struct {
	Contract *TestEncloseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestEnclose creates a new instance of TestEnclose, bound to a specific deployed contract.
func NewTestEnclose(address common.Address, backend bind.ContractBackend) (*TestEnclose, error) {
	contract, err := bindTestEnclose(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestEnclose{TestEncloseCaller: TestEncloseCaller{contract: contract}, TestEncloseTransactor: TestEncloseTransactor{contract: contract}, TestEncloseFilterer: TestEncloseFilterer{contract: contract}}, nil
}

// NewTestEncloseCaller creates a new read-only instance of TestEnclose, bound to a specific deployed contract.
func NewTestEncloseCaller(address common.Address, caller bind.ContractCaller) (*TestEncloseCaller, error) {
	contract, err := bindTestEnclose(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestEncloseCaller{contract: contract}, nil
}

// NewTestEncloseTransactor creates a new write-only instance of TestEnclose, bound to a specific deployed contract.
func NewTestEncloseTransactor(address common.Address, transactor bind.ContractTransactor) (*TestEncloseTransactor, error) {
	contract, err := bindTestEnclose(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestEncloseTransactor{contract: contract}, nil
}

// NewTestEncloseFilterer creates a new log filterer instance of TestEnclose, bound to a specific deployed contract.
func NewTestEncloseFilterer(address common.Address, filterer bind.ContractFilterer) (*TestEncloseFilterer, error) {
	contract, err := bindTestEnclose(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestEncloseFilterer{contract: contract}, nil
}

// bindTestEnclose binds a generic wrapper to an already deployed contract.
func bindTestEnclose(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestEncloseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEnclose *TestEncloseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestEnclose.Contract.TestEncloseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEnclose *TestEncloseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEnclose.Contract.TestEncloseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEnclose *TestEncloseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEnclose.Contract.TestEncloseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEnclose *TestEncloseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestEnclose.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEnclose *TestEncloseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEnclose.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEnclose *TestEncloseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEnclose.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TestEnclose *TestEncloseCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TestEnclose *TestEncloseSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TestEnclose.Contract.DEFAULTADMINROLE(&_TestEnclose.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TestEnclose *TestEncloseCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TestEnclose.Contract.DEFAULTADMINROLE(&_TestEnclose.CallOpts)
}

// CoinUserVault is a free data retrieval call binding the contract method 0xa83b18ec.
//
// Solidity: function coin_user_vault(bytes32 ) view returns(uint256)
func (_TestEnclose *TestEncloseCaller) CoinUserVault(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "coin_user_vault", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinUserVault is a free data retrieval call binding the contract method 0xa83b18ec.
//
// Solidity: function coin_user_vault(bytes32 ) view returns(uint256)
func (_TestEnclose *TestEncloseSession) CoinUserVault(arg0 [32]byte) (*big.Int, error) {
	return _TestEnclose.Contract.CoinUserVault(&_TestEnclose.CallOpts, arg0)
}

// CoinUserVault is a free data retrieval call binding the contract method 0xa83b18ec.
//
// Solidity: function coin_user_vault(bytes32 ) view returns(uint256)
func (_TestEnclose *TestEncloseCallerSession) CoinUserVault(arg0 [32]byte) (*big.Int, error) {
	return _TestEnclose.Contract.CoinUserVault(&_TestEnclose.CallOpts, arg0)
}

// CoinUserVault2 is a free data retrieval call binding the contract method 0x171ea1f8.
//
// Solidity: function coin_user_vault2(address , address ) view returns(uint256)
func (_TestEnclose *TestEncloseCaller) CoinUserVault2(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "coin_user_vault2", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinUserVault2 is a free data retrieval call binding the contract method 0x171ea1f8.
//
// Solidity: function coin_user_vault2(address , address ) view returns(uint256)
func (_TestEnclose *TestEncloseSession) CoinUserVault2(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TestEnclose.Contract.CoinUserVault2(&_TestEnclose.CallOpts, arg0, arg1)
}

// CoinUserVault2 is a free data retrieval call binding the contract method 0x171ea1f8.
//
// Solidity: function coin_user_vault2(address , address ) view returns(uint256)
func (_TestEnclose *TestEncloseCallerSession) CoinUserVault2(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TestEnclose.Contract.CoinUserVault2(&_TestEnclose.CallOpts, arg0, arg1)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TestEnclose *TestEncloseCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TestEnclose *TestEncloseSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TestEnclose.Contract.GetRoleAdmin(&_TestEnclose.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TestEnclose *TestEncloseCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TestEnclose.Contract.GetRoleAdmin(&_TestEnclose.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TestEnclose *TestEncloseCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TestEnclose *TestEncloseSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TestEnclose.Contract.GetRoleMember(&_TestEnclose.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TestEnclose *TestEncloseCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TestEnclose.Contract.GetRoleMember(&_TestEnclose.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TestEnclose *TestEncloseCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TestEnclose *TestEncloseSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TestEnclose.Contract.GetRoleMemberCount(&_TestEnclose.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TestEnclose *TestEncloseCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TestEnclose.Contract.GetRoleMemberCount(&_TestEnclose.CallOpts, role)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(bytes32)
func (_TestEnclose *TestEncloseCaller) Governor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(bytes32)
func (_TestEnclose *TestEncloseSession) Governor() ([32]byte, error) {
	return _TestEnclose.Contract.Governor(&_TestEnclose.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(bytes32)
func (_TestEnclose *TestEncloseCallerSession) Governor() ([32]byte, error) {
	return _TestEnclose.Contract.Governor(&_TestEnclose.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TestEnclose *TestEncloseCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TestEnclose *TestEncloseSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TestEnclose.Contract.HasRole(&_TestEnclose.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TestEnclose *TestEncloseCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TestEnclose.Contract.HasRole(&_TestEnclose.CallOpts, role, account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address account) view returns(bool)
func (_TestEnclose *TestEncloseCaller) IsGovernor(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "isGovernor", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address account) view returns(bool)
func (_TestEnclose *TestEncloseSession) IsGovernor(account common.Address) (bool, error) {
	return _TestEnclose.Contract.IsGovernor(&_TestEnclose.CallOpts, account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address account) view returns(bool)
func (_TestEnclose *TestEncloseCallerSession) IsGovernor(account common.Address) (bool, error) {
	return _TestEnclose.Contract.IsGovernor(&_TestEnclose.CallOpts, account)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_TestEnclose *TestEncloseCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_TestEnclose *TestEncloseSession) IsLocked() (bool, error) {
	return _TestEnclose.Contract.IsLocked(&_TestEnclose.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_TestEnclose *TestEncloseCallerSession) IsLocked() (bool, error) {
	return _TestEnclose.Contract.IsLocked(&_TestEnclose.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TestEnclose *TestEncloseCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TestEnclose *TestEncloseSession) IsOwner() (bool, error) {
	return _TestEnclose.Contract.IsOwner(&_TestEnclose.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TestEnclose *TestEncloseCallerSession) IsOwner() (bool, error) {
	return _TestEnclose.Contract.IsOwner(&_TestEnclose.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_TestEnclose *TestEncloseCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_TestEnclose *TestEncloseSession) IsPaused() (bool, error) {
	return _TestEnclose.Contract.IsPaused(&_TestEnclose.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_TestEnclose *TestEncloseCallerSession) IsPaused() (bool, error) {
	return _TestEnclose.Contract.IsPaused(&_TestEnclose.CallOpts)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_TestEnclose *TestEncloseCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "isWhitelistAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_TestEnclose *TestEncloseSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _TestEnclose.Contract.IsWhitelistAdmin(&_TestEnclose.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_TestEnclose *TestEncloseCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _TestEnclose.Contract.IsWhitelistAdmin(&_TestEnclose.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestEnclose *TestEncloseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestEnclose *TestEncloseSession) Owner() (common.Address, error) {
	return _TestEnclose.Contract.Owner(&_TestEnclose.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestEnclose *TestEncloseCallerSession) Owner() (common.Address, error) {
	return _TestEnclose.Contract.Owner(&_TestEnclose.CallOpts)
}

// Sig is a free data retrieval call binding the contract method 0x7ff7a6e5.
//
// Solidity: function sig(address user, address coin) pure returns(bytes32)
func (_TestEnclose *TestEncloseCaller) Sig(opts *bind.CallOpts, user common.Address, coin common.Address) ([32]byte, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "sig", user, coin)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sig is a free data retrieval call binding the contract method 0x7ff7a6e5.
//
// Solidity: function sig(address user, address coin) pure returns(bytes32)
func (_TestEnclose *TestEncloseSession) Sig(user common.Address, coin common.Address) ([32]byte, error) {
	return _TestEnclose.Contract.Sig(&_TestEnclose.CallOpts, user, coin)
}

// Sig is a free data retrieval call binding the contract method 0x7ff7a6e5.
//
// Solidity: function sig(address user, address coin) pure returns(bytes32)
func (_TestEnclose *TestEncloseCallerSession) Sig(user common.Address, coin common.Address) ([32]byte, error) {
	return _TestEnclose.Contract.Sig(&_TestEnclose.CallOpts, user, coin)
}

// WhitelistAdmins is a free data retrieval call binding the contract method 0x7ef581cc.
//
// Solidity: function whitelistAdmins() view returns(bytes32)
func (_TestEnclose *TestEncloseCaller) WhitelistAdmins(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestEnclose.contract.Call(opts, &out, "whitelistAdmins")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WhitelistAdmins is a free data retrieval call binding the contract method 0x7ef581cc.
//
// Solidity: function whitelistAdmins() view returns(bytes32)
func (_TestEnclose *TestEncloseSession) WhitelistAdmins() ([32]byte, error) {
	return _TestEnclose.Contract.WhitelistAdmins(&_TestEnclose.CallOpts)
}

// WhitelistAdmins is a free data retrieval call binding the contract method 0x7ef581cc.
//
// Solidity: function whitelistAdmins() view returns(bytes32)
func (_TestEnclose *TestEncloseCallerSession) WhitelistAdmins() ([32]byte, error) {
	return _TestEnclose.Contract.WhitelistAdmins(&_TestEnclose.CallOpts)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address account) returns()
func (_TestEnclose *TestEncloseTransactor) AddGovernor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "addGovernor", account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address account) returns()
func (_TestEnclose *TestEncloseSession) AddGovernor(account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.AddGovernor(&_TestEnclose.TransactOpts, account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address account) returns()
func (_TestEnclose *TestEncloseTransactorSession) AddGovernor(account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.AddGovernor(&_TestEnclose.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_TestEnclose *TestEncloseTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_TestEnclose *TestEncloseSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.AddWhitelistAdmin(&_TestEnclose.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_TestEnclose *TestEncloseTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.AddWhitelistAdmin(&_TestEnclose.TransactOpts, account)
}

// DepositErc20 is a paid mutator transaction binding the contract method 0xd16544f0.
//
// Solidity: function deposit_erc20(address coin, uint256 amount) returns()
func (_TestEnclose *TestEncloseTransactor) DepositErc20(opts *bind.TransactOpts, coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "deposit_erc20", coin, amount)
}

// DepositErc20 is a paid mutator transaction binding the contract method 0xd16544f0.
//
// Solidity: function deposit_erc20(address coin, uint256 amount) returns()
func (_TestEnclose *TestEncloseSession) DepositErc20(coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestEnclose.Contract.DepositErc20(&_TestEnclose.TransactOpts, coin, amount)
}

// DepositErc20 is a paid mutator transaction binding the contract method 0xd16544f0.
//
// Solidity: function deposit_erc20(address coin, uint256 amount) returns()
func (_TestEnclose *TestEncloseTransactorSession) DepositErc20(coin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestEnclose.Contract.DepositErc20(&_TestEnclose.TransactOpts, coin, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TestEnclose *TestEncloseTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TestEnclose *TestEncloseSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.GrantRole(&_TestEnclose.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TestEnclose *TestEncloseTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.GrantRole(&_TestEnclose.TransactOpts, role, account)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TestEnclose *TestEncloseTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TestEnclose *TestEncloseSession) Lock() (*types.Transaction, error) {
	return _TestEnclose.Contract.Lock(&_TestEnclose.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_TestEnclose *TestEncloseTransactorSession) Lock() (*types.Transaction, error) {
	return _TestEnclose.Contract.Lock(&_TestEnclose.TransactOpts)
}

// PauseSc is a paid mutator transaction binding the contract method 0x1975ebaf.
//
// Solidity: function pauseSc() returns()
func (_TestEnclose *TestEncloseTransactor) PauseSc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "pauseSc")
}

// PauseSc is a paid mutator transaction binding the contract method 0x1975ebaf.
//
// Solidity: function pauseSc() returns()
func (_TestEnclose *TestEncloseSession) PauseSc() (*types.Transaction, error) {
	return _TestEnclose.Contract.PauseSc(&_TestEnclose.TransactOpts)
}

// PauseSc is a paid mutator transaction binding the contract method 0x1975ebaf.
//
// Solidity: function pauseSc() returns()
func (_TestEnclose *TestEncloseTransactorSession) PauseSc() (*types.Transaction, error) {
	return _TestEnclose.Contract.PauseSc(&_TestEnclose.TransactOpts)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address account) returns()
func (_TestEnclose *TestEncloseTransactor) RemoveGovernor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "removeGovernor", account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address account) returns()
func (_TestEnclose *TestEncloseSession) RemoveGovernor(account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.RemoveGovernor(&_TestEnclose.TransactOpts, account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address account) returns()
func (_TestEnclose *TestEncloseTransactorSession) RemoveGovernor(account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.RemoveGovernor(&_TestEnclose.TransactOpts, account)
}

// RemoveWhitelistAdmin is a paid mutator transaction binding the contract method 0x6897e974.
//
// Solidity: function removeWhitelistAdmin(address account) returns()
func (_TestEnclose *TestEncloseTransactor) RemoveWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "removeWhitelistAdmin", account)
}

// RemoveWhitelistAdmin is a paid mutator transaction binding the contract method 0x6897e974.
//
// Solidity: function removeWhitelistAdmin(address account) returns()
func (_TestEnclose *TestEncloseSession) RemoveWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.RemoveWhitelistAdmin(&_TestEnclose.TransactOpts, account)
}

// RemoveWhitelistAdmin is a paid mutator transaction binding the contract method 0x6897e974.
//
// Solidity: function removeWhitelistAdmin(address account) returns()
func (_TestEnclose *TestEncloseTransactorSession) RemoveWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.RemoveWhitelistAdmin(&_TestEnclose.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestEnclose *TestEncloseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestEnclose *TestEncloseSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestEnclose.Contract.RenounceOwnership(&_TestEnclose.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestEnclose *TestEncloseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestEnclose.Contract.RenounceOwnership(&_TestEnclose.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TestEnclose *TestEncloseTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TestEnclose *TestEncloseSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.RenounceRole(&_TestEnclose.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TestEnclose *TestEncloseTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.RenounceRole(&_TestEnclose.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TestEnclose *TestEncloseTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TestEnclose *TestEncloseSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.RevokeRole(&_TestEnclose.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TestEnclose *TestEncloseTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.RevokeRole(&_TestEnclose.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestEnclose *TestEncloseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestEnclose *TestEncloseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.TransferOwnership(&_TestEnclose.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestEnclose *TestEncloseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestEnclose.Contract.TransferOwnership(&_TestEnclose.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TestEnclose *TestEncloseTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TestEnclose *TestEncloseSession) Unlock() (*types.Transaction, error) {
	return _TestEnclose.Contract.Unlock(&_TestEnclose.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_TestEnclose *TestEncloseTransactorSession) Unlock() (*types.Transaction, error) {
	return _TestEnclose.Contract.Unlock(&_TestEnclose.TransactOpts)
}

// UnpauseSc is a paid mutator transaction binding the contract method 0xdfc877d3.
//
// Solidity: function unpauseSc() returns()
func (_TestEnclose *TestEncloseTransactor) UnpauseSc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEnclose.contract.Transact(opts, "unpauseSc")
}

// UnpauseSc is a paid mutator transaction binding the contract method 0xdfc877d3.
//
// Solidity: function unpauseSc() returns()
func (_TestEnclose *TestEncloseSession) UnpauseSc() (*types.Transaction, error) {
	return _TestEnclose.Contract.UnpauseSc(&_TestEnclose.TransactOpts)
}

// UnpauseSc is a paid mutator transaction binding the contract method 0xdfc877d3.
//
// Solidity: function unpauseSc() returns()
func (_TestEnclose *TestEncloseTransactorSession) UnpauseSc() (*types.Transaction, error) {
	return _TestEnclose.Contract.UnpauseSc(&_TestEnclose.TransactOpts)
}

// TestEncloseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestEnclose contract.
type TestEncloseOwnershipTransferredIterator struct {
	Event *TestEncloseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestEncloseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEncloseOwnershipTransferred)
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
		it.Event = new(TestEncloseOwnershipTransferred)
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
func (it *TestEncloseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEncloseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEncloseOwnershipTransferred represents a OwnershipTransferred event raised by the TestEnclose contract.
type TestEncloseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestEnclose *TestEncloseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestEncloseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestEnclose.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestEncloseOwnershipTransferredIterator{contract: _TestEnclose.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestEnclose *TestEncloseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestEncloseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestEnclose.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEncloseOwnershipTransferred)
				if err := _TestEnclose.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TestEnclose *TestEncloseFilterer) ParseOwnershipTransferred(log types.Log) (*TestEncloseOwnershipTransferred, error) {
	event := new(TestEncloseOwnershipTransferred)
	if err := _TestEnclose.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestEncloseRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TestEnclose contract.
type TestEncloseRoleGrantedIterator struct {
	Event *TestEncloseRoleGranted // Event containing the contract specifics and raw log

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
func (it *TestEncloseRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEncloseRoleGranted)
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
		it.Event = new(TestEncloseRoleGranted)
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
func (it *TestEncloseRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEncloseRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEncloseRoleGranted represents a RoleGranted event raised by the TestEnclose contract.
type TestEncloseRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TestEnclose *TestEncloseFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TestEncloseRoleGrantedIterator, error) {

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

	logs, sub, err := _TestEnclose.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TestEncloseRoleGrantedIterator{contract: _TestEnclose.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TestEnclose *TestEncloseFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TestEncloseRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TestEnclose.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEncloseRoleGranted)
				if err := _TestEnclose.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TestEnclose *TestEncloseFilterer) ParseRoleGranted(log types.Log) (*TestEncloseRoleGranted, error) {
	event := new(TestEncloseRoleGranted)
	if err := _TestEnclose.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestEncloseRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TestEnclose contract.
type TestEncloseRoleRevokedIterator struct {
	Event *TestEncloseRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TestEncloseRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEncloseRoleRevoked)
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
		it.Event = new(TestEncloseRoleRevoked)
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
func (it *TestEncloseRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEncloseRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEncloseRoleRevoked represents a RoleRevoked event raised by the TestEnclose contract.
type TestEncloseRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TestEnclose *TestEncloseFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TestEncloseRoleRevokedIterator, error) {

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

	logs, sub, err := _TestEnclose.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TestEncloseRoleRevokedIterator{contract: _TestEnclose.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TestEnclose *TestEncloseFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TestEncloseRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TestEnclose.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEncloseRoleRevoked)
				if err := _TestEnclose.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TestEnclose *TestEncloseFilterer) ParseRoleRevoked(log types.Log) (*TestEncloseRoleRevoked, error) {
	event := new(TestEncloseRoleRevoked)
	if err := _TestEnclose.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestEncloseUsrDepositIterator is returned from FilterUsrDeposit and is used to iterate over the raw logs and unpacked data for UsrDeposit events raised by the TestEnclose contract.
type TestEncloseUsrDepositIterator struct {
	Event *TestEncloseUsrDeposit // Event containing the contract specifics and raw log

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
func (it *TestEncloseUsrDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEncloseUsrDeposit)
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
		it.Event = new(TestEncloseUsrDeposit)
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
func (it *TestEncloseUsrDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEncloseUsrDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEncloseUsrDeposit represents a UsrDeposit event raised by the TestEnclose contract.
type TestEncloseUsrDeposit struct {
	A   common.Address
	B   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUsrDeposit is a free log retrieval operation binding the contract event 0x089f4869c98e4e8a6236b0a3ea67d4a99eed42aafc725bea496d8c1bd37c4d89.
//
// Solidity: event UsrDeposit(address a, uint256 b)
func (_TestEnclose *TestEncloseFilterer) FilterUsrDeposit(opts *bind.FilterOpts) (*TestEncloseUsrDepositIterator, error) {

	logs, sub, err := _TestEnclose.contract.FilterLogs(opts, "UsrDeposit")
	if err != nil {
		return nil, err
	}
	return &TestEncloseUsrDepositIterator{contract: _TestEnclose.contract, event: "UsrDeposit", logs: logs, sub: sub}, nil
}

// WatchUsrDeposit is a free log subscription operation binding the contract event 0x089f4869c98e4e8a6236b0a3ea67d4a99eed42aafc725bea496d8c1bd37c4d89.
//
// Solidity: event UsrDeposit(address a, uint256 b)
func (_TestEnclose *TestEncloseFilterer) WatchUsrDeposit(opts *bind.WatchOpts, sink chan<- *TestEncloseUsrDeposit) (event.Subscription, error) {

	logs, sub, err := _TestEnclose.contract.WatchLogs(opts, "UsrDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEncloseUsrDeposit)
				if err := _TestEnclose.contract.UnpackLog(event, "UsrDeposit", log); err != nil {
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
// Solidity: event UsrDeposit(address a, uint256 b)
func (_TestEnclose *TestEncloseFilterer) ParseUsrDeposit(log types.Log) (*TestEncloseUsrDeposit, error) {
	event := new(TestEncloseUsrDeposit)
	if err := _TestEnclose.contract.UnpackLog(event, "UsrDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestEncloseContractPausedIterator is returned from FilterContractPaused and is used to iterate over the raw logs and unpacked data for ContractPaused events raised by the TestEnclose contract.
type TestEncloseContractPausedIterator struct {
	Event *TestEncloseContractPaused // Event containing the contract specifics and raw log

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
func (it *TestEncloseContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEncloseContractPaused)
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
		it.Event = new(TestEncloseContractPaused)
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
func (it *TestEncloseContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEncloseContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEncloseContractPaused represents a ContractPaused event raised by the TestEnclose contract.
type TestEncloseContractPaused struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractPaused is a free log retrieval operation binding the contract event 0xf378156ac648519c9cd03e0977fbb5b5c0552e1c0cd584389461f5349ba18657.
//
// Solidity: event contractPaused(bool value)
func (_TestEnclose *TestEncloseFilterer) FilterContractPaused(opts *bind.FilterOpts) (*TestEncloseContractPausedIterator, error) {

	logs, sub, err := _TestEnclose.contract.FilterLogs(opts, "contractPaused")
	if err != nil {
		return nil, err
	}
	return &TestEncloseContractPausedIterator{contract: _TestEnclose.contract, event: "contractPaused", logs: logs, sub: sub}, nil
}

// WatchContractPaused is a free log subscription operation binding the contract event 0xf378156ac648519c9cd03e0977fbb5b5c0552e1c0cd584389461f5349ba18657.
//
// Solidity: event contractPaused(bool value)
func (_TestEnclose *TestEncloseFilterer) WatchContractPaused(opts *bind.WatchOpts, sink chan<- *TestEncloseContractPaused) (event.Subscription, error) {

	logs, sub, err := _TestEnclose.contract.WatchLogs(opts, "contractPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEncloseContractPaused)
				if err := _TestEnclose.contract.UnpackLog(event, "contractPaused", log); err != nil {
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
func (_TestEnclose *TestEncloseFilterer) ParseContractPaused(log types.Log) (*TestEncloseContractPaused, error) {
	event := new(TestEncloseContractPaused)
	if err := _TestEnclose.contract.UnpackLog(event, "contractPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestEncloseTraillockIterator is returned from FilterTraillock and is used to iterate over the raw logs and unpacked data for Traillock events raised by the TestEnclose contract.
type TestEncloseTraillockIterator struct {
	Event *TestEncloseTraillock // Event containing the contract specifics and raw log

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
func (it *TestEncloseTraillockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEncloseTraillock)
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
		it.Event = new(TestEncloseTraillock)
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
func (it *TestEncloseTraillockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEncloseTraillockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEncloseTraillock represents a Traillock event raised by the TestEnclose contract.
type TestEncloseTraillock struct {
	Value uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTraillock is a free log retrieval operation binding the contract event 0xd02d7ece1b390124a8d8d40fc59fc1d6584a261d9a503b71166ec69b3eef00b3.
//
// Solidity: event traillock(uint8 value)
func (_TestEnclose *TestEncloseFilterer) FilterTraillock(opts *bind.FilterOpts) (*TestEncloseTraillockIterator, error) {

	logs, sub, err := _TestEnclose.contract.FilterLogs(opts, "traillock")
	if err != nil {
		return nil, err
	}
	return &TestEncloseTraillockIterator{contract: _TestEnclose.contract, event: "traillock", logs: logs, sub: sub}, nil
}

// WatchTraillock is a free log subscription operation binding the contract event 0xd02d7ece1b390124a8d8d40fc59fc1d6584a261d9a503b71166ec69b3eef00b3.
//
// Solidity: event traillock(uint8 value)
func (_TestEnclose *TestEncloseFilterer) WatchTraillock(opts *bind.WatchOpts, sink chan<- *TestEncloseTraillock) (event.Subscription, error) {

	logs, sub, err := _TestEnclose.contract.WatchLogs(opts, "traillock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEncloseTraillock)
				if err := _TestEnclose.contract.UnpackLog(event, "traillock", log); err != nil {
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
func (_TestEnclose *TestEncloseFilterer) ParseTraillock(log types.Log) (*TestEncloseTraillock, error) {
	event := new(TestEncloseTraillock)
	if err := _TestEnclose.contract.UnpackLog(event, "traillock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

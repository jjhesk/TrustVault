// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault_config_provider

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

// VaultConfigProviderMetaData contains all meta data concerning the VaultConfigProvider contract.
var VaultConfigProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hasProxy\",\"type\":\"bool\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGenesis\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelistAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"}],\"name\":\"setGenesis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistAdmins\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VaultConfigProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultConfigProviderMetaData.ABI instead.
var VaultConfigProviderABI = VaultConfigProviderMetaData.ABI

// VaultConfigProvider is an auto generated Go binding around an Ethereum contract.
type VaultConfigProvider struct {
	VaultConfigProviderCaller     // Read-only binding to the contract
	VaultConfigProviderTransactor // Write-only binding to the contract
	VaultConfigProviderFilterer   // Log filterer for contract events
}

// VaultConfigProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultConfigProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultConfigProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultConfigProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultConfigProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultConfigProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultConfigProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultConfigProviderSession struct {
	Contract     *VaultConfigProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VaultConfigProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultConfigProviderCallerSession struct {
	Contract *VaultConfigProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VaultConfigProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultConfigProviderTransactorSession struct {
	Contract     *VaultConfigProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VaultConfigProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultConfigProviderRaw struct {
	Contract *VaultConfigProvider // Generic contract binding to access the raw methods on
}

// VaultConfigProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultConfigProviderCallerRaw struct {
	Contract *VaultConfigProviderCaller // Generic read-only contract binding to access the raw methods on
}

// VaultConfigProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultConfigProviderTransactorRaw struct {
	Contract *VaultConfigProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultConfigProvider creates a new instance of VaultConfigProvider, bound to a specific deployed contract.
func NewVaultConfigProvider(address common.Address, backend bind.ContractBackend) (*VaultConfigProvider, error) {
	contract, err := bindVaultConfigProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultConfigProvider{VaultConfigProviderCaller: VaultConfigProviderCaller{contract: contract}, VaultConfigProviderTransactor: VaultConfigProviderTransactor{contract: contract}, VaultConfigProviderFilterer: VaultConfigProviderFilterer{contract: contract}}, nil
}

// NewVaultConfigProviderCaller creates a new read-only instance of VaultConfigProvider, bound to a specific deployed contract.
func NewVaultConfigProviderCaller(address common.Address, caller bind.ContractCaller) (*VaultConfigProviderCaller, error) {
	contract, err := bindVaultConfigProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultConfigProviderCaller{contract: contract}, nil
}

// NewVaultConfigProviderTransactor creates a new write-only instance of VaultConfigProvider, bound to a specific deployed contract.
func NewVaultConfigProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultConfigProviderTransactor, error) {
	contract, err := bindVaultConfigProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultConfigProviderTransactor{contract: contract}, nil
}

// NewVaultConfigProviderFilterer creates a new log filterer instance of VaultConfigProvider, bound to a specific deployed contract.
func NewVaultConfigProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultConfigProviderFilterer, error) {
	contract, err := bindVaultConfigProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultConfigProviderFilterer{contract: contract}, nil
}

// bindVaultConfigProvider binds a generic wrapper to an already deployed contract.
func bindVaultConfigProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultConfigProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultConfigProvider *VaultConfigProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultConfigProvider.Contract.VaultConfigProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultConfigProvider *VaultConfigProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.VaultConfigProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultConfigProvider *VaultConfigProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.VaultConfigProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultConfigProvider *VaultConfigProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultConfigProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultConfigProvider *VaultConfigProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultConfigProvider *VaultConfigProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultConfigProvider *VaultConfigProviderCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultConfigProvider *VaultConfigProviderSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultConfigProvider.Contract.DEFAULTADMINROLE(&_VaultConfigProvider.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultConfigProvider.Contract.DEFAULTADMINROLE(&_VaultConfigProvider.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_VaultConfigProvider *VaultConfigProviderCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_VaultConfigProvider *VaultConfigProviderSession) GetAddress(id [32]byte) (common.Address, error) {
	return _VaultConfigProvider.Contract.GetAddress(&_VaultConfigProvider.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _VaultConfigProvider.Contract.GetAddress(&_VaultConfigProvider.CallOpts, id)
}

// GetGenesis is a free data retrieval call binding the contract method 0x1a43bcb5.
//
// Solidity: function getGenesis() view returns(address)
func (_VaultConfigProvider *VaultConfigProviderCaller) GetGenesis(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "getGenesis")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGenesis is a free data retrieval call binding the contract method 0x1a43bcb5.
//
// Solidity: function getGenesis() view returns(address)
func (_VaultConfigProvider *VaultConfigProviderSession) GetGenesis() (common.Address, error) {
	return _VaultConfigProvider.Contract.GetGenesis(&_VaultConfigProvider.CallOpts)
}

// GetGenesis is a free data retrieval call binding the contract method 0x1a43bcb5.
//
// Solidity: function getGenesis() view returns(address)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) GetGenesis() (common.Address, error) {
	return _VaultConfigProvider.Contract.GetGenesis(&_VaultConfigProvider.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultConfigProvider *VaultConfigProviderCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultConfigProvider *VaultConfigProviderSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultConfigProvider.Contract.GetRoleAdmin(&_VaultConfigProvider.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultConfigProvider.Contract.GetRoleAdmin(&_VaultConfigProvider.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultConfigProvider *VaultConfigProviderCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultConfigProvider *VaultConfigProviderSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _VaultConfigProvider.Contract.GetRoleMember(&_VaultConfigProvider.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _VaultConfigProvider.Contract.GetRoleMember(&_VaultConfigProvider.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultConfigProvider *VaultConfigProviderCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultConfigProvider *VaultConfigProviderSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _VaultConfigProvider.Contract.GetRoleMemberCount(&_VaultConfigProvider.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _VaultConfigProvider.Contract.GetRoleMemberCount(&_VaultConfigProvider.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultConfigProvider *VaultConfigProviderCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultConfigProvider *VaultConfigProviderSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultConfigProvider.Contract.HasRole(&_VaultConfigProvider.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultConfigProvider.Contract.HasRole(&_VaultConfigProvider.CallOpts, role, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VaultConfigProvider *VaultConfigProviderCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VaultConfigProvider *VaultConfigProviderSession) IsOwner() (bool, error) {
	return _VaultConfigProvider.Contract.IsOwner(&_VaultConfigProvider.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) IsOwner() (bool, error) {
	return _VaultConfigProvider.Contract.IsOwner(&_VaultConfigProvider.CallOpts)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_VaultConfigProvider *VaultConfigProviderCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "isWhitelistAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_VaultConfigProvider *VaultConfigProviderSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _VaultConfigProvider.Contract.IsWhitelistAdmin(&_VaultConfigProvider.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _VaultConfigProvider.Contract.IsWhitelistAdmin(&_VaultConfigProvider.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VaultConfigProvider *VaultConfigProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VaultConfigProvider *VaultConfigProviderSession) Owner() (common.Address, error) {
	return _VaultConfigProvider.Contract.Owner(&_VaultConfigProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) Owner() (common.Address, error) {
	return _VaultConfigProvider.Contract.Owner(&_VaultConfigProvider.CallOpts)
}

// WhitelistAdmins is a free data retrieval call binding the contract method 0x7ef581cc.
//
// Solidity: function whitelistAdmins() view returns(bytes32)
func (_VaultConfigProvider *VaultConfigProviderCaller) WhitelistAdmins(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultConfigProvider.contract.Call(opts, &out, "whitelistAdmins")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WhitelistAdmins is a free data retrieval call binding the contract method 0x7ef581cc.
//
// Solidity: function whitelistAdmins() view returns(bytes32)
func (_VaultConfigProvider *VaultConfigProviderSession) WhitelistAdmins() ([32]byte, error) {
	return _VaultConfigProvider.Contract.WhitelistAdmins(&_VaultConfigProvider.CallOpts)
}

// WhitelistAdmins is a free data retrieval call binding the contract method 0x7ef581cc.
//
// Solidity: function whitelistAdmins() view returns(bytes32)
func (_VaultConfigProvider *VaultConfigProviderCallerSession) WhitelistAdmins() ([32]byte, error) {
	return _VaultConfigProvider.Contract.WhitelistAdmins(&_VaultConfigProvider.CallOpts)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_VaultConfigProvider *VaultConfigProviderSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.AddWhitelistAdmin(&_VaultConfigProvider.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.AddWhitelistAdmin(&_VaultConfigProvider.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultConfigProvider *VaultConfigProviderSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.GrantRole(&_VaultConfigProvider.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.GrantRole(&_VaultConfigProvider.TransactOpts, role, account)
}

// RemoveWhitelistAdmin is a paid mutator transaction binding the contract method 0x6897e974.
//
// Solidity: function removeWhitelistAdmin(address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactor) RemoveWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.contract.Transact(opts, "removeWhitelistAdmin", account)
}

// RemoveWhitelistAdmin is a paid mutator transaction binding the contract method 0x6897e974.
//
// Solidity: function removeWhitelistAdmin(address account) returns()
func (_VaultConfigProvider *VaultConfigProviderSession) RemoveWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.RemoveWhitelistAdmin(&_VaultConfigProvider.TransactOpts, account)
}

// RemoveWhitelistAdmin is a paid mutator transaction binding the contract method 0x6897e974.
//
// Solidity: function removeWhitelistAdmin(address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactorSession) RemoveWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.RemoveWhitelistAdmin(&_VaultConfigProvider.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VaultConfigProvider *VaultConfigProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultConfigProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VaultConfigProvider *VaultConfigProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.RenounceOwnership(&_VaultConfigProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VaultConfigProvider *VaultConfigProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.RenounceOwnership(&_VaultConfigProvider.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultConfigProvider *VaultConfigProviderSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.RenounceRole(&_VaultConfigProvider.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.RenounceRole(&_VaultConfigProvider.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultConfigProvider *VaultConfigProviderSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.RevokeRole(&_VaultConfigProvider.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.RevokeRole(&_VaultConfigProvider.TransactOpts, role, account)
}

// SetGenesis is a paid mutator transaction binding the contract method 0xeb14b353.
//
// Solidity: function setGenesis(address person) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactor) SetGenesis(opts *bind.TransactOpts, person common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.contract.Transact(opts, "setGenesis", person)
}

// SetGenesis is a paid mutator transaction binding the contract method 0xeb14b353.
//
// Solidity: function setGenesis(address person) returns()
func (_VaultConfigProvider *VaultConfigProviderSession) SetGenesis(person common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.SetGenesis(&_VaultConfigProvider.TransactOpts, person)
}

// SetGenesis is a paid mutator transaction binding the contract method 0xeb14b353.
//
// Solidity: function setGenesis(address person) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactorSession) SetGenesis(person common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.SetGenesis(&_VaultConfigProvider.TransactOpts, person)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VaultConfigProvider *VaultConfigProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.TransferOwnership(&_VaultConfigProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VaultConfigProvider *VaultConfigProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VaultConfigProvider.Contract.TransferOwnership(&_VaultConfigProvider.TransactOpts, newOwner)
}

// VaultConfigProviderAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the VaultConfigProvider contract.
type VaultConfigProviderAddressSetIterator struct {
	Event *VaultConfigProviderAddressSet // Event containing the contract specifics and raw log

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
func (it *VaultConfigProviderAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultConfigProviderAddressSet)
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
		it.Event = new(VaultConfigProviderAddressSet)
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
func (it *VaultConfigProviderAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultConfigProviderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultConfigProviderAddressSet represents a AddressSet event raised by the VaultConfigProvider contract.
type VaultConfigProviderAddressSet struct {
	Id         [32]byte
	NewAddress common.Address
	HasProxy   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_VaultConfigProvider *VaultConfigProviderFilterer) FilterAddressSet(opts *bind.FilterOpts, newAddress []common.Address) (*VaultConfigProviderAddressSetIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _VaultConfigProvider.contract.FilterLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &VaultConfigProviderAddressSetIterator{contract: _VaultConfigProvider.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_VaultConfigProvider *VaultConfigProviderFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *VaultConfigProviderAddressSet, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _VaultConfigProvider.contract.WatchLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultConfigProviderAddressSet)
				if err := _VaultConfigProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_VaultConfigProvider *VaultConfigProviderFilterer) ParseAddressSet(log types.Log) (*VaultConfigProviderAddressSet, error) {
	event := new(VaultConfigProviderAddressSet)
	if err := _VaultConfigProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultConfigProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VaultConfigProvider contract.
type VaultConfigProviderOwnershipTransferredIterator struct {
	Event *VaultConfigProviderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VaultConfigProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultConfigProviderOwnershipTransferred)
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
		it.Event = new(VaultConfigProviderOwnershipTransferred)
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
func (it *VaultConfigProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultConfigProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultConfigProviderOwnershipTransferred represents a OwnershipTransferred event raised by the VaultConfigProvider contract.
type VaultConfigProviderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VaultConfigProvider *VaultConfigProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VaultConfigProviderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VaultConfigProvider.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VaultConfigProviderOwnershipTransferredIterator{contract: _VaultConfigProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VaultConfigProvider *VaultConfigProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VaultConfigProviderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VaultConfigProvider.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultConfigProviderOwnershipTransferred)
				if err := _VaultConfigProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VaultConfigProvider *VaultConfigProviderFilterer) ParseOwnershipTransferred(log types.Log) (*VaultConfigProviderOwnershipTransferred, error) {
	event := new(VaultConfigProviderOwnershipTransferred)
	if err := _VaultConfigProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultConfigProviderRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VaultConfigProvider contract.
type VaultConfigProviderRoleGrantedIterator struct {
	Event *VaultConfigProviderRoleGranted // Event containing the contract specifics and raw log

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
func (it *VaultConfigProviderRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultConfigProviderRoleGranted)
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
		it.Event = new(VaultConfigProviderRoleGranted)
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
func (it *VaultConfigProviderRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultConfigProviderRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultConfigProviderRoleGranted represents a RoleGranted event raised by the VaultConfigProvider contract.
type VaultConfigProviderRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultConfigProvider *VaultConfigProviderFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultConfigProviderRoleGrantedIterator, error) {

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

	logs, sub, err := _VaultConfigProvider.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultConfigProviderRoleGrantedIterator{contract: _VaultConfigProvider.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultConfigProvider *VaultConfigProviderFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VaultConfigProviderRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VaultConfigProvider.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultConfigProviderRoleGranted)
				if err := _VaultConfigProvider.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_VaultConfigProvider *VaultConfigProviderFilterer) ParseRoleGranted(log types.Log) (*VaultConfigProviderRoleGranted, error) {
	event := new(VaultConfigProviderRoleGranted)
	if err := _VaultConfigProvider.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultConfigProviderRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VaultConfigProvider contract.
type VaultConfigProviderRoleRevokedIterator struct {
	Event *VaultConfigProviderRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VaultConfigProviderRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultConfigProviderRoleRevoked)
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
		it.Event = new(VaultConfigProviderRoleRevoked)
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
func (it *VaultConfigProviderRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultConfigProviderRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultConfigProviderRoleRevoked represents a RoleRevoked event raised by the VaultConfigProvider contract.
type VaultConfigProviderRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultConfigProvider *VaultConfigProviderFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultConfigProviderRoleRevokedIterator, error) {

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

	logs, sub, err := _VaultConfigProvider.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultConfigProviderRoleRevokedIterator{contract: _VaultConfigProvider.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultConfigProvider *VaultConfigProviderFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VaultConfigProviderRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VaultConfigProvider.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultConfigProviderRoleRevoked)
				if err := _VaultConfigProvider.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_VaultConfigProvider *VaultConfigProviderFilterer) ParseRoleRevoked(log types.Log) (*VaultConfigProviderRoleRevoked, error) {
	event := new(VaultConfigProviderRoleRevoked)
	if err := _VaultConfigProvider.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

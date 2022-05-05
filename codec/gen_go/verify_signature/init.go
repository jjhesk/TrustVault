// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verify_signature

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

// VerifySignatureMetaData contains all meta data concerning the VerifySignature contract.
var VerifySignatureMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// VerifySignatureABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifySignatureMetaData.ABI instead.
var VerifySignatureABI = VerifySignatureMetaData.ABI

// VerifySignature is an auto generated Go binding around an Ethereum contract.
type VerifySignature struct {
	VerifySignatureCaller     // Read-only binding to the contract
	VerifySignatureTransactor // Write-only binding to the contract
	VerifySignatureFilterer   // Log filterer for contract events
}

// VerifySignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifySignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifySignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySignatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifySignatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifySignatureSession struct {
	Contract     *VerifySignature  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifySignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifySignatureCallerSession struct {
	Contract *VerifySignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VerifySignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifySignatureTransactorSession struct {
	Contract     *VerifySignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VerifySignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifySignatureRaw struct {
	Contract *VerifySignature // Generic contract binding to access the raw methods on
}

// VerifySignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifySignatureCallerRaw struct {
	Contract *VerifySignatureCaller // Generic read-only contract binding to access the raw methods on
}

// VerifySignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifySignatureTransactorRaw struct {
	Contract *VerifySignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifySignature creates a new instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignature(address common.Address, backend bind.ContractBackend) (*VerifySignature, error) {
	contract, err := bindVerifySignature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifySignature{VerifySignatureCaller: VerifySignatureCaller{contract: contract}, VerifySignatureTransactor: VerifySignatureTransactor{contract: contract}, VerifySignatureFilterer: VerifySignatureFilterer{contract: contract}}, nil
}

// NewVerifySignatureCaller creates a new read-only instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignatureCaller(address common.Address, caller bind.ContractCaller) (*VerifySignatureCaller, error) {
	contract, err := bindVerifySignature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifySignatureCaller{contract: contract}, nil
}

// NewVerifySignatureTransactor creates a new write-only instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifySignatureTransactor, error) {
	contract, err := bindVerifySignature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifySignatureTransactor{contract: contract}, nil
}

// NewVerifySignatureFilterer creates a new log filterer instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignatureFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifySignatureFilterer, error) {
	contract, err := bindVerifySignature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifySignatureFilterer{contract: contract}, nil
}

// bindVerifySignature binds a generic wrapper to an already deployed contract.
func bindVerifySignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifySignatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifySignature *VerifySignatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifySignature.Contract.VerifySignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifySignature *VerifySignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifySignature.Contract.VerifySignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifySignature *VerifySignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifySignature.Contract.VerifySignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifySignature *VerifySignatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifySignature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifySignature *VerifySignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifySignature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifySignature *VerifySignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifySignature.Contract.contract.Transact(opts, method, params...)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_VerifySignature *VerifySignatureCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VerifySignature.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_VerifySignature *VerifySignatureSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _VerifySignature.Contract.GetEthSignedMessageHash(&_VerifySignature.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_VerifySignature *VerifySignatureCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _VerifySignature.Contract.GetEthSignedMessageHash(&_VerifySignature.CallOpts, _messageHash)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xa35f8a40.
//
// Solidity: function getMessageHash(address _to, uint256 _amount, string _message, uint256 _nonce) pure returns(bytes32)
func (_VerifySignature *VerifySignatureCaller) GetMessageHash(opts *bind.CallOpts, _to common.Address, _amount *big.Int, _message string, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VerifySignature.contract.Call(opts, &out, "getMessageHash", _to, _amount, _message, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xa35f8a40.
//
// Solidity: function getMessageHash(address _to, uint256 _amount, string _message, uint256 _nonce) pure returns(bytes32)
func (_VerifySignature *VerifySignatureSession) GetMessageHash(_to common.Address, _amount *big.Int, _message string, _nonce *big.Int) ([32]byte, error) {
	return _VerifySignature.Contract.GetMessageHash(&_VerifySignature.CallOpts, _to, _amount, _message, _nonce)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xa35f8a40.
//
// Solidity: function getMessageHash(address _to, uint256 _amount, string _message, uint256 _nonce) pure returns(bytes32)
func (_VerifySignature *VerifySignatureCallerSession) GetMessageHash(_to common.Address, _amount *big.Int, _message string, _nonce *big.Int) ([32]byte, error) {
	return _VerifySignature.Contract.GetMessageHash(&_VerifySignature.CallOpts, _to, _amount, _message, _nonce)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_VerifySignature *VerifySignatureCaller) RecoverSigner(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _VerifySignature.contract.Call(opts, &out, "recoverSigner", _ethSignedMessageHash, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_VerifySignature *VerifySignatureSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _VerifySignature.Contract.RecoverSigner(&_VerifySignature.CallOpts, _ethSignedMessageHash, _signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_VerifySignature *VerifySignatureCallerSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _VerifySignature.Contract.RecoverSigner(&_VerifySignature.CallOpts, _ethSignedMessageHash, _signature)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_VerifySignature *VerifySignatureCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _VerifySignature.contract.Call(opts, &out, "splitSignature", sig)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
		V uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.V = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_VerifySignature *VerifySignatureSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _VerifySignature.Contract.SplitSignature(&_VerifySignature.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_VerifySignature *VerifySignatureCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _VerifySignature.Contract.SplitSignature(&_VerifySignature.CallOpts, sig)
}

// Verify is a free data retrieval call binding the contract method 0xabe5026a.
//
// Solidity: function verify(address _signer, address _to, uint256 _amount, string _message, uint256 _nonce, bytes signature) pure returns(bool)
func (_VerifySignature *VerifySignatureCaller) Verify(opts *bind.CallOpts, _signer common.Address, _to common.Address, _amount *big.Int, _message string, _nonce *big.Int, signature []byte) (bool, error) {
	var out []interface{}
	err := _VerifySignature.contract.Call(opts, &out, "verify", _signer, _to, _amount, _message, _nonce, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xabe5026a.
//
// Solidity: function verify(address _signer, address _to, uint256 _amount, string _message, uint256 _nonce, bytes signature) pure returns(bool)
func (_VerifySignature *VerifySignatureSession) Verify(_signer common.Address, _to common.Address, _amount *big.Int, _message string, _nonce *big.Int, signature []byte) (bool, error) {
	return _VerifySignature.Contract.Verify(&_VerifySignature.CallOpts, _signer, _to, _amount, _message, _nonce, signature)
}

// Verify is a free data retrieval call binding the contract method 0xabe5026a.
//
// Solidity: function verify(address _signer, address _to, uint256 _amount, string _message, uint256 _nonce, bytes signature) pure returns(bool)
func (_VerifySignature *VerifySignatureCallerSession) Verify(_signer common.Address, _to common.Address, _amount *big.Int, _message string, _nonce *big.Int, signature []byte) (bool, error) {
	return _VerifySignature.Contract.Verify(&_VerifySignature.CallOpts, _signer, _to, _amount, _message, _nonce, signature)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package safeproxyfactory

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

// SafeProxyFactorySig is an auto generated low-level Go binding around an user-defined struct.
type SafeProxyFactorySig struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SafeProxyFactoryMetaData contains all meta data concerning the SafeProxyFactory contract.
var SafeProxyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_masterCopy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fallbackHandler\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractGnosisSafe\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ProxyCreation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATE_PROXY_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"computeProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSafeProxyFactory.Sig\",\"name\":\"createSig\",\"type\":\"tuple\"}],\"name\":\"createProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fallbackHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBytecode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSalt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterCopy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyCreationCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// SafeProxyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeProxyFactoryMetaData.ABI instead.
var SafeProxyFactoryABI = SafeProxyFactoryMetaData.ABI

// SafeProxyFactory is an auto generated Go binding around an Ethereum contract.
type SafeProxyFactory struct {
	SafeProxyFactoryCaller     // Read-only binding to the contract
	SafeProxyFactoryTransactor // Write-only binding to the contract
	SafeProxyFactoryFilterer   // Log filterer for contract events
}

// SafeProxyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeProxyFactorySession struct {
	Contract     *SafeProxyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeProxyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeProxyFactoryCallerSession struct {
	Contract *SafeProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SafeProxyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeProxyFactoryTransactorSession struct {
	Contract     *SafeProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SafeProxyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeProxyFactoryRaw struct {
	Contract *SafeProxyFactory // Generic contract binding to access the raw methods on
}

// SafeProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeProxyFactoryCallerRaw struct {
	Contract *SafeProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SafeProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeProxyFactoryTransactorRaw struct {
	Contract *SafeProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeProxyFactory creates a new instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactory(address common.Address, backend bind.ContractBackend) (*SafeProxyFactory, error) {
	contract, err := bindSafeProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactory{SafeProxyFactoryCaller: SafeProxyFactoryCaller{contract: contract}, SafeProxyFactoryTransactor: SafeProxyFactoryTransactor{contract: contract}, SafeProxyFactoryFilterer: SafeProxyFactoryFilterer{contract: contract}}, nil
}

// NewSafeProxyFactoryCaller creates a new read-only instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*SafeProxyFactoryCaller, error) {
	contract, err := bindSafeProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryCaller{contract: contract}, nil
}

// NewSafeProxyFactoryTransactor creates a new write-only instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeProxyFactoryTransactor, error) {
	contract, err := bindSafeProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryTransactor{contract: contract}, nil
}

// NewSafeProxyFactoryFilterer creates a new log filterer instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeProxyFactoryFilterer, error) {
	contract, err := bindSafeProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryFilterer{contract: contract}, nil
}

// bindSafeProxyFactory binds a generic wrapper to an already deployed contract.
func bindSafeProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeProxyFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactory *SafeProxyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactory.Contract.SafeProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactory *SafeProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.SafeProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactory *SafeProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.SafeProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactory *SafeProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactory *SafeProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactory *SafeProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// CREATEPROXYTYPEHASH is a free data retrieval call binding the contract method 0x98827e29.
//
// Solidity: function CREATE_PROXY_TYPEHASH() view returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactoryCaller) CREATEPROXYTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "CREATE_PROXY_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATEPROXYTYPEHASH is a free data retrieval call binding the contract method 0x98827e29.
//
// Solidity: function CREATE_PROXY_TYPEHASH() view returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactorySession) CREATEPROXYTYPEHASH() ([32]byte, error) {
	return _SafeProxyFactory.Contract.CREATEPROXYTYPEHASH(&_SafeProxyFactory.CallOpts)
}

// CREATEPROXYTYPEHASH is a free data retrieval call binding the contract method 0x98827e29.
//
// Solidity: function CREATE_PROXY_TYPEHASH() view returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) CREATEPROXYTYPEHASH() ([32]byte, error) {
	return _SafeProxyFactory.Contract.CREATEPROXYTYPEHASH(&_SafeProxyFactory.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactoryCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactorySession) DOMAINTYPEHASH() ([32]byte, error) {
	return _SafeProxyFactory.Contract.DOMAINTYPEHASH(&_SafeProxyFactory.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _SafeProxyFactory.Contract.DOMAINTYPEHASH(&_SafeProxyFactory.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_SafeProxyFactory *SafeProxyFactoryCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_SafeProxyFactory *SafeProxyFactorySession) NAME() (string, error) {
	return _SafeProxyFactory.Contract.NAME(&_SafeProxyFactory.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) NAME() (string, error) {
	return _SafeProxyFactory.Contract.NAME(&_SafeProxyFactory.CallOpts)
}

// ComputeProxyAddress is a free data retrieval call binding the contract method 0xd600539a.
//
// Solidity: function computeProxyAddress(address user) view returns(address)
func (_SafeProxyFactory *SafeProxyFactoryCaller) ComputeProxyAddress(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "computeProxyAddress", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeProxyAddress is a free data retrieval call binding the contract method 0xd600539a.
//
// Solidity: function computeProxyAddress(address user) view returns(address)
func (_SafeProxyFactory *SafeProxyFactorySession) ComputeProxyAddress(user common.Address) (common.Address, error) {
	return _SafeProxyFactory.Contract.ComputeProxyAddress(&_SafeProxyFactory.CallOpts, user)
}

// ComputeProxyAddress is a free data retrieval call binding the contract method 0xd600539a.
//
// Solidity: function computeProxyAddress(address user) view returns(address)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) ComputeProxyAddress(user common.Address) (common.Address, error) {
	return _SafeProxyFactory.Contract.ComputeProxyAddress(&_SafeProxyFactory.CallOpts, user)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactoryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactorySession) DomainSeparator() ([32]byte, error) {
	return _SafeProxyFactory.Contract.DomainSeparator(&_SafeProxyFactory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) DomainSeparator() ([32]byte, error) {
	return _SafeProxyFactory.Contract.DomainSeparator(&_SafeProxyFactory.CallOpts)
}

// FallbackHandler is a free data retrieval call binding the contract method 0xeed2f252.
//
// Solidity: function fallbackHandler() view returns(address)
func (_SafeProxyFactory *SafeProxyFactoryCaller) FallbackHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "fallbackHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FallbackHandler is a free data retrieval call binding the contract method 0xeed2f252.
//
// Solidity: function fallbackHandler() view returns(address)
func (_SafeProxyFactory *SafeProxyFactorySession) FallbackHandler() (common.Address, error) {
	return _SafeProxyFactory.Contract.FallbackHandler(&_SafeProxyFactory.CallOpts)
}

// FallbackHandler is a free data retrieval call binding the contract method 0xeed2f252.
//
// Solidity: function fallbackHandler() view returns(address)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) FallbackHandler() (common.Address, error) {
	return _SafeProxyFactory.Contract.FallbackHandler(&_SafeProxyFactory.CallOpts)
}

// GetContractBytecode is a free data retrieval call binding the contract method 0xbf82787d.
//
// Solidity: function getContractBytecode() view returns(bytes)
func (_SafeProxyFactory *SafeProxyFactoryCaller) GetContractBytecode(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "getContractBytecode")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetContractBytecode is a free data retrieval call binding the contract method 0xbf82787d.
//
// Solidity: function getContractBytecode() view returns(bytes)
func (_SafeProxyFactory *SafeProxyFactorySession) GetContractBytecode() ([]byte, error) {
	return _SafeProxyFactory.Contract.GetContractBytecode(&_SafeProxyFactory.CallOpts)
}

// GetContractBytecode is a free data retrieval call binding the contract method 0xbf82787d.
//
// Solidity: function getContractBytecode() view returns(bytes)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) GetContractBytecode() ([]byte, error) {
	return _SafeProxyFactory.Contract.GetContractBytecode(&_SafeProxyFactory.CallOpts)
}

// GetSalt is a free data retrieval call binding the contract method 0x59709c45.
//
// Solidity: function getSalt(address user) pure returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactoryCaller) GetSalt(opts *bind.CallOpts, user common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "getSalt", user)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSalt is a free data retrieval call binding the contract method 0x59709c45.
//
// Solidity: function getSalt(address user) pure returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactorySession) GetSalt(user common.Address) ([32]byte, error) {
	return _SafeProxyFactory.Contract.GetSalt(&_SafeProxyFactory.CallOpts, user)
}

// GetSalt is a free data retrieval call binding the contract method 0x59709c45.
//
// Solidity: function getSalt(address user) pure returns(bytes32)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) GetSalt(user common.Address) ([32]byte, error) {
	return _SafeProxyFactory.Contract.GetSalt(&_SafeProxyFactory.CallOpts, user)
}

// MasterCopy is a free data retrieval call binding the contract method 0xa619486e.
//
// Solidity: function masterCopy() view returns(address)
func (_SafeProxyFactory *SafeProxyFactoryCaller) MasterCopy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "masterCopy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterCopy is a free data retrieval call binding the contract method 0xa619486e.
//
// Solidity: function masterCopy() view returns(address)
func (_SafeProxyFactory *SafeProxyFactorySession) MasterCopy() (common.Address, error) {
	return _SafeProxyFactory.Contract.MasterCopy(&_SafeProxyFactory.CallOpts)
}

// MasterCopy is a free data retrieval call binding the contract method 0xa619486e.
//
// Solidity: function masterCopy() view returns(address)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) MasterCopy() (common.Address, error) {
	return _SafeProxyFactory.Contract.MasterCopy(&_SafeProxyFactory.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactory *SafeProxyFactoryCaller) ProxyCreationCode(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "proxyCreationCode")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactory *SafeProxyFactorySession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactory.Contract.ProxyCreationCode(&_SafeProxyFactory.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactory.Contract.ProxyCreationCode(&_SafeProxyFactory.CallOpts)
}

// CreateProxy is a paid mutator transaction binding the contract method 0xa1884d2c.
//
// Solidity: function createProxy(address paymentToken, uint256 payment, address paymentReceiver, (uint8,bytes32,bytes32) createSig) returns()
func (_SafeProxyFactory *SafeProxyFactoryTransactor) CreateProxy(opts *bind.TransactOpts, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address, createSig SafeProxyFactorySig) (*types.Transaction, error) {
	return _SafeProxyFactory.contract.Transact(opts, "createProxy", paymentToken, payment, paymentReceiver, createSig)
}

// CreateProxy is a paid mutator transaction binding the contract method 0xa1884d2c.
//
// Solidity: function createProxy(address paymentToken, uint256 payment, address paymentReceiver, (uint8,bytes32,bytes32) createSig) returns()
func (_SafeProxyFactory *SafeProxyFactorySession) CreateProxy(paymentToken common.Address, payment *big.Int, paymentReceiver common.Address, createSig SafeProxyFactorySig) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxy(&_SafeProxyFactory.TransactOpts, paymentToken, payment, paymentReceiver, createSig)
}

// CreateProxy is a paid mutator transaction binding the contract method 0xa1884d2c.
//
// Solidity: function createProxy(address paymentToken, uint256 payment, address paymentReceiver, (uint8,bytes32,bytes32) createSig) returns()
func (_SafeProxyFactory *SafeProxyFactoryTransactorSession) CreateProxy(paymentToken common.Address, payment *big.Int, paymentReceiver common.Address, createSig SafeProxyFactorySig) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxy(&_SafeProxyFactory.TransactOpts, paymentToken, payment, paymentReceiver, createSig)
}

// SafeProxyFactoryProxyCreationIterator is returned from FilterProxyCreation and is used to iterate over the raw logs and unpacked data for ProxyCreation events raised by the SafeProxyFactory contract.
type SafeProxyFactoryProxyCreationIterator struct {
	Event *SafeProxyFactoryProxyCreation // Event containing the contract specifics and raw log

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
func (it *SafeProxyFactoryProxyCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeProxyFactoryProxyCreation)
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
		it.Event = new(SafeProxyFactoryProxyCreation)
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
func (it *SafeProxyFactoryProxyCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeProxyFactoryProxyCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeProxyFactoryProxyCreation represents a ProxyCreation event raised by the SafeProxyFactory contract.
type SafeProxyFactoryProxyCreation struct {
	Proxy common.Address
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreation is a free log retrieval operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address proxy, address owner)
func (_SafeProxyFactory *SafeProxyFactoryFilterer) FilterProxyCreation(opts *bind.FilterOpts) (*SafeProxyFactoryProxyCreationIterator, error) {

	logs, sub, err := _SafeProxyFactory.contract.FilterLogs(opts, "ProxyCreation")
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryProxyCreationIterator{contract: _SafeProxyFactory.contract, event: "ProxyCreation", logs: logs, sub: sub}, nil
}

// WatchProxyCreation is a free log subscription operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address proxy, address owner)
func (_SafeProxyFactory *SafeProxyFactoryFilterer) WatchProxyCreation(opts *bind.WatchOpts, sink chan<- *SafeProxyFactoryProxyCreation) (event.Subscription, error) {

	logs, sub, err := _SafeProxyFactory.contract.WatchLogs(opts, "ProxyCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeProxyFactoryProxyCreation)
				if err := _SafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
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

// ParseProxyCreation is a log parse operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address proxy, address owner)
func (_SafeProxyFactory *SafeProxyFactoryFilterer) ParseProxyCreation(log types.Log) (*SafeProxyFactoryProxyCreation, error) {
	event := new(SafeProxyFactoryProxyCreation)
	if err := _SafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

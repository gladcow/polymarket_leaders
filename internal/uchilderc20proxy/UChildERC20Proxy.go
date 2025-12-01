// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uchilderc20proxy

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

// UchaildErc20ProxyMetaData contains all meta data concerning the UchaildErc20Proxy contract.
var UchaildErc20ProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyTo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_new\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_old\",\"type\":\"address\"}],\"name\":\"ProxyOwnerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_new\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_old\",\"type\":\"address\"}],\"name\":\"ProxyUpdated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"IMPLEMENTATION_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNER_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proxyTypeId\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferProxyOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProxyTo\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"updateAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProxyTo\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UchaildErc20ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use UchaildErc20ProxyMetaData.ABI instead.
var UchaildErc20ProxyABI = UchaildErc20ProxyMetaData.ABI

// UchaildErc20Proxy is an auto generated Go binding around an Ethereum contract.
type UchaildErc20Proxy struct {
	UchaildErc20ProxyCaller     // Read-only binding to the contract
	UchaildErc20ProxyTransactor // Write-only binding to the contract
	UchaildErc20ProxyFilterer   // Log filterer for contract events
}

// UchaildErc20ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type UchaildErc20ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UchaildErc20ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UchaildErc20ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UchaildErc20ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UchaildErc20ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UchaildErc20ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UchaildErc20ProxySession struct {
	Contract     *UchaildErc20Proxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UchaildErc20ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UchaildErc20ProxyCallerSession struct {
	Contract *UchaildErc20ProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UchaildErc20ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UchaildErc20ProxyTransactorSession struct {
	Contract     *UchaildErc20ProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UchaildErc20ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type UchaildErc20ProxyRaw struct {
	Contract *UchaildErc20Proxy // Generic contract binding to access the raw methods on
}

// UchaildErc20ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UchaildErc20ProxyCallerRaw struct {
	Contract *UchaildErc20ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// UchaildErc20ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UchaildErc20ProxyTransactorRaw struct {
	Contract *UchaildErc20ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUchaildErc20Proxy creates a new instance of UchaildErc20Proxy, bound to a specific deployed contract.
func NewUchaildErc20Proxy(address common.Address, backend bind.ContractBackend) (*UchaildErc20Proxy, error) {
	contract, err := bindUchaildErc20Proxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UchaildErc20Proxy{UchaildErc20ProxyCaller: UchaildErc20ProxyCaller{contract: contract}, UchaildErc20ProxyTransactor: UchaildErc20ProxyTransactor{contract: contract}, UchaildErc20ProxyFilterer: UchaildErc20ProxyFilterer{contract: contract}}, nil
}

// NewUchaildErc20ProxyCaller creates a new read-only instance of UchaildErc20Proxy, bound to a specific deployed contract.
func NewUchaildErc20ProxyCaller(address common.Address, caller bind.ContractCaller) (*UchaildErc20ProxyCaller, error) {
	contract, err := bindUchaildErc20Proxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UchaildErc20ProxyCaller{contract: contract}, nil
}

// NewUchaildErc20ProxyTransactor creates a new write-only instance of UchaildErc20Proxy, bound to a specific deployed contract.
func NewUchaildErc20ProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*UchaildErc20ProxyTransactor, error) {
	contract, err := bindUchaildErc20Proxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UchaildErc20ProxyTransactor{contract: contract}, nil
}

// NewUchaildErc20ProxyFilterer creates a new log filterer instance of UchaildErc20Proxy, bound to a specific deployed contract.
func NewUchaildErc20ProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*UchaildErc20ProxyFilterer, error) {
	contract, err := bindUchaildErc20Proxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UchaildErc20ProxyFilterer{contract: contract}, nil
}

// bindUchaildErc20Proxy binds a generic wrapper to an already deployed contract.
func bindUchaildErc20Proxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UchaildErc20ProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UchaildErc20Proxy *UchaildErc20ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UchaildErc20Proxy.Contract.UchaildErc20ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UchaildErc20Proxy *UchaildErc20ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.UchaildErc20ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UchaildErc20Proxy *UchaildErc20ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.UchaildErc20ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UchaildErc20Proxy *UchaildErc20ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UchaildErc20Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.contract.Transact(opts, method, params...)
}

// IMPLEMENTATIONSLOT is a free data retrieval call binding the contract method 0x086fc0c7.
//
// Solidity: function IMPLEMENTATION_SLOT() view returns(bytes32)
func (_UchaildErc20Proxy *UchaildErc20ProxyCaller) IMPLEMENTATIONSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UchaildErc20Proxy.contract.Call(opts, &out, "IMPLEMENTATION_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IMPLEMENTATIONSLOT is a free data retrieval call binding the contract method 0x086fc0c7.
//
// Solidity: function IMPLEMENTATION_SLOT() view returns(bytes32)
func (_UchaildErc20Proxy *UchaildErc20ProxySession) IMPLEMENTATIONSLOT() ([32]byte, error) {
	return _UchaildErc20Proxy.Contract.IMPLEMENTATIONSLOT(&_UchaildErc20Proxy.CallOpts)
}

// IMPLEMENTATIONSLOT is a free data retrieval call binding the contract method 0x086fc0c7.
//
// Solidity: function IMPLEMENTATION_SLOT() view returns(bytes32)
func (_UchaildErc20Proxy *UchaildErc20ProxyCallerSession) IMPLEMENTATIONSLOT() ([32]byte, error) {
	return _UchaildErc20Proxy.Contract.IMPLEMENTATIONSLOT(&_UchaildErc20Proxy.CallOpts)
}

// OWNERSLOT is a free data retrieval call binding the contract method 0x963949a3.
//
// Solidity: function OWNER_SLOT() view returns(bytes32)
func (_UchaildErc20Proxy *UchaildErc20ProxyCaller) OWNERSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UchaildErc20Proxy.contract.Call(opts, &out, "OWNER_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNERSLOT is a free data retrieval call binding the contract method 0x963949a3.
//
// Solidity: function OWNER_SLOT() view returns(bytes32)
func (_UchaildErc20Proxy *UchaildErc20ProxySession) OWNERSLOT() ([32]byte, error) {
	return _UchaildErc20Proxy.Contract.OWNERSLOT(&_UchaildErc20Proxy.CallOpts)
}

// OWNERSLOT is a free data retrieval call binding the contract method 0x963949a3.
//
// Solidity: function OWNER_SLOT() view returns(bytes32)
func (_UchaildErc20Proxy *UchaildErc20ProxyCallerSession) OWNERSLOT() ([32]byte, error) {
	return _UchaildErc20Proxy.Contract.OWNERSLOT(&_UchaildErc20Proxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UchaildErc20Proxy *UchaildErc20ProxyCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UchaildErc20Proxy.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UchaildErc20Proxy *UchaildErc20ProxySession) Implementation() (common.Address, error) {
	return _UchaildErc20Proxy.Contract.Implementation(&_UchaildErc20Proxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UchaildErc20Proxy *UchaildErc20ProxyCallerSession) Implementation() (common.Address, error) {
	return _UchaildErc20Proxy.Contract.Implementation(&_UchaildErc20Proxy.CallOpts)
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address)
func (_UchaildErc20Proxy *UchaildErc20ProxyCaller) ProxyOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UchaildErc20Proxy.contract.Call(opts, &out, "proxyOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address)
func (_UchaildErc20Proxy *UchaildErc20ProxySession) ProxyOwner() (common.Address, error) {
	return _UchaildErc20Proxy.Contract.ProxyOwner(&_UchaildErc20Proxy.CallOpts)
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address)
func (_UchaildErc20Proxy *UchaildErc20ProxyCallerSession) ProxyOwner() (common.Address, error) {
	return _UchaildErc20Proxy.Contract.ProxyOwner(&_UchaildErc20Proxy.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_UchaildErc20Proxy *UchaildErc20ProxyCaller) ProxyType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UchaildErc20Proxy.contract.Call(opts, &out, "proxyType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_UchaildErc20Proxy *UchaildErc20ProxySession) ProxyType() (*big.Int, error) {
	return _UchaildErc20Proxy.Contract.ProxyType(&_UchaildErc20Proxy.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_UchaildErc20Proxy *UchaildErc20ProxyCallerSession) ProxyType() (*big.Int, error) {
	return _UchaildErc20Proxy.Contract.ProxyType(&_UchaildErc20Proxy.CallOpts)
}

// TransferProxyOwnership is a paid mutator transaction binding the contract method 0xf1739cae.
//
// Solidity: function transferProxyOwnership(address newOwner) returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactor) TransferProxyOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UchaildErc20Proxy.contract.Transact(opts, "transferProxyOwnership", newOwner)
}

// TransferProxyOwnership is a paid mutator transaction binding the contract method 0xf1739cae.
//
// Solidity: function transferProxyOwnership(address newOwner) returns()
func (_UchaildErc20Proxy *UchaildErc20ProxySession) TransferProxyOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.TransferProxyOwnership(&_UchaildErc20Proxy.TransactOpts, newOwner)
}

// TransferProxyOwnership is a paid mutator transaction binding the contract method 0xf1739cae.
//
// Solidity: function transferProxyOwnership(address newOwner) returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactorSession) TransferProxyOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.TransferProxyOwnership(&_UchaildErc20Proxy.TransactOpts, newOwner)
}

// UpdateAndCall is a paid mutator transaction binding the contract method 0xd88ca2c8.
//
// Solidity: function updateAndCall(address _newProxyTo, bytes data) payable returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactor) UpdateAndCall(opts *bind.TransactOpts, _newProxyTo common.Address, data []byte) (*types.Transaction, error) {
	return _UchaildErc20Proxy.contract.Transact(opts, "updateAndCall", _newProxyTo, data)
}

// UpdateAndCall is a paid mutator transaction binding the contract method 0xd88ca2c8.
//
// Solidity: function updateAndCall(address _newProxyTo, bytes data) payable returns()
func (_UchaildErc20Proxy *UchaildErc20ProxySession) UpdateAndCall(_newProxyTo common.Address, data []byte) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.UpdateAndCall(&_UchaildErc20Proxy.TransactOpts, _newProxyTo, data)
}

// UpdateAndCall is a paid mutator transaction binding the contract method 0xd88ca2c8.
//
// Solidity: function updateAndCall(address _newProxyTo, bytes data) payable returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactorSession) UpdateAndCall(_newProxyTo common.Address, data []byte) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.UpdateAndCall(&_UchaildErc20Proxy.TransactOpts, _newProxyTo, data)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _newProxyTo) returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactor) UpdateImplementation(opts *bind.TransactOpts, _newProxyTo common.Address) (*types.Transaction, error) {
	return _UchaildErc20Proxy.contract.Transact(opts, "updateImplementation", _newProxyTo)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _newProxyTo) returns()
func (_UchaildErc20Proxy *UchaildErc20ProxySession) UpdateImplementation(_newProxyTo common.Address) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.UpdateImplementation(&_UchaildErc20Proxy.TransactOpts, _newProxyTo)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _newProxyTo) returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactorSession) UpdateImplementation(_newProxyTo common.Address) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.UpdateImplementation(&_UchaildErc20Proxy.TransactOpts, _newProxyTo)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UchaildErc20Proxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UchaildErc20Proxy *UchaildErc20ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.Fallback(&_UchaildErc20Proxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.Fallback(&_UchaildErc20Proxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UchaildErc20Proxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UchaildErc20Proxy *UchaildErc20ProxySession) Receive() (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.Receive(&_UchaildErc20Proxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UchaildErc20Proxy *UchaildErc20ProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _UchaildErc20Proxy.Contract.Receive(&_UchaildErc20Proxy.TransactOpts)
}

// UchaildErc20ProxyProxyOwnerUpdateIterator is returned from FilterProxyOwnerUpdate and is used to iterate over the raw logs and unpacked data for ProxyOwnerUpdate events raised by the UchaildErc20Proxy contract.
type UchaildErc20ProxyProxyOwnerUpdateIterator struct {
	Event *UchaildErc20ProxyProxyOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *UchaildErc20ProxyProxyOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UchaildErc20ProxyProxyOwnerUpdate)
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
		it.Event = new(UchaildErc20ProxyProxyOwnerUpdate)
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
func (it *UchaildErc20ProxyProxyOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UchaildErc20ProxyProxyOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UchaildErc20ProxyProxyOwnerUpdate represents a ProxyOwnerUpdate event raised by the UchaildErc20Proxy contract.
type UchaildErc20ProxyProxyOwnerUpdate struct {
	New common.Address
	Old common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProxyOwnerUpdate is a free log retrieval operation binding the contract event 0xdbe5fd65bcdbae152f24ab660ea68e72b4d4705b57b16e0caae994e214680ee2.
//
// Solidity: event ProxyOwnerUpdate(address _new, address _old)
func (_UchaildErc20Proxy *UchaildErc20ProxyFilterer) FilterProxyOwnerUpdate(opts *bind.FilterOpts) (*UchaildErc20ProxyProxyOwnerUpdateIterator, error) {

	logs, sub, err := _UchaildErc20Proxy.contract.FilterLogs(opts, "ProxyOwnerUpdate")
	if err != nil {
		return nil, err
	}
	return &UchaildErc20ProxyProxyOwnerUpdateIterator{contract: _UchaildErc20Proxy.contract, event: "ProxyOwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchProxyOwnerUpdate is a free log subscription operation binding the contract event 0xdbe5fd65bcdbae152f24ab660ea68e72b4d4705b57b16e0caae994e214680ee2.
//
// Solidity: event ProxyOwnerUpdate(address _new, address _old)
func (_UchaildErc20Proxy *UchaildErc20ProxyFilterer) WatchProxyOwnerUpdate(opts *bind.WatchOpts, sink chan<- *UchaildErc20ProxyProxyOwnerUpdate) (event.Subscription, error) {

	logs, sub, err := _UchaildErc20Proxy.contract.WatchLogs(opts, "ProxyOwnerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UchaildErc20ProxyProxyOwnerUpdate)
				if err := _UchaildErc20Proxy.contract.UnpackLog(event, "ProxyOwnerUpdate", log); err != nil {
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

// ParseProxyOwnerUpdate is a log parse operation binding the contract event 0xdbe5fd65bcdbae152f24ab660ea68e72b4d4705b57b16e0caae994e214680ee2.
//
// Solidity: event ProxyOwnerUpdate(address _new, address _old)
func (_UchaildErc20Proxy *UchaildErc20ProxyFilterer) ParseProxyOwnerUpdate(log types.Log) (*UchaildErc20ProxyProxyOwnerUpdate, error) {
	event := new(UchaildErc20ProxyProxyOwnerUpdate)
	if err := _UchaildErc20Proxy.contract.UnpackLog(event, "ProxyOwnerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UchaildErc20ProxyProxyUpdatedIterator is returned from FilterProxyUpdated and is used to iterate over the raw logs and unpacked data for ProxyUpdated events raised by the UchaildErc20Proxy contract.
type UchaildErc20ProxyProxyUpdatedIterator struct {
	Event *UchaildErc20ProxyProxyUpdated // Event containing the contract specifics and raw log

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
func (it *UchaildErc20ProxyProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UchaildErc20ProxyProxyUpdated)
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
		it.Event = new(UchaildErc20ProxyProxyUpdated)
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
func (it *UchaildErc20ProxyProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UchaildErc20ProxyProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UchaildErc20ProxyProxyUpdated represents a ProxyUpdated event raised by the UchaildErc20Proxy contract.
type UchaildErc20ProxyProxyUpdated struct {
	New common.Address
	Old common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProxyUpdated is a free log retrieval operation binding the contract event 0xd32d24edea94f55e932d9a008afc425a8561462d1b1f57bc6e508e9a6b9509e1.
//
// Solidity: event ProxyUpdated(address indexed _new, address indexed _old)
func (_UchaildErc20Proxy *UchaildErc20ProxyFilterer) FilterProxyUpdated(opts *bind.FilterOpts, _new []common.Address, _old []common.Address) (*UchaildErc20ProxyProxyUpdatedIterator, error) {

	var _newRule []interface{}
	for _, _newItem := range _new {
		_newRule = append(_newRule, _newItem)
	}
	var _oldRule []interface{}
	for _, _oldItem := range _old {
		_oldRule = append(_oldRule, _oldItem)
	}

	logs, sub, err := _UchaildErc20Proxy.contract.FilterLogs(opts, "ProxyUpdated", _newRule, _oldRule)
	if err != nil {
		return nil, err
	}
	return &UchaildErc20ProxyProxyUpdatedIterator{contract: _UchaildErc20Proxy.contract, event: "ProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyUpdated is a free log subscription operation binding the contract event 0xd32d24edea94f55e932d9a008afc425a8561462d1b1f57bc6e508e9a6b9509e1.
//
// Solidity: event ProxyUpdated(address indexed _new, address indexed _old)
func (_UchaildErc20Proxy *UchaildErc20ProxyFilterer) WatchProxyUpdated(opts *bind.WatchOpts, sink chan<- *UchaildErc20ProxyProxyUpdated, _new []common.Address, _old []common.Address) (event.Subscription, error) {

	var _newRule []interface{}
	for _, _newItem := range _new {
		_newRule = append(_newRule, _newItem)
	}
	var _oldRule []interface{}
	for _, _oldItem := range _old {
		_oldRule = append(_oldRule, _oldItem)
	}

	logs, sub, err := _UchaildErc20Proxy.contract.WatchLogs(opts, "ProxyUpdated", _newRule, _oldRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UchaildErc20ProxyProxyUpdated)
				if err := _UchaildErc20Proxy.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
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

// ParseProxyUpdated is a log parse operation binding the contract event 0xd32d24edea94f55e932d9a008afc425a8561462d1b1f57bc6e508e9a6b9509e1.
//
// Solidity: event ProxyUpdated(address indexed _new, address indexed _old)
func (_UchaildErc20Proxy *UchaildErc20ProxyFilterer) ParseProxyUpdated(log types.Log) (*UchaildErc20ProxyProxyUpdated, error) {
	event := new(UchaildErc20ProxyProxyUpdated)
	if err := _UchaildErc20Proxy.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polymarket_contracts

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	TypeCode uint8
	To       common.Address
	Value    *big.Int
	Data     []byte
}

// ProxyWalletFactoryMetaData contains all meta data concerning the ProxyWalletFactory contract.
var ProxyWalletFactoryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"typeCode\",\"type\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"proxy\",\"outputs\":[{\"name\":\"returnValues\",\"type\":\"bytes[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"cloneConstructor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHubAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"preRelayedCall\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"acceptRelayedCall\",\"outputs\":[{\"name\":\"doCall\",\"type\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayHubVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gsnModule\",\"type\":\"address\"}],\"name\":\"setGSNModule\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"context\",\"type\":\"bytes\"},{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"actualCharge\",\"type\":\"uint256\"},{\"name\":\"preRetVal\",\"type\":\"bytes32\"}],\"name\":\"postRelayedCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGSNModule\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldRelayHub\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newRelayHub\",\"type\":\"address\"}],\"name\":\"RelayHubChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
}

// ProxyWalletFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyWalletFactoryMetaData.ABI instead.
var ProxyWalletFactoryABI = ProxyWalletFactoryMetaData.ABI

// ProxyWalletFactory is an auto generated Go binding around an Ethereum contract.
type ProxyWalletFactory struct {
	ProxyWalletFactoryCaller     // Read-only binding to the contract
	ProxyWalletFactoryTransactor // Write-only binding to the contract
	ProxyWalletFactoryFilterer   // Log filterer for contract events
}

// ProxyWalletFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyWalletFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyWalletFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyWalletFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyWalletFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyWalletFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyWalletFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyWalletFactorySession struct {
	Contract     *ProxyWalletFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProxyWalletFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyWalletFactoryCallerSession struct {
	Contract *ProxyWalletFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ProxyWalletFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyWalletFactoryTransactorSession struct {
	Contract     *ProxyWalletFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ProxyWalletFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyWalletFactoryRaw struct {
	Contract *ProxyWalletFactory // Generic contract binding to access the raw methods on
}

// ProxyWalletFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyWalletFactoryCallerRaw struct {
	Contract *ProxyWalletFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyWalletFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyWalletFactoryTransactorRaw struct {
	Contract *ProxyWalletFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyWalletFactory creates a new instance of ProxyWalletFactory, bound to a specific deployed contract.
func NewProxyWalletFactory(address common.Address, backend bind.ContractBackend) (*ProxyWalletFactory, error) {
	contract, err := bindProxyWalletFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyWalletFactory{ProxyWalletFactoryCaller: ProxyWalletFactoryCaller{contract: contract}, ProxyWalletFactoryTransactor: ProxyWalletFactoryTransactor{contract: contract}, ProxyWalletFactoryFilterer: ProxyWalletFactoryFilterer{contract: contract}}, nil
}

// NewProxyWalletFactoryCaller creates a new read-only instance of ProxyWalletFactory, bound to a specific deployed contract.
func NewProxyWalletFactoryCaller(address common.Address, caller bind.ContractCaller) (*ProxyWalletFactoryCaller, error) {
	contract, err := bindProxyWalletFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyWalletFactoryCaller{contract: contract}, nil
}

// NewProxyWalletFactoryTransactor creates a new write-only instance of ProxyWalletFactory, bound to a specific deployed contract.
func NewProxyWalletFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyWalletFactoryTransactor, error) {
	contract, err := bindProxyWalletFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyWalletFactoryTransactor{contract: contract}, nil
}

// NewProxyWalletFactoryFilterer creates a new log filterer instance of ProxyWalletFactory, bound to a specific deployed contract.
func NewProxyWalletFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyWalletFactoryFilterer, error) {
	contract, err := bindProxyWalletFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyWalletFactoryFilterer{contract: contract}, nil
}

// bindProxyWalletFactory binds a generic wrapper to an already deployed contract.
func bindProxyWalletFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyWalletFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyWalletFactory *ProxyWalletFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyWalletFactory.Contract.ProxyWalletFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyWalletFactory *ProxyWalletFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.ProxyWalletFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyWalletFactory *ProxyWalletFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.ProxyWalletFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyWalletFactory *ProxyWalletFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyWalletFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.contract.Transact(opts, method, params...)
}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address , address , bytes , uint256 , uint256 , uint256 , uint256 , bytes , uint256 ) view returns(uint256 doCall, bytes context)
func (_ProxyWalletFactory *ProxyWalletFactoryCaller) AcceptRelayedCall(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []byte, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte, arg8 *big.Int) (struct {
	DoCall  *big.Int
	Context []byte
}, error) {
	var out []interface{}
	err := _ProxyWalletFactory.contract.Call(opts, &out, "acceptRelayedCall", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)

	outstruct := new(struct {
		DoCall  *big.Int
		Context []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DoCall = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Context = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address , address , bytes , uint256 , uint256 , uint256 , uint256 , bytes , uint256 ) view returns(uint256 doCall, bytes context)
func (_ProxyWalletFactory *ProxyWalletFactorySession) AcceptRelayedCall(arg0 common.Address, arg1 common.Address, arg2 []byte, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte, arg8 *big.Int) (struct {
	DoCall  *big.Int
	Context []byte
}, error) {
	return _ProxyWalletFactory.Contract.AcceptRelayedCall(&_ProxyWalletFactory.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address , address , bytes , uint256 , uint256 , uint256 , uint256 , bytes , uint256 ) view returns(uint256 doCall, bytes context)
func (_ProxyWalletFactory *ProxyWalletFactoryCallerSession) AcceptRelayedCall(arg0 common.Address, arg1 common.Address, arg2 []byte, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte, arg8 *big.Int) (struct {
	DoCall  *big.Int
	Context []byte
}, error) {
	return _ProxyWalletFactory.Contract.AcceptRelayedCall(&_ProxyWalletFactory.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// GetGSNModule is a free data retrieval call binding the contract method 0xfb480edb.
//
// Solidity: function getGSNModule() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactoryCaller) GetGSNModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyWalletFactory.contract.Call(opts, &out, "getGSNModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGSNModule is a free data retrieval call binding the contract method 0xfb480edb.
//
// Solidity: function getGSNModule() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactorySession) GetGSNModule() (common.Address, error) {
	return _ProxyWalletFactory.Contract.GetGSNModule(&_ProxyWalletFactory.CallOpts)
}

// GetGSNModule is a free data retrieval call binding the contract method 0xfb480edb.
//
// Solidity: function getGSNModule() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactoryCallerSession) GetGSNModule() (common.Address, error) {
	return _ProxyWalletFactory.Contract.GetGSNModule(&_ProxyWalletFactory.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactoryCaller) GetHubAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyWalletFactory.contract.Call(opts, &out, "getHubAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactorySession) GetHubAddr() (common.Address, error) {
	return _ProxyWalletFactory.Contract.GetHubAddr(&_ProxyWalletFactory.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactoryCallerSession) GetHubAddr() (common.Address, error) {
	return _ProxyWalletFactory.Contract.GetHubAddr(&_ProxyWalletFactory.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactoryCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyWalletFactory.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactorySession) GetImplementation() (common.Address, error) {
	return _ProxyWalletFactory.Contract.GetImplementation(&_ProxyWalletFactory.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactoryCallerSession) GetImplementation() (common.Address, error) {
	return _ProxyWalletFactory.Contract.GetImplementation(&_ProxyWalletFactory.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ProxyWalletFactory *ProxyWalletFactoryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ProxyWalletFactory.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ProxyWalletFactory *ProxyWalletFactorySession) IsOwner() (bool, error) {
	return _ProxyWalletFactory.Contract.IsOwner(&_ProxyWalletFactory.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ProxyWalletFactory *ProxyWalletFactoryCallerSession) IsOwner() (bool, error) {
	return _ProxyWalletFactory.Contract.IsOwner(&_ProxyWalletFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyWalletFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactorySession) Owner() (common.Address, error) {
	return _ProxyWalletFactory.Contract.Owner(&_ProxyWalletFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyWalletFactory *ProxyWalletFactoryCallerSession) Owner() (common.Address, error) {
	return _ProxyWalletFactory.Contract.Owner(&_ProxyWalletFactory.CallOpts)
}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_ProxyWalletFactory *ProxyWalletFactoryCaller) RelayHubVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProxyWalletFactory.contract.Call(opts, &out, "relayHubVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_ProxyWalletFactory *ProxyWalletFactorySession) RelayHubVersion() (string, error) {
	return _ProxyWalletFactory.Contract.RelayHubVersion(&_ProxyWalletFactory.CallOpts)
}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_ProxyWalletFactory *ProxyWalletFactoryCallerSession) RelayHubVersion() (string, error) {
	return _ProxyWalletFactory.Contract.RelayHubVersion(&_ProxyWalletFactory.CallOpts)
}

// CloneConstructor is a paid mutator transaction binding the contract method 0x52e831dd.
//
// Solidity: function cloneConstructor(bytes ) returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactor) CloneConstructor(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.contract.Transact(opts, "cloneConstructor", arg0)
}

// CloneConstructor is a paid mutator transaction binding the contract method 0x52e831dd.
//
// Solidity: function cloneConstructor(bytes ) returns()
func (_ProxyWalletFactory *ProxyWalletFactorySession) CloneConstructor(arg0 []byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.CloneConstructor(&_ProxyWalletFactory.TransactOpts, arg0)
}

// CloneConstructor is a paid mutator transaction binding the contract method 0x52e831dd.
//
// Solidity: function cloneConstructor(bytes ) returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorSession) CloneConstructor(arg0 []byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.CloneConstructor(&_ProxyWalletFactory.TransactOpts, arg0)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactor) PostRelayedCall(opts *bind.TransactOpts, context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.contract.Transact(opts, "postRelayedCall", context, success, actualCharge, preRetVal)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_ProxyWalletFactory *ProxyWalletFactorySession) PostRelayedCall(context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.PostRelayedCall(&_ProxyWalletFactory.TransactOpts, context, success, actualCharge, preRetVal)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorSession) PostRelayedCall(context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.PostRelayedCall(&_ProxyWalletFactory.TransactOpts, context, success, actualCharge, preRetVal)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_ProxyWalletFactory *ProxyWalletFactoryTransactor) PreRelayedCall(opts *bind.TransactOpts, context []byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.contract.Transact(opts, "preRelayedCall", context)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_ProxyWalletFactory *ProxyWalletFactorySession) PreRelayedCall(context []byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.PreRelayedCall(&_ProxyWalletFactory.TransactOpts, context)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorSession) PreRelayedCall(context []byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.PreRelayedCall(&_ProxyWalletFactory.TransactOpts, context)
}

// Proxy is a paid mutator transaction binding the contract method 0x34ee9791.
//
// Solidity: function proxy((uint8,address,uint256,bytes)[] calls) payable returns(bytes[] returnValues)
func (_ProxyWalletFactory *ProxyWalletFactoryTransactor) Proxy(opts *bind.TransactOpts, calls []Struct0) (*types.Transaction, error) {
	return _ProxyWalletFactory.contract.Transact(opts, "proxy", calls)
}

// Proxy is a paid mutator transaction binding the contract method 0x34ee9791.
//
// Solidity: function proxy((uint8,address,uint256,bytes)[] calls) payable returns(bytes[] returnValues)
func (_ProxyWalletFactory *ProxyWalletFactorySession) Proxy(calls []Struct0) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.Proxy(&_ProxyWalletFactory.TransactOpts, calls)
}

// Proxy is a paid mutator transaction binding the contract method 0x34ee9791.
//
// Solidity: function proxy((uint8,address,uint256,bytes)[] calls) payable returns(bytes[] returnValues)
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorSession) Proxy(calls []Struct0) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.Proxy(&_ProxyWalletFactory.TransactOpts, calls)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyWalletFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyWalletFactory *ProxyWalletFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.RenounceOwnership(&_ProxyWalletFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.RenounceOwnership(&_ProxyWalletFactory.TransactOpts)
}

// SetGSNModule is a paid mutator transaction binding the contract method 0xc40b1532.
//
// Solidity: function setGSNModule(address gsnModule) returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactor) SetGSNModule(opts *bind.TransactOpts, gsnModule common.Address) (*types.Transaction, error) {
	return _ProxyWalletFactory.contract.Transact(opts, "setGSNModule", gsnModule)
}

// SetGSNModule is a paid mutator transaction binding the contract method 0xc40b1532.
//
// Solidity: function setGSNModule(address gsnModule) returns()
func (_ProxyWalletFactory *ProxyWalletFactorySession) SetGSNModule(gsnModule common.Address) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.SetGSNModule(&_ProxyWalletFactory.TransactOpts, gsnModule)
}

// SetGSNModule is a paid mutator transaction binding the contract method 0xc40b1532.
//
// Solidity: function setGSNModule(address gsnModule) returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorSession) SetGSNModule(gsnModule common.Address) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.SetGSNModule(&_ProxyWalletFactory.TransactOpts, gsnModule)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProxyWalletFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyWalletFactory *ProxyWalletFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.TransferOwnership(&_ProxyWalletFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.TransferOwnership(&_ProxyWalletFactory.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ProxyWalletFactory *ProxyWalletFactorySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.Fallback(&_ProxyWalletFactory.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ProxyWalletFactory *ProxyWalletFactoryTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ProxyWalletFactory.Contract.Fallback(&_ProxyWalletFactory.TransactOpts, calldata)
}

// ProxyWalletFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProxyWalletFactory contract.
type ProxyWalletFactoryOwnershipTransferredIterator struct {
	Event *ProxyWalletFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProxyWalletFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyWalletFactoryOwnershipTransferred)
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
		it.Event = new(ProxyWalletFactoryOwnershipTransferred)
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
func (it *ProxyWalletFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyWalletFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyWalletFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the ProxyWalletFactory contract.
type ProxyWalletFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyWalletFactory *ProxyWalletFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProxyWalletFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyWalletFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyWalletFactoryOwnershipTransferredIterator{contract: _ProxyWalletFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyWalletFactory *ProxyWalletFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyWalletFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyWalletFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyWalletFactoryOwnershipTransferred)
				if err := _ProxyWalletFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProxyWalletFactory *ProxyWalletFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*ProxyWalletFactoryOwnershipTransferred, error) {
	event := new(ProxyWalletFactoryOwnershipTransferred)
	if err := _ProxyWalletFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyWalletFactoryRelayHubChangedIterator is returned from FilterRelayHubChanged and is used to iterate over the raw logs and unpacked data for RelayHubChanged events raised by the ProxyWalletFactory contract.
type ProxyWalletFactoryRelayHubChangedIterator struct {
	Event *ProxyWalletFactoryRelayHubChanged // Event containing the contract specifics and raw log

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
func (it *ProxyWalletFactoryRelayHubChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyWalletFactoryRelayHubChanged)
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
		it.Event = new(ProxyWalletFactoryRelayHubChanged)
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
func (it *ProxyWalletFactoryRelayHubChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyWalletFactoryRelayHubChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyWalletFactoryRelayHubChanged represents a RelayHubChanged event raised by the ProxyWalletFactory contract.
type ProxyWalletFactoryRelayHubChanged struct {
	OldRelayHub common.Address
	NewRelayHub common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayHubChanged is a free log retrieval operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_ProxyWalletFactory *ProxyWalletFactoryFilterer) FilterRelayHubChanged(opts *bind.FilterOpts, oldRelayHub []common.Address, newRelayHub []common.Address) (*ProxyWalletFactoryRelayHubChangedIterator, error) {

	var oldRelayHubRule []interface{}
	for _, oldRelayHubItem := range oldRelayHub {
		oldRelayHubRule = append(oldRelayHubRule, oldRelayHubItem)
	}
	var newRelayHubRule []interface{}
	for _, newRelayHubItem := range newRelayHub {
		newRelayHubRule = append(newRelayHubRule, newRelayHubItem)
	}

	logs, sub, err := _ProxyWalletFactory.contract.FilterLogs(opts, "RelayHubChanged", oldRelayHubRule, newRelayHubRule)
	if err != nil {
		return nil, err
	}
	return &ProxyWalletFactoryRelayHubChangedIterator{contract: _ProxyWalletFactory.contract, event: "RelayHubChanged", logs: logs, sub: sub}, nil
}

// WatchRelayHubChanged is a free log subscription operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_ProxyWalletFactory *ProxyWalletFactoryFilterer) WatchRelayHubChanged(opts *bind.WatchOpts, sink chan<- *ProxyWalletFactoryRelayHubChanged, oldRelayHub []common.Address, newRelayHub []common.Address) (event.Subscription, error) {

	var oldRelayHubRule []interface{}
	for _, oldRelayHubItem := range oldRelayHub {
		oldRelayHubRule = append(oldRelayHubRule, oldRelayHubItem)
	}
	var newRelayHubRule []interface{}
	for _, newRelayHubItem := range newRelayHub {
		newRelayHubRule = append(newRelayHubRule, newRelayHubItem)
	}

	logs, sub, err := _ProxyWalletFactory.contract.WatchLogs(opts, "RelayHubChanged", oldRelayHubRule, newRelayHubRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyWalletFactoryRelayHubChanged)
				if err := _ProxyWalletFactory.contract.UnpackLog(event, "RelayHubChanged", log); err != nil {
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

// ParseRelayHubChanged is a log parse operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_ProxyWalletFactory *ProxyWalletFactoryFilterer) ParseRelayHubChanged(log types.Log) (*ProxyWalletFactoryRelayHubChanged, error) {
	event := new(ProxyWalletFactoryRelayHubChanged)
	if err := _ProxyWalletFactory.contract.UnpackLog(event, "RelayHubChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

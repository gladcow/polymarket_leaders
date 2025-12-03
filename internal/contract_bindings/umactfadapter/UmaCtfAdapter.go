// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package umactfadapter

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

// AncillaryDataUpdate is an auto generated low-level Go binding around an user-defined struct.
type AncillaryDataUpdate struct {
	Timestamp *big.Int
	Update    []byte
}

// QuestionData2 is an auto generated low-level Go binding around an user-defined struct.
type QuestionData struct {
	RequestTimestamp          *big.Int
	Reward                    *big.Int
	ProposalBond              *big.Int
	Liveness                  *big.Int
	ManualResolutionTimestamp *big.Int
	Resolved                  bool
	Paused                    bool
	Reset                     bool
	Refund                    bool
	RewardToken               common.Address
	Creator                   common.Address
	AncillaryData             []byte
}

// UmaCtfAdapterMetaData contains all meta data concerning the UmaCtfAdapter contract.
var UmaCtfAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ctf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_finder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Flagged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAncillaryData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOOPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayouts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFlagged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOptimisticOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotReadyToResolve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Resolved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafetyPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafetyPeriodPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"update\",\"type\":\"bytes\"}],\"name\":\"AncillaryDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"QuestionFlagged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ancillaryData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalBond\",\"type\":\"uint256\"}],\"name\":\"QuestionInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payouts\",\"type\":\"uint256[]\"}],\"name\":\"QuestionManuallyResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"QuestionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"QuestionReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"settledPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payouts\",\"type\":\"uint256[]\"}],\"name\":\"QuestionResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"QuestionUnflagged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"QuestionUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ANCILLARY_DATA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SAFETY_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"YES_OR_NO_IDENTIFIER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralWhitelist\",\"outputs\":[{\"internalType\":\"contractIAddressWhitelist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctf\",\"outputs\":[{\"internalType\":\"contractIConditionalTokens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"flag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"getExpectedPayouts\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getLatestUpdate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"update\",\"type\":\"bytes\"}],\"internalType\":\"structAncillaryDataUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"getQuestion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liveness\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"manualResolutionTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resolved\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reset\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"refund\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ancillaryData\",\"type\":\"bytes\"}],\"internalType\":\"structQuestionData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getUpdates\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"update\",\"type\":\"bytes\"}],\"internalType\":\"structAncillaryDataUpdate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"ancillaryData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liveness\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"isFlagged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticOracle\",\"outputs\":[{\"internalType\":\"contractIOptimisticOracleV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"update\",\"type\":\"bytes\"}],\"name\":\"postUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"ancillaryData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"priceDisputed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"questions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liveness\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"manualResolutionTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resolved\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reset\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"refund\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ancillaryData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"resolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"payouts\",\"type\":\"uint256[]\"}],\"name\":\"resolveManually\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"unflag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"updates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"update\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UmaCtfAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use UmaCtfAdapterMetaData.ABI instead.
var UmaCtfAdapterABI = UmaCtfAdapterMetaData.ABI

// UmaCtfAdapter is an auto generated Go binding around an Ethereum contract.
type UmaCtfAdapter struct {
	UmaCtfAdapterCaller     // Read-only binding to the contract
	UmaCtfAdapterTransactor // Write-only binding to the contract
	UmaCtfAdapterFilterer   // Log filterer for contract events
}

// UmaCtfAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UmaCtfAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UmaCtfAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UmaCtfAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UmaCtfAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UmaCtfAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UmaCtfAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UmaCtfAdapterSession struct {
	Contract     *UmaCtfAdapter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UmaCtfAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UmaCtfAdapterCallerSession struct {
	Contract *UmaCtfAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UmaCtfAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UmaCtfAdapterTransactorSession struct {
	Contract     *UmaCtfAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UmaCtfAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UmaCtfAdapterRaw struct {
	Contract *UmaCtfAdapter // Generic contract binding to access the raw methods on
}

// UmaCtfAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UmaCtfAdapterCallerRaw struct {
	Contract *UmaCtfAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// UmaCtfAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UmaCtfAdapterTransactorRaw struct {
	Contract *UmaCtfAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUmaCtfAdapter creates a new instance of UmaCtfAdapter, bound to a specific deployed contract.
func NewUmaCtfAdapter(address common.Address, backend bind.ContractBackend) (*UmaCtfAdapter, error) {
	contract, err := bindUmaCtfAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapter{UmaCtfAdapterCaller: UmaCtfAdapterCaller{contract: contract}, UmaCtfAdapterTransactor: UmaCtfAdapterTransactor{contract: contract}, UmaCtfAdapterFilterer: UmaCtfAdapterFilterer{contract: contract}}, nil
}

// NewUmaCtfAdapterCaller creates a new read-only instance of UmaCtfAdapter, bound to a specific deployed contract.
func NewUmaCtfAdapterCaller(address common.Address, caller bind.ContractCaller) (*UmaCtfAdapterCaller, error) {
	contract, err := bindUmaCtfAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterCaller{contract: contract}, nil
}

// NewUmaCtfAdapterTransactor creates a new write-only instance of UmaCtfAdapter, bound to a specific deployed contract.
func NewUmaCtfAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*UmaCtfAdapterTransactor, error) {
	contract, err := bindUmaCtfAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterTransactor{contract: contract}, nil
}

// NewUmaCtfAdapterFilterer creates a new log filterer instance of UmaCtfAdapter, bound to a specific deployed contract.
func NewUmaCtfAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*UmaCtfAdapterFilterer, error) {
	contract, err := bindUmaCtfAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterFilterer{contract: contract}, nil
}

// bindUmaCtfAdapter binds a generic wrapper to an already deployed contract.
func bindUmaCtfAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UmaCtfAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UmaCtfAdapter *UmaCtfAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UmaCtfAdapter.Contract.UmaCtfAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UmaCtfAdapter *UmaCtfAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.UmaCtfAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UmaCtfAdapter *UmaCtfAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.UmaCtfAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UmaCtfAdapter *UmaCtfAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UmaCtfAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UmaCtfAdapter *UmaCtfAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UmaCtfAdapter *UmaCtfAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.contract.Transact(opts, method, params...)
}

// MAXANCILLARYDATA is a free data retrieval call binding the contract method 0x27f8feac.
//
// Solidity: function MAX_ANCILLARY_DATA() view returns(uint256)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) MAXANCILLARYDATA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "MAX_ANCILLARY_DATA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXANCILLARYDATA is a free data retrieval call binding the contract method 0x27f8feac.
//
// Solidity: function MAX_ANCILLARY_DATA() view returns(uint256)
func (_UmaCtfAdapter *UmaCtfAdapterSession) MAXANCILLARYDATA() (*big.Int, error) {
	return _UmaCtfAdapter.Contract.MAXANCILLARYDATA(&_UmaCtfAdapter.CallOpts)
}

// MAXANCILLARYDATA is a free data retrieval call binding the contract method 0x27f8feac.
//
// Solidity: function MAX_ANCILLARY_DATA() view returns(uint256)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) MAXANCILLARYDATA() (*big.Int, error) {
	return _UmaCtfAdapter.Contract.MAXANCILLARYDATA(&_UmaCtfAdapter.CallOpts)
}

// SAFETYPERIOD is a free data retrieval call binding the contract method 0xd1dfb2e9.
//
// Solidity: function SAFETY_PERIOD() view returns(uint256)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) SAFETYPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "SAFETY_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SAFETYPERIOD is a free data retrieval call binding the contract method 0xd1dfb2e9.
//
// Solidity: function SAFETY_PERIOD() view returns(uint256)
func (_UmaCtfAdapter *UmaCtfAdapterSession) SAFETYPERIOD() (*big.Int, error) {
	return _UmaCtfAdapter.Contract.SAFETYPERIOD(&_UmaCtfAdapter.CallOpts)
}

// SAFETYPERIOD is a free data retrieval call binding the contract method 0xd1dfb2e9.
//
// Solidity: function SAFETY_PERIOD() view returns(uint256)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) SAFETYPERIOD() (*big.Int, error) {
	return _UmaCtfAdapter.Contract.SAFETYPERIOD(&_UmaCtfAdapter.CallOpts)
}

// YESORNOIDENTIFIER is a free data retrieval call binding the contract method 0x6c66f07d.
//
// Solidity: function YES_OR_NO_IDENTIFIER() view returns(bytes32)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) YESORNOIDENTIFIER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "YES_OR_NO_IDENTIFIER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// YESORNOIDENTIFIER is a free data retrieval call binding the contract method 0x6c66f07d.
//
// Solidity: function YES_OR_NO_IDENTIFIER() view returns(bytes32)
func (_UmaCtfAdapter *UmaCtfAdapterSession) YESORNOIDENTIFIER() ([32]byte, error) {
	return _UmaCtfAdapter.Contract.YESORNOIDENTIFIER(&_UmaCtfAdapter.CallOpts)
}

// YESORNOIDENTIFIER is a free data retrieval call binding the contract method 0x6c66f07d.
//
// Solidity: function YES_OR_NO_IDENTIFIER() view returns(bytes32)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) YESORNOIDENTIFIER() ([32]byte, error) {
	return _UmaCtfAdapter.Contract.YESORNOIDENTIFIER(&_UmaCtfAdapter.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_UmaCtfAdapter *UmaCtfAdapterSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _UmaCtfAdapter.Contract.Admins(&_UmaCtfAdapter.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _UmaCtfAdapter.Contract.Admins(&_UmaCtfAdapter.CallOpts, arg0)
}

// CollateralWhitelist is a free data retrieval call binding the contract method 0xe4ee614a.
//
// Solidity: function collateralWhitelist() view returns(address)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) CollateralWhitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "collateralWhitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralWhitelist is a free data retrieval call binding the contract method 0xe4ee614a.
//
// Solidity: function collateralWhitelist() view returns(address)
func (_UmaCtfAdapter *UmaCtfAdapterSession) CollateralWhitelist() (common.Address, error) {
	return _UmaCtfAdapter.Contract.CollateralWhitelist(&_UmaCtfAdapter.CallOpts)
}

// CollateralWhitelist is a free data retrieval call binding the contract method 0xe4ee614a.
//
// Solidity: function collateralWhitelist() view returns(address)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) CollateralWhitelist() (common.Address, error) {
	return _UmaCtfAdapter.Contract.CollateralWhitelist(&_UmaCtfAdapter.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_UmaCtfAdapter *UmaCtfAdapterSession) Ctf() (common.Address, error) {
	return _UmaCtfAdapter.Contract.Ctf(&_UmaCtfAdapter.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) Ctf() (common.Address, error) {
	return _UmaCtfAdapter.Contract.Ctf(&_UmaCtfAdapter.CallOpts)
}

// GetExpectedPayouts is a free data retrieval call binding the contract method 0x34e5e28e.
//
// Solidity: function getExpectedPayouts(bytes32 questionID) view returns(uint256[])
func (_UmaCtfAdapter *UmaCtfAdapterCaller) GetExpectedPayouts(opts *bind.CallOpts, questionID [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "getExpectedPayouts", questionID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetExpectedPayouts is a free data retrieval call binding the contract method 0x34e5e28e.
//
// Solidity: function getExpectedPayouts(bytes32 questionID) view returns(uint256[])
func (_UmaCtfAdapter *UmaCtfAdapterSession) GetExpectedPayouts(questionID [32]byte) ([]*big.Int, error) {
	return _UmaCtfAdapter.Contract.GetExpectedPayouts(&_UmaCtfAdapter.CallOpts, questionID)
}

// GetExpectedPayouts is a free data retrieval call binding the contract method 0x34e5e28e.
//
// Solidity: function getExpectedPayouts(bytes32 questionID) view returns(uint256[])
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) GetExpectedPayouts(questionID [32]byte) ([]*big.Int, error) {
	return _UmaCtfAdapter.Contract.GetExpectedPayouts(&_UmaCtfAdapter.CallOpts, questionID)
}

// GetLatestUpdate is a free data retrieval call binding the contract method 0xc0cab0a2.
//
// Solidity: function getLatestUpdate(bytes32 questionID, address owner) view returns((uint256,bytes))
func (_UmaCtfAdapter *UmaCtfAdapterCaller) GetLatestUpdate(opts *bind.CallOpts, questionID [32]byte, owner common.Address) (AncillaryDataUpdate, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "getLatestUpdate", questionID, owner)

	if err != nil {
		return *new(AncillaryDataUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(AncillaryDataUpdate)).(*AncillaryDataUpdate)

	return out0, err

}

// GetLatestUpdate is a free data retrieval call binding the contract method 0xc0cab0a2.
//
// Solidity: function getLatestUpdate(bytes32 questionID, address owner) view returns((uint256,bytes))
func (_UmaCtfAdapter *UmaCtfAdapterSession) GetLatestUpdate(questionID [32]byte, owner common.Address) (AncillaryDataUpdate, error) {
	return _UmaCtfAdapter.Contract.GetLatestUpdate(&_UmaCtfAdapter.CallOpts, questionID, owner)
}

// GetLatestUpdate is a free data retrieval call binding the contract method 0xc0cab0a2.
//
// Solidity: function getLatestUpdate(bytes32 questionID, address owner) view returns((uint256,bytes))
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) GetLatestUpdate(questionID [32]byte, owner common.Address) (AncillaryDataUpdate, error) {
	return _UmaCtfAdapter.Contract.GetLatestUpdate(&_UmaCtfAdapter.CallOpts, questionID, owner)
}

// GetQuestion is a free data retrieval call binding the contract method 0x58c039cd.
//
// Solidity: function getQuestion(bytes32 questionID) view returns((uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,address,address,bytes))
func (_UmaCtfAdapter *UmaCtfAdapterCaller) GetQuestion(opts *bind.CallOpts, questionID [32]byte) (QuestionData, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "getQuestion", questionID)

	if err != nil {
		return *new(QuestionData), err
	}

	out0 := *abi.ConvertType(out[0], new(QuestionData)).(*QuestionData)

	return out0, err

}

// GetQuestion is a free data retrieval call binding the contract method 0x58c039cd.
//
// Solidity: function getQuestion(bytes32 questionID) view returns((uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,address,address,bytes))
func (_UmaCtfAdapter *UmaCtfAdapterSession) GetQuestion(questionID [32]byte) (QuestionData, error) {
	return _UmaCtfAdapter.Contract.GetQuestion(&_UmaCtfAdapter.CallOpts, questionID)
}

// GetQuestion is a free data retrieval call binding the contract method 0x58c039cd.
//
// Solidity: function getQuestion(bytes32 questionID) view returns((uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,address,address,bytes))
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) GetQuestion(questionID [32]byte) (QuestionData, error) {
	return _UmaCtfAdapter.Contract.GetQuestion(&_UmaCtfAdapter.CallOpts, questionID)
}

// GetUpdates is a free data retrieval call binding the contract method 0x555c56fc.
//
// Solidity: function getUpdates(bytes32 questionID, address owner) view returns((uint256,bytes)[])
func (_UmaCtfAdapter *UmaCtfAdapterCaller) GetUpdates(opts *bind.CallOpts, questionID [32]byte, owner common.Address) ([]AncillaryDataUpdate, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "getUpdates", questionID, owner)

	if err != nil {
		return *new([]AncillaryDataUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new([]AncillaryDataUpdate)).(*[]AncillaryDataUpdate)

	return out0, err

}

// GetUpdates is a free data retrieval call binding the contract method 0x555c56fc.
//
// Solidity: function getUpdates(bytes32 questionID, address owner) view returns((uint256,bytes)[])
func (_UmaCtfAdapter *UmaCtfAdapterSession) GetUpdates(questionID [32]byte, owner common.Address) ([]AncillaryDataUpdate, error) {
	return _UmaCtfAdapter.Contract.GetUpdates(&_UmaCtfAdapter.CallOpts, questionID, owner)
}

// GetUpdates is a free data retrieval call binding the contract method 0x555c56fc.
//
// Solidity: function getUpdates(bytes32 questionID, address owner) view returns((uint256,bytes)[])
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) GetUpdates(questionID [32]byte, owner common.Address) ([]AncillaryDataUpdate, error) {
	return _UmaCtfAdapter.Contract.GetUpdates(&_UmaCtfAdapter.CallOpts, questionID, owner)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "isAdmin", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterSession) IsAdmin(addr common.Address) (bool, error) {
	return _UmaCtfAdapter.Contract.IsAdmin(&_UmaCtfAdapter.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _UmaCtfAdapter.Contract.IsAdmin(&_UmaCtfAdapter.CallOpts, addr)
}

// IsFlagged is a free data retrieval call binding the contract method 0xbf2dde38.
//
// Solidity: function isFlagged(bytes32 questionID) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) IsFlagged(opts *bind.CallOpts, questionID [32]byte) (bool, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "isFlagged", questionID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFlagged is a free data retrieval call binding the contract method 0xbf2dde38.
//
// Solidity: function isFlagged(bytes32 questionID) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterSession) IsFlagged(questionID [32]byte) (bool, error) {
	return _UmaCtfAdapter.Contract.IsFlagged(&_UmaCtfAdapter.CallOpts, questionID)
}

// IsFlagged is a free data retrieval call binding the contract method 0xbf2dde38.
//
// Solidity: function isFlagged(bytes32 questionID) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) IsFlagged(questionID [32]byte) (bool, error) {
	return _UmaCtfAdapter.Contract.IsFlagged(&_UmaCtfAdapter.CallOpts, questionID)
}

// IsInitialized is a free data retrieval call binding the contract method 0xf7b637bb.
//
// Solidity: function isInitialized(bytes32 questionID) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) IsInitialized(opts *bind.CallOpts, questionID [32]byte) (bool, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "isInitialized", questionID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0xf7b637bb.
//
// Solidity: function isInitialized(bytes32 questionID) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterSession) IsInitialized(questionID [32]byte) (bool, error) {
	return _UmaCtfAdapter.Contract.IsInitialized(&_UmaCtfAdapter.CallOpts, questionID)
}

// IsInitialized is a free data retrieval call binding the contract method 0xf7b637bb.
//
// Solidity: function isInitialized(bytes32 questionID) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) IsInitialized(questionID [32]byte) (bool, error) {
	return _UmaCtfAdapter.Contract.IsInitialized(&_UmaCtfAdapter.CallOpts, questionID)
}

// OptimisticOracle is a free data retrieval call binding the contract method 0x22302922.
//
// Solidity: function optimisticOracle() view returns(address)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) OptimisticOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "optimisticOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimisticOracle is a free data retrieval call binding the contract method 0x22302922.
//
// Solidity: function optimisticOracle() view returns(address)
func (_UmaCtfAdapter *UmaCtfAdapterSession) OptimisticOracle() (common.Address, error) {
	return _UmaCtfAdapter.Contract.OptimisticOracle(&_UmaCtfAdapter.CallOpts)
}

// OptimisticOracle is a free data retrieval call binding the contract method 0x22302922.
//
// Solidity: function optimisticOracle() view returns(address)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) OptimisticOracle() (common.Address, error) {
	return _UmaCtfAdapter.Contract.OptimisticOracle(&_UmaCtfAdapter.CallOpts)
}

// Questions is a free data retrieval call binding the contract method 0x95addb90.
//
// Solidity: function questions(bytes32 ) view returns(uint256 requestTimestamp, uint256 reward, uint256 proposalBond, uint256 liveness, uint256 manualResolutionTimestamp, bool resolved, bool paused, bool reset, bool refund, address rewardToken, address creator, bytes ancillaryData)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) Questions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	RequestTimestamp          *big.Int
	Reward                    *big.Int
	ProposalBond              *big.Int
	Liveness                  *big.Int
	ManualResolutionTimestamp *big.Int
	Resolved                  bool
	Paused                    bool
	Reset                     bool
	Refund                    bool
	RewardToken               common.Address
	Creator                   common.Address
	AncillaryData             []byte
}, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "questions", arg0)

	outstruct := new(struct {
		RequestTimestamp          *big.Int
		Reward                    *big.Int
		ProposalBond              *big.Int
		Liveness                  *big.Int
		ManualResolutionTimestamp *big.Int
		Resolved                  bool
		Paused                    bool
		Reset                     bool
		Refund                    bool
		RewardToken               common.Address
		Creator                   common.Address
		AncillaryData             []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProposalBond = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Liveness = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ManualResolutionTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Resolved = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Paused = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Reset = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.Refund = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.RewardToken = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.Creator = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)
	outstruct.AncillaryData = *abi.ConvertType(out[11], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Questions is a free data retrieval call binding the contract method 0x95addb90.
//
// Solidity: function questions(bytes32 ) view returns(uint256 requestTimestamp, uint256 reward, uint256 proposalBond, uint256 liveness, uint256 manualResolutionTimestamp, bool resolved, bool paused, bool reset, bool refund, address rewardToken, address creator, bytes ancillaryData)
func (_UmaCtfAdapter *UmaCtfAdapterSession) Questions(arg0 [32]byte) (struct {
	RequestTimestamp          *big.Int
	Reward                    *big.Int
	ProposalBond              *big.Int
	Liveness                  *big.Int
	ManualResolutionTimestamp *big.Int
	Resolved                  bool
	Paused                    bool
	Reset                     bool
	Refund                    bool
	RewardToken               common.Address
	Creator                   common.Address
	AncillaryData             []byte
}, error) {
	return _UmaCtfAdapter.Contract.Questions(&_UmaCtfAdapter.CallOpts, arg0)
}

// Questions is a free data retrieval call binding the contract method 0x95addb90.
//
// Solidity: function questions(bytes32 ) view returns(uint256 requestTimestamp, uint256 reward, uint256 proposalBond, uint256 liveness, uint256 manualResolutionTimestamp, bool resolved, bool paused, bool reset, bool refund, address rewardToken, address creator, bytes ancillaryData)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) Questions(arg0 [32]byte) (struct {
	RequestTimestamp          *big.Int
	Reward                    *big.Int
	ProposalBond              *big.Int
	Liveness                  *big.Int
	ManualResolutionTimestamp *big.Int
	Resolved                  bool
	Paused                    bool
	Reset                     bool
	Refund                    bool
	RewardToken               common.Address
	Creator                   common.Address
	AncillaryData             []byte
}, error) {
	return _UmaCtfAdapter.Contract.Questions(&_UmaCtfAdapter.CallOpts, arg0)
}

// Ready is a free data retrieval call binding the contract method 0xfcac49a2.
//
// Solidity: function ready(bytes32 questionID) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) Ready(opts *bind.CallOpts, questionID [32]byte) (bool, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "ready", questionID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ready is a free data retrieval call binding the contract method 0xfcac49a2.
//
// Solidity: function ready(bytes32 questionID) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterSession) Ready(questionID [32]byte) (bool, error) {
	return _UmaCtfAdapter.Contract.Ready(&_UmaCtfAdapter.CallOpts, questionID)
}

// Ready is a free data retrieval call binding the contract method 0xfcac49a2.
//
// Solidity: function ready(bytes32 questionID) view returns(bool)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) Ready(questionID [32]byte) (bool, error) {
	return _UmaCtfAdapter.Contract.Ready(&_UmaCtfAdapter.CallOpts, questionID)
}

// Updates is a free data retrieval call binding the contract method 0x89ab0871.
//
// Solidity: function updates(bytes32 , uint256 ) view returns(uint256 timestamp, bytes update)
func (_UmaCtfAdapter *UmaCtfAdapterCaller) Updates(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Update    []byte
}, error) {
	var out []interface{}
	err := _UmaCtfAdapter.contract.Call(opts, &out, "updates", arg0, arg1)

	outstruct := new(struct {
		Timestamp *big.Int
		Update    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Update = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Updates is a free data retrieval call binding the contract method 0x89ab0871.
//
// Solidity: function updates(bytes32 , uint256 ) view returns(uint256 timestamp, bytes update)
func (_UmaCtfAdapter *UmaCtfAdapterSession) Updates(arg0 [32]byte, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Update    []byte
}, error) {
	return _UmaCtfAdapter.Contract.Updates(&_UmaCtfAdapter.CallOpts, arg0, arg1)
}

// Updates is a free data retrieval call binding the contract method 0x89ab0871.
//
// Solidity: function updates(bytes32 , uint256 ) view returns(uint256 timestamp, bytes update)
func (_UmaCtfAdapter *UmaCtfAdapterCallerSession) Updates(arg0 [32]byte, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Update    []byte
}, error) {
	return _UmaCtfAdapter.Contract.Updates(&_UmaCtfAdapter.CallOpts, arg0, arg1)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.AddAdmin(&_UmaCtfAdapter.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.AddAdmin(&_UmaCtfAdapter.TransactOpts, admin)
}

// Flag is a paid mutator transaction binding the contract method 0x78165a48.
//
// Solidity: function flag(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) Flag(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "flag", questionID)
}

// Flag is a paid mutator transaction binding the contract method 0x78165a48.
//
// Solidity: function flag(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) Flag(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Flag(&_UmaCtfAdapter.TransactOpts, questionID)
}

// Flag is a paid mutator transaction binding the contract method 0x78165a48.
//
// Solidity: function flag(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) Flag(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Flag(&_UmaCtfAdapter.TransactOpts, questionID)
}

// Initialize is a paid mutator transaction binding the contract method 0x185d1646.
//
// Solidity: function initialize(bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond, uint256 liveness) returns(bytes32 questionID)
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) Initialize(opts *bind.TransactOpts, ancillaryData []byte, rewardToken common.Address, reward *big.Int, proposalBond *big.Int, liveness *big.Int) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "initialize", ancillaryData, rewardToken, reward, proposalBond, liveness)
}

// Initialize is a paid mutator transaction binding the contract method 0x185d1646.
//
// Solidity: function initialize(bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond, uint256 liveness) returns(bytes32 questionID)
func (_UmaCtfAdapter *UmaCtfAdapterSession) Initialize(ancillaryData []byte, rewardToken common.Address, reward *big.Int, proposalBond *big.Int, liveness *big.Int) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Initialize(&_UmaCtfAdapter.TransactOpts, ancillaryData, rewardToken, reward, proposalBond, liveness)
}

// Initialize is a paid mutator transaction binding the contract method 0x185d1646.
//
// Solidity: function initialize(bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond, uint256 liveness) returns(bytes32 questionID)
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) Initialize(ancillaryData []byte, rewardToken common.Address, reward *big.Int, proposalBond *big.Int, liveness *big.Int) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Initialize(&_UmaCtfAdapter.TransactOpts, ancillaryData, rewardToken, reward, proposalBond, liveness)
}

// Pause is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) Pause(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "pause", questionID)
}

// Pause is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) Pause(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Pause(&_UmaCtfAdapter.TransactOpts, questionID)
}

// Pause is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) Pause(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Pause(&_UmaCtfAdapter.TransactOpts, questionID)
}

// PostUpdate is a paid mutator transaction binding the contract method 0x072d1259.
//
// Solidity: function postUpdate(bytes32 questionID, bytes update) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) PostUpdate(opts *bind.TransactOpts, questionID [32]byte, update []byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "postUpdate", questionID, update)
}

// PostUpdate is a paid mutator transaction binding the contract method 0x072d1259.
//
// Solidity: function postUpdate(bytes32 questionID, bytes update) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) PostUpdate(questionID [32]byte, update []byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.PostUpdate(&_UmaCtfAdapter.TransactOpts, questionID, update)
}

// PostUpdate is a paid mutator transaction binding the contract method 0x072d1259.
//
// Solidity: function postUpdate(bytes32 questionID, bytes update) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) PostUpdate(questionID [32]byte, update []byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.PostUpdate(&_UmaCtfAdapter.TransactOpts, questionID, update)
}

// PriceDisputed is a paid mutator transaction binding the contract method 0x0d8f2372.
//
// Solidity: function priceDisputed(bytes32 , uint256 , bytes ancillaryData, uint256 ) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) PriceDisputed(opts *bind.TransactOpts, arg0 [32]byte, arg1 *big.Int, ancillaryData []byte, arg3 *big.Int) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "priceDisputed", arg0, arg1, ancillaryData, arg3)
}

// PriceDisputed is a paid mutator transaction binding the contract method 0x0d8f2372.
//
// Solidity: function priceDisputed(bytes32 , uint256 , bytes ancillaryData, uint256 ) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) PriceDisputed(arg0 [32]byte, arg1 *big.Int, ancillaryData []byte, arg3 *big.Int) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.PriceDisputed(&_UmaCtfAdapter.TransactOpts, arg0, arg1, ancillaryData, arg3)
}

// PriceDisputed is a paid mutator transaction binding the contract method 0x0d8f2372.
//
// Solidity: function priceDisputed(bytes32 , uint256 , bytes ancillaryData, uint256 ) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) PriceDisputed(arg0 [32]byte, arg1 *big.Int, ancillaryData []byte, arg3 *big.Int) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.PriceDisputed(&_UmaCtfAdapter.TransactOpts, arg0, arg1, ancillaryData, arg3)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.RemoveAdmin(&_UmaCtfAdapter.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.RemoveAdmin(&_UmaCtfAdapter.TransactOpts, admin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) RenounceAdmin() (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.RenounceAdmin(&_UmaCtfAdapter.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.RenounceAdmin(&_UmaCtfAdapter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xed3c7d40.
//
// Solidity: function reset(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) Reset(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "reset", questionID)
}

// Reset is a paid mutator transaction binding the contract method 0xed3c7d40.
//
// Solidity: function reset(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) Reset(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Reset(&_UmaCtfAdapter.TransactOpts, questionID)
}

// Reset is a paid mutator transaction binding the contract method 0xed3c7d40.
//
// Solidity: function reset(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) Reset(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Reset(&_UmaCtfAdapter.TransactOpts, questionID)
}

// Resolve is a paid mutator transaction binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) Resolve(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "resolve", questionID)
}

// Resolve is a paid mutator transaction binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) Resolve(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Resolve(&_UmaCtfAdapter.TransactOpts, questionID)
}

// Resolve is a paid mutator transaction binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) Resolve(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Resolve(&_UmaCtfAdapter.TransactOpts, questionID)
}

// ResolveManually is a paid mutator transaction binding the contract method 0x80696d85.
//
// Solidity: function resolveManually(bytes32 questionID, uint256[] payouts) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) ResolveManually(opts *bind.TransactOpts, questionID [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "resolveManually", questionID, payouts)
}

// ResolveManually is a paid mutator transaction binding the contract method 0x80696d85.
//
// Solidity: function resolveManually(bytes32 questionID, uint256[] payouts) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) ResolveManually(questionID [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.ResolveManually(&_UmaCtfAdapter.TransactOpts, questionID, payouts)
}

// ResolveManually is a paid mutator transaction binding the contract method 0x80696d85.
//
// Solidity: function resolveManually(bytes32 questionID, uint256[] payouts) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) ResolveManually(questionID [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.ResolveManually(&_UmaCtfAdapter.TransactOpts, questionID, payouts)
}

// Unflag is a paid mutator transaction binding the contract method 0x88697de4.
//
// Solidity: function unflag(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) Unflag(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "unflag", questionID)
}

// Unflag is a paid mutator transaction binding the contract method 0x88697de4.
//
// Solidity: function unflag(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) Unflag(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Unflag(&_UmaCtfAdapter.TransactOpts, questionID)
}

// Unflag is a paid mutator transaction binding the contract method 0x88697de4.
//
// Solidity: function unflag(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) Unflag(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Unflag(&_UmaCtfAdapter.TransactOpts, questionID)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactor) Unpause(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.contract.Transact(opts, "unpause", questionID)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterSession) Unpause(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Unpause(&_UmaCtfAdapter.TransactOpts, questionID)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 questionID) returns()
func (_UmaCtfAdapter *UmaCtfAdapterTransactorSession) Unpause(questionID [32]byte) (*types.Transaction, error) {
	return _UmaCtfAdapter.Contract.Unpause(&_UmaCtfAdapter.TransactOpts, questionID)
}

// UmaCtfAdapterAncillaryDataUpdatedIterator is returned from FilterAncillaryDataUpdated and is used to iterate over the raw logs and unpacked data for AncillaryDataUpdated events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterAncillaryDataUpdatedIterator struct {
	Event *UmaCtfAdapterAncillaryDataUpdated // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterAncillaryDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterAncillaryDataUpdated)
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
		it.Event = new(UmaCtfAdapterAncillaryDataUpdated)
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
func (it *UmaCtfAdapterAncillaryDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterAncillaryDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterAncillaryDataUpdated represents a AncillaryDataUpdated event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterAncillaryDataUpdated struct {
	QuestionID [32]byte
	Owner      common.Address
	Update     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAncillaryDataUpdated is a free log retrieval operation binding the contract event 0x0059e11815211969c0c4aaf3f498b52b6c2f2d14f286275d0862d70de22a836b.
//
// Solidity: event AncillaryDataUpdated(bytes32 indexed questionID, address indexed owner, bytes update)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterAncillaryDataUpdated(opts *bind.FilterOpts, questionID [][32]byte, owner []common.Address) (*UmaCtfAdapterAncillaryDataUpdatedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "AncillaryDataUpdated", questionIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterAncillaryDataUpdatedIterator{contract: _UmaCtfAdapter.contract, event: "AncillaryDataUpdated", logs: logs, sub: sub}, nil
}

// WatchAncillaryDataUpdated is a free log subscription operation binding the contract event 0x0059e11815211969c0c4aaf3f498b52b6c2f2d14f286275d0862d70de22a836b.
//
// Solidity: event AncillaryDataUpdated(bytes32 indexed questionID, address indexed owner, bytes update)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchAncillaryDataUpdated(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterAncillaryDataUpdated, questionID [][32]byte, owner []common.Address) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "AncillaryDataUpdated", questionIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterAncillaryDataUpdated)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "AncillaryDataUpdated", log); err != nil {
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

// ParseAncillaryDataUpdated is a log parse operation binding the contract event 0x0059e11815211969c0c4aaf3f498b52b6c2f2d14f286275d0862d70de22a836b.
//
// Solidity: event AncillaryDataUpdated(bytes32 indexed questionID, address indexed owner, bytes update)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseAncillaryDataUpdated(log types.Log) (*UmaCtfAdapterAncillaryDataUpdated, error) {
	event := new(UmaCtfAdapterAncillaryDataUpdated)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "AncillaryDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterNewAdminIterator struct {
	Event *UmaCtfAdapterNewAdmin // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterNewAdmin)
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
		it.Event = new(UmaCtfAdapterNewAdmin)
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
func (it *UmaCtfAdapterNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterNewAdmin represents a NewAdmin event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address, newAdminAddress []common.Address) (*UmaCtfAdapterNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterNewAdminIterator{contract: _UmaCtfAdapter.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterNewAdmin, admin []common.Address, newAdminAddress []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterNewAdmin)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseNewAdmin(log types.Log) (*UmaCtfAdapterNewAdmin, error) {
	event := new(UmaCtfAdapterNewAdmin)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterQuestionFlaggedIterator is returned from FilterQuestionFlagged and is used to iterate over the raw logs and unpacked data for QuestionFlagged events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionFlaggedIterator struct {
	Event *UmaCtfAdapterQuestionFlagged // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterQuestionFlaggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterQuestionFlagged)
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
		it.Event = new(UmaCtfAdapterQuestionFlagged)
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
func (it *UmaCtfAdapterQuestionFlaggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterQuestionFlaggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterQuestionFlagged represents a QuestionFlagged event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionFlagged struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionFlagged is a free log retrieval operation binding the contract event 0x2435a0347185933b12027c6f394a5fd9c03646dba233e956f50658719dfc0b35.
//
// Solidity: event QuestionFlagged(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterQuestionFlagged(opts *bind.FilterOpts, questionID [][32]byte) (*UmaCtfAdapterQuestionFlaggedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "QuestionFlagged", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterQuestionFlaggedIterator{contract: _UmaCtfAdapter.contract, event: "QuestionFlagged", logs: logs, sub: sub}, nil
}

// WatchQuestionFlagged is a free log subscription operation binding the contract event 0x2435a0347185933b12027c6f394a5fd9c03646dba233e956f50658719dfc0b35.
//
// Solidity: event QuestionFlagged(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchQuestionFlagged(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterQuestionFlagged, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "QuestionFlagged", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterQuestionFlagged)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionFlagged", log); err != nil {
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

// ParseQuestionFlagged is a log parse operation binding the contract event 0x2435a0347185933b12027c6f394a5fd9c03646dba233e956f50658719dfc0b35.
//
// Solidity: event QuestionFlagged(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseQuestionFlagged(log types.Log) (*UmaCtfAdapterQuestionFlagged, error) {
	event := new(UmaCtfAdapterQuestionFlagged)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionFlagged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterQuestionInitializedIterator is returned from FilterQuestionInitialized and is used to iterate over the raw logs and unpacked data for QuestionInitialized events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionInitializedIterator struct {
	Event *UmaCtfAdapterQuestionInitialized // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterQuestionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterQuestionInitialized)
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
		it.Event = new(UmaCtfAdapterQuestionInitialized)
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
func (it *UmaCtfAdapterQuestionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterQuestionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterQuestionInitialized represents a QuestionInitialized event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionInitialized struct {
	QuestionID       [32]byte
	RequestTimestamp *big.Int
	Creator          common.Address
	AncillaryData    []byte
	RewardToken      common.Address
	Reward           *big.Int
	ProposalBond     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterQuestionInitialized is a free log retrieval operation binding the contract event 0xeee0897acd6893adcaf2ba5158191b3601098ab6bece35c5d57874340b64c5b7.
//
// Solidity: event QuestionInitialized(bytes32 indexed questionID, uint256 indexed requestTimestamp, address indexed creator, bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterQuestionInitialized(opts *bind.FilterOpts, questionID [][32]byte, requestTimestamp []*big.Int, creator []common.Address) (*UmaCtfAdapterQuestionInitializedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var requestTimestampRule []interface{}
	for _, requestTimestampItem := range requestTimestamp {
		requestTimestampRule = append(requestTimestampRule, requestTimestampItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "QuestionInitialized", questionIDRule, requestTimestampRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterQuestionInitializedIterator{contract: _UmaCtfAdapter.contract, event: "QuestionInitialized", logs: logs, sub: sub}, nil
}

// WatchQuestionInitialized is a free log subscription operation binding the contract event 0xeee0897acd6893adcaf2ba5158191b3601098ab6bece35c5d57874340b64c5b7.
//
// Solidity: event QuestionInitialized(bytes32 indexed questionID, uint256 indexed requestTimestamp, address indexed creator, bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchQuestionInitialized(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterQuestionInitialized, questionID [][32]byte, requestTimestamp []*big.Int, creator []common.Address) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var requestTimestampRule []interface{}
	for _, requestTimestampItem := range requestTimestamp {
		requestTimestampRule = append(requestTimestampRule, requestTimestampItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "QuestionInitialized", questionIDRule, requestTimestampRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterQuestionInitialized)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionInitialized", log); err != nil {
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

// ParseQuestionInitialized is a log parse operation binding the contract event 0xeee0897acd6893adcaf2ba5158191b3601098ab6bece35c5d57874340b64c5b7.
//
// Solidity: event QuestionInitialized(bytes32 indexed questionID, uint256 indexed requestTimestamp, address indexed creator, bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseQuestionInitialized(log types.Log) (*UmaCtfAdapterQuestionInitialized, error) {
	event := new(UmaCtfAdapterQuestionInitialized)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterQuestionManuallyResolvedIterator is returned from FilterQuestionManuallyResolved and is used to iterate over the raw logs and unpacked data for QuestionManuallyResolved events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionManuallyResolvedIterator struct {
	Event *UmaCtfAdapterQuestionManuallyResolved // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterQuestionManuallyResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterQuestionManuallyResolved)
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
		it.Event = new(UmaCtfAdapterQuestionManuallyResolved)
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
func (it *UmaCtfAdapterQuestionManuallyResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterQuestionManuallyResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterQuestionManuallyResolved represents a QuestionManuallyResolved event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionManuallyResolved struct {
	QuestionID [32]byte
	Payouts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionManuallyResolved is a free log retrieval operation binding the contract event 0x5909815fe7fe0a550d5fcb95fbf33821b580521d3c97089c6ce12808d1cd0566.
//
// Solidity: event QuestionManuallyResolved(bytes32 indexed questionID, uint256[] payouts)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterQuestionManuallyResolved(opts *bind.FilterOpts, questionID [][32]byte) (*UmaCtfAdapterQuestionManuallyResolvedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "QuestionManuallyResolved", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterQuestionManuallyResolvedIterator{contract: _UmaCtfAdapter.contract, event: "QuestionManuallyResolved", logs: logs, sub: sub}, nil
}

// WatchQuestionManuallyResolved is a free log subscription operation binding the contract event 0x5909815fe7fe0a550d5fcb95fbf33821b580521d3c97089c6ce12808d1cd0566.
//
// Solidity: event QuestionManuallyResolved(bytes32 indexed questionID, uint256[] payouts)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchQuestionManuallyResolved(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterQuestionManuallyResolved, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "QuestionManuallyResolved", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterQuestionManuallyResolved)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionManuallyResolved", log); err != nil {
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

// ParseQuestionManuallyResolved is a log parse operation binding the contract event 0x5909815fe7fe0a550d5fcb95fbf33821b580521d3c97089c6ce12808d1cd0566.
//
// Solidity: event QuestionManuallyResolved(bytes32 indexed questionID, uint256[] payouts)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseQuestionManuallyResolved(log types.Log) (*UmaCtfAdapterQuestionManuallyResolved, error) {
	event := new(UmaCtfAdapterQuestionManuallyResolved)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionManuallyResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterQuestionPausedIterator is returned from FilterQuestionPaused and is used to iterate over the raw logs and unpacked data for QuestionPaused events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionPausedIterator struct {
	Event *UmaCtfAdapterQuestionPaused // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterQuestionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterQuestionPaused)
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
		it.Event = new(UmaCtfAdapterQuestionPaused)
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
func (it *UmaCtfAdapterQuestionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterQuestionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterQuestionPaused represents a QuestionPaused event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionPaused struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionPaused is a free log retrieval operation binding the contract event 0x6ded7250a9d5f79aef5add44600fc20a74a0af6f4730baa4fc4ab87bf484b812.
//
// Solidity: event QuestionPaused(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterQuestionPaused(opts *bind.FilterOpts, questionID [][32]byte) (*UmaCtfAdapterQuestionPausedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "QuestionPaused", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterQuestionPausedIterator{contract: _UmaCtfAdapter.contract, event: "QuestionPaused", logs: logs, sub: sub}, nil
}

// WatchQuestionPaused is a free log subscription operation binding the contract event 0x6ded7250a9d5f79aef5add44600fc20a74a0af6f4730baa4fc4ab87bf484b812.
//
// Solidity: event QuestionPaused(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchQuestionPaused(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterQuestionPaused, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "QuestionPaused", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterQuestionPaused)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionPaused", log); err != nil {
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

// ParseQuestionPaused is a log parse operation binding the contract event 0x6ded7250a9d5f79aef5add44600fc20a74a0af6f4730baa4fc4ab87bf484b812.
//
// Solidity: event QuestionPaused(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseQuestionPaused(log types.Log) (*UmaCtfAdapterQuestionPaused, error) {
	event := new(UmaCtfAdapterQuestionPaused)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterQuestionResetIterator is returned from FilterQuestionReset and is used to iterate over the raw logs and unpacked data for QuestionReset events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionResetIterator struct {
	Event *UmaCtfAdapterQuestionReset // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterQuestionResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterQuestionReset)
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
		it.Event = new(UmaCtfAdapterQuestionReset)
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
func (it *UmaCtfAdapterQuestionResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterQuestionResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterQuestionReset represents a QuestionReset event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionReset struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionReset is a free log retrieval operation binding the contract event 0x7981b5832932948db4e32a4a16a0f44b2ce7ff088574afb9364b313f70f82e8f.
//
// Solidity: event QuestionReset(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterQuestionReset(opts *bind.FilterOpts, questionID [][32]byte) (*UmaCtfAdapterQuestionResetIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "QuestionReset", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterQuestionResetIterator{contract: _UmaCtfAdapter.contract, event: "QuestionReset", logs: logs, sub: sub}, nil
}

// WatchQuestionReset is a free log subscription operation binding the contract event 0x7981b5832932948db4e32a4a16a0f44b2ce7ff088574afb9364b313f70f82e8f.
//
// Solidity: event QuestionReset(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchQuestionReset(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterQuestionReset, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "QuestionReset", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterQuestionReset)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionReset", log); err != nil {
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

// ParseQuestionReset is a log parse operation binding the contract event 0x7981b5832932948db4e32a4a16a0f44b2ce7ff088574afb9364b313f70f82e8f.
//
// Solidity: event QuestionReset(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseQuestionReset(log types.Log) (*UmaCtfAdapterQuestionReset, error) {
	event := new(UmaCtfAdapterQuestionReset)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterQuestionResolvedIterator is returned from FilterQuestionResolved and is used to iterate over the raw logs and unpacked data for QuestionResolved events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionResolvedIterator struct {
	Event *UmaCtfAdapterQuestionResolved // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterQuestionResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterQuestionResolved)
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
		it.Event = new(UmaCtfAdapterQuestionResolved)
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
func (it *UmaCtfAdapterQuestionResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterQuestionResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterQuestionResolved represents a QuestionResolved event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionResolved struct {
	QuestionID   [32]byte
	SettledPrice *big.Int
	Payouts      []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuestionResolved is a free log retrieval operation binding the contract event 0x566c3fbdd12dd86bb341787f6d531f79fd7ad4ce7e3ae2d15ac0ca1b601af9df.
//
// Solidity: event QuestionResolved(bytes32 indexed questionID, int256 indexed settledPrice, uint256[] payouts)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterQuestionResolved(opts *bind.FilterOpts, questionID [][32]byte, settledPrice []*big.Int) (*UmaCtfAdapterQuestionResolvedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var settledPriceRule []interface{}
	for _, settledPriceItem := range settledPrice {
		settledPriceRule = append(settledPriceRule, settledPriceItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "QuestionResolved", questionIDRule, settledPriceRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterQuestionResolvedIterator{contract: _UmaCtfAdapter.contract, event: "QuestionResolved", logs: logs, sub: sub}, nil
}

// WatchQuestionResolved is a free log subscription operation binding the contract event 0x566c3fbdd12dd86bb341787f6d531f79fd7ad4ce7e3ae2d15ac0ca1b601af9df.
//
// Solidity: event QuestionResolved(bytes32 indexed questionID, int256 indexed settledPrice, uint256[] payouts)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchQuestionResolved(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterQuestionResolved, questionID [][32]byte, settledPrice []*big.Int) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var settledPriceRule []interface{}
	for _, settledPriceItem := range settledPrice {
		settledPriceRule = append(settledPriceRule, settledPriceItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "QuestionResolved", questionIDRule, settledPriceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterQuestionResolved)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionResolved", log); err != nil {
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

// ParseQuestionResolved is a log parse operation binding the contract event 0x566c3fbdd12dd86bb341787f6d531f79fd7ad4ce7e3ae2d15ac0ca1b601af9df.
//
// Solidity: event QuestionResolved(bytes32 indexed questionID, int256 indexed settledPrice, uint256[] payouts)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseQuestionResolved(log types.Log) (*UmaCtfAdapterQuestionResolved, error) {
	event := new(UmaCtfAdapterQuestionResolved)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterQuestionUnflaggedIterator is returned from FilterQuestionUnflagged and is used to iterate over the raw logs and unpacked data for QuestionUnflagged events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionUnflaggedIterator struct {
	Event *UmaCtfAdapterQuestionUnflagged // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterQuestionUnflaggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterQuestionUnflagged)
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
		it.Event = new(UmaCtfAdapterQuestionUnflagged)
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
func (it *UmaCtfAdapterQuestionUnflaggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterQuestionUnflaggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterQuestionUnflagged represents a QuestionUnflagged event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionUnflagged struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionUnflagged is a free log retrieval operation binding the contract event 0x052435bc04fc49113a7bfd9198a92c0852ca622a621800f6da66d4b29b786c05.
//
// Solidity: event QuestionUnflagged(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterQuestionUnflagged(opts *bind.FilterOpts, questionID [][32]byte) (*UmaCtfAdapterQuestionUnflaggedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "QuestionUnflagged", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterQuestionUnflaggedIterator{contract: _UmaCtfAdapter.contract, event: "QuestionUnflagged", logs: logs, sub: sub}, nil
}

// WatchQuestionUnflagged is a free log subscription operation binding the contract event 0x052435bc04fc49113a7bfd9198a92c0852ca622a621800f6da66d4b29b786c05.
//
// Solidity: event QuestionUnflagged(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchQuestionUnflagged(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterQuestionUnflagged, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "QuestionUnflagged", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterQuestionUnflagged)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionUnflagged", log); err != nil {
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

// ParseQuestionUnflagged is a log parse operation binding the contract event 0x052435bc04fc49113a7bfd9198a92c0852ca622a621800f6da66d4b29b786c05.
//
// Solidity: event QuestionUnflagged(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseQuestionUnflagged(log types.Log) (*UmaCtfAdapterQuestionUnflagged, error) {
	event := new(UmaCtfAdapterQuestionUnflagged)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionUnflagged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterQuestionUnpausedIterator is returned from FilterQuestionUnpaused and is used to iterate over the raw logs and unpacked data for QuestionUnpaused events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionUnpausedIterator struct {
	Event *UmaCtfAdapterQuestionUnpaused // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterQuestionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterQuestionUnpaused)
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
		it.Event = new(UmaCtfAdapterQuestionUnpaused)
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
func (it *UmaCtfAdapterQuestionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterQuestionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterQuestionUnpaused represents a QuestionUnpaused event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterQuestionUnpaused struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionUnpaused is a free log retrieval operation binding the contract event 0x92d28918c5574e7fc0f4f948c39502682c81cfb4089b07b83f95b3264e5e5e06.
//
// Solidity: event QuestionUnpaused(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterQuestionUnpaused(opts *bind.FilterOpts, questionID [][32]byte) (*UmaCtfAdapterQuestionUnpausedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "QuestionUnpaused", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterQuestionUnpausedIterator{contract: _UmaCtfAdapter.contract, event: "QuestionUnpaused", logs: logs, sub: sub}, nil
}

// WatchQuestionUnpaused is a free log subscription operation binding the contract event 0x92d28918c5574e7fc0f4f948c39502682c81cfb4089b07b83f95b3264e5e5e06.
//
// Solidity: event QuestionUnpaused(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchQuestionUnpaused(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterQuestionUnpaused, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "QuestionUnpaused", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterQuestionUnpaused)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionUnpaused", log); err != nil {
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

// ParseQuestionUnpaused is a log parse operation binding the contract event 0x92d28918c5574e7fc0f4f948c39502682c81cfb4089b07b83f95b3264e5e5e06.
//
// Solidity: event QuestionUnpaused(bytes32 indexed questionID)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseQuestionUnpaused(log types.Log) (*UmaCtfAdapterQuestionUnpaused, error) {
	event := new(UmaCtfAdapterQuestionUnpaused)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "QuestionUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmaCtfAdapterRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the UmaCtfAdapter contract.
type UmaCtfAdapterRemovedAdminIterator struct {
	Event *UmaCtfAdapterRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *UmaCtfAdapterRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmaCtfAdapterRemovedAdmin)
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
		it.Event = new(UmaCtfAdapterRemovedAdmin)
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
func (it *UmaCtfAdapterRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmaCtfAdapterRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmaCtfAdapterRemovedAdmin represents a RemovedAdmin event raised by the UmaCtfAdapter contract.
type UmaCtfAdapterRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, admin []common.Address, removedAdmin []common.Address) (*UmaCtfAdapterRemovedAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.FilterLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return &UmaCtfAdapterRemovedAdminIterator{contract: _UmaCtfAdapter.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *UmaCtfAdapterRemovedAdmin, admin []common.Address, removedAdmin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _UmaCtfAdapter.contract.WatchLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmaCtfAdapterRemovedAdmin)
				if err := _UmaCtfAdapter.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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

// ParseRemovedAdmin is a log parse operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_UmaCtfAdapter *UmaCtfAdapterFilterer) ParseRemovedAdmin(log types.Log) (*UmaCtfAdapterRemovedAdmin, error) {
	event := new(UmaCtfAdapterRemovedAdmin)
	if err := _UmaCtfAdapter.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package negriskadapter

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

// NegRiskAdapterMetaData contains all meta data concerning the NegRiskAdapter contract.
var NegRiskAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ctf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DeterminedFlagAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeBipsOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIndexSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketAlreadyDetermined\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketAlreadyPrepared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketNotPrepared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoConvertiblePositions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedForAll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedCollateralToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MarketPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"outcome\",\"type\":\"bool\"}],\"name\":\"OutcomeReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"name\":\"PayoutRedemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionSplit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"indexSet\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionsConverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionsMerge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"QuestionPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NO_TOKEN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"col\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_indexSet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertPositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctf\",\"outputs\":[{\"internalType\":\"contractIConditionalTokens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"}],\"name\":\"getConditionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getDetermined\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getFeeBips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getMarketData\",\"outputs\":[{\"internalType\":\"MarketData\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_outcome\",\"type\":\"bool\"}],\"name\":\"getPositionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getQuestionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mergePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mergePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeBips\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"}],\"name\":\"prepareMarket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"}],\"name\":\"prepareQuestion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"redeemPositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_outcome\",\"type\":\"bool\"}],\"name\":\"reportOutcome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"splitPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"splitPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wcol\",\"outputs\":[{\"internalType\":\"contractWrappedCollateral\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NegRiskAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use NegRiskAdapterMetaData.ABI instead.
var NegRiskAdapterABI = NegRiskAdapterMetaData.ABI

// NegRiskAdapter is an auto generated Go binding around an Ethereum contract.
type NegRiskAdapter struct {
	NegRiskAdapterCaller     // Read-only binding to the contract
	NegRiskAdapterTransactor // Write-only binding to the contract
	NegRiskAdapterFilterer   // Log filterer for contract events
}

// NegRiskAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type NegRiskAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NegRiskAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NegRiskAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NegRiskAdapterSession struct {
	Contract     *NegRiskAdapter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NegRiskAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NegRiskAdapterCallerSession struct {
	Contract *NegRiskAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NegRiskAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NegRiskAdapterTransactorSession struct {
	Contract     *NegRiskAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NegRiskAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type NegRiskAdapterRaw struct {
	Contract *NegRiskAdapter // Generic contract binding to access the raw methods on
}

// NegRiskAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NegRiskAdapterCallerRaw struct {
	Contract *NegRiskAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// NegRiskAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NegRiskAdapterTransactorRaw struct {
	Contract *NegRiskAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNegRiskAdapter creates a new instance of NegRiskAdapter, bound to a specific deployed contract.
func NewNegRiskAdapter(address common.Address, backend bind.ContractBackend) (*NegRiskAdapter, error) {
	contract, err := bindNegRiskAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapter{NegRiskAdapterCaller: NegRiskAdapterCaller{contract: contract}, NegRiskAdapterTransactor: NegRiskAdapterTransactor{contract: contract}, NegRiskAdapterFilterer: NegRiskAdapterFilterer{contract: contract}}, nil
}

// NewNegRiskAdapterCaller creates a new read-only instance of NegRiskAdapter, bound to a specific deployed contract.
func NewNegRiskAdapterCaller(address common.Address, caller bind.ContractCaller) (*NegRiskAdapterCaller, error) {
	contract, err := bindNegRiskAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterCaller{contract: contract}, nil
}

// NewNegRiskAdapterTransactor creates a new write-only instance of NegRiskAdapter, bound to a specific deployed contract.
func NewNegRiskAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*NegRiskAdapterTransactor, error) {
	contract, err := bindNegRiskAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterTransactor{contract: contract}, nil
}

// NewNegRiskAdapterFilterer creates a new log filterer instance of NegRiskAdapter, bound to a specific deployed contract.
func NewNegRiskAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*NegRiskAdapterFilterer, error) {
	contract, err := bindNegRiskAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterFilterer{contract: contract}, nil
}

// bindNegRiskAdapter binds a generic wrapper to an already deployed contract.
func bindNegRiskAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskAdapter *NegRiskAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskAdapter.Contract.NegRiskAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskAdapter *NegRiskAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.NegRiskAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskAdapter *NegRiskAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.NegRiskAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskAdapter *NegRiskAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskAdapter *NegRiskAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskAdapter *NegRiskAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.contract.Transact(opts, method, params...)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterSession) FEEDENOMINATOR() (*big.Int, error) {
	return _NegRiskAdapter.Contract.FEEDENOMINATOR(&_NegRiskAdapter.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _NegRiskAdapter.Contract.FEEDENOMINATOR(&_NegRiskAdapter.CallOpts)
}

// NOTOKENBURNADDRESS is a free data retrieval call binding the contract method 0x7ad7fe36.
//
// Solidity: function NO_TOKEN_BURN_ADDRESS() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCaller) NOTOKENBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "NO_TOKEN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NOTOKENBURNADDRESS is a free data retrieval call binding the contract method 0x7ad7fe36.
//
// Solidity: function NO_TOKEN_BURN_ADDRESS() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterSession) NOTOKENBURNADDRESS() (common.Address, error) {
	return _NegRiskAdapter.Contract.NOTOKENBURNADDRESS(&_NegRiskAdapter.CallOpts)
}

// NOTOKENBURNADDRESS is a free data retrieval call binding the contract method 0x7ad7fe36.
//
// Solidity: function NO_TOKEN_BURN_ADDRESS() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) NOTOKENBURNADDRESS() (common.Address, error) {
	return _NegRiskAdapter.Contract.NOTOKENBURNADDRESS(&_NegRiskAdapter.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _NegRiskAdapter.Contract.Admins(&_NegRiskAdapter.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _NegRiskAdapter.Contract.Admins(&_NegRiskAdapter.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "balanceOf", _owner, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _NegRiskAdapter.Contract.BalanceOf(&_NegRiskAdapter.CallOpts, _owner, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _NegRiskAdapter.Contract.BalanceOf(&_NegRiskAdapter.CallOpts, _owner, _id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_NegRiskAdapter *NegRiskAdapterCaller) BalanceOfBatch(opts *bind.CallOpts, _owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "balanceOfBatch", _owners, _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_NegRiskAdapter *NegRiskAdapterSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _NegRiskAdapter.Contract.BalanceOfBatch(&_NegRiskAdapter.CallOpts, _owners, _ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_NegRiskAdapter *NegRiskAdapterCallerSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _NegRiskAdapter.Contract.BalanceOfBatch(&_NegRiskAdapter.CallOpts, _owners, _ids)
}

// Col is a free data retrieval call binding the contract method 0xa78695b0.
//
// Solidity: function col() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCaller) Col(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "col")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Col is a free data retrieval call binding the contract method 0xa78695b0.
//
// Solidity: function col() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterSession) Col() (common.Address, error) {
	return _NegRiskAdapter.Contract.Col(&_NegRiskAdapter.CallOpts)
}

// Col is a free data retrieval call binding the contract method 0xa78695b0.
//
// Solidity: function col() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) Col() (common.Address, error) {
	return _NegRiskAdapter.Contract.Col(&_NegRiskAdapter.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCaller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterSession) Ctf() (common.Address, error) {
	return _NegRiskAdapter.Contract.Ctf(&_NegRiskAdapter.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) Ctf() (common.Address, error) {
	return _NegRiskAdapter.Contract.Ctf(&_NegRiskAdapter.CallOpts)
}

// GetConditionId is a free data retrieval call binding the contract method 0x04329c03.
//
// Solidity: function getConditionId(bytes32 _questionId) view returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterCaller) GetConditionId(opts *bind.CallOpts, _questionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "getConditionId", _questionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConditionId is a free data retrieval call binding the contract method 0x04329c03.
//
// Solidity: function getConditionId(bytes32 _questionId) view returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterSession) GetConditionId(_questionId [32]byte) ([32]byte, error) {
	return _NegRiskAdapter.Contract.GetConditionId(&_NegRiskAdapter.CallOpts, _questionId)
}

// GetConditionId is a free data retrieval call binding the contract method 0x04329c03.
//
// Solidity: function getConditionId(bytes32 _questionId) view returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) GetConditionId(_questionId [32]byte) ([32]byte, error) {
	return _NegRiskAdapter.Contract.GetConditionId(&_NegRiskAdapter.CallOpts, _questionId)
}

// GetDetermined is a free data retrieval call binding the contract method 0x7ae2e67b.
//
// Solidity: function getDetermined(bytes32 _marketId) view returns(bool)
func (_NegRiskAdapter *NegRiskAdapterCaller) GetDetermined(opts *bind.CallOpts, _marketId [32]byte) (bool, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "getDetermined", _marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetDetermined is a free data retrieval call binding the contract method 0x7ae2e67b.
//
// Solidity: function getDetermined(bytes32 _marketId) view returns(bool)
func (_NegRiskAdapter *NegRiskAdapterSession) GetDetermined(_marketId [32]byte) (bool, error) {
	return _NegRiskAdapter.Contract.GetDetermined(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetDetermined is a free data retrieval call binding the contract method 0x7ae2e67b.
//
// Solidity: function getDetermined(bytes32 _marketId) view returns(bool)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) GetDetermined(_marketId [32]byte) (bool, error) {
	return _NegRiskAdapter.Contract.GetDetermined(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetFeeBips is a free data retrieval call binding the contract method 0x2582cb5e.
//
// Solidity: function getFeeBips(bytes32 _marketId) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCaller) GetFeeBips(opts *bind.CallOpts, _marketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "getFeeBips", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBips is a free data retrieval call binding the contract method 0x2582cb5e.
//
// Solidity: function getFeeBips(bytes32 _marketId) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterSession) GetFeeBips(_marketId [32]byte) (*big.Int, error) {
	return _NegRiskAdapter.Contract.GetFeeBips(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetFeeBips is a free data retrieval call binding the contract method 0x2582cb5e.
//
// Solidity: function getFeeBips(bytes32 _marketId) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) GetFeeBips(_marketId [32]byte) (*big.Int, error) {
	return _NegRiskAdapter.Contract.GetFeeBips(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetMarketData is a free data retrieval call binding the contract method 0x30f4f4bb.
//
// Solidity: function getMarketData(bytes32 _marketId) view returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterCaller) GetMarketData(opts *bind.CallOpts, _marketId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "getMarketData", _marketId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMarketData is a free data retrieval call binding the contract method 0x30f4f4bb.
//
// Solidity: function getMarketData(bytes32 _marketId) view returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterSession) GetMarketData(_marketId [32]byte) ([32]byte, error) {
	return _NegRiskAdapter.Contract.GetMarketData(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetMarketData is a free data retrieval call binding the contract method 0x30f4f4bb.
//
// Solidity: function getMarketData(bytes32 _marketId) view returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) GetMarketData(_marketId [32]byte) ([32]byte, error) {
	return _NegRiskAdapter.Contract.GetMarketData(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _marketId) view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCaller) GetOracle(opts *bind.CallOpts, _marketId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "getOracle", _marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _marketId) view returns(address)
func (_NegRiskAdapter *NegRiskAdapterSession) GetOracle(_marketId [32]byte) (common.Address, error) {
	return _NegRiskAdapter.Contract.GetOracle(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _marketId) view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) GetOracle(_marketId [32]byte) (common.Address, error) {
	return _NegRiskAdapter.Contract.GetOracle(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetPositionId is a free data retrieval call binding the contract method 0x752b5ba5.
//
// Solidity: function getPositionId(bytes32 _questionId, bool _outcome) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCaller) GetPositionId(opts *bind.CallOpts, _questionId [32]byte, _outcome bool) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "getPositionId", _questionId, _outcome)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionId is a free data retrieval call binding the contract method 0x752b5ba5.
//
// Solidity: function getPositionId(bytes32 _questionId, bool _outcome) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterSession) GetPositionId(_questionId [32]byte, _outcome bool) (*big.Int, error) {
	return _NegRiskAdapter.Contract.GetPositionId(&_NegRiskAdapter.CallOpts, _questionId, _outcome)
}

// GetPositionId is a free data retrieval call binding the contract method 0x752b5ba5.
//
// Solidity: function getPositionId(bytes32 _questionId, bool _outcome) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) GetPositionId(_questionId [32]byte, _outcome bool) (*big.Int, error) {
	return _NegRiskAdapter.Contract.GetPositionId(&_NegRiskAdapter.CallOpts, _questionId, _outcome)
}

// GetQuestionCount is a free data retrieval call binding the contract method 0xb7f75d2c.
//
// Solidity: function getQuestionCount(bytes32 _marketId) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCaller) GetQuestionCount(opts *bind.CallOpts, _marketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "getQuestionCount", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuestionCount is a free data retrieval call binding the contract method 0xb7f75d2c.
//
// Solidity: function getQuestionCount(bytes32 _marketId) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterSession) GetQuestionCount(_marketId [32]byte) (*big.Int, error) {
	return _NegRiskAdapter.Contract.GetQuestionCount(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetQuestionCount is a free data retrieval call binding the contract method 0xb7f75d2c.
//
// Solidity: function getQuestionCount(bytes32 _marketId) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) GetQuestionCount(_marketId [32]byte) (*big.Int, error) {
	return _NegRiskAdapter.Contract.GetQuestionCount(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _marketId) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCaller) GetResult(opts *bind.CallOpts, _marketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "getResult", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _marketId) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterSession) GetResult(_marketId [32]byte) (*big.Int, error) {
	return _NegRiskAdapter.Contract.GetResult(&_NegRiskAdapter.CallOpts, _marketId)
}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _marketId) view returns(uint256)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) GetResult(_marketId [32]byte) (*big.Int, error) {
	return _NegRiskAdapter.Contract.GetResult(&_NegRiskAdapter.CallOpts, _marketId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_NegRiskAdapter *NegRiskAdapterCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "isAdmin", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_NegRiskAdapter *NegRiskAdapterSession) IsAdmin(addr common.Address) (bool, error) {
	return _NegRiskAdapter.Contract.IsAdmin(&_NegRiskAdapter.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _NegRiskAdapter.Contract.IsAdmin(&_NegRiskAdapter.CallOpts, addr)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterSession) Vault() (common.Address, error) {
	return _NegRiskAdapter.Contract.Vault(&_NegRiskAdapter.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) Vault() (common.Address, error) {
	return _NegRiskAdapter.Contract.Vault(&_NegRiskAdapter.CallOpts)
}

// Wcol is a free data retrieval call binding the contract method 0x7e3b74c3.
//
// Solidity: function wcol() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCaller) Wcol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskAdapter.contract.Call(opts, &out, "wcol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wcol is a free data retrieval call binding the contract method 0x7e3b74c3.
//
// Solidity: function wcol() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterSession) Wcol() (common.Address, error) {
	return _NegRiskAdapter.Contract.Wcol(&_NegRiskAdapter.CallOpts)
}

// Wcol is a free data retrieval call binding the contract method 0x7e3b74c3.
//
// Solidity: function wcol() view returns(address)
func (_NegRiskAdapter *NegRiskAdapterCallerSession) Wcol() (common.Address, error) {
	return _NegRiskAdapter.Contract.Wcol(&_NegRiskAdapter.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.AddAdmin(&_NegRiskAdapter.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.AddAdmin(&_NegRiskAdapter.TransactOpts, admin)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) ConvertPositions(opts *bind.TransactOpts, _marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "convertPositions", _marketId, _indexSet, _amount)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) ConvertPositions(_marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.ConvertPositions(&_NegRiskAdapter.TransactOpts, _marketId, _indexSet, _amount)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) ConvertPositions(_marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.ConvertPositions(&_NegRiskAdapter.TransactOpts, _marketId, _indexSet, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) MergePositions(opts *bind.TransactOpts, _collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "mergePositions", _collateralToken, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) MergePositions(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.MergePositions(&_NegRiskAdapter.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) MergePositions(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.MergePositions(&_NegRiskAdapter.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// MergePositions0 is a paid mutator transaction binding the contract method 0xb10c5c17.
//
// Solidity: function mergePositions(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) MergePositions0(opts *bind.TransactOpts, _conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "mergePositions0", _conditionId, _amount)
}

// MergePositions0 is a paid mutator transaction binding the contract method 0xb10c5c17.
//
// Solidity: function mergePositions(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) MergePositions0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.MergePositions0(&_NegRiskAdapter.TransactOpts, _conditionId, _amount)
}

// MergePositions0 is a paid mutator transaction binding the contract method 0xb10c5c17.
//
// Solidity: function mergePositions(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) MergePositions0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.MergePositions0(&_NegRiskAdapter.TransactOpts, _conditionId, _amount)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskAdapter *NegRiskAdapterTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskAdapter *NegRiskAdapterSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.OnERC1155BatchReceived(&_NegRiskAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.OnERC1155BatchReceived(&_NegRiskAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskAdapter *NegRiskAdapterTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskAdapter *NegRiskAdapterSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.OnERC1155Received(&_NegRiskAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.OnERC1155Received(&_NegRiskAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PrepareMarket is a paid mutator transaction binding the contract method 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _metadata) returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterTransactor) PrepareMarket(opts *bind.TransactOpts, _feeBips *big.Int, _metadata []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "prepareMarket", _feeBips, _metadata)
}

// PrepareMarket is a paid mutator transaction binding the contract method 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _metadata) returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterSession) PrepareMarket(_feeBips *big.Int, _metadata []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.PrepareMarket(&_NegRiskAdapter.TransactOpts, _feeBips, _metadata)
}

// PrepareMarket is a paid mutator transaction binding the contract method 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _metadata) returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) PrepareMarket(_feeBips *big.Int, _metadata []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.PrepareMarket(&_NegRiskAdapter.TransactOpts, _feeBips, _metadata)
}

// PrepareQuestion is a paid mutator transaction binding the contract method 0x1d69b48d.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _metadata) returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterTransactor) PrepareQuestion(opts *bind.TransactOpts, _marketId [32]byte, _metadata []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "prepareQuestion", _marketId, _metadata)
}

// PrepareQuestion is a paid mutator transaction binding the contract method 0x1d69b48d.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _metadata) returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterSession) PrepareQuestion(_marketId [32]byte, _metadata []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.PrepareQuestion(&_NegRiskAdapter.TransactOpts, _marketId, _metadata)
}

// PrepareQuestion is a paid mutator transaction binding the contract method 0x1d69b48d.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _metadata) returns(bytes32)
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) PrepareQuestion(_marketId [32]byte, _metadata []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.PrepareQuestion(&_NegRiskAdapter.TransactOpts, _marketId, _metadata)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0xdbeccb23.
//
// Solidity: function redeemPositions(bytes32 _conditionId, uint256[] _amounts) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) RedeemPositions(opts *bind.TransactOpts, _conditionId [32]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "redeemPositions", _conditionId, _amounts)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0xdbeccb23.
//
// Solidity: function redeemPositions(bytes32 _conditionId, uint256[] _amounts) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) RedeemPositions(_conditionId [32]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.RedeemPositions(&_NegRiskAdapter.TransactOpts, _conditionId, _amounts)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0xdbeccb23.
//
// Solidity: function redeemPositions(bytes32 _conditionId, uint256[] _amounts) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) RedeemPositions(_conditionId [32]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.RedeemPositions(&_NegRiskAdapter.TransactOpts, _conditionId, _amounts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.RemoveAdmin(&_NegRiskAdapter.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.RemoveAdmin(&_NegRiskAdapter.TransactOpts, admin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_NegRiskAdapter *NegRiskAdapterSession) RenounceAdmin() (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.RenounceAdmin(&_NegRiskAdapter.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.RenounceAdmin(&_NegRiskAdapter.TransactOpts)
}

// ReportOutcome is a paid mutator transaction binding the contract method 0xe200affd.
//
// Solidity: function reportOutcome(bytes32 _questionId, bool _outcome) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) ReportOutcome(opts *bind.TransactOpts, _questionId [32]byte, _outcome bool) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "reportOutcome", _questionId, _outcome)
}

// ReportOutcome is a paid mutator transaction binding the contract method 0xe200affd.
//
// Solidity: function reportOutcome(bytes32 _questionId, bool _outcome) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) ReportOutcome(_questionId [32]byte, _outcome bool) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.ReportOutcome(&_NegRiskAdapter.TransactOpts, _questionId, _outcome)
}

// ReportOutcome is a paid mutator transaction binding the contract method 0xe200affd.
//
// Solidity: function reportOutcome(bytes32 _questionId, bool _outcome) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) ReportOutcome(_questionId [32]byte, _outcome bool) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.ReportOutcome(&_NegRiskAdapter.TransactOpts, _questionId, _outcome)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "safeTransferFrom", _from, _to, _id, _value, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.SafeTransferFrom(&_NegRiskAdapter.TransactOpts, _from, _to, _id, _value, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.SafeTransferFrom(&_NegRiskAdapter.TransactOpts, _from, _to, _id, _value, _data)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) SplitPosition(opts *bind.TransactOpts, _collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "splitPosition", _collateralToken, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) SplitPosition(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.SplitPosition(&_NegRiskAdapter.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) SplitPosition(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.SplitPosition(&_NegRiskAdapter.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// SplitPosition0 is a paid mutator transaction binding the contract method 0xa3d7da1d.
//
// Solidity: function splitPosition(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactor) SplitPosition0(opts *bind.TransactOpts, _conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.contract.Transact(opts, "splitPosition0", _conditionId, _amount)
}

// SplitPosition0 is a paid mutator transaction binding the contract method 0xa3d7da1d.
//
// Solidity: function splitPosition(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterSession) SplitPosition0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.SplitPosition0(&_NegRiskAdapter.TransactOpts, _conditionId, _amount)
}

// SplitPosition0 is a paid mutator transaction binding the contract method 0xa3d7da1d.
//
// Solidity: function splitPosition(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRiskAdapter *NegRiskAdapterTransactorSession) SplitPosition0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskAdapter.Contract.SplitPosition0(&_NegRiskAdapter.TransactOpts, _conditionId, _amount)
}

// NegRiskAdapterMarketPreparedIterator is returned from FilterMarketPrepared and is used to iterate over the raw logs and unpacked data for MarketPrepared events raised by the NegRiskAdapter contract.
type NegRiskAdapterMarketPreparedIterator struct {
	Event *NegRiskAdapterMarketPrepared // Event containing the contract specifics and raw log

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
func (it *NegRiskAdapterMarketPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskAdapterMarketPrepared)
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
		it.Event = new(NegRiskAdapterMarketPrepared)
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
func (it *NegRiskAdapterMarketPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskAdapterMarketPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskAdapterMarketPrepared represents a MarketPrepared event raised by the NegRiskAdapter contract.
type NegRiskAdapterMarketPrepared struct {
	MarketId [32]byte
	Oracle   common.Address
	FeeBips  *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketPrepared is a free log retrieval operation binding the contract event 0xf059ab16d1ca60e123eab60e3c02b68faf060347c701a5d14885a8e1def7b3a8.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data)
func (_NegRiskAdapter *NegRiskAdapterFilterer) FilterMarketPrepared(opts *bind.FilterOpts, marketId [][32]byte, oracle []common.Address) (*NegRiskAdapterMarketPreparedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.FilterLogs(opts, "MarketPrepared", marketIdRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterMarketPreparedIterator{contract: _NegRiskAdapter.contract, event: "MarketPrepared", logs: logs, sub: sub}, nil
}

// WatchMarketPrepared is a free log subscription operation binding the contract event 0xf059ab16d1ca60e123eab60e3c02b68faf060347c701a5d14885a8e1def7b3a8.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data)
func (_NegRiskAdapter *NegRiskAdapterFilterer) WatchMarketPrepared(opts *bind.WatchOpts, sink chan<- *NegRiskAdapterMarketPrepared, marketId [][32]byte, oracle []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.WatchLogs(opts, "MarketPrepared", marketIdRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskAdapterMarketPrepared)
				if err := _NegRiskAdapter.contract.UnpackLog(event, "MarketPrepared", log); err != nil {
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

// ParseMarketPrepared is a log parse operation binding the contract event 0xf059ab16d1ca60e123eab60e3c02b68faf060347c701a5d14885a8e1def7b3a8.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data)
func (_NegRiskAdapter *NegRiskAdapterFilterer) ParseMarketPrepared(log types.Log) (*NegRiskAdapterMarketPrepared, error) {
	event := new(NegRiskAdapterMarketPrepared)
	if err := _NegRiskAdapter.contract.UnpackLog(event, "MarketPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskAdapterNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the NegRiskAdapter contract.
type NegRiskAdapterNewAdminIterator struct {
	Event *NegRiskAdapterNewAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskAdapterNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskAdapterNewAdmin)
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
		it.Event = new(NegRiskAdapterNewAdmin)
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
func (it *NegRiskAdapterNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskAdapterNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskAdapterNewAdmin represents a NewAdmin event raised by the NegRiskAdapter contract.
type NegRiskAdapterNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_NegRiskAdapter *NegRiskAdapterFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address, newAdminAddress []common.Address) (*NegRiskAdapterNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.FilterLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterNewAdminIterator{contract: _NegRiskAdapter.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_NegRiskAdapter *NegRiskAdapterFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskAdapterNewAdmin, admin []common.Address, newAdminAddress []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.WatchLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskAdapterNewAdmin)
				if err := _NegRiskAdapter.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_NegRiskAdapter *NegRiskAdapterFilterer) ParseNewAdmin(log types.Log) (*NegRiskAdapterNewAdmin, error) {
	event := new(NegRiskAdapterNewAdmin)
	if err := _NegRiskAdapter.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskAdapterOutcomeReportedIterator is returned from FilterOutcomeReported and is used to iterate over the raw logs and unpacked data for OutcomeReported events raised by the NegRiskAdapter contract.
type NegRiskAdapterOutcomeReportedIterator struct {
	Event *NegRiskAdapterOutcomeReported // Event containing the contract specifics and raw log

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
func (it *NegRiskAdapterOutcomeReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskAdapterOutcomeReported)
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
		it.Event = new(NegRiskAdapterOutcomeReported)
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
func (it *NegRiskAdapterOutcomeReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskAdapterOutcomeReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskAdapterOutcomeReported represents a OutcomeReported event raised by the NegRiskAdapter contract.
type NegRiskAdapterOutcomeReported struct {
	MarketId   [32]byte
	QuestionId [32]byte
	Outcome    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOutcomeReported is a free log retrieval operation binding the contract event 0x9e9fa7fd355160bd4cd3f22d4333519354beff1f5689bde87f2c5e63d8d484b2.
//
// Solidity: event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome)
func (_NegRiskAdapter *NegRiskAdapterFilterer) FilterOutcomeReported(opts *bind.FilterOpts, marketId [][32]byte, questionId [][32]byte) (*NegRiskAdapterOutcomeReportedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.FilterLogs(opts, "OutcomeReported", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterOutcomeReportedIterator{contract: _NegRiskAdapter.contract, event: "OutcomeReported", logs: logs, sub: sub}, nil
}

// WatchOutcomeReported is a free log subscription operation binding the contract event 0x9e9fa7fd355160bd4cd3f22d4333519354beff1f5689bde87f2c5e63d8d484b2.
//
// Solidity: event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome)
func (_NegRiskAdapter *NegRiskAdapterFilterer) WatchOutcomeReported(opts *bind.WatchOpts, sink chan<- *NegRiskAdapterOutcomeReported, marketId [][32]byte, questionId [][32]byte) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.WatchLogs(opts, "OutcomeReported", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskAdapterOutcomeReported)
				if err := _NegRiskAdapter.contract.UnpackLog(event, "OutcomeReported", log); err != nil {
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

// ParseOutcomeReported is a log parse operation binding the contract event 0x9e9fa7fd355160bd4cd3f22d4333519354beff1f5689bde87f2c5e63d8d484b2.
//
// Solidity: event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome)
func (_NegRiskAdapter *NegRiskAdapterFilterer) ParseOutcomeReported(log types.Log) (*NegRiskAdapterOutcomeReported, error) {
	event := new(NegRiskAdapterOutcomeReported)
	if err := _NegRiskAdapter.contract.UnpackLog(event, "OutcomeReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskAdapterPayoutRedemptionIterator is returned from FilterPayoutRedemption and is used to iterate over the raw logs and unpacked data for PayoutRedemption events raised by the NegRiskAdapter contract.
type NegRiskAdapterPayoutRedemptionIterator struct {
	Event *NegRiskAdapterPayoutRedemption // Event containing the contract specifics and raw log

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
func (it *NegRiskAdapterPayoutRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskAdapterPayoutRedemption)
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
		it.Event = new(NegRiskAdapterPayoutRedemption)
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
func (it *NegRiskAdapterPayoutRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskAdapterPayoutRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskAdapterPayoutRedemption represents a PayoutRedemption event raised by the NegRiskAdapter contract.
type NegRiskAdapterPayoutRedemption struct {
	Redeemer    common.Address
	ConditionId [32]byte
	Amounts     []*big.Int
	Payout      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayoutRedemption is a free log retrieval operation binding the contract event 0x9140a6a270ef945260c03894b3c6b3b2695e9d5101feef0ff24fec960cfd3224.
//
// Solidity: event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout)
func (_NegRiskAdapter *NegRiskAdapterFilterer) FilterPayoutRedemption(opts *bind.FilterOpts, redeemer []common.Address, conditionId [][32]byte) (*NegRiskAdapterPayoutRedemptionIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.FilterLogs(opts, "PayoutRedemption", redeemerRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterPayoutRedemptionIterator{contract: _NegRiskAdapter.contract, event: "PayoutRedemption", logs: logs, sub: sub}, nil
}

// WatchPayoutRedemption is a free log subscription operation binding the contract event 0x9140a6a270ef945260c03894b3c6b3b2695e9d5101feef0ff24fec960cfd3224.
//
// Solidity: event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout)
func (_NegRiskAdapter *NegRiskAdapterFilterer) WatchPayoutRedemption(opts *bind.WatchOpts, sink chan<- *NegRiskAdapterPayoutRedemption, redeemer []common.Address, conditionId [][32]byte) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.WatchLogs(opts, "PayoutRedemption", redeemerRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskAdapterPayoutRedemption)
				if err := _NegRiskAdapter.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
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

// ParsePayoutRedemption is a log parse operation binding the contract event 0x9140a6a270ef945260c03894b3c6b3b2695e9d5101feef0ff24fec960cfd3224.
//
// Solidity: event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout)
func (_NegRiskAdapter *NegRiskAdapterFilterer) ParsePayoutRedemption(log types.Log) (*NegRiskAdapterPayoutRedemption, error) {
	event := new(NegRiskAdapterPayoutRedemption)
	if err := _NegRiskAdapter.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskAdapterPositionSplitIterator is returned from FilterPositionSplit and is used to iterate over the raw logs and unpacked data for PositionSplit events raised by the NegRiskAdapter contract.
type NegRiskAdapterPositionSplitIterator struct {
	Event *NegRiskAdapterPositionSplit // Event containing the contract specifics and raw log

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
func (it *NegRiskAdapterPositionSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskAdapterPositionSplit)
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
		it.Event = new(NegRiskAdapterPositionSplit)
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
func (it *NegRiskAdapterPositionSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskAdapterPositionSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskAdapterPositionSplit represents a PositionSplit event raised by the NegRiskAdapter contract.
type NegRiskAdapterPositionSplit struct {
	Stakeholder common.Address
	ConditionId [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionSplit is a free log retrieval operation binding the contract event 0xbbed930dbfb7907ae2d60ddf78345610214f26419a0128df39b6cc3d9e5df9b0.
//
// Solidity: event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRiskAdapter *NegRiskAdapterFilterer) FilterPositionSplit(opts *bind.FilterOpts, stakeholder []common.Address, conditionId [][32]byte) (*NegRiskAdapterPositionSplitIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.FilterLogs(opts, "PositionSplit", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterPositionSplitIterator{contract: _NegRiskAdapter.contract, event: "PositionSplit", logs: logs, sub: sub}, nil
}

// WatchPositionSplit is a free log subscription operation binding the contract event 0xbbed930dbfb7907ae2d60ddf78345610214f26419a0128df39b6cc3d9e5df9b0.
//
// Solidity: event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRiskAdapter *NegRiskAdapterFilterer) WatchPositionSplit(opts *bind.WatchOpts, sink chan<- *NegRiskAdapterPositionSplit, stakeholder []common.Address, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.WatchLogs(opts, "PositionSplit", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskAdapterPositionSplit)
				if err := _NegRiskAdapter.contract.UnpackLog(event, "PositionSplit", log); err != nil {
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

// ParsePositionSplit is a log parse operation binding the contract event 0xbbed930dbfb7907ae2d60ddf78345610214f26419a0128df39b6cc3d9e5df9b0.
//
// Solidity: event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRiskAdapter *NegRiskAdapterFilterer) ParsePositionSplit(log types.Log) (*NegRiskAdapterPositionSplit, error) {
	event := new(NegRiskAdapterPositionSplit)
	if err := _NegRiskAdapter.contract.UnpackLog(event, "PositionSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskAdapterPositionsConvertedIterator is returned from FilterPositionsConverted and is used to iterate over the raw logs and unpacked data for PositionsConverted events raised by the NegRiskAdapter contract.
type NegRiskAdapterPositionsConvertedIterator struct {
	Event *NegRiskAdapterPositionsConverted // Event containing the contract specifics and raw log

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
func (it *NegRiskAdapterPositionsConvertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskAdapterPositionsConverted)
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
		it.Event = new(NegRiskAdapterPositionsConverted)
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
func (it *NegRiskAdapterPositionsConvertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskAdapterPositionsConvertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskAdapterPositionsConverted represents a PositionsConverted event raised by the NegRiskAdapter contract.
type NegRiskAdapterPositionsConverted struct {
	Stakeholder common.Address
	MarketId    [32]byte
	IndexSet    *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionsConverted is a free log retrieval operation binding the contract event 0xb03d19dddbc72a87e735ff0ea3b57bef133ebe44e1894284916a84044deb367e.
//
// Solidity: event PositionsConverted(address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount)
func (_NegRiskAdapter *NegRiskAdapterFilterer) FilterPositionsConverted(opts *bind.FilterOpts, stakeholder []common.Address, marketId [][32]byte, indexSet []*big.Int) (*NegRiskAdapterPositionsConvertedIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var indexSetRule []interface{}
	for _, indexSetItem := range indexSet {
		indexSetRule = append(indexSetRule, indexSetItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.FilterLogs(opts, "PositionsConverted", stakeholderRule, marketIdRule, indexSetRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterPositionsConvertedIterator{contract: _NegRiskAdapter.contract, event: "PositionsConverted", logs: logs, sub: sub}, nil
}

// WatchPositionsConverted is a free log subscription operation binding the contract event 0xb03d19dddbc72a87e735ff0ea3b57bef133ebe44e1894284916a84044deb367e.
//
// Solidity: event PositionsConverted(address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount)
func (_NegRiskAdapter *NegRiskAdapterFilterer) WatchPositionsConverted(opts *bind.WatchOpts, sink chan<- *NegRiskAdapterPositionsConverted, stakeholder []common.Address, marketId [][32]byte, indexSet []*big.Int) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var indexSetRule []interface{}
	for _, indexSetItem := range indexSet {
		indexSetRule = append(indexSetRule, indexSetItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.WatchLogs(opts, "PositionsConverted", stakeholderRule, marketIdRule, indexSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskAdapterPositionsConverted)
				if err := _NegRiskAdapter.contract.UnpackLog(event, "PositionsConverted", log); err != nil {
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

// ParsePositionsConverted is a log parse operation binding the contract event 0xb03d19dddbc72a87e735ff0ea3b57bef133ebe44e1894284916a84044deb367e.
//
// Solidity: event PositionsConverted(address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount)
func (_NegRiskAdapter *NegRiskAdapterFilterer) ParsePositionsConverted(log types.Log) (*NegRiskAdapterPositionsConverted, error) {
	event := new(NegRiskAdapterPositionsConverted)
	if err := _NegRiskAdapter.contract.UnpackLog(event, "PositionsConverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskAdapterPositionsMergeIterator is returned from FilterPositionsMerge and is used to iterate over the raw logs and unpacked data for PositionsMerge events raised by the NegRiskAdapter contract.
type NegRiskAdapterPositionsMergeIterator struct {
	Event *NegRiskAdapterPositionsMerge // Event containing the contract specifics and raw log

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
func (it *NegRiskAdapterPositionsMergeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskAdapterPositionsMerge)
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
		it.Event = new(NegRiskAdapterPositionsMerge)
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
func (it *NegRiskAdapterPositionsMergeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskAdapterPositionsMergeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskAdapterPositionsMerge represents a PositionsMerge event raised by the NegRiskAdapter contract.
type NegRiskAdapterPositionsMerge struct {
	Stakeholder common.Address
	ConditionId [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionsMerge is a free log retrieval operation binding the contract event 0xba33ac50d8894676597e6e35dc09cff59854708b642cd069d21eb9c7ca072a04.
//
// Solidity: event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRiskAdapter *NegRiskAdapterFilterer) FilterPositionsMerge(opts *bind.FilterOpts, stakeholder []common.Address, conditionId [][32]byte) (*NegRiskAdapterPositionsMergeIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.FilterLogs(opts, "PositionsMerge", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterPositionsMergeIterator{contract: _NegRiskAdapter.contract, event: "PositionsMerge", logs: logs, sub: sub}, nil
}

// WatchPositionsMerge is a free log subscription operation binding the contract event 0xba33ac50d8894676597e6e35dc09cff59854708b642cd069d21eb9c7ca072a04.
//
// Solidity: event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRiskAdapter *NegRiskAdapterFilterer) WatchPositionsMerge(opts *bind.WatchOpts, sink chan<- *NegRiskAdapterPositionsMerge, stakeholder []common.Address, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.WatchLogs(opts, "PositionsMerge", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskAdapterPositionsMerge)
				if err := _NegRiskAdapter.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
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

// ParsePositionsMerge is a log parse operation binding the contract event 0xba33ac50d8894676597e6e35dc09cff59854708b642cd069d21eb9c7ca072a04.
//
// Solidity: event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRiskAdapter *NegRiskAdapterFilterer) ParsePositionsMerge(log types.Log) (*NegRiskAdapterPositionsMerge, error) {
	event := new(NegRiskAdapterPositionsMerge)
	if err := _NegRiskAdapter.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskAdapterQuestionPreparedIterator is returned from FilterQuestionPrepared and is used to iterate over the raw logs and unpacked data for QuestionPrepared events raised by the NegRiskAdapter contract.
type NegRiskAdapterQuestionPreparedIterator struct {
	Event *NegRiskAdapterQuestionPrepared // Event containing the contract specifics and raw log

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
func (it *NegRiskAdapterQuestionPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskAdapterQuestionPrepared)
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
		it.Event = new(NegRiskAdapterQuestionPrepared)
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
func (it *NegRiskAdapterQuestionPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskAdapterQuestionPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskAdapterQuestionPrepared represents a QuestionPrepared event raised by the NegRiskAdapter contract.
type NegRiskAdapterQuestionPrepared struct {
	MarketId   [32]byte
	QuestionId [32]byte
	Index      *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionPrepared is a free log retrieval operation binding the contract event 0xaac410f87d423a922a7b226ac68f0c2eaf5bf6d15e644ac0758c7f96e2c253f7.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data)
func (_NegRiskAdapter *NegRiskAdapterFilterer) FilterQuestionPrepared(opts *bind.FilterOpts, marketId [][32]byte, questionId [][32]byte) (*NegRiskAdapterQuestionPreparedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.FilterLogs(opts, "QuestionPrepared", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterQuestionPreparedIterator{contract: _NegRiskAdapter.contract, event: "QuestionPrepared", logs: logs, sub: sub}, nil
}

// WatchQuestionPrepared is a free log subscription operation binding the contract event 0xaac410f87d423a922a7b226ac68f0c2eaf5bf6d15e644ac0758c7f96e2c253f7.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data)
func (_NegRiskAdapter *NegRiskAdapterFilterer) WatchQuestionPrepared(opts *bind.WatchOpts, sink chan<- *NegRiskAdapterQuestionPrepared, marketId [][32]byte, questionId [][32]byte) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.WatchLogs(opts, "QuestionPrepared", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskAdapterQuestionPrepared)
				if err := _NegRiskAdapter.contract.UnpackLog(event, "QuestionPrepared", log); err != nil {
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

// ParseQuestionPrepared is a log parse operation binding the contract event 0xaac410f87d423a922a7b226ac68f0c2eaf5bf6d15e644ac0758c7f96e2c253f7.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data)
func (_NegRiskAdapter *NegRiskAdapterFilterer) ParseQuestionPrepared(log types.Log) (*NegRiskAdapterQuestionPrepared, error) {
	event := new(NegRiskAdapterQuestionPrepared)
	if err := _NegRiskAdapter.contract.UnpackLog(event, "QuestionPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskAdapterRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the NegRiskAdapter contract.
type NegRiskAdapterRemovedAdminIterator struct {
	Event *NegRiskAdapterRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskAdapterRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskAdapterRemovedAdmin)
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
		it.Event = new(NegRiskAdapterRemovedAdmin)
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
func (it *NegRiskAdapterRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskAdapterRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskAdapterRemovedAdmin represents a RemovedAdmin event raised by the NegRiskAdapter contract.
type NegRiskAdapterRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_NegRiskAdapter *NegRiskAdapterFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, admin []common.Address, removedAdmin []common.Address) (*NegRiskAdapterRemovedAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.FilterLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskAdapterRemovedAdminIterator{contract: _NegRiskAdapter.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_NegRiskAdapter *NegRiskAdapterFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskAdapterRemovedAdmin, admin []common.Address, removedAdmin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _NegRiskAdapter.contract.WatchLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskAdapterRemovedAdmin)
				if err := _NegRiskAdapter.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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
func (_NegRiskAdapter *NegRiskAdapterFilterer) ParseRemovedAdmin(log types.Log) (*NegRiskAdapterRemovedAdmin, error) {
	event := new(NegRiskAdapterRemovedAdmin)
	if err := _NegRiskAdapter.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

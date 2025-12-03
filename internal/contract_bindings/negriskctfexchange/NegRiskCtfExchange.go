// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package negriskctfexchange

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

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Salt          *big.Int
	Maker         common.Address
	Signer        common.Address
	Taker         common.Address
	TokenId       *big.Int
	MakerAmount   *big.Int
	TakerAmount   *big.Int
	Expiration    *big.Int
	Nonce         *big.Int
	FeeRateBps    *big.Int
	Side          uint8
	SignatureType uint8
	Signature     []byte
}

// OrderStatus is an auto generated low-level Go binding around an user-defined struct.
type OrderStatus struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}

// NegRiskCtfExchangeMetaData contains all meta data concerning the NegRiskCtfExchange contract.
var NegRiskCtfExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_negRiskAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_safeFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidComplement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingGtRemaining\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedTokenIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCrossing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderFilledOrCancelled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLittleTokensReceived\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"takerOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"takerOrderMaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProxyFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProxyFactory\",\"type\":\"address\"}],\"name\":\"ProxyFactoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSafeFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSafeFactory\",\"type\":\"address\"}],\"name\":\"SafeFactoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"token0\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"token1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingUnpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fillAmounts\",\"type\":\"uint256[]\"}],\"name\":\"fillOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"getComplement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"getConditionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCtf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isFilledOrCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"internalType\":\"structOrderStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolyProxyFactoryImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPolyProxyWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getSafeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeFactoryImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"isValidNonce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"makerOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFillAmounts\",\"type\":\"uint256[]\"}],\"name\":\"matchOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFilledOrCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentCollectionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProxyFactory\",\"type\":\"address\"}],\"name\":\"setProxyFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSafeFactory\",\"type\":\"address\"}],\"name\":\"setSafeFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"}],\"name\":\"validateComplement\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"validateTokenId\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NegRiskCtfExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use NegRiskCtfExchangeMetaData.ABI instead.
var NegRiskCtfExchangeABI = NegRiskCtfExchangeMetaData.ABI

// NegRiskCtfExchange is an auto generated Go binding around an Ethereum contract.
type NegRiskCtfExchange struct {
	NegRiskCtfExchangeCaller     // Read-only binding to the contract
	NegRiskCtfExchangeTransactor // Write-only binding to the contract
	NegRiskCtfExchangeFilterer   // Log filterer for contract events
}

// NegRiskCtfExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NegRiskCtfExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskCtfExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NegRiskCtfExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskCtfExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NegRiskCtfExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskCtfExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NegRiskCtfExchangeSession struct {
	Contract     *NegRiskCtfExchange // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NegRiskCtfExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NegRiskCtfExchangeCallerSession struct {
	Contract *NegRiskCtfExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// NegRiskCtfExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NegRiskCtfExchangeTransactorSession struct {
	Contract     *NegRiskCtfExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// NegRiskCtfExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NegRiskCtfExchangeRaw struct {
	Contract *NegRiskCtfExchange // Generic contract binding to access the raw methods on
}

// NegRiskCtfExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NegRiskCtfExchangeCallerRaw struct {
	Contract *NegRiskCtfExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// NegRiskCtfExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NegRiskCtfExchangeTransactorRaw struct {
	Contract *NegRiskCtfExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNegRiskCtfExchange creates a new instance of NegRiskCtfExchange, bound to a specific deployed contract.
func NewNegRiskCtfExchange(address common.Address, backend bind.ContractBackend) (*NegRiskCtfExchange, error) {
	contract, err := bindNegRiskCtfExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchange{NegRiskCtfExchangeCaller: NegRiskCtfExchangeCaller{contract: contract}, NegRiskCtfExchangeTransactor: NegRiskCtfExchangeTransactor{contract: contract}, NegRiskCtfExchangeFilterer: NegRiskCtfExchangeFilterer{contract: contract}}, nil
}

// NewNegRiskCtfExchangeCaller creates a new read-only instance of NegRiskCtfExchange, bound to a specific deployed contract.
func NewNegRiskCtfExchangeCaller(address common.Address, caller bind.ContractCaller) (*NegRiskCtfExchangeCaller, error) {
	contract, err := bindNegRiskCtfExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeCaller{contract: contract}, nil
}

// NewNegRiskCtfExchangeTransactor creates a new write-only instance of NegRiskCtfExchange, bound to a specific deployed contract.
func NewNegRiskCtfExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*NegRiskCtfExchangeTransactor, error) {
	contract, err := bindNegRiskCtfExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeTransactor{contract: contract}, nil
}

// NewNegRiskCtfExchangeFilterer creates a new log filterer instance of NegRiskCtfExchange, bound to a specific deployed contract.
func NewNegRiskCtfExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*NegRiskCtfExchangeFilterer, error) {
	contract, err := bindNegRiskCtfExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeFilterer{contract: contract}, nil
}

// bindNegRiskCtfExchange binds a generic wrapper to an already deployed contract.
func bindNegRiskCtfExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NegRiskCtfExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskCtfExchange *NegRiskCtfExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskCtfExchange.Contract.NegRiskCtfExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskCtfExchange *NegRiskCtfExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.NegRiskCtfExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskCtfExchange *NegRiskCtfExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.NegRiskCtfExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskCtfExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.Admins(&_NegRiskCtfExchange.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.Admins(&_NegRiskCtfExchange.CallOpts, arg0)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) DomainSeparator() ([32]byte, error) {
	return _NegRiskCtfExchange.Contract.DomainSeparator(&_NegRiskCtfExchange.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) DomainSeparator() ([32]byte, error) {
	return _NegRiskCtfExchange.Contract.DomainSeparator(&_NegRiskCtfExchange.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetCollateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getCollateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetCollateral() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetCollateral(&_NegRiskCtfExchange.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetCollateral() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetCollateral(&_NegRiskCtfExchange.CallOpts)
}

// GetComplement is a free data retrieval call binding the contract method 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetComplement(opts *bind.CallOpts, token *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getComplement", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetComplement is a free data retrieval call binding the contract method 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetComplement(token *big.Int) (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.GetComplement(&_NegRiskCtfExchange.CallOpts, token)
}

// GetComplement is a free data retrieval call binding the contract method 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetComplement(token *big.Int) (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.GetComplement(&_NegRiskCtfExchange.CallOpts, token)
}

// GetConditionId is a free data retrieval call binding the contract method 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetConditionId(opts *bind.CallOpts, token *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getConditionId", token)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConditionId is a free data retrieval call binding the contract method 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetConditionId(token *big.Int) ([32]byte, error) {
	return _NegRiskCtfExchange.Contract.GetConditionId(&_NegRiskCtfExchange.CallOpts, token)
}

// GetConditionId is a free data retrieval call binding the contract method 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetConditionId(token *big.Int) ([32]byte, error) {
	return _NegRiskCtfExchange.Contract.GetConditionId(&_NegRiskCtfExchange.CallOpts, token)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetCtf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getCtf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetCtf() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetCtf(&_NegRiskCtfExchange.CallOpts)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetCtf() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetCtf(&_NegRiskCtfExchange.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetMaxFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getMaxFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetMaxFeeRate() (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.GetMaxFeeRate(&_NegRiskCtfExchange.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetMaxFeeRate() (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.GetMaxFeeRate(&_NegRiskCtfExchange.CallOpts)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (OrderStatus, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getOrderStatus", orderHash)

	if err != nil {
		return *new(OrderStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderStatus)).(*OrderStatus)

	return out0, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _NegRiskCtfExchange.Contract.GetOrderStatus(&_NegRiskCtfExchange.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _NegRiskCtfExchange.Contract.GetOrderStatus(&_NegRiskCtfExchange.CallOpts, orderHash)
}

// GetPolyProxyFactoryImplementation is a free data retrieval call binding the contract method 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetPolyProxyFactoryImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getPolyProxyFactoryImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPolyProxyFactoryImplementation is a free data retrieval call binding the contract method 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetPolyProxyFactoryImplementation() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetPolyProxyFactoryImplementation(&_NegRiskCtfExchange.CallOpts)
}

// GetPolyProxyFactoryImplementation is a free data retrieval call binding the contract method 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetPolyProxyFactoryImplementation() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetPolyProxyFactoryImplementation(&_NegRiskCtfExchange.CallOpts)
}

// GetPolyProxyWalletAddress is a free data retrieval call binding the contract method 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetPolyProxyWalletAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getPolyProxyWalletAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPolyProxyWalletAddress is a free data retrieval call binding the contract method 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetPolyProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetPolyProxyWalletAddress(&_NegRiskCtfExchange.CallOpts, _addr)
}

// GetPolyProxyWalletAddress is a free data retrieval call binding the contract method 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetPolyProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetPolyProxyWalletAddress(&_NegRiskCtfExchange.CallOpts, _addr)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getProxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetProxyFactory() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetProxyFactory(&_NegRiskCtfExchange.CallOpts)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetProxyFactory() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetProxyFactory(&_NegRiskCtfExchange.CallOpts)
}

// GetSafeAddress is a free data retrieval call binding the contract method 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetSafeAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getSafeAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeAddress is a free data retrieval call binding the contract method 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetSafeAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetSafeAddress(&_NegRiskCtfExchange.CallOpts, _addr)
}

// GetSafeAddress is a free data retrieval call binding the contract method 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetSafeAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetSafeAddress(&_NegRiskCtfExchange.CallOpts, _addr)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetSafeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getSafeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetSafeFactory() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetSafeFactory(&_NegRiskCtfExchange.CallOpts)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetSafeFactory() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetSafeFactory(&_NegRiskCtfExchange.CallOpts)
}

// GetSafeFactoryImplementation is a free data retrieval call binding the contract method 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) GetSafeFactoryImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "getSafeFactoryImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeFactoryImplementation is a free data retrieval call binding the contract method 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) GetSafeFactoryImplementation() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetSafeFactoryImplementation(&_NegRiskCtfExchange.CallOpts)
}

// GetSafeFactoryImplementation is a free data retrieval call binding the contract method 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) GetSafeFactoryImplementation() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.GetSafeFactoryImplementation(&_NegRiskCtfExchange.CallOpts)
}

// HashOrder is a free data retrieval call binding the contract method 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) HashOrder(opts *bind.CallOpts, order Order) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) HashOrder(order Order) ([32]byte, error) {
	return _NegRiskCtfExchange.Contract.HashOrder(&_NegRiskCtfExchange.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) HashOrder(order Order) ([32]byte, error) {
	return _NegRiskCtfExchange.Contract.HashOrder(&_NegRiskCtfExchange.CallOpts, order)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) IsAdmin(opts *bind.CallOpts, usr common.Address) (bool, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "isAdmin", usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) IsAdmin(usr common.Address) (bool, error) {
	return _NegRiskCtfExchange.Contract.IsAdmin(&_NegRiskCtfExchange.CallOpts, usr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) IsAdmin(usr common.Address) (bool, error) {
	return _NegRiskCtfExchange.Contract.IsAdmin(&_NegRiskCtfExchange.CallOpts, usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) IsOperator(opts *bind.CallOpts, usr common.Address) (bool, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "isOperator", usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) IsOperator(usr common.Address) (bool, error) {
	return _NegRiskCtfExchange.Contract.IsOperator(&_NegRiskCtfExchange.CallOpts, usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) IsOperator(usr common.Address) (bool, error) {
	return _NegRiskCtfExchange.Contract.IsOperator(&_NegRiskCtfExchange.CallOpts, usr)
}

// IsValidNonce is a free data retrieval call binding the contract method 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) IsValidNonce(opts *bind.CallOpts, usr common.Address, nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "isValidNonce", usr, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidNonce is a free data retrieval call binding the contract method 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) IsValidNonce(usr common.Address, nonce *big.Int) (bool, error) {
	return _NegRiskCtfExchange.Contract.IsValidNonce(&_NegRiskCtfExchange.CallOpts, usr, nonce)
}

// IsValidNonce is a free data retrieval call binding the contract method 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) IsValidNonce(usr common.Address, nonce *big.Int) (bool, error) {
	return _NegRiskCtfExchange.Contract.IsValidNonce(&_NegRiskCtfExchange.CallOpts, usr, nonce)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.Nonces(&_NegRiskCtfExchange.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.Nonces(&_NegRiskCtfExchange.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.Operators(&_NegRiskCtfExchange.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCtfExchange.Contract.Operators(&_NegRiskCtfExchange.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) OrderStatus(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "orderStatus", arg0)

	outstruct := new(struct {
		IsFilledOrCancelled bool
		Remaining           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsFilledOrCancelled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Remaining = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) OrderStatus(arg0 [32]byte) (struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}, error) {
	return _NegRiskCtfExchange.Contract.OrderStatus(&_NegRiskCtfExchange.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) OrderStatus(arg0 [32]byte) (struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}, error) {
	return _NegRiskCtfExchange.Contract.OrderStatus(&_NegRiskCtfExchange.CallOpts, arg0)
}

// ParentCollectionId is a free data retrieval call binding the contract method 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) ParentCollectionId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "parentCollectionId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParentCollectionId is a free data retrieval call binding the contract method 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) ParentCollectionId() ([32]byte, error) {
	return _NegRiskCtfExchange.Contract.ParentCollectionId(&_NegRiskCtfExchange.CallOpts)
}

// ParentCollectionId is a free data retrieval call binding the contract method 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) ParentCollectionId() ([32]byte, error) {
	return _NegRiskCtfExchange.Contract.ParentCollectionId(&_NegRiskCtfExchange.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) Paused() (bool, error) {
	return _NegRiskCtfExchange.Contract.Paused(&_NegRiskCtfExchange.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) Paused() (bool, error) {
	return _NegRiskCtfExchange.Contract.Paused(&_NegRiskCtfExchange.CallOpts)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) ProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "proxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) ProxyFactory() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.ProxyFactory(&_NegRiskCtfExchange.CallOpts)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) ProxyFactory() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.ProxyFactory(&_NegRiskCtfExchange.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) Registry(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Complement  *big.Int
	ConditionId [32]byte
}, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "registry", arg0)

	outstruct := new(struct {
		Complement  *big.Int
		ConditionId [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Complement = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ConditionId = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Registry is a free data retrieval call binding the contract method 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) Registry(arg0 *big.Int) (struct {
	Complement  *big.Int
	ConditionId [32]byte
}, error) {
	return _NegRiskCtfExchange.Contract.Registry(&_NegRiskCtfExchange.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) Registry(arg0 *big.Int) (struct {
	Complement  *big.Int
	ConditionId [32]byte
}, error) {
	return _NegRiskCtfExchange.Contract.Registry(&_NegRiskCtfExchange.CallOpts, arg0)
}

// SafeFactory is a free data retrieval call binding the contract method 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) SafeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "safeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeFactory is a free data retrieval call binding the contract method 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) SafeFactory() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.SafeFactory(&_NegRiskCtfExchange.CallOpts)
}

// SafeFactory is a free data retrieval call binding the contract method 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) SafeFactory() (common.Address, error) {
	return _NegRiskCtfExchange.Contract.SafeFactory(&_NegRiskCtfExchange.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NegRiskCtfExchange.Contract.SupportsInterface(&_NegRiskCtfExchange.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NegRiskCtfExchange.Contract.SupportsInterface(&_NegRiskCtfExchange.CallOpts, interfaceId)
}

// ValidateComplement is a free data retrieval call binding the contract method 0xd82da838.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) ValidateComplement(opts *bind.CallOpts, token *big.Int, complement *big.Int) error {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "validateComplement", token, complement)

	if err != nil {
		return err
	}

	return err

}

// ValidateComplement is a free data retrieval call binding the contract method 0xd82da838.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) ValidateComplement(token *big.Int, complement *big.Int) error {
	return _NegRiskCtfExchange.Contract.ValidateComplement(&_NegRiskCtfExchange.CallOpts, token, complement)
}

// ValidateComplement is a free data retrieval call binding the contract method 0xd82da838.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) ValidateComplement(token *big.Int, complement *big.Int) error {
	return _NegRiskCtfExchange.Contract.ValidateComplement(&_NegRiskCtfExchange.CallOpts, token, complement)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x654f0ce4.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) ValidateOrder(opts *bind.CallOpts, order Order) error {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "validateOrder", order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrder is a free data retrieval call binding the contract method 0x654f0ce4.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) ValidateOrder(order Order) error {
	return _NegRiskCtfExchange.Contract.ValidateOrder(&_NegRiskCtfExchange.CallOpts, order)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x654f0ce4.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) ValidateOrder(order Order) error {
	return _NegRiskCtfExchange.Contract.ValidateOrder(&_NegRiskCtfExchange.CallOpts, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xe2eec405.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) ValidateOrderSignature(opts *bind.CallOpts, orderHash [32]byte, order Order) error {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "validateOrderSignature", orderHash, order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xe2eec405.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _NegRiskCtfExchange.Contract.ValidateOrderSignature(&_NegRiskCtfExchange.CallOpts, orderHash, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xe2eec405.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _NegRiskCtfExchange.Contract.ValidateOrderSignature(&_NegRiskCtfExchange.CallOpts, orderHash, order)
}

// ValidateTokenId is a free data retrieval call binding the contract method 0x34600901.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeCaller) ValidateTokenId(opts *bind.CallOpts, tokenId *big.Int) error {
	var out []interface{}
	err := _NegRiskCtfExchange.contract.Call(opts, &out, "validateTokenId", tokenId)

	if err != nil {
		return err
	}

	return err

}

// ValidateTokenId is a free data retrieval call binding the contract method 0x34600901.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) ValidateTokenId(tokenId *big.Int) error {
	return _NegRiskCtfExchange.Contract.ValidateTokenId(&_NegRiskCtfExchange.CallOpts, tokenId)
}

// ValidateTokenId is a free data retrieval call binding the contract method 0x34600901.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeCallerSession) ValidateTokenId(tokenId *big.Int) error {
	return _NegRiskCtfExchange.Contract.ValidateTokenId(&_NegRiskCtfExchange.CallOpts, tokenId)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) AddAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "addAdmin", admin_)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) AddAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.AddAdmin(&_NegRiskCtfExchange.TransactOpts, admin_)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) AddAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.AddAdmin(&_NegRiskCtfExchange.TransactOpts, admin_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) AddOperator(opts *bind.TransactOpts, operator_ common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "addOperator", operator_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) AddOperator(operator_ common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.AddOperator(&_NegRiskCtfExchange.TransactOpts, operator_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) AddOperator(operator_ common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.AddOperator(&_NegRiskCtfExchange.TransactOpts, operator_)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa6dfcf86.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa6dfcf86.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.CancelOrder(&_NegRiskCtfExchange.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa6dfcf86.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.CancelOrder(&_NegRiskCtfExchange.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xfa950b48.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) CancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xfa950b48.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.CancelOrders(&_NegRiskCtfExchange.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xfa950b48.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.CancelOrders(&_NegRiskCtfExchange.TransactOpts, orders)
}

// FillOrder is a paid mutator transaction binding the contract method 0xfe729aaf.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) FillOrder(opts *bind.TransactOpts, order Order, fillAmount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "fillOrder", order, fillAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0xfe729aaf.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) FillOrder(order Order, fillAmount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.FillOrder(&_NegRiskCtfExchange.TransactOpts, order, fillAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0xfe729aaf.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) FillOrder(order Order, fillAmount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.FillOrder(&_NegRiskCtfExchange.TransactOpts, order, fillAmount)
}

// FillOrders is a paid mutator transaction binding the contract method 0xd798eff6.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) FillOrders(opts *bind.TransactOpts, orders []Order, fillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "fillOrders", orders, fillAmounts)
}

// FillOrders is a paid mutator transaction binding the contract method 0xd798eff6.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) FillOrders(orders []Order, fillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.FillOrders(&_NegRiskCtfExchange.TransactOpts, orders, fillAmounts)
}

// FillOrders is a paid mutator transaction binding the contract method 0xd798eff6.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) FillOrders(orders []Order, fillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.FillOrders(&_NegRiskCtfExchange.TransactOpts, orders, fillAmounts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) IncrementNonce() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.IncrementNonce(&_NegRiskCtfExchange.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.IncrementNonce(&_NegRiskCtfExchange.TransactOpts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xe60f0c05.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) MatchOrders(opts *bind.TransactOpts, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "matchOrders", takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xe60f0c05.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.MatchOrders(&_NegRiskCtfExchange.TransactOpts, takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xe60f0c05.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.MatchOrders(&_NegRiskCtfExchange.TransactOpts, takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.OnERC1155BatchReceived(&_NegRiskCtfExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.OnERC1155BatchReceived(&_NegRiskCtfExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.OnERC1155Received(&_NegRiskCtfExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.OnERC1155Received(&_NegRiskCtfExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) PauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "pauseTrading")
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) PauseTrading() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.PauseTrading(&_NegRiskCtfExchange.TransactOpts)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) PauseTrading() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.PauseTrading(&_NegRiskCtfExchange.TransactOpts)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x68c7450f.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) RegisterToken(opts *bind.TransactOpts, token *big.Int, complement *big.Int, conditionId [32]byte) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "registerToken", token, complement, conditionId)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x68c7450f.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) RegisterToken(token *big.Int, complement *big.Int, conditionId [32]byte) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RegisterToken(&_NegRiskCtfExchange.TransactOpts, token, complement, conditionId)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x68c7450f.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) RegisterToken(token *big.Int, complement *big.Int, conditionId [32]byte) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RegisterToken(&_NegRiskCtfExchange.TransactOpts, token, complement, conditionId)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RemoveAdmin(&_NegRiskCtfExchange.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RemoveAdmin(&_NegRiskCtfExchange.TransactOpts, admin)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RemoveOperator(&_NegRiskCtfExchange.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RemoveOperator(&_NegRiskCtfExchange.TransactOpts, operator)
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) RenounceAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "renounceAdminRole")
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) RenounceAdminRole() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RenounceAdminRole(&_NegRiskCtfExchange.TransactOpts)
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) RenounceAdminRole() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RenounceAdminRole(&_NegRiskCtfExchange.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) RenounceOperatorRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "renounceOperatorRole")
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RenounceOperatorRole(&_NegRiskCtfExchange.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.RenounceOperatorRole(&_NegRiskCtfExchange.TransactOpts)
}

// SetProxyFactory is a paid mutator transaction binding the contract method 0xfbddd751.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) SetProxyFactory(opts *bind.TransactOpts, _newProxyFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "setProxyFactory", _newProxyFactory)
}

// SetProxyFactory is a paid mutator transaction binding the contract method 0xfbddd751.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) SetProxyFactory(_newProxyFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.SetProxyFactory(&_NegRiskCtfExchange.TransactOpts, _newProxyFactory)
}

// SetProxyFactory is a paid mutator transaction binding the contract method 0xfbddd751.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) SetProxyFactory(_newProxyFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.SetProxyFactory(&_NegRiskCtfExchange.TransactOpts, _newProxyFactory)
}

// SetSafeFactory is a paid mutator transaction binding the contract method 0x4544f055.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) SetSafeFactory(opts *bind.TransactOpts, _newSafeFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "setSafeFactory", _newSafeFactory)
}

// SetSafeFactory is a paid mutator transaction binding the contract method 0x4544f055.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) SetSafeFactory(_newSafeFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.SetSafeFactory(&_NegRiskCtfExchange.TransactOpts, _newSafeFactory)
}

// SetSafeFactory is a paid mutator transaction binding the contract method 0x4544f055.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) SetSafeFactory(_newSafeFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.SetSafeFactory(&_NegRiskCtfExchange.TransactOpts, _newSafeFactory)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactor) UnpauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfExchange.contract.Transact(opts, "unpauseTrading")
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeSession) UnpauseTrading() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.UnpauseTrading(&_NegRiskCtfExchange.TransactOpts)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_NegRiskCtfExchange *NegRiskCtfExchangeTransactorSession) UnpauseTrading() (*types.Transaction, error) {
	return _NegRiskCtfExchange.Contract.UnpauseTrading(&_NegRiskCtfExchange.TransactOpts)
}

// NegRiskCtfExchangeFeeChargedIterator is returned from FilterFeeCharged and is used to iterate over the raw logs and unpacked data for FeeCharged events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeFeeChargedIterator struct {
	Event *NegRiskCtfExchangeFeeCharged // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeFeeCharged)
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
		it.Event = new(NegRiskCtfExchangeFeeCharged)
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
func (it *NegRiskCtfExchangeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeFeeCharged represents a FeeCharged event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeFeeCharged struct {
	Receiver common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeCharged is a free log retrieval operation binding the contract event 0xacffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterFeeCharged(opts *bind.FilterOpts, receiver []common.Address) (*NegRiskCtfExchangeFeeChargedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeFeeChargedIterator{contract: _NegRiskCtfExchange.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

// WatchFeeCharged is a free log subscription operation binding the contract event 0xacffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeFeeCharged, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeFeeCharged)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

// ParseFeeCharged is a log parse operation binding the contract event 0xacffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseFeeCharged(log types.Log) (*NegRiskCtfExchangeFeeCharged, error) {
	event := new(NegRiskCtfExchangeFeeCharged)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNewAdminIterator struct {
	Event *NegRiskCtfExchangeNewAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeNewAdmin)
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
		it.Event = new(NegRiskCtfExchangeNewAdmin)
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
func (it *NegRiskCtfExchangeNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeNewAdmin represents a NewAdmin event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNewAdmin struct {
	NewAdminAddress common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterNewAdmin(opts *bind.FilterOpts, newAdminAddress []common.Address, admin []common.Address) (*NegRiskCtfExchangeNewAdminIterator, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeNewAdminIterator{contract: _NegRiskCtfExchange.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeNewAdmin, newAdminAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeNewAdmin)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseNewAdmin(log types.Log) (*NegRiskCtfExchangeNewAdmin, error) {
	event := new(NegRiskCtfExchangeNewAdmin)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeNewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNewOperatorIterator struct {
	Event *NegRiskCtfExchangeNewOperator // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeNewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeNewOperator)
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
		it.Event = new(NegRiskCtfExchangeNewOperator)
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
func (it *NegRiskCtfExchangeNewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeNewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeNewOperator represents a NewOperator event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNewOperator struct {
	NewOperatorAddress common.Address
	Admin              common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterNewOperator(opts *bind.FilterOpts, newOperatorAddress []common.Address, admin []common.Address) (*NegRiskCtfExchangeNewOperatorIterator, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeNewOperatorIterator{contract: _NegRiskCtfExchange.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeNewOperator, newOperatorAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeNewOperator)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "NewOperator", log); err != nil {
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

// ParseNewOperator is a log parse operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseNewOperator(log types.Log) (*NegRiskCtfExchangeNewOperator, error) {
	event := new(NegRiskCtfExchangeNewOperator)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrderCancelledIterator struct {
	Event *NegRiskCtfExchangeOrderCancelled // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeOrderCancelled)
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
		it.Event = new(NegRiskCtfExchangeOrderCancelled)
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
func (it *NegRiskCtfExchangeOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeOrderCancelled represents a OrderCancelled event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrderCancelled struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderHash [][32]byte) (*NegRiskCtfExchangeOrderCancelledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "OrderCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeOrderCancelledIterator{contract: _NegRiskCtfExchange.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeOrderCancelled, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "OrderCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeOrderCancelled)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseOrderCancelled(log types.Log) (*NegRiskCtfExchangeOrderCancelled, error) {
	event := new(NegRiskCtfExchangeOrderCancelled)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrderFilledIterator struct {
	Event *NegRiskCtfExchangeOrderFilled // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeOrderFilled)
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
		it.Event = new(NegRiskCtfExchangeOrderFilled)
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
func (it *NegRiskCtfExchangeOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeOrderFilled represents a OrderFilled event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrderFilled struct {
	OrderHash         [32]byte
	Maker             common.Address
	Taker             common.Address
	MakerAssetId      *big.Int
	TakerAssetId      *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Fee               *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterOrderFilled(opts *bind.FilterOpts, orderHash [][32]byte, maker []common.Address, taker []common.Address) (*NegRiskCtfExchangeOrderFilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeOrderFilledIterator{contract: _NegRiskCtfExchange.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeOrderFilled, orderHash [][32]byte, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeOrderFilled)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseOrderFilled(log types.Log) (*NegRiskCtfExchangeOrderFilled, error) {
	event := new(NegRiskCtfExchangeOrderFilled)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrdersMatchedIterator struct {
	Event *NegRiskCtfExchangeOrdersMatched // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeOrdersMatched)
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
		it.Event = new(NegRiskCtfExchangeOrdersMatched)
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
func (it *NegRiskCtfExchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeOrdersMatched represents a OrdersMatched event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrdersMatched struct {
	TakerOrderHash    [32]byte
	TakerOrderMaker   common.Address
	MakerAssetId      *big.Int
	TakerAssetId      *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (*NegRiskCtfExchangeOrdersMatchedIterator, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeOrdersMatchedIterator{contract: _NegRiskCtfExchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeOrdersMatched, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (event.Subscription, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeOrdersMatched)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseOrdersMatched(log types.Log) (*NegRiskCtfExchangeOrdersMatched, error) {
	event := new(NegRiskCtfExchangeOrdersMatched)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeProxyFactoryUpdatedIterator is returned from FilterProxyFactoryUpdated and is used to iterate over the raw logs and unpacked data for ProxyFactoryUpdated events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeProxyFactoryUpdatedIterator struct {
	Event *NegRiskCtfExchangeProxyFactoryUpdated // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeProxyFactoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeProxyFactoryUpdated)
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
		it.Event = new(NegRiskCtfExchangeProxyFactoryUpdated)
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
func (it *NegRiskCtfExchangeProxyFactoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeProxyFactoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeProxyFactoryUpdated represents a ProxyFactoryUpdated event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeProxyFactoryUpdated struct {
	OldProxyFactory common.Address
	NewProxyFactory common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProxyFactoryUpdated is a free log retrieval operation binding the contract event 0x3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a37.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterProxyFactoryUpdated(opts *bind.FilterOpts, oldProxyFactory []common.Address, newProxyFactory []common.Address) (*NegRiskCtfExchangeProxyFactoryUpdatedIterator, error) {

	var oldProxyFactoryRule []interface{}
	for _, oldProxyFactoryItem := range oldProxyFactory {
		oldProxyFactoryRule = append(oldProxyFactoryRule, oldProxyFactoryItem)
	}
	var newProxyFactoryRule []interface{}
	for _, newProxyFactoryItem := range newProxyFactory {
		newProxyFactoryRule = append(newProxyFactoryRule, newProxyFactoryItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "ProxyFactoryUpdated", oldProxyFactoryRule, newProxyFactoryRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeProxyFactoryUpdatedIterator{contract: _NegRiskCtfExchange.contract, event: "ProxyFactoryUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyFactoryUpdated is a free log subscription operation binding the contract event 0x3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a37.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchProxyFactoryUpdated(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeProxyFactoryUpdated, oldProxyFactory []common.Address, newProxyFactory []common.Address) (event.Subscription, error) {

	var oldProxyFactoryRule []interface{}
	for _, oldProxyFactoryItem := range oldProxyFactory {
		oldProxyFactoryRule = append(oldProxyFactoryRule, oldProxyFactoryItem)
	}
	var newProxyFactoryRule []interface{}
	for _, newProxyFactoryItem := range newProxyFactory {
		newProxyFactoryRule = append(newProxyFactoryRule, newProxyFactoryItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "ProxyFactoryUpdated", oldProxyFactoryRule, newProxyFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeProxyFactoryUpdated)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "ProxyFactoryUpdated", log); err != nil {
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

// ParseProxyFactoryUpdated is a log parse operation binding the contract event 0x3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a37.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseProxyFactoryUpdated(log types.Log) (*NegRiskCtfExchangeProxyFactoryUpdated, error) {
	event := new(NegRiskCtfExchangeProxyFactoryUpdated)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "ProxyFactoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeRemovedAdminIterator struct {
	Event *NegRiskCtfExchangeRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeRemovedAdmin)
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
		it.Event = new(NegRiskCtfExchangeRemovedAdmin)
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
func (it *NegRiskCtfExchangeRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeRemovedAdmin represents a RemovedAdmin event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeRemovedAdmin struct {
	RemovedAdmin common.Address
	Admin        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, removedAdmin []common.Address, admin []common.Address) (*NegRiskCtfExchangeRemovedAdminIterator, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeRemovedAdminIterator{contract: _NegRiskCtfExchange.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeRemovedAdmin, removedAdmin []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeRemovedAdmin)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseRemovedAdmin(log types.Log) (*NegRiskCtfExchangeRemovedAdmin, error) {
	event := new(NegRiskCtfExchangeRemovedAdmin)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeRemovedOperatorIterator is returned from FilterRemovedOperator and is used to iterate over the raw logs and unpacked data for RemovedOperator events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeRemovedOperatorIterator struct {
	Event *NegRiskCtfExchangeRemovedOperator // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeRemovedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeRemovedOperator)
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
		it.Event = new(NegRiskCtfExchangeRemovedOperator)
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
func (it *NegRiskCtfExchangeRemovedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeRemovedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeRemovedOperator represents a RemovedOperator event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeRemovedOperator struct {
	RemovedOperator common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRemovedOperator is a free log retrieval operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterRemovedOperator(opts *bind.FilterOpts, removedOperator []common.Address, admin []common.Address) (*NegRiskCtfExchangeRemovedOperatorIterator, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeRemovedOperatorIterator{contract: _NegRiskCtfExchange.contract, event: "RemovedOperator", logs: logs, sub: sub}, nil
}

// WatchRemovedOperator is a free log subscription operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchRemovedOperator(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeRemovedOperator, removedOperator []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeRemovedOperator)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
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

// ParseRemovedOperator is a log parse operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseRemovedOperator(log types.Log) (*NegRiskCtfExchangeRemovedOperator, error) {
	event := new(NegRiskCtfExchangeRemovedOperator)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeSafeFactoryUpdatedIterator is returned from FilterSafeFactoryUpdated and is used to iterate over the raw logs and unpacked data for SafeFactoryUpdated events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeSafeFactoryUpdatedIterator struct {
	Event *NegRiskCtfExchangeSafeFactoryUpdated // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeSafeFactoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeSafeFactoryUpdated)
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
		it.Event = new(NegRiskCtfExchangeSafeFactoryUpdated)
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
func (it *NegRiskCtfExchangeSafeFactoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeSafeFactoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeSafeFactoryUpdated represents a SafeFactoryUpdated event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeSafeFactoryUpdated struct {
	OldSafeFactory common.Address
	NewSafeFactory common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSafeFactoryUpdated is a free log retrieval operation binding the contract event 0x9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterSafeFactoryUpdated(opts *bind.FilterOpts, oldSafeFactory []common.Address, newSafeFactory []common.Address) (*NegRiskCtfExchangeSafeFactoryUpdatedIterator, error) {

	var oldSafeFactoryRule []interface{}
	for _, oldSafeFactoryItem := range oldSafeFactory {
		oldSafeFactoryRule = append(oldSafeFactoryRule, oldSafeFactoryItem)
	}
	var newSafeFactoryRule []interface{}
	for _, newSafeFactoryItem := range newSafeFactory {
		newSafeFactoryRule = append(newSafeFactoryRule, newSafeFactoryItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "SafeFactoryUpdated", oldSafeFactoryRule, newSafeFactoryRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeSafeFactoryUpdatedIterator{contract: _NegRiskCtfExchange.contract, event: "SafeFactoryUpdated", logs: logs, sub: sub}, nil
}

// WatchSafeFactoryUpdated is a free log subscription operation binding the contract event 0x9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchSafeFactoryUpdated(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeSafeFactoryUpdated, oldSafeFactory []common.Address, newSafeFactory []common.Address) (event.Subscription, error) {

	var oldSafeFactoryRule []interface{}
	for _, oldSafeFactoryItem := range oldSafeFactory {
		oldSafeFactoryRule = append(oldSafeFactoryRule, oldSafeFactoryItem)
	}
	var newSafeFactoryRule []interface{}
	for _, newSafeFactoryItem := range newSafeFactory {
		newSafeFactoryRule = append(newSafeFactoryRule, newSafeFactoryItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "SafeFactoryUpdated", oldSafeFactoryRule, newSafeFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeSafeFactoryUpdated)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "SafeFactoryUpdated", log); err != nil {
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

// ParseSafeFactoryUpdated is a log parse operation binding the contract event 0x9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseSafeFactoryUpdated(log types.Log) (*NegRiskCtfExchangeSafeFactoryUpdated, error) {
	event := new(NegRiskCtfExchangeSafeFactoryUpdated)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "SafeFactoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTokenRegisteredIterator struct {
	Event *NegRiskCtfExchangeTokenRegistered // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeTokenRegistered)
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
		it.Event = new(NegRiskCtfExchangeTokenRegistered)
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
func (it *NegRiskCtfExchangeTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeTokenRegistered represents a TokenRegistered event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTokenRegistered struct {
	Token0      *big.Int
	Token1      *big.Int
	ConditionId [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0xbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterTokenRegistered(opts *bind.FilterOpts, token0 []*big.Int, token1 []*big.Int, conditionId [][32]byte) (*NegRiskCtfExchangeTokenRegisteredIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "TokenRegistered", token0Rule, token1Rule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeTokenRegisteredIterator{contract: _NegRiskCtfExchange.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeTokenRegistered, token0 []*big.Int, token1 []*big.Int, conditionId [][32]byte) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "TokenRegistered", token0Rule, token1Rule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeTokenRegistered)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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

// ParseTokenRegistered is a log parse operation binding the contract event 0xbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseTokenRegistered(log types.Log) (*NegRiskCtfExchangeTokenRegistered, error) {
	event := new(NegRiskCtfExchangeTokenRegistered)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeTradingPausedIterator is returned from FilterTradingPaused and is used to iterate over the raw logs and unpacked data for TradingPaused events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTradingPausedIterator struct {
	Event *NegRiskCtfExchangeTradingPaused // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeTradingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeTradingPaused)
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
		it.Event = new(NegRiskCtfExchangeTradingPaused)
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
func (it *NegRiskCtfExchangeTradingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeTradingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeTradingPaused represents a TradingPaused event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTradingPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingPaused is a free log retrieval operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterTradingPaused(opts *bind.FilterOpts, pauser []common.Address) (*NegRiskCtfExchangeTradingPausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeTradingPausedIterator{contract: _NegRiskCtfExchange.contract, event: "TradingPaused", logs: logs, sub: sub}, nil
}

// WatchTradingPaused is a free log subscription operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchTradingPaused(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeTradingPaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeTradingPaused)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "TradingPaused", log); err != nil {
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

// ParseTradingPaused is a log parse operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseTradingPaused(log types.Log) (*NegRiskCtfExchangeTradingPaused, error) {
	event := new(NegRiskCtfExchangeTradingPaused)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "TradingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfExchangeTradingUnpausedIterator is returned from FilterTradingUnpaused and is used to iterate over the raw logs and unpacked data for TradingUnpaused events raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTradingUnpausedIterator struct {
	Event *NegRiskCtfExchangeTradingUnpaused // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfExchangeTradingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfExchangeTradingUnpaused)
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
		it.Event = new(NegRiskCtfExchangeTradingUnpaused)
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
func (it *NegRiskCtfExchangeTradingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfExchangeTradingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfExchangeTradingUnpaused represents a TradingUnpaused event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTradingUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingUnpaused is a free log retrieval operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) FilterTradingUnpaused(opts *bind.FilterOpts, pauser []common.Address) (*NegRiskCtfExchangeTradingUnpausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.FilterLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfExchangeTradingUnpausedIterator{contract: _NegRiskCtfExchange.contract, event: "TradingUnpaused", logs: logs, sub: sub}, nil
}

// WatchTradingUnpaused is a free log subscription operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) WatchTradingUnpaused(opts *bind.WatchOpts, sink chan<- *NegRiskCtfExchangeTradingUnpaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskCtfExchange.contract.WatchLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfExchangeTradingUnpaused)
				if err := _NegRiskCtfExchange.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
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

// ParseTradingUnpaused is a log parse operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_NegRiskCtfExchange *NegRiskCtfExchangeFilterer) ParseTradingUnpaused(log types.Log) (*NegRiskCtfExchangeTradingUnpaused, error) {
	event := new(NegRiskCtfExchangeTradingUnpaused)
	if err := _NegRiskCtfExchange.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

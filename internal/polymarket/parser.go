package polymarket

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"polymarket_leaders/internal/contract_bindings/conditionaltokens"
	"polymarket_leaders/internal/contract_bindings/ctfexchange"
	"polymarket_leaders/internal/contract_bindings/negriskadapter"
	"polymarket_leaders/internal/contract_bindings/negriskctfexchange"
)

// ContractInfo contains contract name and ABI
type ContractInfo struct {
	Name string
	ABI  *abi.ABI
}

// ParsedEvent contains event name and parsed data
type ParsedEvent struct {
	Name string
	Data interface{}
}

// TransactionWithEvents contains a transaction and its parsed events
type TransactionWithEvents struct {
	Tx     *types.Transaction
	Events []ParsedEvent
}

// ProcessBlock processes a single block and returns Polymarket transactions with events
func ProcessBlock(ctx context.Context, client *ethclient.Client, blockNumber uint64, contracts map[common.Address]ContractInfo, signer types.Signer) ([]TransactionWithEvents, bool) {
	// Get the full block with transactions
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, false
	}

	// Initialize contract filterers for event parsing
	ctfExchangeFilterer, _ := ctfexchange.NewCtfExchangeFilterer(common.Address{}, nil)
	negRiskCtfExchangeFilterer, _ := negriskctfexchange.NewNegRiskCtfExchangeFilterer(common.Address{}, nil)
	conditionalTokensFilterer, _ := conditionaltokens.NewConditionalTokensFilterer(common.Address{}, nil)
	negRiskAdapterFilterer, _ := negriskadapter.NewNegRiskAdapterFilterer(common.Address{}, nil)

	var polymarketTxs []TransactionWithEvents

	for _, tx := range block.Transactions() {
		// Get transaction receipt to access logs
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			continue
		}

		var foundEvents []ParsedEvent

		// Check logs for events we're interested in
		for _, log := range receipt.Logs {
			// Check for OrderFilled event (CTFExchange)
			if event, err := ctfExchangeFilterer.ParseOrderFilled(*log); err == nil {
				foundEvents = append(foundEvents, ParsedEvent{
					Name: "OrderFilled (CTFExchange)",
					Data: event,
				})
			}
			// Check for OrderFilled event (NegRiskCtfExchange)
			if event, err := negRiskCtfExchangeFilterer.ParseOrderFilled(*log); err == nil {
				foundEvents = append(foundEvents, ParsedEvent{
					Name: "OrderFilled (NegRiskCtfExchange)",
					Data: event,
				})
			}

			// Check for OrdersMatched event (CTFExchange)
			if event, err := ctfExchangeFilterer.ParseOrdersMatched(*log); err == nil {
				foundEvents = append(foundEvents, ParsedEvent{
					Name: "OrdersMatched (CTFExchange)",
					Data: event,
				})
			}
			// Check for OrdersMatched event (NegRiskCtfExchange)
			if event, err := negRiskCtfExchangeFilterer.ParseOrdersMatched(*log); err == nil {
				foundEvents = append(foundEvents, ParsedEvent{
					Name: "OrdersMatched (NegRiskCtfExchange)",
					Data: event,
				})
			}

			// Check for PositionSplit event (ConditionalTokens)
			if event, err := conditionalTokensFilterer.ParsePositionSplit(*log); err == nil {
				foundEvents = append(foundEvents, ParsedEvent{
					Name: "PositionSplit (ConditionalTokens)",
					Data: event,
				})
			}
			// Check for PositionSplit event (NegRiskAdapter)
			if event, err := negRiskAdapterFilterer.ParsePositionSplit(*log); err == nil {
				foundEvents = append(foundEvents, ParsedEvent{
					Name: "PositionSplit (NegRiskAdapter)",
					Data: event,
				})
			}

			// Check for PositionsMerge event (ConditionalTokens)
			if event, err := conditionalTokensFilterer.ParsePositionsMerge(*log); err == nil {
				foundEvents = append(foundEvents, ParsedEvent{
					Name: "PositionsMerge (ConditionalTokens)",
					Data: event,
				})
			}
			// Check for PositionsMerge event (NegRiskAdapter)
			if event, err := negRiskAdapterFilterer.ParsePositionsMerge(*log); err == nil {
				foundEvents = append(foundEvents, ParsedEvent{
					Name: "PositionsMerge (NegRiskAdapter)",
					Data: event,
				})
			}

			// Check for PositionsConverted event (NegRiskAdapter)
			if event, err := negRiskAdapterFilterer.ParsePositionsConverted(*log); err == nil {
				foundEvents = append(foundEvents, ParsedEvent{
					Name: "PositionsConverted (NegRiskAdapter)",
					Data: event,
				})
			}
		}

		// If we found any of the events we're interested in, include this transaction
		if len(foundEvents) > 0 {
			polymarketTxs = append(polymarketTxs, TransactionWithEvents{Tx: tx, Events: foundEvents})
		}
	}

	return polymarketTxs, true
}

// DecodeTransactionData decodes transaction data using the contract ABI
func DecodeTransactionData(data []byte, contractABI *abi.ABI) (string, []string) {
	if contractABI == nil || len(data) < 4 {
		return "", nil
	}

	// Extract method selector (first 4 bytes)
	methodSelector := data[:4]

	// Find the method by selector
	var method *abi.Method
	for _, m := range contractABI.Methods {
		if len(m.ID) == 4 && string(m.ID) == string(methodSelector) {
			method = &m
			break
		}
	}

	if method == nil {
		return "", nil
	}

	// Decode the arguments
	args, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return method.Name, []string{fmt.Sprintf("(decoding error: %v)", err)}
	}

	// Format arguments
	argStrings := make([]string, 0, len(method.Inputs))
	for i, input := range method.Inputs {
		if i < len(args) {
			argValue := formatArgumentValue(args[i], input.Type)
			argStrings = append(argStrings, fmt.Sprintf("%s: %s", input.Name, argValue))
		}
	}

	return method.Name, argStrings
}

// PrintEventData prints the parsed event data in a readable format
func PrintEventData(eventData interface{}, indent string) {
	switch e := eventData.(type) {
	// CTFExchange events
	case *ctfexchange.CtfExchangeOrderFilled:
		fmt.Printf("%s  OrderHash: 0x%x\n", indent, e.OrderHash)
		fmt.Printf("%s  Maker: %s\n", indent, e.Maker.Hex())
		fmt.Printf("%s  Taker: %s\n", indent, e.Taker.Hex())
		fmt.Printf("%s  MakerAssetId: %s\n", indent, e.MakerAssetId.String())
		fmt.Printf("%s  TakerAssetId: %s\n", indent, e.TakerAssetId.String())
		fmt.Printf("%s  MakerAmountFilled: %s\n", indent, e.MakerAmountFilled.String())
		fmt.Printf("%s  TakerAmountFilled: %s\n", indent, e.TakerAmountFilled.String())
		fmt.Printf("%s  Fee: %s\n", indent, e.Fee.String())
	case *ctfexchange.CtfExchangeOrdersMatched:
		fmt.Printf("%s  TakerOrderHash: 0x%x\n", indent, e.TakerOrderHash)
		fmt.Printf("%s  TakerOrderMaker: %s\n", indent, e.TakerOrderMaker.Hex())
		fmt.Printf("%s  MakerAssetId: %s\n", indent, e.MakerAssetId.String())
		fmt.Printf("%s  TakerAssetId: %s\n", indent, e.TakerAssetId.String())
		fmt.Printf("%s  MakerAmountFilled: %s\n", indent, e.MakerAmountFilled.String())
		fmt.Printf("%s  TakerAmountFilled: %s\n", indent, e.TakerAmountFilled.String())

	// NegRiskCtfExchange events
	case *negriskctfexchange.NegRiskCtfExchangeOrderFilled:
		fmt.Printf("%s  OrderHash: 0x%x\n", indent, e.OrderHash)
		fmt.Printf("%s  Maker: %s\n", indent, e.Maker.Hex())
		fmt.Printf("%s  Taker: %s\n", indent, e.Taker.Hex())
		fmt.Printf("%s  MakerAssetId: %s\n", indent, e.MakerAssetId.String())
		fmt.Printf("%s  TakerAssetId: %s\n", indent, e.TakerAssetId.String())
		fmt.Printf("%s  MakerAmountFilled: %s\n", indent, e.MakerAmountFilled.String())
		fmt.Printf("%s  TakerAmountFilled: %s\n", indent, e.TakerAmountFilled.String())
		fmt.Printf("%s  Fee: %s\n", indent, e.Fee.String())
	case *negriskctfexchange.NegRiskCtfExchangeOrdersMatched:
		fmt.Printf("%s  TakerOrderHash: 0x%x\n", indent, e.TakerOrderHash)
		fmt.Printf("%s  TakerOrderMaker: %s\n", indent, e.TakerOrderMaker.Hex())
		fmt.Printf("%s  MakerAssetId: %s\n", indent, e.MakerAssetId.String())
		fmt.Printf("%s  TakerAssetId: %s\n", indent, e.TakerAssetId.String())
		fmt.Printf("%s  MakerAmountFilled: %s\n", indent, e.MakerAmountFilled.String())
		fmt.Printf("%s  TakerAmountFilled: %s\n", indent, e.TakerAmountFilled.String())

	// ConditionalTokens events
	case *conditionaltokens.ConditionalTokensPositionSplit:
		fmt.Printf("%s  Stakeholder: %s\n", indent, e.Stakeholder.Hex())
		fmt.Printf("%s  CollateralToken: %s\n", indent, e.CollateralToken.Hex())
		fmt.Printf("%s  ParentCollectionId: 0x%x\n", indent, e.ParentCollectionId)
		fmt.Printf("%s  ConditionId: 0x%x\n", indent, e.ConditionId)
		fmt.Printf("%s  Partition: %s\n", indent, formatBigIntArray(e.Partition))
		fmt.Printf("%s  Amount: %s\n", indent, e.Amount.String())
	case *conditionaltokens.ConditionalTokensPositionsMerge:
		fmt.Printf("%s  Stakeholder: %s\n", indent, e.Stakeholder.Hex())
		fmt.Printf("%s  CollateralToken: %s\n", indent, e.CollateralToken.Hex())
		fmt.Printf("%s  ParentCollectionId: 0x%x\n", indent, e.ParentCollectionId)
		fmt.Printf("%s  ConditionId: 0x%x\n", indent, e.ConditionId)
		fmt.Printf("%s  Partition: %s\n", indent, formatBigIntArray(e.Partition))
		fmt.Printf("%s  Amount: %s\n", indent, e.Amount.String())

	// NegRiskAdapter events
	case *negriskadapter.NegRiskAdapterPositionSplit:
		fmt.Printf("%s  Stakeholder: %s\n", indent, e.Stakeholder.Hex())
		fmt.Printf("%s  ConditionId: 0x%x\n", indent, e.ConditionId)
		fmt.Printf("%s  Amount: %s\n", indent, e.Amount.String())
	case *negriskadapter.NegRiskAdapterPositionsMerge:
		fmt.Printf("%s  Stakeholder: %s\n", indent, e.Stakeholder.Hex())
		fmt.Printf("%s  ConditionId: 0x%x\n", indent, e.ConditionId)
		fmt.Printf("%s  Amount: %s\n", indent, e.Amount.String())
	case *negriskadapter.NegRiskAdapterPositionsConverted:
		fmt.Printf("%s  Stakeholder: %s\n", indent, e.Stakeholder.Hex())
		fmt.Printf("%s  MarketId: 0x%x\n", indent, e.MarketId)
		fmt.Printf("%s  IndexSet: %s\n", indent, e.IndexSet.String())
		fmt.Printf("%s  Amount: %s\n", indent, e.Amount.String())
	}
}

// formatArgumentValue formats an argument value for display
func formatArgumentValue(value interface{}, typ abi.Type) string {
	switch v := value.(type) {
	case *big.Int:
		return v.String()
	case common.Address:
		return v.Hex()
	case []byte:
		return fmt.Sprintf("0x%x", v)
	case []interface{}:
		// Handle arrays
		items := make([]string, 0, len(v))
		for _, item := range v {
			items = append(items, formatArgumentValue(item, abi.Type{}))
		}
		return fmt.Sprintf("[%s]", strings.Join(items, ", "))
	case bool:
		return fmt.Sprintf("%t", v)
	case string:
		return fmt.Sprintf("\"%s\"", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// formatBigIntArray formats an array of big.Int for display
func formatBigIntArray(arr []*big.Int) string {
	if len(arr) == 0 {
		return "[]"
	}
	items := make([]string, 0, len(arr))
	for _, item := range arr {
		items = append(items, item.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(items, ", "))
}

// GetContractABI extracts the ABI from contract metadata
func GetContractABI(metadata *bind.MetaData) *abi.ABI {
	if metadata == nil {
		return nil
	}
	contractABI, err := metadata.GetAbi()
	if err != nil {
		log.Printf("Failed to get ABI: %v", err)
		return nil
	}
	return contractABI
}


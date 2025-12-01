package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"polymarket_leaders/internal/conditionaltokens"
	"polymarket_leaders/internal/ctfexchange"
	"polymarket_leaders/internal/umactfadapter"
	"polymarket_leaders/internal/negriskctfexchange"
	"polymarket_leaders/internal/negriskadapter"
	"polymarket_leaders/internal/uchilderc20proxy"
	"polymarket_leaders/internal/umactfadapter2"
	"polymarket_leaders/internal/safeproxyfactory"
	"polymarket_leaders/internal/proxywalletfactory"
	"polymarket_leaders/internal/negriskfeemodule"
)

type contractInfo struct {
	name string
	abi  *abi.ABI
}

func main() {
	// Polygon RPC endpoint (using public endpoint)
	rpcURL := "https://polygon-rpc.com"

	// Connect to Polygon RPC
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Polygon RPC: %v", err)
	}
	defer client.Close()

	fmt.Println("Successfully connected to Polygon RPC")
	fmt.Println("Starting continuous block monitoring...\n")

	// Define Polymarket contract addresses and their ABIs
	polymarketContracts := map[common.Address]contractInfo{
		common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"): {
			name: "ConditionalTokens",
			abi:  getContractABI(conditionaltokens.ConditionalTokensMetaData),
		},
		common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"): {
			name: "CTFExchange",
			abi:  getContractABI(ctfexchange.CtfExchangeMetaData),
		},
		common.HexToAddress("0x65070BE91477460D8A7AeEb94ef92fe056C2f2A7"): {
			name: "UmaCtfAdapter",
			abi:  getContractABI(umactfadapter.UmaCtfAdapterMetaData),
		},
		common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"): {
			name: "NegRiskCtfExchange",
			abi:  getContractABI(negriskctfexchange.NegRiskCtfExchangeMetaData),
		},
		common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"): {
			name: "NegRiskAdapter",
			abi:  getContractABI(negriskadapter.NegRiskAdapterMetaData),
		},
		common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"): {
			name: "UChildERC20Proxy",
			abi:  getContractABI(uchilderc20proxy.UchaildErc20ProxyMetaData),
		},
		common.HexToAddress("0x6A9D222616C90FcA5754cd1333cFD9b7fb6a4F74"): {
			name: "UmaCtfAdapter2",
			abi:  getContractABI(umactfadapter2.UmaCtfAdapter2MetaData),
		},
		common.HexToAddress("0xaacfeea03eb1561c4e67d661e40682bd20e3541b"): {
			name: "SafeProxyFactory",
			abi:  getContractABI(safeproxyfactory.SafeProxyFactoryMetaData),
		},
		common.HexToAddress("0xaB45c5A4B0c941a2F231C04C3f49182e1A254052"): {
			name: "ProxyWalletFactory",
			abi:  getContractABI(proxywalletfactory.ProxyWalletFactoryMetaData),
		},
		common.HexToAddress("0x78769D50Be1763ed1CA0D5E878D93f05aabff29e"): {
			name: "NegRiskFeeModule",
			abi:  getContractABI(negriskfeemodule.NegRiskFeeModuleMetaData),
		},
	}

	// Polygon chain ID is 137
	chainID := big.NewInt(137)
	signer := types.LatestSignerForChainID(chainID)

	// Get current block number to start from
	ctx := context.Background()
	currentBlock, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatalf("Failed to get current block number: %v", err)
	}

	fmt.Printf("Starting from block: %d\n", currentBlock)
	fmt.Println("Monitoring for new blocks...\n")

	// Endless loop to monitor blocks
	for {
		// Process current block
		for !processBlock(ctx, client, currentBlock, polymarketContracts, signer) {
			// Wait before checking again
			time.Sleep(1 * time.Second)
		}

		// Wait for next block
		currentBlock++
		fmt.Printf("Block: %d\n", currentBlock)

		// Poll for next block (Polygon blocks are ~2 seconds, check every 1 second)
		// Wait before checking again
		time.Sleep(1 * time.Second)
	}
}

// processBlock processes a single block and prints Polymarket transactions
func processBlock(ctx context.Context, client *ethclient.Client, blockNumber uint64, polymarketContracts map[common.Address]contractInfo, signer types.Signer) bool {
	// Get the full block with transactions
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return false
	}

	// Initialize contract filterers for event parsing
	ctfExchangeFilterer, _ := ctfexchange.NewCtfExchangeFilterer(common.Address{}, nil)
	negRiskCtfExchangeFilterer, _ := negriskctfexchange.NewNegRiskCtfExchangeFilterer(common.Address{}, nil)
	conditionalTokensFilterer, _ := conditionaltokens.NewConditionalTokensFilterer(common.Address{}, nil)
	negRiskAdapterFilterer, _ := negriskadapter.NewNegRiskAdapterFilterer(common.Address{}, nil)

	// Filter transactions by checking their logs for specific events
	type parsedEvent struct {
		name string
		data interface{}
	}
	type txWithEvents struct {
		tx     *types.Transaction
		events []parsedEvent
	}
	var polymarketTxs []txWithEvents

	for _, tx := range block.Transactions() {
		// Get transaction receipt to access logs
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			continue
		}

		var foundEvents []parsedEvent

		// Check logs for events we're interested in
		for _, log := range receipt.Logs {
			// Check for OrderFilled event (CTFExchange)
			if event, err := ctfExchangeFilterer.ParseOrderFilled(*log); err == nil {
				foundEvents = append(foundEvents, parsedEvent{
					name: "OrderFilled (CTFExchange)",
					data: event,
				})
			}
			// Check for OrderFilled event (NegRiskCtfExchange)
			if event, err := negRiskCtfExchangeFilterer.ParseOrderFilled(*log); err == nil {
				foundEvents = append(foundEvents, parsedEvent{
					name: "OrderFilled (NegRiskCtfExchange)",
					data: event,
				})
			}

			// Check for OrdersMatched event (CTFExchange)
			if event, err := ctfExchangeFilterer.ParseOrdersMatched(*log); err == nil {
				foundEvents = append(foundEvents, parsedEvent{
					name: "OrdersMatched (CTFExchange)",
					data: event,
				})
			}
			// Check for OrdersMatched event (NegRiskCtfExchange)
			if event, err := negRiskCtfExchangeFilterer.ParseOrdersMatched(*log); err == nil {
				foundEvents = append(foundEvents, parsedEvent{
					name: "OrdersMatched (NegRiskCtfExchange)",
					data: event,
				})
			}

			// Check for PositionSplit event (ConditionalTokens)
			if event, err := conditionalTokensFilterer.ParsePositionSplit(*log); err == nil {
				foundEvents = append(foundEvents, parsedEvent{
					name: "PositionSplit (ConditionalTokens)",
					data: event,
				})
			}
			// Check for PositionSplit event (NegRiskAdapter)
			if event, err := negRiskAdapterFilterer.ParsePositionSplit(*log); err == nil {
				foundEvents = append(foundEvents, parsedEvent{
					name: "PositionSplit (NegRiskAdapter)",
					data: event,
				})
			}

			// Check for PositionsMerge event (ConditionalTokens)
			if event, err := conditionalTokensFilterer.ParsePositionsMerge(*log); err == nil {
				foundEvents = append(foundEvents, parsedEvent{
					name: "PositionsMerge (ConditionalTokens)",
					data: event,
				})
			}
			// Check for PositionsMerge event (NegRiskAdapter)
			if event, err := negRiskAdapterFilterer.ParsePositionsMerge(*log); err == nil {
				foundEvents = append(foundEvents, parsedEvent{
					name: "PositionsMerge (NegRiskAdapter)",
					data: event,
				})
			}

			// Check for PositionsConverted event (NegRiskAdapter)
			if event, err := negRiskAdapterFilterer.ParsePositionsConverted(*log); err == nil {
				foundEvents = append(foundEvents, parsedEvent{
					name: "PositionsConverted (NegRiskAdapter)",
					data: event,
				})
			}
		}

		// If we found any of the events we're interested in, include this transaction
		if len(foundEvents) > 0 {
			polymarketTxs = append(polymarketTxs, txWithEvents{tx: tx, events: foundEvents})
		}
	}

	// Print only Polymarket transactions
	if len(polymarketTxs) == 0 {
		return true
	}

	for i, txData := range polymarketTxs {
		tx := txData.tx
		fmt.Printf("\n[Polymarket Transaction %d]\n", i+1)
		fmt.Printf("  Block #%d\n", blockNumber)
		fmt.Printf("  Hash: %s\n", tx.Hash().Hex())

		// Get sender address from transaction
		sender, err := types.Sender(signer, tx)
		if err != nil {
			fmt.Printf("  From: <unable to determine>\n")
		} else {
			fmt.Printf("  From: %s\n", sender.Hex())
		}

		if tx.To() != nil {
			fmt.Printf("  To: %s\n", tx.To().Hex())
			if contractInfo, ok := polymarketContracts[*tx.To()]; ok {
				fmt.Printf("  Contract: %s\n", contractInfo.name)
			}
		} else {
			fmt.Printf("  To: <Contract Creation>\n")
		}

		fmt.Printf("  Value: %s MATIC\n", formatWei(tx.Value()))
		fmt.Printf("  Gas: %d\n", tx.Gas())
		fmt.Printf("  Gas Price: %s Gwei\n", formatGwei(tx.GasPrice()))
		fmt.Printf("  Nonce: %d\n", tx.Nonce())

		// Print events found in logs with parsed data
		fmt.Printf("  Events:\n")
		for _, event := range txData.events {
			fmt.Printf("    - %s\n", event.name)
			printEventData(event.data, "      ")
		}

		// Decode transaction data if we can identify the contract
		if tx.To() != nil {
			if contractInfo, ok := polymarketContracts[*tx.To()]; ok {
				if tx.Data() != nil && len(tx.Data()) > 0 {
					methodName, args := decodeTransactionData(tx.Data(), contractInfo.abi)
					if methodName != "" {
						fmt.Printf("  Method: %s\n", methodName)
						if len(args) > 0 {
							fmt.Printf("  Arguments:\n")
							for _, arg := range args {
								fmt.Printf("    %s\n", arg)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println()
	return true
}

// formatWei converts wei to MATIC (1 MATIC = 10^18 wei)
func formatWei(wei *big.Int) string {
	matic := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return matic.Text('f', 6)
}

// formatGwei converts wei to Gwei (1 Gwei = 10^9 wei)
func formatGwei(wei *big.Int) string {
	gwei := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e9))
	return gwei.Text('f', 2)
}

// getContractABI extracts the ABI from contract metadata
func getContractABI(metadata *bind.MetaData) *abi.ABI {
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

// decodeTransactionData decodes transaction data using the contract ABI
func decodeTransactionData(data []byte, contractABI *abi.ABI) (string, []string) {
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

// printEventData prints the parsed event data in a readable format
func printEventData(eventData interface{}, indent string) {
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

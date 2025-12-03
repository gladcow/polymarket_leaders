package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"polymarket_leaders/internal/contract_bindings/conditionaltokens"
	"polymarket_leaders/internal/contract_bindings/ctfexchange"
	"polymarket_leaders/internal/contract_bindings/negriskadapter"
	"polymarket_leaders/internal/contract_bindings/negriskctfexchange"
	"polymarket_leaders/internal/contract_bindings/negriskfeemodule"
	"polymarket_leaders/internal/contract_bindings/proxywalletfactory"
	"polymarket_leaders/internal/contract_bindings/safeproxyfactory"
	"polymarket_leaders/internal/contract_bindings/uchilderc20proxy"
	"polymarket_leaders/internal/contract_bindings/umactfadapter"
	"polymarket_leaders/internal/contract_bindings/umactfadapter2"
	"polymarket_leaders/internal/polymarket"
)

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
	polymarketContracts := map[common.Address]polymarket.ContractInfo{
		common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"): {
			Name: "ConditionalTokens",
			ABI:  polymarket.GetContractABI(conditionaltokens.ConditionalTokensMetaData),
		},
		common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"): {
			Name: "CTFExchange",
			ABI:  polymarket.GetContractABI(ctfexchange.CtfExchangeMetaData),
		},
		common.HexToAddress("0x65070BE91477460D8A7AeEb94ef92fe056C2f2A7"): {
			Name: "UmaCtfAdapter",
			ABI:  polymarket.GetContractABI(umactfadapter.UmaCtfAdapterMetaData),
		},
		common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"): {
			Name: "NegRiskCtfExchange",
			ABI:  polymarket.GetContractABI(negriskctfexchange.NegRiskCtfExchangeMetaData),
		},
		common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"): {
			Name: "NegRiskAdapter",
			ABI:  polymarket.GetContractABI(negriskadapter.NegRiskAdapterMetaData),
		},
		common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"): {
			Name: "UChildERC20Proxy",
			ABI:  polymarket.GetContractABI(uchilderc20proxy.UchaildErc20ProxyMetaData),
		},
		common.HexToAddress("0x6A9D222616C90FcA5754cd1333cFD9b7fb6a4F74"): {
			Name: "UmaCtfAdapter2",
			ABI:  polymarket.GetContractABI(umactfadapter2.UmaCtfAdapter2MetaData),
		},
		common.HexToAddress("0xaacfeea03eb1561c4e67d661e40682bd20e3541b"): {
			Name: "SafeProxyFactory",
			ABI:  polymarket.GetContractABI(safeproxyfactory.SafeProxyFactoryMetaData),
		},
		common.HexToAddress("0xaB45c5A4B0c941a2F231C04C3f49182e1A254052"): {
			Name: "ProxyWalletFactory",
			ABI:  polymarket.GetContractABI(proxywalletfactory.ProxyWalletFactoryMetaData),
		},
		common.HexToAddress("0x78769D50Be1763ed1CA0D5E878D93f05aabff29e"): {
			Name: "NegRiskFeeModule",
			ABI:  polymarket.GetContractABI(negriskfeemodule.NegRiskFeeModuleMetaData),
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
func processBlock(ctx context.Context, client *ethclient.Client, blockNumber uint64, contracts map[common.Address]polymarket.ContractInfo, signer types.Signer) bool {
	// Use the polymarket package to process the block
	polymarketTxs, success := polymarket.ProcessBlock(ctx, client, blockNumber, contracts, signer)
	if !success {
		return false
	}

	// Print only Polymarket transactions
	if len(polymarketTxs) == 0 {
		return true
	}

	for i, txData := range polymarketTxs {
		tx := txData.Tx
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
			if contractInfo, ok := contracts[*tx.To()]; ok {
				fmt.Printf("  Contract: %s\n", contractInfo.Name)
			}
		} else {
			fmt.Printf("  To: <Contract Creation>\n")
		}

		fmt.Printf("  Value: %s MATIC\n", polymarket.FormatWei(tx.Value()))
		fmt.Printf("  Gas: %d\n", tx.Gas())
		fmt.Printf("  Gas Price: %s Gwei\n", polymarket.FormatGwei(tx.GasPrice()))
		fmt.Printf("  Nonce: %d\n", tx.Nonce())

		// Print events found in logs with parsed data
		fmt.Printf("  Events:\n")
		for _, event := range txData.Events {
			fmt.Printf("    - %s\n", event.Name)
			polymarket.PrintEventData(event.Data, "      ")
		}

		// Decode transaction data if we can identify the contract
		if tx.To() != nil {
			if contractInfo, ok := contracts[*tx.To()]; ok {
				if tx.Data() != nil && len(tx.Data()) > 0 {
					methodName, args := polymarket.DecodeTransactionData(tx.Data(), contractInfo.ABI)
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

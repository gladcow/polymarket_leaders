package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

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

	// Map to track address counts
	addressCounts := make(map[common.Address]int)
	startBlock := currentBlock

	// Endless loop to monitor blocks
	for {
		// Process current block
		for !processBlock(ctx, client, currentBlock, signer, addressCounts) {
			// Wait before checking again
			time.Sleep(1 * time.Second)
		}

		// Print top-10 addresses every 100 blocks
		if (currentBlock-startBlock+1)%100 == 0 {
			printTopAddresses(addressCounts, currentBlock)
		}

		// Wait for next block
		currentBlock++

		// Poll for next block (Polygon blocks are ~2 seconds, check every 1 second)
		time.Sleep(1 * time.Second)
	}
}

// processBlock processes a single block and collects addresses from Polymarket events
func processBlock(ctx context.Context, client *ethclient.Client, blockNumber uint64, signer types.Signer, addressCounts map[common.Address]int) bool {
	// Use the polymarket package to process the block
	polymarketTxs, success := polymarket.ProcessBlock(ctx, client, blockNumber, signer)
	if !success {
		return false
	}

	// Extract addresses from all events and update counts
	for _, txData := range polymarketTxs {
		for _, event := range txData.Events {
			addresses := polymarket.ExtractAddressesFromEvent(event.Data)
			for _, addr := range addresses {
				addressCounts[addr]++
			}
		}
	}

	return true
}

// printTopAddresses prints the top-10 addresses by count
func printTopAddresses(addressCounts map[common.Address]int, currentBlock uint64) {
	// Create a slice of address-count pairs for sorting
	type addressCount struct {
		address common.Address
		count   int
	}

	var pairs []addressCount
	for addr, count := range addressCounts {
		pairs = append(pairs, addressCount{address: addr, count: count})
	}

	// Sort by count (descending)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	// Print top-10
	fmt.Printf("\n=== Top 10 Addresses (Block %d) ===\n", currentBlock)
	topN := 10
	if len(pairs) < topN {
		topN = len(pairs)
	}
	for i := 0; i < topN; i++ {
		fmt.Printf("%2d. %s: %d\n", i+1, pairs[i].address.Hex(), pairs[i].count)
	}
	fmt.Println()
}

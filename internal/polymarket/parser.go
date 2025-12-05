// Copyright 2025 gladcow
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package polymarket

import (
	"context"
	"log"
	"math/big"
	"polymarket_leaders/internal/contract_bindings/negriskfeemodule"
	"polymarket_leaders/internal/contract_bindings/proxywalletfactory"
	"polymarket_leaders/internal/contract_bindings/safeproxyfactory"
	"polymarket_leaders/internal/contract_bindings/uchilderc20proxy"
	"polymarket_leaders/internal/contract_bindings/umactfadapter"
	"polymarket_leaders/internal/contract_bindings/umactfadapter2"

	"github.com/ethereum/go-ethereum"
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
func ProcessBlock(ctx context.Context, client *ethclient.Client, lastProcessedBlockNumber uint64) ([]ParsedEvent, uint64, error) {
	var foundEvents []ParsedEvent

	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return foundEvents, lastProcessedBlockNumber, err
	}
	if blockNumber == lastProcessedBlockNumber {
		return foundEvents, lastProcessedBlockNumber, nil
	}
	log.Printf("Start process block number: %v", blockNumber)

	// get all events since last lastProcessedBlockNumber
	filter := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(lastProcessedBlockNumber + 1)),
		ToBlock:   big.NewInt(int64(blockNumber)),
	}
	logs, err := client.FilterLogs(ctx, filter)
	if err != nil {
		return foundEvents, lastProcessedBlockNumber, err
	}

	// Initialize contract filterers for event parsing
	ctfExchangeFilterer, _ := ctfexchange.NewCtfExchangeFilterer(common.Address{}, nil)
	negRiskCtfExchangeFilterer, _ := negriskctfexchange.NewNegRiskCtfExchangeFilterer(common.Address{}, nil)
	conditionalTokensFilterer, _ := conditionaltokens.NewConditionalTokensFilterer(common.Address{}, nil)
	negRiskAdapterFilterer, _ := negriskadapter.NewNegRiskAdapterFilterer(common.Address{}, nil)

	// Check logs for events we're interested in
	for _, log := range logs {
		// Check for OrderFilled event (CTFExchange)
		if event, err := ctfExchangeFilterer.ParseOrderFilled(log); err == nil {
			foundEvents = append(foundEvents, ParsedEvent{
				Name: "OrderFilled (CTFExchange)",
				Data: event,
			})
		}
		// Check for OrderFilled event (NegRiskCtfExchange)
		if event, err := negRiskCtfExchangeFilterer.ParseOrderFilled(log); err == nil {
			foundEvents = append(foundEvents, ParsedEvent{
				Name: "OrderFilled (NegRiskCtfExchange)",
				Data: event,
			})
		}

		// Check for OrdersMatched event (CTFExchange)
		if event, err := ctfExchangeFilterer.ParseOrdersMatched(log); err == nil {
			foundEvents = append(foundEvents, ParsedEvent{
				Name: "OrdersMatched (CTFExchange)",
				Data: event,
			})
		}
		// Check for OrdersMatched event (NegRiskCtfExchange)
		if event, err := negRiskCtfExchangeFilterer.ParseOrdersMatched(log); err == nil {
			foundEvents = append(foundEvents, ParsedEvent{
				Name: "OrdersMatched (NegRiskCtfExchange)",
				Data: event,
			})
		}

		// Check for PositionSplit event (ConditionalTokens)
		if event, err := conditionalTokensFilterer.ParsePositionSplit(log); err == nil {
			foundEvents = append(foundEvents, ParsedEvent{
				Name: "PositionSplit (ConditionalTokens)",
				Data: event,
			})
		}
		// Check for PositionSplit event (NegRiskAdapter)
		if event, err := negRiskAdapterFilterer.ParsePositionSplit(log); err == nil {
			foundEvents = append(foundEvents, ParsedEvent{
				Name: "PositionSplit (NegRiskAdapter)",
				Data: event,
			})
		}

		// Check for PositionsMerge event (ConditionalTokens)
		if event, err := conditionalTokensFilterer.ParsePositionsMerge(log); err == nil {
			foundEvents = append(foundEvents, ParsedEvent{
				Name: "PositionsMerge (ConditionalTokens)",
				Data: event,
			})
		}
		// Check for PositionsMerge event (NegRiskAdapter)
		if event, err := negRiskAdapterFilterer.ParsePositionsMerge(log); err == nil {
			foundEvents = append(foundEvents, ParsedEvent{
				Name: "PositionsMerge (NegRiskAdapter)",
				Data: event,
			})
		}

		// Check for PositionsConverted event (NegRiskAdapter)
		if event, err := negRiskAdapterFilterer.ParsePositionsConverted(log); err == nil {
			foundEvents = append(foundEvents, ParsedEvent{
				Name: "PositionsConverted (NegRiskAdapter)",
				Data: event,
			})
		}
	}

	return foundEvents, blockNumber, nil
}

// ExtractAddressesFromEvent extracts addresses (Maker, Taker, TakerOrderMaker, Stakeholder) from an event
// It excludes addresses that are Polymarket contract addresses
func ExtractAddressesFromEvent(eventData interface{}) []common.Address {
	var addresses []common.Address

	switch e := eventData.(type) {
	// CTFExchange events
	case *ctfexchange.CtfExchangeOrderFilled:
		addresses = append(addresses, e.Maker, e.Taker)
	case *ctfexchange.CtfExchangeOrdersMatched:
		addresses = append(addresses, e.TakerOrderMaker)

	// NegRiskCtfExchange events
	case *negriskctfexchange.NegRiskCtfExchangeOrderFilled:
		addresses = append(addresses, e.Maker, e.Taker)
	case *negriskctfexchange.NegRiskCtfExchangeOrdersMatched:
		addresses = append(addresses, e.TakerOrderMaker)

	// ConditionalTokens events
	case *conditionaltokens.ConditionalTokensPositionSplit:
		addresses = append(addresses, e.Stakeholder)
	case *conditionaltokens.ConditionalTokensPositionsMerge:
		addresses = append(addresses, e.Stakeholder)

	// NegRiskAdapter events
	case *negriskadapter.NegRiskAdapterPositionSplit:
		addresses = append(addresses, e.Stakeholder)
	case *negriskadapter.NegRiskAdapterPositionsMerge:
		addresses = append(addresses, e.Stakeholder)
	case *negriskadapter.NegRiskAdapterPositionsConverted:
		addresses = append(addresses, e.Stakeholder)
	}

	// Filter out contract addresses
	var filtered []common.Address
	for _, addr := range addresses {
		if _, isContract := contracts[addr]; !isContract {
			filtered = append(filtered, addr)
		}
	}

	return filtered
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

// contracts is a package-level constant map of Polymarket contract addresses to their contract info
var contracts map[common.Address]ContractInfo

func init() {
	contracts = map[common.Address]ContractInfo{
		common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"): {
			Name: "ConditionalTokens",
			ABI:  GetContractABI(conditionaltokens.ConditionalTokensMetaData),
		},
		common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"): {
			Name: "CTFExchange",
			ABI:  GetContractABI(ctfexchange.CtfExchangeMetaData),
		},
		common.HexToAddress("0x65070BE91477460D8A7AeEb94ef92fe056C2f2A7"): {
			Name: "UmaCtfAdapter",
			ABI:  GetContractABI(umactfadapter.UmaCtfAdapterMetaData),
		},
		common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"): {
			Name: "NegRiskCtfExchange",
			ABI:  GetContractABI(negriskctfexchange.NegRiskCtfExchangeMetaData),
		},
		common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"): {
			Name: "NegRiskAdapter",
			ABI:  GetContractABI(negriskadapter.NegRiskAdapterMetaData),
		},
		common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"): {
			Name: "UChildERC20Proxy",
			ABI:  GetContractABI(uchilderc20proxy.UchaildErc20ProxyMetaData),
		},
		common.HexToAddress("0x6A9D222616C90FcA5754cd1333cFD9b7fb6a4F74"): {
			Name: "UmaCtfAdapter2",
			ABI:  GetContractABI(umactfadapter2.UmaCtfAdapter2MetaData),
		},
		common.HexToAddress("0xaacfeea03eb1561c4e67d661e40682bd20e3541b"): {
			Name: "SafeProxyFactory",
			ABI:  GetContractABI(safeproxyfactory.SafeProxyFactoryMetaData),
		},
		common.HexToAddress("0xaB45c5A4B0c941a2F231C04C3f49182e1A254052"): {
			Name: "ProxyWalletFactory",
			ABI:  GetContractABI(proxywalletfactory.ProxyWalletFactoryMetaData),
		},
		common.HexToAddress("0x78769D50Be1763ed1CA0D5E878D93f05aabff29e"): {
			Name: "NegRiskFeeModule",
			ABI:  GetContractABI(negriskfeemodule.NegRiskFeeModuleMetaData),
		},
	}
}

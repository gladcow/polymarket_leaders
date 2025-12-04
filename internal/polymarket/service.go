package polymarket

import (
	"context"
	"log"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const TOP_ADDRESS_COUNT = 10

type addressCount struct {
	address common.Address
	count   int
}

type Service struct {
	client          *ethclient.Client
	signer          types.Signer
	requestInterval time.Duration
	addressCounts   map[common.Address]int
	addresses       [TOP_ADDRESS_COUNT]addressCount
	startBlock      uint64
	currentBlock    uint64
	lastUpdateBlock uint64
}

func NewService(
	rpcUrl string,
	chainID *big.Int,
	interval time.Duration,
) (*Service, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	return &Service{
		client:          client,
		signer:          types.LatestSignerForChainID(chainID),
		requestInterval: interval,
		addressCounts:   make(map[common.Address]int),
	}, nil
}

func (s *Service) Start(ctx context.Context) error {
	blockNumber, err := s.client.BlockNumber(ctx)
	if err != nil {
		return err
	}
	log.Printf("Start block number: %v", blockNumber)
	s.startBlock = blockNumber
	s.currentBlock = blockNumber

	ticker := time.NewTicker(s.requestInterval)
	defer ticker.Stop()
	if err := s.request(ctx); err != nil {
		log.Printf("Request failed: %v", err)
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := s.request(ctx); err != nil {
				log.Printf("Request failed: %v", err)
			}
		}
	}
}

func (s *Service) Close() {
	s.client.Close()
}

// AddressInfo represents an address with its count
type AddressInfo struct {
	Address string
	Count   int
}

// GetAddresses returns the top addresses with their counts
func (s *Service) GetAddresses() []AddressInfo {
	result := make([]AddressInfo, 0, TOP_ADDRESS_COUNT)
	for i := 0; i < TOP_ADDRESS_COUNT; i++ {
		if s.addresses[i].address != (common.Address{}) {
			result = append(result, AddressInfo{
				Address: s.addresses[i].address.String(),
				Count:   s.addresses[i].count,
			})
		}
	}
	return result
}

func (s *Service) request(ctx context.Context) error {
	foundEvents, currentBlock, err := ProcessBlock(ctx, s.client, s.currentBlock)
	if err != nil {
		return err
	}

	// Extract addresses from all events and update counts
	for _, event := range foundEvents {
		addresses := ExtractAddressesFromEvent(event.Data)
		for _, addr := range addresses {
			s.addressCounts[addr]++
		}
	}

	// Calculate top-10 addresses every 10 blocks
	if currentBlock-s.lastUpdateBlock > 10 {
		var pairs []addressCount
		for addr, count := range s.addressCounts {
			pairs = append(pairs, addressCount{address: addr, count: count})
		}

		// Sort by count (descending)
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].count > pairs[j].count
		})

		topN := TOP_ADDRESS_COUNT
		if len(pairs) < topN {
			topN = len(pairs)
		}
		log.Printf("\nBlock %d:", s.currentBlock)
		for i := 0; i < topN; i++ {
			s.addresses[i].address = pairs[i].address
			s.addresses[i].count = pairs[i].count
			log.Printf("%s: %d", s.addresses[i].address.String(), s.addresses[i].count)
		}
		s.lastUpdateBlock = currentBlock
	}

	// Wait for next block
	s.currentBlock = currentBlock

	return nil
}

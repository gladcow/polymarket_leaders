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

package config

import (
	"errors"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

// default values
const (
	defaultRPCURL           = "https://polygon-rpc.com"
	defaultChainId          = 137
	defaultRequestInterval  = 2 * time.Second
	defaultDashboardAddress = ":8080"
)

type Config struct {
	RPCURL           string
	ChainId          *big.Int
	RequestInterval  time.Duration
	DashboardAddress string
}

func LoadConfig(path string) (*Config, error) {
	_ = loadEnvFile()

	config := &Config{
		RPCURL:           defaultRPCURL,
		ChainId:          big.NewInt(defaultChainId),
		RequestInterval:  defaultRequestInterval,
		DashboardAddress: defaultDashboardAddress,
	}

	// Override with environment variables
	if rpcURL := os.Getenv("RPC_URL"); rpcURL != "" {
		config.RPCURL = rpcURL
	}

	if chainId := os.Getenv("CHAIN_ID"); chainId != "" {
		_, ok := config.ChainId.SetString(chainId, 10)
		if !ok {
			return nil, errors.New("could not convert string to big.Int")
		}
	}

	if requestIntervalStr := os.Getenv("REQUEST_INTERVAL"); requestIntervalStr != "" {
		if requestInterval, err := time.ParseDuration(requestIntervalStr); err == nil {
			config.RequestInterval = requestInterval
		}
	}

	if dashboardAddr := os.Getenv("DASHBOARD_ADDRESS"); dashboardAddr != "" {
		config.DashboardAddress = dashboardAddr
	}

	log.Printf("Config loaded: RPC_URL=%s, CHAIN_ID=%s, REQUEST_INTERVAL=%s, DASHBOARD_ADDRESS=%s",
		config.RPCURL, config.ChainId, config.RequestInterval, config.DashboardAddress)

	return config, nil
}

// loadEnvFile attempts to load a .env file but doesn't fail if it's missing
func loadEnvFile() error {
	if !fileExists(".env") {
		return nil // not an error
	}

	content, err := os.ReadFile(".env")
	if err != nil {
		log.Printf("Warning: could not read .env file: %v", err)
		return nil
	}

	lines := string(content)
	for _, line := range strings.Split(lines, "\n") {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if err := os.Setenv(key, value); err != nil {
			log.Printf("Warning: error setting environment variable %s: %v", key, err)
		}
	}

	log.Println("Successfully loaded configuration from .env file")
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

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
	defaultRPCURL          = "https://polygon-rpc.com"
	defaultChainId         = 137
	defaultRequestInterval = 2 * time.Second
)

type Config struct {
	RPCURL          string
	ChainId         *big.Int
	RequestInterval time.Duration
}

func LoadConfig(path string) (*Config, error) {
	_ = loadEnvFile()

	config := &Config{
		RPCURL:          defaultRPCURL,
		ChainId:         big.NewInt(defaultChainId),
		RequestInterval: defaultRequestInterval,
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

	log.Printf("Config loaded: RPC_URL=%s, CHAIN_ID=%s, REQUEST_INTERVAL=%s",
		config.RPCURL, config.ChainId, config.RequestInterval)

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

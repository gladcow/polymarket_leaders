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

package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"polymarket_leaders/internal/config"
	"polymarket_leaders/internal/dashboard"
	"polymarket_leaders/internal/polymarket"
	"syscall"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	configPath := flag.String("config", "config/config.yaml", "Path to configuration file")
	//dataDir := flag.String("data-dir", "data", "Directory for persistent data")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	service, err := polymarket.NewService(
		cfg.RPCURL, cfg.ChainId, cfg.RequestInterval, cfg.ResetIntervalBlocks,
	)
	if err != nil {
		log.Fatalf("Failed to create polymarket service: %v", err)
	}
	defer service.Close()

	errChan := make(chan error, 1) // err chan for service errors

	uiSrv := dashboard.NewService(cfg.DashboardAddress, service)
	go func() {
		if err := uiSrv.Start(); err != nil {
			errChan <- err
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Starting polymarket service")

	go func() {
		if err := service.Start(ctx); err != nil {
			errChan <- err
		}
	}()

	for {
		select {
		case sig := <-sigChan:
			log.Printf("Received signal: %v", sig)
			log.Println("Gracefully shutting down...")
			uiSrv.Close(ctx)
			cancel()
			service.Close()
			return
		case err := <-errChan:
			log.Printf("Service error: %v", err)
			cancel()
			os.Exit(1)
		case <-ctx.Done():
			log.Println("Context cancelled, shutting down...")
			return
		}
	}
}

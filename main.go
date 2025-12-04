package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"polymarket_leaders/internal/config"
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
		cfg.RPCURL, cfg.ChainId, cfg.RequestInterval,
	)
	if err != nil {
		log.Fatalf("Failed to create polymarket service: %v", err)
	}
	defer service.Close()

	errChan := make(chan error, 1) // err chan for service errors

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

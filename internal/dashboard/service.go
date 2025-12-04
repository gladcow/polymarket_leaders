package dashboard

import (
	"context"
	"encoding/json"
	"net/http"
	"polymarket_leaders/internal/polymarket"
	"time"
)

type Service struct {
	server            *http.Server
	polymarketService *polymarket.Service
}

func NewService(dashboardAddress string, polymarketService *polymarket.Service) *Service {
	s := &Service{
		polymarketService: polymarketService,
	}
	server := &http.Server{
		Addr:    dashboardAddress,
		Handler: http.HandlerFunc(s.handleRequest),
	}
	s.server = server
	return s
}

func (s *Service) handleRequest(w http.ResponseWriter, r *http.Request) {
	addresses := s.polymarketService.GetAddresses()

	type AddressResponse struct {
		Address string `json:"address"`
		Count   int    `json:"count"`
	}

	type Response struct {
		Service   string            `json:"service"`
		Addresses []AddressResponse `json:"addresses"`
	}

	response := Response{
		Service:   "polymarket-activity-leaders",
		Addresses: make([]AddressResponse, 0, len(addresses)),
	}

	for _, addr := range addresses {
		response.Addresses = append(response.Addresses, AddressResponse{
			Address: addr.Address,
			Count:   addr.Count,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (s *Service) Start() error {
	return s.server.ListenAndServe()
}

func (s *Service) Close(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

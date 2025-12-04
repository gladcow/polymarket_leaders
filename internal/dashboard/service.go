package dashboard

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"polymarket_leaders/internal/polymarket"
	"time"
)

//go:embed templates/*.html
var templatesFS embed.FS

type Service struct {
	server            *http.Server
	polymarketService *polymarket.Service
}

type AddressResponse struct {
	Address string `json:"address"`
	Count   int    `json:"count"`
}

type Response struct {
	Service    string            `json:"service"`
	StartBlock uint64            `json:"startBlock"`
	LastBlock  uint64            `json:"lastBlock"`
	Addresses  []AddressResponse `json:"addresses"`
}

func NewService(dashboardAddress string, polymarketService *polymarket.Service) *Service {
	s := &Service{
		polymarketService: polymarketService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleHTML)
	mux.HandleFunc("/api/leaders", s.handleJSON)

	server := &http.Server{
		Addr:    dashboardAddress,
		Handler: mux,
	}
	s.server = server
	return s
}

func (s *Service) getData() Response {
	addresses, lastBlock, startBlock := s.polymarketService.GetAddresses()

	response := Response{
		Service:    "polymarket-activity-leaders",
		StartBlock: startBlock,
		LastBlock:  lastBlock,
		Addresses:  make([]AddressResponse, 0, len(addresses)),
	}

	for _, addr := range addresses {
		response.Addresses = append(response.Addresses, AddressResponse{
			Address: addr.Address,
			Count:   addr.Count,
		})
	}

	return response
}

func (s *Service) handleJSON(w http.ResponseWriter, r *http.Request) {
	response := s.getData()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (s *Service) handleHTML(w http.ResponseWriter, r *http.Request) {
	response := s.getData()

	tmpl, err := template.New("dashboard.html").Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b uint64) uint64 {
			if a > b {
				return a - b
			}
			return 0
		},
		"avg": func(count int, blocks uint64) string {
			if blocks == 0 {
				return "0.00"
			}
			avg := float64(count) / float64(blocks)
			return fmt.Sprintf("%.2f", avg)
		},
	}).ParseFS(templatesFS, "templates/dashboard.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := tmpl.Execute(w, response); err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func (s *Service) Start() error {
	return s.server.ListenAndServe()
}

func (s *Service) Close(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

package service

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Jurupoc/PriceAgregator/ingestion/internal/domain"
)

// IngestionService gerencia a coleta e armazenamento de preços
type IngestionService struct {
	providers []domain.PriceProvider
	cache     *PriceCache
	interval  time.Duration
	stopCh    chan struct{}
	wg        sync.WaitGroup
}

// PriceCache armazena os preços em memória
type PriceCache struct {
	mu     sync.RWMutex
	prices []domain.PriceSnapshot
}

// NewIngestionService cria uma nova instância do serviço de ingestion
func NewIngestionService(providers []domain.PriceProvider, interval time.Duration) *IngestionService {
	return &IngestionService{
		providers: providers,
		cache:     NewPriceCache(),
		interval:  interval,
		stopCh:    make(chan struct{}),
	}
}

// NewPriceCache cria uma nova instância do cache
func NewPriceCache() *PriceCache {
	return &PriceCache{
		prices: make([]domain.PriceSnapshot, 0),
	}
}

// Start inicia o loop de ingestion
func (s *IngestionService) Start(ctx context.Context) {
	s.wg.Add(1)
	go s.run(ctx)
}

// Stop para o loop de ingestion
func (s *IngestionService) Stop() {
	close(s.stopCh)
	s.wg.Wait()
}

// run executa o loop principal de coleta de preços
func (s *IngestionService) run(ctx context.Context) {
	defer s.wg.Done()

	// Executa imediatamente na primeira vez
	s.fetchAllProviders(ctx)

	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Ingestion service: context cancelled")
			return
		case <-s.stopCh:
			log.Println("Ingestion service: stopping")
			return
		case <-ticker.C:
			s.fetchAllProviders(ctx)
		}
	}
}

// fetchAllProviders coleta preços de todos os providers
func (s *IngestionService) fetchAllProviders(ctx context.Context) {
	log.Println("Ingestion service: fetching prices from all providers")

	var allPrices []domain.PriceSnapshot

	for _, provider := range s.providers {
		prices, err := provider.FetchPrices(ctx)
		if err != nil {
			log.Printf("Error fetching prices from %s: %v", provider.Name(), err)
			continue
		}

		log.Printf("Fetched %d prices from %s", len(prices), provider.Name())
		allPrices = append(allPrices, prices...)
	}

	if len(allPrices) > 0 {
		s.cache.Update(allPrices)
		log.Printf("Cache updated with %d total prices", len(allPrices))
	}
}

// GetLatestPrices retorna todos os preços armazenados no cache
func (s *IngestionService) GetLatestPrices() []domain.PriceSnapshot {
	return s.cache.GetAll()
}

// Update atualiza o cache com novos preços
func (c *PriceCache) Update(prices []domain.PriceSnapshot) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.prices = prices
}

// GetAll retorna todos os preços do cache
func (c *PriceCache) GetAll() []domain.PriceSnapshot {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Retorna uma cópia para evitar race conditions
	result := make([]domain.PriceSnapshot, len(c.prices))
	copy(result, c.prices)
	return result
}


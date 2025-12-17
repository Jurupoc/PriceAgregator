package service

import (
	"context"
	"testing"
	"time"

	"github.com/Jurupoc/PriceAgregator/ingestion/internal/domain"
)

// mockProvider é um provider mock para testes
type mockProvider struct {
	name   string
	prices []domain.PriceSnapshot
	err    error
}

func (m *mockProvider) Name() string {
	return m.name
}

func (m *mockProvider) FetchPrices(ctx context.Context) ([]domain.PriceSnapshot, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.prices, nil
}

func TestIngestionService_GetLatestPrices(t *testing.T) {
	now := time.Now()
	providers := []domain.PriceProvider{
		&mockProvider{
			name: "MockProvider1",
			prices: []domain.PriceSnapshot{
				{Symbol: "BTC", PriceUSD: 50000, Source: "MockProvider1", Timestamp: now},
			},
		},
	}

	service := NewIngestionService(providers, 1*time.Second)

	// Inicia o serviço
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	service.Start(ctx)

	// Aguarda um pouco para garantir que o fetch foi executado
	time.Sleep(100 * time.Millisecond)

	prices := service.GetLatestPrices()
	if len(prices) == 0 {
		t.Error("Expected at least one price after service start")
	}

	service.Stop()
}

func TestPriceCache_UpdateAndGetAll(t *testing.T) {
	cache := NewPriceCache()
	now := time.Now()

	prices := []domain.PriceSnapshot{
		{Symbol: "BTC", PriceUSD: 50000, Source: "Test", Timestamp: now},
		{Symbol: "ETH", PriceUSD: 3000, Source: "Test", Timestamp: now},
	}

	cache.Update(prices)
	allPrices := cache.GetAll()

	if len(allPrices) != len(prices) {
		t.Errorf("Expected %d prices, got %d", len(prices), len(allPrices))
	}

	// Verifica que retorna uma cópia
	allPrices[0].Symbol = "MODIFIED"
	allPrices2 := cache.GetAll()
	if allPrices2[0].Symbol == "MODIFIED" {
		t.Error("Expected cache to return a copy, but modification affected cache")
	}
}

func TestIngestionService_MultipleProviders(t *testing.T) {
	now := time.Now()
	providers := []domain.PriceProvider{
		&mockProvider{
			name: "Provider1",
			prices: []domain.PriceSnapshot{
				{Symbol: "BTC", PriceUSD: 50000, Source: "Provider1", Timestamp: now},
			},
		},
		&mockProvider{
			name: "Provider2",
			prices: []domain.PriceSnapshot{
				{Symbol: "ETH", PriceUSD: 3000, Source: "Provider2", Timestamp: now},
			},
		},
	}

	service := NewIngestionService(providers, 1*time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	service.Start(ctx)

	// Aguarda um pouco para garantir que o fetch foi executado
	time.Sleep(100 * time.Millisecond)

	prices := service.GetLatestPrices()
	if len(prices) < 2 {
		t.Errorf("Expected at least 2 prices from 2 providers, got %d", len(prices))
	}

	service.Stop()
}

func TestIngestionService_ProviderError(t *testing.T) {
	now := time.Now()
	providers := []domain.PriceProvider{
		&mockProvider{
			name: "ErrorProvider",
			err:  context.DeadlineExceeded,
		},
		&mockProvider{
			name: "GoodProvider",
			prices: []domain.PriceSnapshot{
				{Symbol: "BTC", PriceUSD: 50000, Source: "GoodProvider", Timestamp: now},
			},
		},
	}

	service := NewIngestionService(providers, 1*time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	service.Start(ctx)

	// Aguarda um pouco para garantir que o fetch foi executado
	time.Sleep(100 * time.Millisecond)

	prices := service.GetLatestPrices()
	// Deve ter pelo menos um preço do provider que funcionou
	if len(prices) == 0 {
		t.Error("Expected at least one price from the working provider")
	}

	service.Stop()
}


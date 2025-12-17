package coingecko

import (
	"context"
	"testing"
	"time"
)

func TestCoinGeckoProvider_Name(t *testing.T) {
	provider := NewCoinGeckoProvider()
	if provider.Name() != ProviderName {
		t.Errorf("Expected name %s, got %s", ProviderName, provider.Name())
	}
}

func TestCoinGeckoProvider_FetchPrices(t *testing.T) {
	provider := NewCoinGeckoProvider()
	ctx := context.Background()

	prices, err := provider.FetchPrices(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(prices) == 0 {
		t.Fatal("Expected at least one price, got none")
	}

	// Verifica estrutura dos dados
	for _, price := range prices {
		if price.Symbol == "" {
			t.Error("Expected symbol to be non-empty")
		}
		if price.PriceUSD <= 0 {
			t.Errorf("Expected priceUSD to be positive, got %f", price.PriceUSD)
		}
		if price.Source != ProviderName {
			t.Errorf("Expected source to be %s, got %s", ProviderName, price.Source)
		}
		if price.Timestamp.IsZero() {
			t.Error("Expected timestamp to be set")
		}
		if price.Timestamp.After(time.Now()) {
			t.Error("Expected timestamp to be in the past or present")
		}
	}

	// Verifica que retorna dados consistentes
	prices2, err := provider.FetchPrices(ctx)
	if err != nil {
		t.Fatalf("Expected no error on second call, got %v", err)
	}

	if len(prices) != len(prices2) {
		t.Errorf("Expected same number of prices, got %d and %d", len(prices), len(prices2))
	}

	// Verifica símbolos esperados
	expectedSymbols := map[string]bool{
		"BTC": true,
		"ETH": true,
		"BNB": true,
		"SOL": true,
		"ADA": true,
	}

	foundSymbols := make(map[string]bool)
	for _, price := range prices {
		foundSymbols[price.Symbol] = true
	}

	for symbol := range expectedSymbols {
		if !foundSymbols[symbol] {
			t.Errorf("Expected symbol %s to be present", symbol)
		}
	}
}

func TestCoinGeckoProvider_FetchPrices_ContextCancellation(t *testing.T) {
	provider := NewCoinGeckoProvider()
	ctx, cancel := context.WithCancel(context.Background())

	// Cancela o contexto imediatamente
	cancel()

	prices, err := provider.FetchPrices(ctx)
	if err == nil {
		t.Error("Expected error when context is cancelled")
	}
	if prices != nil {
		t.Error("Expected nil prices when context is cancelled")
	}
}


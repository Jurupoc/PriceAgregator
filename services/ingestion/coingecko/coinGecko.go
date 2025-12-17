package coingecko

import (
	"context"
	"time"

	"github.com/Jurupoc/PriceAgregator/ingestion/internal/domain"
)

const (
	ProviderName = "CoinGecko"
)

// CoinGeckoProvider implementa PriceProvider com dados mockados
type CoinGeckoProvider struct{}

// NewCoinGeckoProvider cria uma nova instância do provider CoinGecko
func NewCoinGeckoProvider() domain.PriceProvider {
	return &CoinGeckoProvider{}
}

// Name retorna o nome do provider
func (p *CoinGeckoProvider) Name() string {
	return ProviderName
}

// FetchPrices retorna dados mockados de preços de criptomoedas
func (p *CoinGeckoProvider) FetchPrices(ctx context.Context) ([]domain.PriceSnapshot, error) {
	// Simula um pequeno delay de rede
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(50 * time.Millisecond):
	}

	// Dados mockados consistentes
	now := time.Now()
	return []domain.PriceSnapshot{
		{
			Symbol:    "BTC",
			PriceUSD:  43250.50,
			Source:    ProviderName,
			Timestamp: now,
		},
		{
			Symbol:    "ETH",
			PriceUSD:  2650.75,
			Source:    ProviderName,
			Timestamp: now,
		},
		{
			Symbol:    "BNB",
			PriceUSD:  315.20,
			Source:    ProviderName,
			Timestamp: now,
		},
		{
			Symbol:    "SOL",
			PriceUSD:  98.45,
			Source:    ProviderName,
			Timestamp: now,
		},
		{
			Symbol:    "ADA",
			PriceUSD:  0.52,
			Source:    ProviderName,
			Timestamp: now,
		},
	}, nil
}

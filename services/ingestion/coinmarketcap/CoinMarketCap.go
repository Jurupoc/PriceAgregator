package coinmarketcap

import (
	"context"
	"time"

	"github.com/Jurupoc/PriceAgregator/ingestion/internal/domain"
)

const (
	ProviderName = "CoinMarketCap"
)

// CoinMarketCapProvider implementa PriceProvider com dados mockados
type CoinMarketCapProvider struct{}

// NewCoinMarketCapProvider cria uma nova instância do provider CoinMarketCap
func NewCoinMarketCapProvider() domain.PriceProvider {
	return &CoinMarketCapProvider{}
}

// Name retorna o nome do provider
func (p *CoinMarketCapProvider) Name() string {
	return ProviderName
}

// FetchPrices retorna dados mockados de preços de criptomoedas
func (p *CoinMarketCapProvider) FetchPrices(ctx context.Context) ([]domain.PriceSnapshot, error) {
	// Simula um pequeno delay de rede
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(50 * time.Millisecond):
	}

	// Dados mockados consistentes (valores ligeiramente diferentes do CoinGecko para simular diferenças entre providers)
	now := time.Now()
	return []domain.PriceSnapshot{
		{
			Symbol:    "BTC",
			PriceUSD:  43280.25,
			Source:    ProviderName,
			Timestamp: now,
		},
		{
			Symbol:    "ETH",
			PriceUSD:  2652.10,
			Source:    ProviderName,
			Timestamp: now,
		},
		{
			Symbol:    "BNB",
			PriceUSD:  315.80,
			Source:    ProviderName,
			Timestamp: now,
		},
		{
			Symbol:    "SOL",
			PriceUSD:  98.90,
			Source:    ProviderName,
			Timestamp: now,
		},
		{
			Symbol:    "ADA",
			PriceUSD:  0.525,
			Source:    ProviderName,
			Timestamp: now,
		},
	}, nil
}

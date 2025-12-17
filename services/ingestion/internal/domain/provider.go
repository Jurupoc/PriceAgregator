package domain

import (
	"context"
)

// PriceProvider define a interface comum para todos os providers de preço
type PriceProvider interface {
	Name() string
	FetchPrices(ctx context.Context) ([]PriceSnapshot, error)
}

// DataProvider mantido para compatibilidade durante refatoração
type DataProvider interface {
	FetchPrice() (*PriceData, error)
}

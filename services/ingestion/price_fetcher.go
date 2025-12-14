package ingestion

import (
    "log"

    "github.com/Jurupoc/PriceAgregator/ingestion/internal/domain"
)

type PriceFetcher struct {
    providers []domain.DataProvider
}

func NewPriceFetcher(providers []domain.DataProvider) *PriceFetcher {
    return &PriceFetcher{providers: providers}
}

func (f *PriceFetcher) FetchAll() ([]*domain.PriceData, error) {
    var all []*domain.PriceData

    for _, p := range f.providers {
        price, err := p.FetchPrice()
        if err != nil {
            log.Printf("provider failed: %v", err)
            continue // degrade gracefully
        }
        all = append(all, price)
    }

    return all, nil
}

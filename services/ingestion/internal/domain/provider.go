package domain

import (
	"net/http"
	"github.com/Jurupoc/PriceAgregator/ingestion/internal/config"
)

type DataProvider interface {
	FetchPrice() (*PriceData, error)
}

type provider struct {
	config config.Config
	client http.Client
}

func NewDataProvider(cf config.ConfigProvider, cl http.Client) DataProvider {
	config := cf.Load()
	return &provider{
		config: config,
		client: cl,
	}
}

func (i *provider) FetchPrice() (*PriceData, error) {
	return nil, nil
}

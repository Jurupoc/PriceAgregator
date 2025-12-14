package coinmarketcap

import "time"

const (
	COIN_MARKET_CAP_API_URL        = "https://sandbox-api.coinmarketcap.com/v1/"
	COIN_MARKET_CAP_API_KEY_HEADER = "X-CMC_PRO_API_KEY"

	COIN_MARKETCAP_API_ENDPOINT = "listings/latest"
)

type PriceDataResponse struct {
	Data   []CryptoAsset
	Status ResponseStatus `json:"status"`
}

type ResponseStatus struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    *int      `json:"error_code"`
	ErrorMessage *string   `json:"error_message"`
}

type CryptoAsset struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Symbol          string           `json:"symbol"`
	Slug            string           `json:"slug"`
	CMCRank         int              `json:"cmc_rank"`
	NumMarketPairs  int              `json:"num_market_pairs"`
	Circulating     float64          `json:"circulating_supply"`
	TotalSupply     float64          `json:"total_supply"`
	MaxSupply       float64          `json:"max_supply"`
	LastUpdated     time.Time        `json:"last_updated"`
	DateAdded       time.Time        `json:"date_added"`
	Tags            []string         `json:"tags"`
	Platform        *Platform        `json:"platform"`
	Quote           map[string]Quote `json:"quote"`     // USD, BTC, etc
}

type Quote struct {
	Price            float64   `json:"price"`
	Volume24h        float64   `json:"volume_24h"`
	PercentChange1h  float64   `json:"percent_change_1h"`
	PercentChange24h float64   `json:"percent_change_24h"`
	PercentChange7d  float64   `json:"percent_change_7d"`
	MarketCap        float64   `json:"market_cap"`
	LastUpdated      time.Time `json:"last_updated"`
}

type Platform struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
}

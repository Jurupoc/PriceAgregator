package ingestion

const (
	COIN_GECKO_API_URL        = "https://api.coingecko.com/api/v3"
	COIN_GECKO_API_KEY_HEADER = "x-cg-demo-api-key"
)

type PriceDataResponse struct {
	CoinID CoinData
}

type CoinData struct {
	Usd           float64 `json:"usd"`
	UsdMarketCap  float64 `json:"usd_market_cap"`
	Usd24hVol     float64 `json:"usd_24h_vol"`
	Usd24hChange  float64 `json:"usd_24h_change"`
	LastUpdatedAt int64   `json:"last_updated_at"`
}

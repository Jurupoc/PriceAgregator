package config

import "time"

type Config struct {
	CoinGeckoAPIKey     string
	CoinMarketCapAPIKey string
	DBHost              string
	DBUser              string
	DBPassword          string
	Interval            time.Duration
}

type ConfigProvider interface {
	Load() Config
}

func NewConfigProvider() ConfigProvider {
	return configProvider{}
}

type configProvider struct{}

func (cp configProvider) Load() Config {
	return Config{
		CoinGeckoAPIKey:     mustEnv("COINGECKO_API_KEY"),
		CoinMarketCapAPIKey: mustEnv("COIN_MARKET_CAP_API_KEY"),
		DBHost:              mustEnv("DB_HOST"),
		DBUser:              mustEnv("DB_USER"),
		DBPassword:          mustEnv("DB_PASSWORD"),
		Interval:            mustDuration("FETCH_INTERVAL", 5*time.Minute),
	}
}

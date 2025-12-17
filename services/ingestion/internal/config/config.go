package config

import "time"

type Config struct {
	CoinGeckoAPIKey     string
	CoinMarketCapAPIKey string
	DBHost              string
	DBUser              string
	DBPassword          string
	Interval            time.Duration
	GRPCPort            string
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
		CoinGeckoAPIKey:     optionalEnv("COINGECKO_API_KEY", ""),
		CoinMarketCapAPIKey: optionalEnv("COIN_MARKET_CAP_API_KEY", ""),
		DBHost:              optionalEnv("DB_HOST", ""),
		DBUser:              optionalEnv("DB_USER", ""),
		DBPassword:          optionalEnv("DB_PASSWORD", ""),
		Interval:            mustDuration("FETCH_INTERVAL", 30*time.Second),
		GRPCPort:            optionalEnv("GRPC_PORT", "50051"),
	}
}

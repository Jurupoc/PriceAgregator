package config

import (
	"os"
	"time"
)

// Config contém todas as configurações do API Gateway
type Config struct {
	GraphQLPort     string
	IngestionAddr   string
	GRPCTimeout     time.Duration
}

// Load carrega a configuração das variáveis de ambiente
func Load() Config {
	return Config{
		GraphQLPort:   getEnv("GRAPHQL_PORT", "8080"),
		IngestionAddr: getEnv("INGESTION_ADDR", "ingestion:50051"),
		GRPCTimeout:   getDurationEnv("GRPC_TIMEOUT", 5*time.Second),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}


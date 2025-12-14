package config

import (
	"log"
	"os"
	"time"
)

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing env var: %s", key)
	}
	return v
}

func mustDuration(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			log.Fatalf("invalid duration %s: %v", key, err)
		}
		return d
	}
	return def
}

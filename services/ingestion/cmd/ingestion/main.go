package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jurupoc/PriceAgregator/ingestion/coingecko"
	"github.com/Jurupoc/PriceAgregator/ingestion/coinmarketcap"
	"github.com/Jurupoc/PriceAgregator/ingestion/internal/config"
	"github.com/Jurupoc/PriceAgregator/ingestion/internal/domain"
	"github.com/Jurupoc/PriceAgregator/ingestion/internal/grpc"
	"github.com/Jurupoc/PriceAgregator/ingestion/internal/service"
)

func main() {
	// Carrega configuração
	cfgProvider := config.NewConfigProvider()
	cfg := cfgProvider.Load()

	// Cria providers
	providers := []domain.PriceProvider{
		coingecko.NewCoinGeckoProvider(),
		coinmarketcap.NewCoinMarketCapProvider(),
	}

	// Cria serviço de ingestion
	ingestionService := service.NewIngestionService(providers, cfg.Interval)

	// Cria contexto para graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Inicia o serviço de ingestion
	ingestionService.Start(ctx)

	// Inicia servidor gRPC em goroutine
	go func() {
		if err := grpc.Start(cfg.GRPCPort, ingestionService); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Aguarda sinal de interrupção
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down...")
	cancel()
	ingestionService.Stop()
	log.Println("Shutdown complete")
}

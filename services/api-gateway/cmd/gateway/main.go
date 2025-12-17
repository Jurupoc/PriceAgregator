package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/Jurupoc/PriceAgregator/api-gateway/graph"
	"github.com/Jurupoc/PriceAgregator/api-gateway/internal/config"
	"github.com/Jurupoc/PriceAgregator/api-gateway/internal/grpc"
)

func main() {
	// Carrega configuração
	cfg := config.Load()

	// Cria cliente gRPC
	grpcClient, err := grpc.NewClient(cfg.IngestionAddr)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer grpcClient.Close()

	// Cria resolver
	resolver := graph.NewResolver(grpcClient)

	// Cria servidor GraphQL
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	// Configura mux
	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	// Inicia servidor HTTP
	server := &http.Server{
		Addr:    ":" + cfg.GraphQLPort,
		Handler: mux,
	}

	go func() {
		log.Printf("GraphQL server listening on port %s", cfg.GraphQLPort)
		log.Printf("Playground available at http://localhost:%s/", cfg.GraphQLPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Aguarda sinal de interrupção
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}

	log.Println("Shutdown complete")
}


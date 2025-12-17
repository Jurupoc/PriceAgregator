package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/Jurupoc/PriceAgregator/api-gateway/graph/model"
	pb "github.com/Jurupoc/PriceAgregator/api-gateway/proto/gen/price"
)

// Resolver implementa os resolvers GraphQL
type Resolver struct {
	grpcClient GRPCClient
}

// GRPCClient define a interface para o cliente gRPC
type GRPCClient interface {
	GetLatestPrices(ctx context.Context) ([]*pb.Price, error)
}

// NewResolver cria um novo resolver
func NewResolver(grpcClient GRPCClient) *Resolver {
	return &Resolver{
		grpcClient: grpcClient,
	}
}

// Query retorna o resolver de Query
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct {
	*Resolver
}

// Prices retorna todos os preços disponíveis
func (r *queryResolver) Prices(ctx context.Context) ([]*model.Price, error) {
	pbPrices, err := r.grpcClient.GetLatestPrices(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prices from ingestion service: %w", err)
	}

	prices := make([]*model.Price, 0, len(pbPrices))
	for _, pbPrice := range pbPrices {
		prices = append(prices, &model.Price{
			Symbol:    pbPrice.Symbol,
			PriceUSD:  pbPrice.PriceUsd,
			Source:    pbPrice.Source,
			Timestamp: time.Unix(pbPrice.Timestamp, 0).Format(time.RFC3339),
		})
	}

	return prices, nil
}


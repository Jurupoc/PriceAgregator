package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Jurupoc/PriceAgregator/api-gateway/proto/gen/price"
)

// Client representa um cliente gRPC para o serviço de ingestion
type Client struct {
	conn   *grpc.ClientConn
	client pb.PriceServiceClient
	addr   string
}

// NewClient cria um novo cliente gRPC
func NewClient(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ingestion service: %w", err)
	}

	return &Client{
		conn:   conn,
		client: pb.NewPriceServiceClient(conn),
		addr:   addr,
	}, nil
}

// GetLatestPrices obtém os últimos preços do serviço de ingestion
func (c *Client) GetLatestPrices(ctx context.Context) ([]*pb.Price, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := c.client.GetLatestPrices(ctx, &pb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get latest prices: %w", err)
	}

	return resp.Prices, nil
}

// Close fecha a conexão gRPC
func (c *Client) Close() error {
	return c.conn.Close()
}


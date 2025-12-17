package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Jurupoc/PriceAgregator/ingestion/internal/service"
	pb "github.com/Jurupoc/PriceAgregator/ingestion/proto/gen/price"
)

// PriceServer implementa o servidor gRPC
type PriceServer struct {
	pb.UnimplementedPriceServiceServer
	ingestionService *service.IngestionService
}

// NewPriceServer cria uma nova instância do servidor gRPC
func NewPriceServer(ingestionService *service.IngestionService) *PriceServer {
	return &PriceServer{
		ingestionService: ingestionService,
	}
}

// GetLatestPrices implementa o método gRPC para obter os últimos preços
func (s *PriceServer) GetLatestPrices(ctx context.Context, req *pb.Empty) (*pb.PriceList, error) {
	snapshots := s.ingestionService.GetLatestPrices()

	prices := make([]*pb.Price, 0, len(snapshots))
	for _, snapshot := range snapshots {
		prices = append(prices, &pb.Price{
			Symbol:    snapshot.Symbol,
			PriceUsd:  snapshot.PriceUSD,
			Source:    snapshot.Source,
			Timestamp: snapshot.Timestamp.Unix(),
		})
	}

	return &pb.PriceList{
		Prices: prices,
	}, nil
}

// Start inicia o servidor gRPC
func Start(port string, ingestionService *service.IngestionService) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterPriceServiceServer(s, NewPriceServer(ingestionService))

	log.Printf("gRPC server listening on port %s", port)
	return s.Serve(lis)
}


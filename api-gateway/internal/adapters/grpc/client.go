package grpcclient

import (
	"fmt"

	pb "movies-api/api-gateway/gen/movies"
	"movies-api/api-gateway/internal/ports"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewMovieGRPCClient creates a MovieClient that communicates with the movies-service at addr.
func NewMovieGRPCClient(addr string) (ports.MovieClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("dial movies-service: %w", err)
	}
	return &movieGRPCClient{client: pb.NewMovieServiceClient(conn)}, nil
}

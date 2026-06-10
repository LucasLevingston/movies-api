package grpcserver

import (
	pb "movies-api/movies-service/gen/movies"
	"movies-api/movies-service/internal/ports"
)

// NewMovieServer creates a gRPC MovieServiceServer backed by the given service.
func NewMovieServer(service ports.MovieService) pb.MovieServiceServer {
	return &movieServer{service: service}
}

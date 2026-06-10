package grpcserver

import (
	pb "movies-api/movies-service/gen/movies"
	"movies-api/movies-service/internal/ports"
)

type movieServer struct {
	pb.UnimplementedMovieServiceServer
	service ports.MovieService
}

func NewMovieServer(service ports.MovieService) pb.MovieServiceServer {
	return &movieServer{service: service}
}

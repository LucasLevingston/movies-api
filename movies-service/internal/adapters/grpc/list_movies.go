package grpcserver

import (
	"context"

	pb "movies-api/movies-service/gen/movies"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *movieServer) ListMovies(ctx context.Context, _ *pb.ListMoviesRequest) (*pb.ListMoviesResponse, error) {
	movies, err := server.service.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list movies: %v", err)
	}

	pbMovies := make([]*pb.Movie, 0, len(movies))
	for _, movie := range movies {
		pbMovies = append(pbMovies, &pb.Movie{
			Id:         movie.ID,
			ExternalId: movie.ExternalID,
			Title:      movie.Title,
			Year:       movie.Year,
		})
	}
	return &pb.ListMoviesResponse{Movies: pbMovies}, nil
}

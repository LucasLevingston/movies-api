package grpcserver

import (
	"context"
	"errors"

	pb "movies-api/movies-service/gen/movies"
	"movies-api/movies-service/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *movieServer) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	movie, err := server.service.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "movie not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get movie: %v", err)
	}

	return &pb.GetMovieResponse{
		Movie: &pb.Movie{
			Id:         movie.ID,
			ExternalId: movie.ExternalID,
			Title:      movie.Title,
			Year:       movie.Year,
		},
	}, nil
}

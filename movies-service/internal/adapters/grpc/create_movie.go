package grpcserver

import (
	"context"

	pb "movies-api/movies-service/gen/movies"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *movieServer) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	movie, err := s.service.Create(ctx, req.ExternalId, req.Title, req.Year)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create movie: %v", err)
	}

	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:         movie.ID.Hex(),
			ExternalId: movie.ExternalID,
			Title:      movie.Title,
			Year:       movie.Year,
		},
	}, nil
}

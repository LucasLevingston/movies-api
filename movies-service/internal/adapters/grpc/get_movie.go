package grpcserver

import (
	"context"

	pb "movies-api/movies-service/gen/movies"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *movieServer) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	movie, err := s.service.GetByID(ctx, req.Id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "movie not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get movie: %v", err)
	}

	return &pb.GetMovieResponse{
		Movie: &pb.Movie{
			Id:         movie.ID.Hex(),
			ExternalId: movie.ExternalID,
			Title:      movie.Title,
			Year:       movie.Year,
		},
	}, nil
}

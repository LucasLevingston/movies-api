package grpcserver

import (
	"context"
	"errors"

	pb "movies-api/movies-service/gen/movies"
	"movies-api/movies-service/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *movieServer) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	if err := s.service.Delete(ctx, req.Id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "movie not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete movie: %v", err)
	}
	return &pb.DeleteMovieResponse{Success: true, Message: "movie deleted successfully"}, nil
}

package grpcserver

import (
	"context"

	pb "movies-api/movies-service/gen/movies"
	"movies-api/movies-service/internal/ports"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type movieServer struct {
	pb.UnimplementedMovieServiceServer
	service ports.MovieService
}

func NewMovieServer(service ports.MovieService) pb.MovieServiceServer {
	return &movieServer{service: service}
}

func (s *movieServer) ListMovies(ctx context.Context, _ *pb.ListMoviesRequest) (*pb.ListMoviesResponse, error) {
	movies, err := s.service.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list movies: %v", err)
	}

	pbMovies := make([]*pb.Movie, 0, len(movies))
	for _, m := range movies {
		pbMovies = append(pbMovies, &pb.Movie{
			Id:         m.ID.Hex(),
			ExternalId: m.ExternalID,
			Title:      m.Title,
			Year:       m.Year,
		})
	}
	return &pb.ListMoviesResponse{Movies: pbMovies}, nil
}

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

func (s *movieServer) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	if err := s.service.Delete(ctx, req.Id); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "movie not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete movie: %v", err)
	}
	return &pb.DeleteMovieResponse{Success: true, Message: "movie deleted successfully"}, nil
}

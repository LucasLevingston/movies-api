package grpcclient

import (
	"context"
	"fmt"

	pb "movies-api/api-gateway/gen/movies"
	"movies-api/api-gateway/internal/domain"
	"movies-api/api-gateway/internal/ports"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type movieGRPCClient struct {
	client pb.MovieServiceClient
}

func NewMovieGRPCClient(addr string) (ports.MovieClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("dial movies-service: %w", err)
	}
	return &movieGRPCClient{client: pb.NewMovieServiceClient(conn)}, nil
}

func (c *movieGRPCClient) ListMovies(ctx context.Context) ([]domain.Movie, error) {
	resp, err := c.client.ListMovies(ctx, &pb.ListMoviesRequest{})
	if err != nil {
		return nil, err
	}

	movies := make([]domain.Movie, 0, len(resp.Movies))
	for _, m := range resp.Movies {
		movies = append(movies, domain.Movie{
			ID:         m.Id,
			ExternalID: m.ExternalId,
			Title:      m.Title,
			Year:       m.Year,
		})
	}
	return movies, nil
}

func (c *movieGRPCClient) GetMovie(ctx context.Context, id string) (*domain.Movie, error) {
	resp, err := c.client.GetMovie(ctx, &pb.GetMovieRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &domain.Movie{
		ID:         resp.Movie.Id,
		ExternalID: resp.Movie.ExternalId,
		Title:      resp.Movie.Title,
		Year:       resp.Movie.Year,
	}, nil
}

func (c *movieGRPCClient) CreateMovie(ctx context.Context, externalID int32, title, year string) (*domain.Movie, error) {
	resp, err := c.client.CreateMovie(ctx, &pb.CreateMovieRequest{
		ExternalId: externalID,
		Title:      title,
		Year:       year,
	})
	if err != nil {
		return nil, err
	}
	return &domain.Movie{
		ID:         resp.Movie.Id,
		ExternalID: resp.Movie.ExternalId,
		Title:      resp.Movie.Title,
		Year:       resp.Movie.Year,
	}, nil
}

func (c *movieGRPCClient) DeleteMovie(ctx context.Context, id string) error {
	_, err := c.client.DeleteMovie(ctx, &pb.DeleteMovieRequest{Id: id})
	return err
}

package grpcclient

import (
	"context"

	pb "movies-api/api-gateway/gen/movies"
	"movies-api/api-gateway/internal/domain"
)

func (grpcClient *movieGRPCClient) ListMovies(ctx context.Context) ([]domain.Movie, error) {
	resp, err := grpcClient.client.ListMovies(ctx, &pb.ListMoviesRequest{})
	if err != nil {
		return nil, err
	}

	movies := make([]domain.Movie, 0, len(resp.Movies))
	for _, movie := range resp.Movies {
		movies = append(movies, domain.Movie{
			ID:         movie.Id,
			ExternalID: movie.ExternalId,
			Title:      movie.Title,
			Year:       movie.Year,
		})
	}
	return movies, nil
}

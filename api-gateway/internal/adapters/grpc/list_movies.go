package grpcclient

import (
	"context"

	pb "movies-api/api-gateway/gen/movies"
	"movies-api/api-gateway/internal/domain"
)

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

package grpcclient

import (
	"context"

	pb "movies-api/api-gateway/gen/movies"
	"movies-api/api-gateway/internal/domain"
)

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

package grpcclient

import (
	"context"

	pb "movies-api/api-gateway/gen/movies"
	"movies-api/api-gateway/internal/domain"
)

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

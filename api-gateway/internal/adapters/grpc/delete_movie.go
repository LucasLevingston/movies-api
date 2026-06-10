package grpcclient

import (
	"context"

	pb "movies-api/api-gateway/gen/movies"
)

func (c *movieGRPCClient) DeleteMovie(ctx context.Context, id string) error {
	_, err := c.client.DeleteMovie(ctx, &pb.DeleteMovieRequest{Id: id})
	return err
}

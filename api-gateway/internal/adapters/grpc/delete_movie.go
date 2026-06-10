package grpcclient

import (
	"context"

	pb "movies-api/api-gateway/gen/movies"
)

func (grpcClient *movieGRPCClient) DeleteMovie(ctx context.Context, id string) error {
	_, err := grpcClient.client.DeleteMovie(ctx, &pb.DeleteMovieRequest{Id: id})
	return err
}

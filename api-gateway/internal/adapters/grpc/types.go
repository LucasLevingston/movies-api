package grpcclient

import pb "movies-api/api-gateway/gen/movies"

type movieGRPCClient struct {
	client pb.MovieServiceClient
}

package main

import (
	"log"
	"net"

	grpcserver "movies-api/movies-service/internal/adapters/grpc"
	"movies-api/movies-service/internal/ports"
	"movies-api/movies-service/internal/usecase"
	pb "movies-api/movies-service/gen/movies"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcPort := getEnv("GRPC_PORT", "50051")
	storageBackend := getEnv("STORAGE_BACKEND", "mongodb")

	var repo ports.MovieRepository
	switch storageBackend {
	case "dynamodb":
		repo = newDynamoDBRepository()
	default:
		repo = newMongoDBRepository()
	}

	publisher := newEventPublisher()
	svc := usecase.NewMovieUseCase(repo, publisher)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterMovieServiceServer(grpcSrv, grpcserver.NewMovieServer(svc))
	reflection.Register(grpcSrv)

	log.Printf("movies gRPC server on :%s", grpcPort)
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}

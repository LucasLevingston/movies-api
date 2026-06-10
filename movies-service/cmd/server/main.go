package main

import (
	"context"
	"log"
	"net"
	"time"

	grpcserver "movies-api/movies-service/internal/adapters/grpc"
	"movies-api/movies-service/internal/adapters/mongodb"
	"movies-api/movies-service/internal/usecase"
	pb "movies-api/movies-service/gen/movies"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const dbConnectTimeout = 30 * time.Second

func main() {
	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	grpcPort := getEnv("GRPC_PORT", "50051")

	ctx, cancel := context.WithTimeout(context.Background(), dbConnectTimeout)
	defer cancel()

	client, err := mongodriver.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("mongodb connect: %v", err)
	}
	defer client.Disconnect(context.Background())

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("mongodb ping: %v", err)
	}
	log.Println("connected to MongoDB")

	db := client.Database("moviesdb")
	repo := mongodb.NewMovieRepository(db)
	svc := usecase.NewMovieUseCase(repo)

	seedMovies(db)

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

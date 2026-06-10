package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	grpcserver "movies-api/movies-service/internal/adapters/grpc"
	"movies-api/movies-service/internal/adapters/mongodb"
	"movies-api/movies-service/internal/domain"
	"movies-api/movies-service/internal/usecase"
	pb "movies-api/movies-service/gen/movies"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	grpcPort := getEnv("GRPC_PORT", "50051")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

func seedMovies(db *mongodriver.Database) {
	collection := db.Collection("movies")
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil || count > 0 {
		return
	}

	data, err := os.ReadFile("movies.json")
	if err != nil {
		log.Printf("movies.json not found, skipping seed: %v", err)
		return
	}

	var raw []struct {
		ID    int32  `json:"id"`
		Title string `json:"title"`
		Year  string `json:"year"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Printf("parse movies.json: %v", err)
		return
	}

	docs := make([]interface{}, 0, len(raw))
	for _, m := range raw {
		docs = append(docs, domain.Movie{
			ExternalID: m.ID,
			Title:      m.Title,
			Year:       m.Year,
		})
	}

	if _, err := collection.InsertMany(context.Background(), docs); err != nil {
		log.Printf("seed error: %v", err)
		return
	}
	log.Printf("seeded %d movies", len(docs))
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

package main

import (
	"context"
	"log"
	"time"

	"movies-api/movies-service/internal/adapters/mongodb"
	"movies-api/movies-service/internal/ports"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoConnectTimeout = 30 * time.Second

func newMongoDBRepository() ports.MovieRepository {
	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), mongoConnectTimeout)
	defer cancel()

	client, err := mongodriver.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("mongodb connect: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("mongodb ping: %v", err)
	}
	log.Println("connected to MongoDB")

	database := client.Database("moviesdb")
	seedMovies(database)

	return mongodb.NewMovieRepository(database)
}

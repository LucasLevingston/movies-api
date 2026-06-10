package mongodb

import (
	"movies-api/movies-service/internal/ports"

	"go.mongodb.org/mongo-driver/mongo"
)

type movieRepository struct {
	collection *mongo.Collection
}

func NewMovieRepository(db *mongo.Database) ports.MovieRepository {
	return &movieRepository{
		collection: db.Collection("movies"),
	}
}

package mongodb

import (
	"movies-api/movies-service/internal/ports"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewMovieRepository(db *mongo.Database) ports.MovieRepository {
	return &movieRepository{
		collection: db.Collection("movies"),
	}
}

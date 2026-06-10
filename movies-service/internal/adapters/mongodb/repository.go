package mongodb

import (
	"movies-api/movies-service/internal/ports"

	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "movies"

// NewMovieRepository creates a MongoDB-backed MovieRepository.
func NewMovieRepository(db *mongo.Database) ports.MovieRepository {
	return &movieRepository{
		collection: db.Collection(collectionName),
	}
}

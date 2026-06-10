package mongodb

import (
	"context"

	"movies-api/movies-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *movieRepository) FindAll(ctx context.Context) ([]domain.Movie, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var movies []domain.Movie
	if err := cursor.All(ctx, &movies); err != nil {
		return nil, err
	}
	if movies == nil {
		movies = []domain.Movie{}
	}
	return movies, nil
}

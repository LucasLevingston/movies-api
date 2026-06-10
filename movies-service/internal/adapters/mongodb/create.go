package mongodb

import (
	"context"

	"movies-api/movies-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *movieRepository) Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	result, err := r.collection.InsertOne(ctx, movie)
	if err != nil {
		return nil, err
	}
	movie.ID = result.InsertedID.(primitive.ObjectID)
	return movie, nil
}

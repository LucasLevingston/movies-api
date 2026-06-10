package mongodb

import (
	"context"

	"movies-api/movies-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *movieRepository) Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	doc := movieDocument{
		ExternalID: movie.ExternalID,
		Title:      movie.Title,
		Year:       movie.Year,
	}

	result, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	movie.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return movie, nil
}

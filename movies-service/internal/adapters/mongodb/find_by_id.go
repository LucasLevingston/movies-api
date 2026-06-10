package mongodb

import (
	"context"
	"fmt"

	"movies-api/movies-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *movieRepository) FindByID(ctx context.Context, id string) (*domain.Movie, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %w", err)
	}

	var movie domain.Movie
	if err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&movie); err != nil {
		return nil, err
	}
	return &movie, nil
}

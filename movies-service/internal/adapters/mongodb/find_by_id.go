package mongodb

import (
	"context"
	"fmt"

	"movies-api/movies-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repository *movieRepository) FindByID(ctx context.Context, id string) (*domain.Movie, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %w", err)
	}

	var doc movieDocument
	if err := repository.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &domain.Movie{
		ID:         doc.ID.Hex(),
		ExternalID: doc.ExternalID,
		Title:      doc.Title,
		Year:       doc.Year,
	}, nil
}

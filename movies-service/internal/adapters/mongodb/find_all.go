package mongodb

import (
	"context"

	"movies-api/movies-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (repository *movieRepository) FindAll(ctx context.Context) ([]domain.Movie, error) {
	cursor, err := repository.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []movieDocument
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	movies := make([]domain.Movie, len(docs))
	for index, document := range docs {
		movies[index] = domain.Movie{
			ID:         document.ID.Hex(),
			ExternalID: document.ExternalID,
			Title:      document.Title,
			Year:       document.Year,
		}
	}
	return movies, nil
}

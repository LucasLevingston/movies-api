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

	var docs []movieDocument
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	movies := make([]domain.Movie, len(docs))
	for i, d := range docs {
		movies[i] = domain.Movie{
			ID:         d.ID.Hex(),
			ExternalID: d.ExternalID,
			Title:      d.Title,
			Year:       d.Year,
		}
	}
	return movies, nil
}

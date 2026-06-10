package mongodb

import (
	"context"
	"fmt"

	"movies-api/movies-service/internal/domain"
	"movies-api/movies-service/internal/ports"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *movieRepository) Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	result, err := r.collection.InsertOne(ctx, movie)
	if err != nil {
		return nil, err
	}
	movie.ID = result.InsertedID.(primitive.ObjectID)
	return movie, nil
}

func (r *movieRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id format: %w", err)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

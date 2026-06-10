package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type movieRepository struct {
	collection *mongo.Collection
}

// movieDocument is the internal BSON mapping — never exposed outside this package.
type movieDocument struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ExternalID int32              `bson:"external_id"`
	Title      string             `bson:"title"`
	Year       string             `bson:"year"`
}

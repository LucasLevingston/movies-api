package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type movieRepository struct {
	collection *mongo.Collection
}

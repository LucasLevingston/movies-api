package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ExternalID int32              `bson:"external_id" json:"external_id"`
	Title      string             `bson:"title" json:"title"`
	Year       string             `bson:"year" json:"year"`
}

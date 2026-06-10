package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
)

// seedDoc matches the MongoDB document shape without exposing domain types to the seed layer.
type seedDoc struct {
	ExternalID int32  `bson:"external_id"`
	Title      string `bson:"title"`
	Year       string `bson:"year"`
}

func seedMovies(db *mongodriver.Database) {
	collection := db.Collection("movies")
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil || count > 0 {
		return
	}

	data, err := os.ReadFile("movies.json")
	if err != nil {
		log.Printf("movies.json not found, skipping seed: %v", err)
		return
	}

	var raw []struct {
		ID    int32  `json:"id"`
		Title string `json:"title"`
		Year  string `json:"year"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Printf("parse movies.json: %v", err)
		return
	}

	docs := make([]interface{}, 0, len(raw))
	for _, m := range raw {
		docs = append(docs, seedDoc{
			ExternalID: m.ID,
			Title:      m.Title,
			Year:       m.Year,
		})
	}

	if _, err := collection.InsertMany(context.Background(), docs); err != nil {
		log.Printf("seed error: %v", err)
		return
	}
	log.Printf("seeded %d movies", len(docs))
}

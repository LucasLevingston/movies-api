package domain_test

import (
	"encoding/json"
	"testing"

	"movies-api/api-gateway/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestMovieJSONSerialization(t *testing.T) {
	movie := domain.Movie{
		ID:         "507f1f77bcf86cd799439011",
		ExternalID: 8,
		Title:      "Edison Kinetoscopic Record of a Sneeze (1894)",
		Year:       "1894",
	}

	data, err := json.Marshal(movie)
	assert.NoError(t, err)

	var result domain.Movie
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)
	assert.Equal(t, movie.ID, result.ID)
	assert.Equal(t, movie.ExternalID, result.ExternalID)
	assert.Equal(t, movie.Title, result.Title)
	assert.Equal(t, movie.Year, result.Year)
}

func TestMovieJSONKeys(t *testing.T) {
	movie := domain.Movie{ID: "abc", ExternalID: 1, Title: "T", Year: "2024"}
	data, _ := json.Marshal(movie)
	jsonString := string(data)

	assert.Contains(t, jsonString, `"id"`)
	assert.Contains(t, jsonString, `"external_id"`)
	assert.Contains(t, jsonString, `"title"`)
	assert.Contains(t, jsonString, `"year"`)
}

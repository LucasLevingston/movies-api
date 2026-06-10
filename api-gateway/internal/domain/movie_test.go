package domain_test

import (
	"encoding/json"
	"testing"

	"movies-api/api-gateway/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestMovieJSONSerialization(t *testing.T) {
	m := domain.Movie{
		ID:         "507f1f77bcf86cd799439011",
		ExternalID: 8,
		Title:      "Edison Kinetoscopic Record of a Sneeze (1894)",
		Year:       "1894",
	}

	data, err := json.Marshal(m)
	assert.NoError(t, err)

	var result domain.Movie
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)
	assert.Equal(t, m.ID, result.ID)
	assert.Equal(t, m.ExternalID, result.ExternalID)
	assert.Equal(t, m.Title, result.Title)
	assert.Equal(t, m.Year, result.Year)
}

func TestMovieJSONKeys(t *testing.T) {
	m := domain.Movie{ID: "abc", ExternalID: 1, Title: "T", Year: "2024"}
	data, _ := json.Marshal(m)
	s := string(data)

	assert.Contains(t, s, `"id"`)
	assert.Contains(t, s, `"external_id"`)
	assert.Contains(t, s, `"title"`)
	assert.Contains(t, s, `"year"`)
}

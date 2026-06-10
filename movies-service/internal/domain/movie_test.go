package domain_test

import (
	"testing"

	"movies-api/movies-service/internal/domain"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMovieFields(t *testing.T) {
	id := primitive.NewObjectID()
	m := domain.Movie{
		ID:         id,
		ExternalID: 8,
		Title:      "Edison Kinetoscopic Record of a Sneeze (1894)",
		Year:       "1894",
	}

	assert.Equal(t, id, m.ID)
	assert.Equal(t, int32(8), m.ExternalID)
	assert.Equal(t, "Edison Kinetoscopic Record of a Sneeze (1894)", m.Title)
	assert.Equal(t, "1894", m.Year)
}

func TestMovieIDHex(t *testing.T) {
	id := primitive.NewObjectID()
	m := domain.Movie{ID: id}
	assert.Len(t, m.ID.Hex(), 24)
	assert.Equal(t, id.Hex(), m.ID.Hex())
}

func TestZeroValueMovie(t *testing.T) {
	m := domain.Movie{}
	assert.True(t, m.ID.IsZero())
	assert.Equal(t, int32(0), m.ExternalID)
	assert.Empty(t, m.Title)
	assert.Empty(t, m.Year)
}

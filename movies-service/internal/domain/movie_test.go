package domain_test

import (
	"testing"

	"movies-api/movies-service/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestMovieFields(t *testing.T) {
	m := domain.Movie{
		ID:         "507f1f77bcf86cd799439011",
		ExternalID: 8,
		Title:      "Edison Kinetoscopic Record of a Sneeze (1894)",
		Year:       "1894",
	}

	assert.Equal(t, "507f1f77bcf86cd799439011", m.ID)
	assert.Equal(t, int32(8), m.ExternalID)
	assert.Equal(t, "Edison Kinetoscopic Record of a Sneeze (1894)", m.Title)
	assert.Equal(t, "1894", m.Year)
}

func TestMovieIDIsString(t *testing.T) {
	m := domain.Movie{ID: "507f1f77bcf86cd799439011"}
	assert.Len(t, m.ID, 24)
	assert.IsType(t, "", m.ID)
}

func TestZeroValueMovie(t *testing.T) {
	m := domain.Movie{}
	assert.Empty(t, m.ID)
	assert.Equal(t, int32(0), m.ExternalID)
	assert.Empty(t, m.Title)
	assert.Empty(t, m.Year)
}

func TestErrNotFound(t *testing.T) {
	assert.Error(t, domain.ErrNotFound)
	assert.EqualError(t, domain.ErrNotFound, "not found")
}

package domain_test

import (
	"testing"

	"movies-api/api-gateway/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestMovieFields(t *testing.T) {
	movie := domain.Movie{
		ID:         "507f1f77bcf86cd799439011",
		ExternalID: 8,
		Title:      "Edison Kinetoscopic Record of a Sneeze (1894)",
		Year:       "1894",
	}

	assert.Equal(t, "507f1f77bcf86cd799439011", movie.ID)
	assert.Equal(t, int32(8), movie.ExternalID)
	assert.Equal(t, "Edison Kinetoscopic Record of a Sneeze (1894)", movie.Title)
	assert.Equal(t, "1894", movie.Year)
}

func TestMovieZeroValue(t *testing.T) {
	movie := domain.Movie{}
	assert.Empty(t, movie.ID)
	assert.Equal(t, int32(0), movie.ExternalID)
	assert.Empty(t, movie.Title)
	assert.Empty(t, movie.Year)
}

func TestErrNotFound(t *testing.T) {
	assert.Error(t, domain.ErrNotFound)
	assert.EqualError(t, domain.ErrNotFound, "not found")
}

func TestErrBadRequest(t *testing.T) {
	assert.Error(t, domain.ErrBadRequest)
	assert.EqualError(t, domain.ErrBadRequest, "bad request")
}

package httphandler

import (
	"movies-api/api-gateway/internal/ports"
)

// MovieHandler handles HTTP requests for movies.
type MovieHandler struct {
	client ports.MovieClient
}

func NewMovieHandler(client ports.MovieClient) *MovieHandler {
	return &MovieHandler{client: client}
}

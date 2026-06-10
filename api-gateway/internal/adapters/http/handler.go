package httphandler

import "movies-api/api-gateway/internal/ports"

// NewMovieHandler creates an HTTP handler wired to the given MovieClient.
func NewMovieHandler(client ports.MovieClient) *MovieHandler {
	return &MovieHandler{client: client}
}

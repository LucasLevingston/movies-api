package httphandler

import "movies-api/api-gateway/internal/ports"

func NewMovieHandler(client ports.MovieClient) *MovieHandler {
	return &MovieHandler{client: client}
}

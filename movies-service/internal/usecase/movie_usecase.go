package usecase

import "movies-api/movies-service/internal/ports"

// NewMovieUseCase creates a MovieService backed by the given repository.
func NewMovieUseCase(repo ports.MovieRepository) ports.MovieService {
	return &movieUseCase{repo: repo}
}

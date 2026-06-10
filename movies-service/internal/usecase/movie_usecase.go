package usecase

import "movies-api/movies-service/internal/ports"

// NewMovieUseCase creates a MovieService backed by the given repository.
// publisher may be nil; events are skipped when it is.
func NewMovieUseCase(repo ports.MovieRepository, publisher ports.EventPublisher) ports.MovieService {
	return &movieUseCase{repo: repo, publisher: publisher}
}

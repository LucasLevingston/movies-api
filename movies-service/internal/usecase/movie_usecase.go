package usecase

import "movies-api/movies-service/internal/ports"

func NewMovieUseCase(repo ports.MovieRepository) ports.MovieService {
	return &movieUseCase{repo: repo}
}

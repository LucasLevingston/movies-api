package usecase

import (
	"movies-api/movies-service/internal/ports"
)

type movieUseCase struct {
	repo ports.MovieRepository
}

func NewMovieUseCase(repo ports.MovieRepository) ports.MovieService {
	return &movieUseCase{repo: repo}
}

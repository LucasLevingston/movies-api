package usecase

import "movies-api/movies-service/internal/ports"

type movieUseCase struct {
	repo ports.MovieRepository
}

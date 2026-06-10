package usecase

import (
	"context"

	"movies-api/movies-service/internal/domain"
)

func (useCase *movieUseCase) GetAll(ctx context.Context) ([]domain.Movie, error) {
	return useCase.repo.FindAll(ctx)
}

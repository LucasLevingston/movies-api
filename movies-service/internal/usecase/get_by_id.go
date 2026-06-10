package usecase

import (
	"context"

	"movies-api/movies-service/internal/domain"
)

func (useCase *movieUseCase) GetByID(ctx context.Context, id string) (*domain.Movie, error) {
	return useCase.repo.FindByID(ctx, id)
}

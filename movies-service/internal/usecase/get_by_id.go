package usecase

import (
	"context"

	"movies-api/movies-service/internal/domain"
)

func (u *movieUseCase) GetByID(ctx context.Context, id string) (*domain.Movie, error) {
	return u.repo.FindByID(ctx, id)
}

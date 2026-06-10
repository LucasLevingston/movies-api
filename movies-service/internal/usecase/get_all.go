package usecase

import (
	"context"

	"movies-api/movies-service/internal/domain"
)

func (u *movieUseCase) GetAll(ctx context.Context) ([]domain.Movie, error) {
	return u.repo.FindAll(ctx)
}

package usecase

import (
	"context"

	"movies-api/movies-service/internal/domain"
)

func (u *movieUseCase) Create(ctx context.Context, externalID int32, title, year string) (*domain.Movie, error) {
	movie := &domain.Movie{
		ExternalID: externalID,
		Title:      title,
		Year:       year,
	}
	return u.repo.Create(ctx, movie)
}

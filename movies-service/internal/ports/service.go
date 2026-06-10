package ports

import (
	"context"

	"movies-api/movies-service/internal/domain"
)

type MovieService interface {
	GetAll(ctx context.Context) ([]domain.Movie, error)
	GetByID(ctx context.Context, id string) (*domain.Movie, error)
	Create(ctx context.Context, externalID int32, title, year string) (*domain.Movie, error)
	Delete(ctx context.Context, id string) error
}

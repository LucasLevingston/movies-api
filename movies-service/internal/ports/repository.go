package ports

import (
	"context"

	"movies-api/movies-service/internal/domain"
)

// MovieRepository defines persistence operations for movies.
type MovieRepository interface {
	FindAll(ctx context.Context) ([]domain.Movie, error)
	FindByID(ctx context.Context, id string) (*domain.Movie, error)
	Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error)
	Delete(ctx context.Context, id string) error
}

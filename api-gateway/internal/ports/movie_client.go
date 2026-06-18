package ports

import (
	"context"

	"movies-api/api-gateway/internal/domain"
)


type MovieClient interface {
	ListMovies(ctx context.Context) ([]domain.Movie, error)
	GetMovie(ctx context.Context, id string) (*domain.Movie, error)
	CreateMovie(ctx context.Context, externalID int32, title, year string) (*domain.Movie, error)
	DeleteMovie(ctx context.Context, id string) error
}

package usecase

import (
	"context"

	"movies-api/movies-service/internal/domain"
	"movies-api/movies-service/internal/ports"
)

type movieUseCase struct {
	repo ports.MovieRepository
}

func NewMovieUseCase(repo ports.MovieRepository) ports.MovieService {
	return &movieUseCase{repo: repo}
}

func (u *movieUseCase) GetAll(ctx context.Context) ([]domain.Movie, error) {
	return u.repo.FindAll(ctx)
}

func (u *movieUseCase) GetByID(ctx context.Context, id string) (*domain.Movie, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *movieUseCase) Create(ctx context.Context, externalID int32, title, year string) (*domain.Movie, error) {
	movie := &domain.Movie{
		ExternalID: externalID,
		Title:      title,
		Year:       year,
	}
	return u.repo.Create(ctx, movie)
}

func (u *movieUseCase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

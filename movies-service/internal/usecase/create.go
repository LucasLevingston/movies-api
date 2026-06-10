package usecase

import (
	"context"
	"log"

	"movies-api/movies-service/internal/domain"
)

func (useCase *movieUseCase) Create(ctx context.Context, externalID int32, title, year string) (*domain.Movie, error) {
	movie := &domain.Movie{
		ExternalID: externalID,
		Title:      title,
		Year:       year,
	}
	created, err := useCase.repo.Create(ctx, movie)
	if err != nil {
		return nil, err
	}
	if useCase.publisher != nil {
		if publishErr := useCase.publisher.PublishMovieCreated(ctx, *created); publishErr != nil {
			log.Printf("publish movie.created event: %v", publishErr)
		}
	}
	return created, nil
}

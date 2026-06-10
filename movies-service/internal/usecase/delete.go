package usecase

import (
	"context"
	"log"
)

func (useCase *movieUseCase) Delete(ctx context.Context, id string) error {
	if err := useCase.repo.Delete(ctx, id); err != nil {
		return err
	}
	if useCase.publisher != nil {
		if publishErr := useCase.publisher.PublishMovieDeleted(ctx, id); publishErr != nil {
			log.Printf("publish movie.deleted event: %v", publishErr)
		}
	}
	return nil
}

package ports

import (
	"context"

	"movies-api/movies-service/internal/domain"
)

// EventPublisher defines the contract for publishing domain events.
type EventPublisher interface {
	PublishMovieCreated(ctx context.Context, movie domain.Movie) error
	PublishMovieDeleted(ctx context.Context, id string) error
}

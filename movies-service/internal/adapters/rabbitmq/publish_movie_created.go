package rabbitmq

import (
	"context"
	"encoding/json"

	"movies-api/movies-service/internal/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (publisher *movieEventPublisher) PublishMovieCreated(ctx context.Context, movie domain.Movie) error {
	body, err := json.Marshal(movie)
	if err != nil {
		return err
	}
	return publisher.channel.PublishWithContext(ctx, exchangeName, routingKeyCreated, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

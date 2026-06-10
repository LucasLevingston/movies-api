package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (publisher *movieEventPublisher) PublishMovieDeleted(ctx context.Context, id string) error {
	body, err := json.Marshal(map[string]string{"id": id})
	if err != nil {
		return err
	}
	return publisher.channel.PublishWithContext(ctx, exchangeName, routingKeyDeleted, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

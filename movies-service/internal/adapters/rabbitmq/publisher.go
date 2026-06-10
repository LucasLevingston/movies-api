package rabbitmq

import (
	"movies-api/movies-service/internal/ports"

	amqp "github.com/rabbitmq/amqp091-go"
)

// NewMovieEventPublisher creates an EventPublisher backed by RabbitMQ.
func NewMovieEventPublisher(uri string) (ports.EventPublisher, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err := channel.ExchangeDeclare(
		exchangeName,
		exchangeKind,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	return &movieEventPublisher{channel: channel}, nil
}

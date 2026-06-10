package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

const (
	exchangeName      = "movies"
	exchangeKind      = "topic"
	routingKeyCreated = "movie.created"
	routingKeyDeleted = "movie.deleted"
)

type movieEventPublisher struct {
	channel *amqp.Channel
}

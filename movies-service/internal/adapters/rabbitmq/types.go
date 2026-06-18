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

type movieCreatedPayload struct {
	ID         string `json:"id"`
	ExternalID int32  `json:"external_id"`
	Title      string `json:"title"`
	Year       string `json:"year"`
}

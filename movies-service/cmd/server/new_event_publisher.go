package main

import (
	"log"

	"movies-api/movies-service/internal/adapters/rabbitmq"
	"movies-api/movies-service/internal/ports"
)

func newEventPublisher() ports.EventPublisher {
	rabbitURI := getEnv("RABBITMQ_URI", "")
	if rabbitURI == "" {
		log.Println("RABBITMQ_URI not set, event publishing disabled")
		return nil
	}

	publisher, err := rabbitmq.NewMovieEventPublisher(rabbitURI)
	if err != nil {
		log.Printf("rabbitmq connect failed, event publishing disabled: %v", err)
		return nil
	}

	log.Println("connected to RabbitMQ")
	return publisher
}

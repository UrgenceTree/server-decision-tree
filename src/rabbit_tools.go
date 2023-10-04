package main

import (
	amqp "github.com/streadway/amqp"
)

// RabbitMQService struct is a service that encapsulates operations for RabbitMQ.
// It contains a pointer to the RabbitMQ Connection, Channel and the name of the Queue.
type RabbitMQService struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
}

// NewRabbitMQService is a constructor function that returns a new instance of RabbitMQService.
// It takes as arguments the URI of the RabbitMQ server and the name of the queue.
// It initializes a connection, a channel and a queue on RabbitMQ and returns a pointer to the instance of RabbitMQService.
func NewRabbitMQService(uri string, queueName string) (*RabbitMQService, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQService{
		Connection: conn,
		Channel:    ch,
		QueueName:  queueName,
	}, nil
}

// Publish method of RabbitMQService publishes a message to the queue.
// It takes as argument the message to be published as a string.
// It converts the string message to a byte slice and publishes it to the queue.
func (s *RabbitMQService) Publish(message string) error {
	err := s.Channel.Publish(
		"",
		s.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}
	return nil
}

// Consume method of RabbitMQService consumes a message from the queue.
// It returns a channel of Delivery instances, each of which represents an AMQP message.
func (s *RabbitMQService) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := s.Channel.Consume(
		s.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

// Close method of RabbitMQService closes the channel and the connection of RabbitMQ.
func (s *RabbitMQService) Close() {
	s.Channel.Close()
	s.Connection.Close()
}

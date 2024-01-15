package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

type Patient struct {
	PhoneNumber string `json:"phone_number"`
	Question    string `json:"question"`
}

const (
	rabbitMQURI   = "amqp://guest:guest@localhost:5672/"
	recvQueueName = "server_data"
)

func main() {

	conn, err := amqp.Dial(rabbitMQURI)
	if err != nil {
		log.Fatalf("[SERVER DATA] : Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("[SERVER DATA] : Failed to open a channel: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(recvQueueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("[SERVER DATA] : Failed to declare recv queue: %v", err)
	}

	// Start a goroutine to receive messages
	go func() {

		msgs, err := ch.Consume(recvQueueName, "", true, false, false, false, nil)
		if err != nil {
			log.Fatalf("[SERVER DATA] : Failed to register a consumer: %v", err)
		}

		for msg := range msgs {
			log.Printf("[SERVER DATA] : Received a message: %s", msg.Body)
		}
	}()

	// Signal handling for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signals:
		log.Println("[SERVER DATA] : Interrupt received, shutting down...")
		return
	}
}

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

const (
	rabbitMQURI  = "amqp://guest:guest@localhost:5672/"
	queueName    = "decision_tree"
	messageCount = 10000 // number of messages to send
)

func main() {

	conn, err := amqp.Dial(rabbitMQURI)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclarePassive(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("QueueDeclarePassive failed: %v", err)
	}

	start := time.Now()

	for i := 0; i < messageCount; i++ {
		body := fmt.Sprintf("message %d", i)
		err = ch.Publish("", queueName, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
		if err != nil {
			log.Fatalf("Failed to publish a message: %v", err)
		}
	}

	duration := time.Since(start)

	log.Printf("Sent %d messages in %v", messageCount, duration)
}

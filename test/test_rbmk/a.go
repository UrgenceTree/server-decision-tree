package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {

	message := flag.String("message", "Hello, World!", "Message to send to RabbitMQ")
	flag.Parse()

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Publish a message
	err = ch.Publish(
		"",              // exchange
		"decision_tree", // routing key (queue name)
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(*message),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	fmt.Println("Sent message:", *message)
}

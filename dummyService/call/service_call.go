package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/streadway/amqp"
)

type UserMessage struct {
	PhoneNumber string `json:"phone_number"`
	Text        string `json:"text"`
}

type Question struct {
	PhoneNumber string `json:"phone_number"`
	Question    string `json:"question"`
}

const (
	rabbitMQURI   = "amqp://guest:guest@localhost:5672/"
	sendQueueName = "decision_tree"
	recvQueueName = "server_call"
)

func main() {

	conn, err := amqp.Dial(rabbitMQURI)
	if err != nil {
		log.Fatalf("[SERVER CALL] : Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("[SERVER CALL] : Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the queues
	_, err = ch.QueueDeclare(sendQueueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("[SERVER CALL] : Failed to declare send queue: %v", err)
	}

	_, err = ch.QueueDeclare(recvQueueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("[SERVER CALL] : Failed to declare recv queue: %v", err)
	}

	// Start a goroutine to receive messages
	go func() {

		msgs, err := ch.Consume(recvQueueName, "", true, false, false, false, nil)
		if err != nil {
			log.Fatalf("[SERVER CALL] : Failed to register a consumer: %v", err)
		}

		for msg := range msgs {
			log.Printf("[SERVER CALL] : Received a message: %s", msg.Body)
		}
	}()

	// Signal handling for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Send messages in a loop
	i := 0
	for {
		select {

		case <-signals:
			log.Println("[SERVER CALL] : Interrupt received, shutting down...")
			return

		default:
			userMessage := UserMessage{
				PhoneNumber: "0101010101",
				Text:        fmt.Sprintf("Message %d", i),
			}
			body, err := json.Marshal(userMessage)
			if err != nil {
				log.Fatalf("[SERVER CALL] : Failed to marshal user message: %v", err)
			}
			err = ch.Publish("", sendQueueName, false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
			if err != nil {
				log.Fatalf("[SERVER CALL] : Failed to publish a message: %v", err)
			} else {
				log.Printf("[SERVER CALL] : Message %d published : ", i, userMessage)
			}
			i++
			time.Sleep(time.Millisecond * 1000) // Add a slight delay to prevent overwhelming the queue
		}
	}
}

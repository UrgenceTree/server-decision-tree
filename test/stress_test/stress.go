package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/streadway/amqp"
)

type UserMessage struct {
	PhoneNumber string `json:"phone_number"`
	Text        string `json:"text"`
}

const (
	rabbitMQURI   = "amqp://guest:guest@localhost:5672/"
	sendQueueName = "decision_tree"
	recvQueueName = "server_call"
)

func main() {

	var medianTime time.Duration
	var medianCount float32

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

	_, err = ch.QueueDeclare(sendQueueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare send queue: %v", err)
	}

	_, err = ch.QueueDeclare(recvQueueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare recv queue: %v", err)
	}

	// Signal handling for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Variables for tracking messages
	var numberOfMessages int
	var receivedCount int
	var wg sync.WaitGroup

	// Start a goroutine to receive messages
	go func() {
		msgs, err := ch.Consume(recvQueueName, "", true, false, false, false, nil)
		if err != nil {
			log.Fatalf("Failed to register a consumer: %v", err)
		}

		wg.Add(1)
		defer wg.Done()

		for msg := range msgs {
			//log.Printf("Received a message: %s", msg.Body)
			msg.Ack(false)
			receivedCount++
			if receivedCount == numberOfMessages {
				break
			}
		}
	}()

	for {
		select {
		case <-signals:
			log.Println("Interrupt received, shutting down...")
			return
		default:
		}

		fmt.Print("Enter the number of messages to send: ")
		_, err = fmt.Scanln(&numberOfMessages)
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}

		if numberOfMessages == 0 {
			return
		}

		start := time.Now() // Start timer for sending messages

		for i := 0; i < numberOfMessages; i++ {
			select {
			case <-signals:
				log.Println("Interrupt received, shutting down...")
				return
			default:
				userMessage := UserMessage{
					PhoneNumber: "test",
					Text:        fmt.Sprintf("Message %d", i),
				}
				body, err := json.Marshal(userMessage)
				if err != nil {
					log.Fatalf("Failed to marshal user message: %v", err)
				}
				err = ch.Publish("", sendQueueName, false, false, amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
				if err != nil {
					log.Fatalf("Failed to publish a message: %v", err)
				}
			}
		}

		// Wait for all messages to be received
		wg.Wait()
		elapsed := time.Since(start)
		fmt.Printf("Time taken to send and receive %d messages: %s\n", numberOfMessages, elapsed)
		medianCount++
		medianTime += time.Duration(float64(elapsed) / float64(medianCount))

		fmt.Printf("Time median: %s\n", medianTime)
	}
}

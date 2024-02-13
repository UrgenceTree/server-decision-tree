package main

import (
	"encoding/json"

	amqp "github.com/streadway/amqp"
)

type UserAPI struct {
	decisionTree DecisionTree
	conn         *amqp.Connection
	ch           *amqp.Channel
}

func NewUserAPI() *UserAPI {

	return &UserAPI{
		decisionTree: DecisionTree{},
	}
}

func (api *UserAPI) rabbitMQInit() error {

	LogInfo("function=UserAPI::rabbitMQInit, message=Initializing RabbitMQ...")

	var err error

	LogInfo("function=UserAPI::rabbitMQInit, message=Connecting to RabbitMQ...")
	if api.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/"); err != nil {
		LogFatal("function=UserAPI::connectToRabbitMQ, error=%v", err.Error())
		return err
	}
	LogSuccess("function=UserAPI::rabbitMQInit, message=Connected to RabbitMQ")

	LogInfo("function=UserAPI::rabbitMQInit, message=Creating channel...")
	if api.ch, err = api.conn.Channel(); err != nil {
		LogFatal("function=UserAPI::createChannel, error=%v", err.Error())
		return err
	}
	LogSuccess("function=UserAPI::rabbitMQInit, message=Channel created")

	LogInfo("function=UserAPI::rabbitMQInit, message=Declaring queue...")
	_, err = api.ch.QueueDeclarePassive(
		"decision_tree", // name
		false,           // durable?
		false,           // delete when unused?
		false,           // exclusive?
		false,           // no-wait?
		nil,             // arguments
	)
	if err != nil {
		LogError("function=UserAPI::declareQueue, error=%v", err.Error())
		return err
	}
	LogSuccess("function=UserAPI::rabbitMQInit, message=Queue declared")

	return nil
}

func (api *UserAPI) LoadTree(treeConfFilepath string) error {

	LogInfo("function=UserAPI::LoadTree, message=Loading decision tree from %s", treeConfFilepath)

	tree := parser(treeConfFilepath)

	LogDebug("function=UserAPI::LoadTree, message=Decision tree loaded: %+v", tree)

	return nil
}

type Patient struct {
	PhoneNumber string `json:"phone_number"`
	rank        int    `json:"rank"`
}

type Question_ struct {
	PhoneNumber string `json:"phone_number"`
	Question    string `json:"question"`
}

func patientCreator() Patient {

	return Patient{
		PhoneNumber: "0123456789",
		rank:        0,
	}
}

func questionCreator() Question_ {

	return Question_{
		PhoneNumber: "0123456789",
		Question:    "La victime respire ?",
	}
}

// func (api *UserAPI) HandleUser(userMsg *UserMessage) error {

// 	LogInfo("function=UserAPI::HandleUser, message=Handling user: %s", userMsg)

// 	// SEND TO SERVER DATA //

// 	// create Patient, will be algorithm
// 	body, err := json.Marshal(patientCreator())
// 	if err != nil {
// 		log.Fatalf("Failed to marshal user message: %v", err)
// 	}

// 	err = api.ch.Publish("", "server_data", false, false, amqp.Publishing{
// 		ContentType: "text/plain",
// 		Body:        body,
// 	})
// 	if err != nil {
// 		LogFatal("Failed to publish a message: %v", err)
// 		if userMsg.PhoneNumber != "test" {
// 			return err
// 		}
// 	} else {
// 		LogSuccess("function=UserAPI::HandleUser, message=Message published to server_data")
// 	}

// 	// SEND TO SERVER CALL //

// 	if body, err = json.Marshal(questionCreator()); err != nil {
// 		log.Fatalf("Failed to marshal user message: %v", err)
// 	}

// 	err = api.ch.Publish("", "server_call", false, false, amqp.Publishing{
// 		ContentType: "text/plain",
// 		Body:        body,
// 	})
// 	if err != nil {
// 		LogFatal("Failed to publish a message: %v", err)
// 		if userMsg.PhoneNumber != "test" {
// 			return err
// 		}
// 	} else {
// 		LogSuccess("function=UserAPI::HandleUser, message=Message published to server_call")
// 	}

// 	return nil
// }

func (api *UserAPI) HandleUser(userMsg *UserMessage) error {

	LogInfo("function=UserAPI::HandleUser, message=Handling user: %s", userMsg)

	// SEND TO SERVER DATA //
	if err := api.publishToQueue(patientCreator(), "server_data", userMsg.PhoneNumber); err != nil {
		return err
	}

	// SEND TO SERVER CALL //
	if err := api.publishToQueue(questionCreator(), "server_call", userMsg.PhoneNumber); err != nil {
		return err
	}

	return nil
}

func (api *UserAPI) publishToQueue(data interface{}, queueName string, phoneNumber string) error {

	body, err := json.Marshal(data)
	if err != nil {
		LogError("Failed to marshal message: %v", err)
		return err
	}

	err = api.ch.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})

	if err != nil {
		LogError("Failed to publish a message: %v", err)
		if phoneNumber != "test" {
			return err
		}
	} else {
		LogSuccess("function=UserAPI::publishToQueue, message=Message published to %s", queueName)
	}

	return nil
}

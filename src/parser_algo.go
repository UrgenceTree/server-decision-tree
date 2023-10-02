package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	//"strconv"
	//"flag"


	"github.com/streadway/amqp"
)

type jsonContent struct {
	FirstQuestion string
	Questions     []string
	YesScore      []float64
	NoScore       []float64
	Link          []string
	Score			float64
}

type RabbitMQ struct {
	connection *amqp.Connection
	channel   *amqp.Channel
}

func initRabbitMQ(mq *RabbitMQ)  {
	err := error(nil)

	mq.connection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    //defer mq.connection.Close()

    mq.channel, err = mq.connection.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    //defer mq.channel.Close()
}

func sendToRabbitMQ(mq *RabbitMQ, message string) {
	err := mq.channel.Publish(
		"",              // exchange
		"decision_tree", // routing key (queue name)
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	fmt.Println("Sent message:", message)
}

func fillScore(questions []interface{}, elem *jsonContent) {
	tmpYes := []float64{}
	tmpNo := []float64{}

	for i := 0; i < len(questions); i++ {
		firstQuestion := questions[i].(map[string]interface{})
		yesScore := firstQuestion["yes_score"].(float64)
		noScore := firstQuestion["no_score"].(float64)
		tmpYes = append(tmpYes, yesScore)
		tmpNo = append(tmpNo, noScore)
	}

	elem.YesScore = tmpYes
	elem.NoScore = tmpNo
}

func fillQuestions(questions []interface{}, elem *jsonContent) {
	tmpQuestions := []string{}

	for i := 0; i < len(questions); i++ {
		currQuestion := questions[i].(map[string]interface{})
		tmp := currQuestion["question"].(string)
		tmpQuestions = append(tmpQuestions, tmp)
	}

	elem.Questions = tmpQuestions
}

func fillLinks(questions []interface{}, elem *jsonContent) {
	tmpLinks := []string{}

	for i := 0; i < len(questions); i++ {
		currLink := questions[i].(map[string]interface{})
		tmp := currLink["link"].(string)
		tmpLinks = append(tmpLinks, tmp)
	}

	elem.Link = tmpLinks
}

func orderLists(firstQuestion string, elem *jsonContent) {
	newQuestions := []string{}
	newyesScore := []float64{}
	newnoScore := []float64{}
	newLink := []string{}
	target := ""

	for i := 0; i < len(elem.Questions); i++ {
		curr := elem.Questions[i]
		if curr == firstQuestion {
			newQuestions = append(newQuestions, curr)
			newyesScore = append(newyesScore, elem.YesScore[i])
			newnoScore = append(newnoScore, elem.NoScore[i])
			newLink = append(newLink, elem.Link[i])
			target = elem.Link[i]
			break
		}
	}

	for i := 0; i < len(elem.Questions); i++ {
		curr := elem.Questions[i]
		if curr == target {
			newQuestions = append(newQuestions, curr)
			newyesScore = append(newyesScore, elem.YesScore[i])
			newnoScore = append(newnoScore, elem.NoScore[i])
			newLink = append(newLink, elem.Link[i])
			target = elem.Link[i]
			i = -1
		}
	}

	elem.Questions = newQuestions
	elem.YesScore = newyesScore
	elem.NoScore = newnoScore
	elem.Link = newLink	
}


func main() {

	mq := RabbitMQ{}
	
	initRabbitMQ(&mq)
	
	jsonFile, err := os.Open("adc.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	questions := result["questions"].([]interface{})
	elem := jsonContent{}

	elem.FirstQuestion = result["starting_question"].(string)
	fillScore(questions, &elem)
	fillQuestions(questions, &elem)
	fillLinks(questions, &elem)
	
	orderLists(elem.FirstQuestion, &elem)

	var responses []string

	i := 0
	for i < len(elem.Questions) {
		fmt.Printf("Question %d: %s\n", i, elem.Questions[i])
		var response string
		fmt.Scanln(&response)

		if response == "yes" {
			elem.Score = elem.Score + elem.YesScore[i]
		} else if response == "no" {
			elem.Score = elem.Score + elem.NoScore[i]
		}
		
		sendToRabbitMQ(&mq, response)
		responses = append(responses, response)

		if elem.Link[i] == "null" {
			fmt.Println("\nEnd of questions.\n")
			break
		}
		i++
	}

	for i, response := range responses {
		fmt.Printf("Response %d: %s\n", i, response)
	}

	fmt.Printf("Score: %.0f\n", elem.Score)
}
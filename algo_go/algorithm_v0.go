package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type json_content struct {
    FirstQuestion string
    Questions     []string
    YesScore      []float64
    NoScore       []float64
    Link          []string
}


func fill_score(questions []interface{}, elem *json_content) {
    
    tmp_yes := [] float64{}
    tmp_no := [] float64{}
    
    for i := 0; i < len(questions); i++ {
        firstQuestion := questions[i].(map[string]interface{})
        yesScore := firstQuestion["yes_score"].(float64)
        noScore := firstQuestion["no_score"].(float64)
        tmp_yes = append(tmp_yes, yesScore)
        tmp_no = append(tmp_no, noScore)
    }
    
    elem.YesScore = tmp_yes
    elem.NoScore = tmp_no
}

func fill_questions(questions []interface{}, elem *json_content) {
    
    tmp_questions := [] string{}
    
    for i := 0; i < len(questions); i++ {
        curr_question := questions[i].(map[string]interface{})
        tmp := curr_question["question"].(string)
        tmp_questions = append(tmp_questions, tmp)
    }

    elem.Questions = tmp_questions
}

func fill_links(questions []interface{}, elem *json_content) {
    
    tmp_links := [] string{}
    
    for i := 0; i < len(questions); i++ {
        curr_link := questions[i].(map[string]interface{})
        tmp := curr_link["link"].(string)
        tmp_links = append(tmp_links, tmp)
    }

    elem.Link = tmp_links
}

//func print_all_questions(questions []interface{}) {
//    for i := 0; i < len(questions); i++ {
//        fmt.Println(questions[i].(map[string]interface{})["question"])
//    }
//}

func main() {
	jsonFile, err := os.Open("adc.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	questions := result["questions"].([]interface{})
	elem := json_content{}

	elem.FirstQuestion = result["starting_question"].(string)
	fill_score(questions, &elem)
	fill_questions(questions, &elem)
	fill_links(questions, &elem)

	var responses []string
	var firstResponse string
	var secondResponse string
	
	fmt.Printf("Question 1: %s\n", elem.FirstQuestion)
	fmt.Scanln(&firstResponse)
	responses = append(responses, firstResponse)

	fmt.Printf("\nQuestion %d: %s\n", 2, elem.Questions[0])
	fmt.Scanln(&secondResponse)
	responses = append(responses, secondResponse)



	// Ask questions
	i := 2
	for i < len(elem.Questions) {
		fmt.Printf("\nQuestion %d: %s\n", i+1, elem.Questions[i])
		//fmt.Print("Enter your response (Yes/No): ")

		var response string
		fmt.Scanln(&response)
		responses = append(responses, response)

		//var nextQuestionIndex int
		
		//if response == "Yes" {
		//	fmt.Printf("Yes Score: %.0f\n", elem.YesScore[i])
		//	nextQuestionIndex = int(elem.YesScore[i])
		//} else if response == "No" {
		//	fmt.Printf("No Score: %.0f\n", elem.NoScore[i])
		//	nextQuestionIndex = int(elem.NoScore[i])
		//} else {
		//	fmt.Println("Invalid response. Please enter either 'Yes' or 'No'.")
		//	continue
		//}

		if elem.Link[i] == "null" {
			fmt.Println("\nEnd of questions.\n")
			break
		}

		//if nextQuestionIndex < 0 || nextQuestionIndex >= len(elem.Questions) {
		//	fmt.Println("End of questions.")
		//	break
		//}

		//fmt.Printf("Next Question: %s\n", elem.Questions[i+1])
		//i = nextQuestionIndex
		i++
	}

	for i, response := range responses {
		fmt.Printf("Response %d: %s\n", i, response)
	}
}
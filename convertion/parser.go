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

func isInList(curr_target string, sorted_list []string) bool {
    found := false

    for _, item := range sorted_list {
        if item == curr_target {
            found = true
            break
        }
    }

    if !found {
        return false
    }
    return true
}


func sortByLink( elem *json_content) {
    sorted_list := [] string{}
    sorted_list_yes := [] float64{}
    sorted_list_no := [] float64{}
    curr_target := "What is your name?"

    for i := 0; i != len(elem.Questions); i++ {
		for j := 0; j != len(elem.Questions); j++ {
            if curr_target == elem.Questions[j] {
                sorted_list = append(sorted_list, elem.Questions[j])
                sorted_list_yes = append(sorted_list_yes, elem.YesScore[j])
                sorted_list_no = append(sorted_list_no, elem.NoScore[j])
				curr_target = elem.Link[j]
				i = -1
				break
				}        
        }
    }
       
    for i := 0; i < len(sorted_list); i++ {
        fmt.Printf("Question : %s, Yes: %.0f, No: %.0f\n", sorted_list[i], sorted_list_yes[i], sorted_list_no[i])
    }
}

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

    //fmt.Println(elem.FirstQuestion)
    //fmt.Println(elem.YesScore)
    //fmt.Println(elem.NoScore)
    //fmt.Println(elem.Questions)
    //fmt.Println(elem.Link)

    sortByLink(&elem)

    // fmt.Printf("Question 1: %s\n", elem.FirstQuestion)
    // for i := 0; i < len(elem.Questions); i++ {
    //     fmt.Printf("Question %d: %s, Yes: %.0f, No: %.0f, Link: %s\n", i+2, elem.Questions[i], elem.YesScore[i], elem.NoScore[i], elem.Link[i])
    // }
}

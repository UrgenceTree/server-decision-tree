package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

type Questions struct {
    answers_yes []string
    answers_no  []string
}

type Tree struct {
    score       int
    last_action string
}

type Request struct{}


func NewQuestions() *Questions {
    return &Questions{
        answers_yes: []string{"yes", "Yes.", "absolutely", "Absolutely.", "indeed", "Indeed.", "totally", "Totally."},
        answers_no:  []string{"no", "No.", "not", "Not.", "is not", "He's not.", "he's not"},
    }
}

func (q *Questions) Malaise(tree *Tree, rq *Request) {
    for {
        fmt.Println("Did the victim fainted ?")
        line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        line = strings.TrimSpace(line)
        if contains(q.answers_yes, line) {
            tree.score += 10
            tree.last_action = "Yes."
            break
        } else if contains(q.answers_no, line) {
            tree.last_action = "No."
            break
        }
    }
}

func (q *Questions) Cardiac_arrest(tree *Tree, rq *Request) {
    for {
        fmt.Println("\nIs the victim in cardiac arrest ?")
        line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        line = strings.TrimSpace(line)
        if contains(q.answers_yes, line) {
            tree.score += 100
            tree.last_action = "Yes."
            break
        } else if contains(q.answers_no, line) {
            tree.last_action = "No."
            break
        }
    }
}

func (q *Questions) Symptome(tree *Tree, rq *Request) {
    for {
        fmt.Println("\nDoes the victim have any of the following symptoms ?\n" +
            "\t- Unconscious, don't speak anymore, don't open your eyes., don't watch, respond when you speak to him, reacts\n" +
            "\t- Difficulty breathing, to other BP related to breathing\n" +
            "\t- Signs of shock, pallor, sweating")
        line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        line = strings.TrimSpace(line)
        if contains(q.answers_yes, line) {
            tree.score += 10
            tree.last_action = "Yes."
            break
        } else if contains(q.answers_no, line) {
            tree.last_action = "No."
            break
        }
    }
}

func contains(slice []string, s string) bool {
    for _, str := range slice {
        if s == str {
            return true
        }
    }
    return false
}

func main() {
    questions := NewQuestions()
    tree := &Tree{}
    rq := &Request{}

    questions.Malaise(tree, rq)
    questions.Cardiac_arrest(tree, rq)
    questions.Symptome(tree, rq)
    fmt.Println("Score final du patient :", tree.score)
}

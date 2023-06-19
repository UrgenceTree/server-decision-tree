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

func (q *Questions) Malaise(tree *Tree, rq *Request, input string) {
	line := strings.TrimSpace(input)
	if contains(q.answers_yes, line) {
		tree.score += 10
		tree.last_action = "Yes."
	} else if contains(q.answers_no, line) {
		tree.last_action = "No."
	}
}

func (q *Questions) Cardiac_arrest(tree *Tree, rq *Request, input string) {
	line := strings.TrimSpace(input)
	if contains(q.answers_yes, line) {
		tree.score += 100
		tree.last_action = "Yes."
	} else if contains(q.answers_no, line) {
		tree.last_action = "No."
	}
}

func (q *Questions) Symptome(tree *Tree, rq *Request, input string) {
	line := strings.TrimSpace(input)
	if contains(q.answers_yes, line) {
		tree.score += 10
		tree.last_action = "Yes."
	} else if contains(q.answers_no, line) {
		tree.last_action = "No."
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
	if len(os.Args) < 2 {
		fmt.Println("Veuillez spécifier un fichier en paramètre.")
		return
	}
	filename := os.Args[1] 
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	questions := NewQuestions()
	tree := &Tree{}
	rq := &Request{}

	questions.Malaise(tree, rq, lines[0])
	questions.Cardiac_arrest(tree, rq, lines[1])
	questions.Symptome(tree, rq, lines[2])

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	fmt.Println("Score final du patient :", tree.score)
}

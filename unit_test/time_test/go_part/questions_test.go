package main

import (
	"testing"
	"os"
	"fmt"
	"io"
	"bytes"
	"strings"
)

func TestQuestions_Malaise(t *testing.T) {
	testCases := []struct {
		input         string
		expectedScore int
		expectedAction string
	}{
		{"No.", 0, "No."},
		{"Yes.", 10, "Yes."},
	}

	q := NewQuestions()
	tree := &Tree{}
	rq := &Request{}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			q.Malaise(tree, rq, tc.input)
			if tree.score != tc.expectedScore {
				t.Errorf("Got score %d, expected %d", tree.score, tc.expectedScore)
			}
			if tree.last_action != tc.expectedAction {
				t.Errorf("Got last action %s, expected %s", tree.last_action, tc.expectedAction)
			}
		})
	}
}

func TestQuestions_Cardiac_arrest(t *testing.T) {
	testCases := []struct {
		input         string
		expectedScore int
		expectedAction string
	}{
		{"No.", 0, "No."},
		{"Yes.", 100, "Yes."},
	}

	q := NewQuestions()
	tree := &Tree{}
	rq := &Request{}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			q.Cardiac_arrest(tree, rq, tc.input)
			if tree.score != tc.expectedScore {
				t.Errorf("Got score %d, expected %d", tree.score, tc.expectedScore)
			}
			if tree.last_action != tc.expectedAction {
				t.Errorf("Got last action %s, expected %s", tree.last_action, tc.expectedAction)
			}
		})
	}
}

func TestQuestions_Symptome(t *testing.T) {
	testCases := []struct {
		input         string
		expectedScore int
		expectedAction string
	}{
		{"No.", 0, "No."},
		{"Yes.", 10, "Yes."},
	}

	q := NewQuestions()
	tree := &Tree{}
	rq := &Request{}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			q.Symptome(tree, rq, tc.input)
			if tree.score != tc.expectedScore {
				t.Errorf("Got score %d, expected %d", tree.score, tc.expectedScore)
			}
			if tree.last_action != tc.expectedAction {
				t.Errorf("Got last action %s, expected %s", tree.last_action, tc.expectedAction)
			}
		})
	}
}

//func TestMain_FileNotFound(t *testing.T) {
//	// Simulate a scenario where the file does not exist.
//	os.Args = []string{"program", "file_1.txt"}
//	main()
//	// Add assertions to check for error handling.
//	// For example, check that an error message is printed.
//	// You can use a buffer to capture standard output and then check it.
//}

func TestQuestions_EmptyInput(t *testing.T) {
	q := NewQuestions()
	tree := &Tree{}
	rq := &Request{}

	tests := []struct {
		input         string
		expectedScore int
	}{
		{"", 0},
		// Add more test cases for empty input or unexpected characters.
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			q.Malaise(tree, rq, test.input)
			if tree.score != test.expectedScore {
				t.Errorf("Got score %d, expected %d", tree.score, test.expectedScore)
			}
		})
	}
}


func TestQuestions_Contains(t *testing.T) {
	testCases := []struct {
		slice    []string
		s        string
		expected bool
	}{
		{[]string{"yes", "Yes.", "absolutely"}, "Yes.", true},
		{[]string{"no", "No.", "not"}, "Maybe.", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v contains %s", tc.slice, tc.s), func(t *testing.T) {
			result := contains(tc.slice, tc.s)
			if result != tc.expected {
				t.Errorf("Got %v, expected %v", result, tc.expected)
			}
		})
	}
}

func TestMain_FileNotFound(t *testing.T) {
	// Simulate a scenario where the file does not exist.
	os.Args = []string{"program", "non_existent_file.txt"}

	// Redirect standard error to a buffer to capture error output.
	oldStderr := os.Stderr
	defer func() { os.Stderr = oldStderr }()
	r, w, _ := os.Pipe()
	os.Stderr = w

	main()

	w.Close()

	// Read the captured error message from the buffer.
	var stderr bytes.Buffer
	io.Copy(&stderr, r)
	errorMessage := stderr.String()

	// Assert that the error message contains the expected text.
	expectedErrorMessage := "Erreur lors de l'ouverture du fichier :"
	if !strings.Contains(errorMessage, expectedErrorMessage) {
		t.Errorf("Expected error message to contain '%s', but got '%s'", expectedErrorMessage, errorMessage)
	}
}

func TestMain_EmptyArgs(t *testing.T) {
	// Simulate running the program without specifying a file.
	os.Args = []string{"program"}

	// Redirect standard output to a buffer to capture output.
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()

	// Read the captured output from the buffer.
	var stdout bytes.Buffer
	io.Copy(&stdout, r)
	output := stdout.String()

	// Assert that the output contains the expected text.
	expectedOutput := "Veuillez spécifier un fichier en paramètre."
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, output)
	}
}

func TestMain_FileFound(t *testing.T) {
	// Simulate a scenario where the file does not exist.
	os.Args = []string{"program", "file_1.txt"}

	// Redirect standard error to a buffer to capture error output.
	oldStderr := os.Stderr
	defer func() { os.Stderr = oldStderr }()
	r, w, _ := os.Pipe()
	os.Stderr = w

	main()

	w.Close()

	// Read the captured error message from the buffer.
	var stderr bytes.Buffer
	io.Copy(&stderr, r)
	errorMessage := stderr.String()

	// Assert that the error message contains the expected text.
	expectedErrorMessage := ""
	if !strings.Contains(errorMessage, expectedErrorMessage) {
		t.Errorf("Expected error message to contain '%s', but got '%s'", expectedErrorMessage, errorMessage)
	}
}

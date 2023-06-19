package main

import (
	"io/ioutil"
	//"log"
	"os"
	//"strconv"
	"strings"
	"testing"
)

func createTestEnvFile(t *testing.T) {
	content := []byte(`NB_COMMANDS 3
TEST 2 YES: CARDIAC_ARREST NO: NOTHING
CARDIAC_ARREST 1 YES: NOTHING NO: SYMPTOME
SYMPTOME 0 YES: NOTHING NO: NOTHING`)
	err := ioutil.WriteFile(".env", content, 0644)
	if err != nil {
		t.Fatal(err)
	}
}

//func TestInitEnv(t *testing.T) {
//	createTestEnvFile(t)
//	defer os.Remove(".env")
//
//	tree := &DecisionalTree{}
//	tree.initEnv()
//
//	if tree.servDataIP != "SERV_DATA_IP_VALUE" {
//		t.Errorf("Expected SERV_DATA_IP_VALUE, but got %s", tree.servDataIP)
//	}
//
//	if tree.servDataPort != SERV_DATA_PORT_VALUE {
//		t.Errorf("Expected SERV_DATA_PORT_VALUE, but got %d", tree.servDataPort)
//	}
//}

func TestHandleEnv(t *testing.T) {
	tree := &DecisionalTree{}
	tree.handleEnv()

	if len(tree.listOfCommands) != 0 {
		t.Errorf("Expected empty listOfCommands, but got %v", tree.listOfCommands)
	}

	if tree.name != "" {
		t.Errorf("Expected empty name, but got %s", tree.name)
	}

	if tree.surname != "" {
		t.Errorf("Expected empty surname, but got %s", tree.surname)
	}
}

func TestParseConf(t *testing.T) {
	createTestEnvFile(t)
	defer os.Remove(".env")

	tree := &DecisionalTree{}
	tree.parseConf()

	if tree.numberOfSteps != 3 {
		t.Errorf("Expected 3 numberOfSteps, but got %d", tree.numberOfSteps)
	}

	expectedCommands := []map[string]string{
		{"Question": "TEST", "Order": "2", "Yes.": "CARDIAC_ARREST", "No.": "NOTHING"},
		{"Question": "CARDIAC_ARREST", "Order": "1", "Yes.": "NOTHING", "No.": "SYMPTOME"},
		{"Question": "SYMPTOME", "Order": "0", "Yes.": "NOTHING", "No.": "NOTHING"},
	}

	for i, command := range tree.listOfCommands {
		expectedCommand := expectedCommands[i]
		if command["Question"] != expectedCommand["Question"] {
			t.Errorf("Expected question: %s, but got: %s", expectedCommand["Question"], command["Question"])
		}

		if command["Order"] != expectedCommand["Order"] {
			t.Errorf("Expected order: %s, but got: %s", expectedCommand["Order"], command["Order"])
		}

		if command["Yes."] != expectedCommand["Yes."] {
			t.Errorf("Expected yes link: %s, but got: %s", expectedCommand["Yes."], command["Yes."])
		}

		if command["No."] != expectedCommand["No."] {
			t.Errorf("Expected no link: %s, but got: %s", expectedCommand["No."], command["No."])
		}
	}
}

func TestMainFunction(t *testing.T) {
	createTestEnvFile(t)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	expectedOutput := `Question:
TEST
CARDIAC_ARREST
SYMPTOME

Ordre:
2
1
0

Link Yes:
CARDIAC_ARREST
NOTHING
NOTHING

Link No:
NOTHING
SYMPTOME
NOTHING`

	if strings.TrimSpace(string(out)) != expectedOutput {
		t.Errorf("Expected output:\n%s\nBut got:\n%s", expectedOutput, out)
	}
}

package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "strings"
)

type DecisionalTree struct {
    step             int
    score            int
    numberOfSteps    int
    lastAction       string
    listOfCommands   []map[string]string
    actionOrder      []string
    servDataIP       string
    servDataPort     int
	name             string
	surname          string
}

func (tree *DecisionalTree) initEnv() {
    data, err := ioutil.ReadFile(".env")
    if err != nil {
        log.Fatal(err)
    }
    envVars := strings.Split(string(data), "\n")
    for _, envVar := range envVars {
        splitVar := strings.Split(envVar, "=")
        if len(splitVar) == 2 {
            switch splitVar[0] {
            case "SERV_DATA_IP":
                tree.servDataIP = splitVar[1]
            case "SERV_DATA_PORT":
                port, err := strconv.Atoi(splitVar[1])
                if err != nil {
                    log.Fatal(err)
                }
                if port < 1025 || port > 65535 {
                    fmt.Fprintln(os.Stderr, "ERROR: The PORT in the env must be between 1025 and 65535")
                    os.Exit(84)
                }
                tree.servDataPort = port
            }
        }
    }
    if tree.servDataIP == "" {
        fmt.Fprintln(os.Stderr, "ERROR: Bad ip adress")
        os.Exit(84)
    }
}

func (tree *DecisionalTree) handleEnv() {
    tree.listOfCommands = make([]map[string]string, 0)
    tree.name = ""
    tree.surname = ""
}

func (tree *DecisionalTree) parseConf() {
    confFile, err := os.Open(".env")
    if err != nil {
        log.Fatal(err)
    }
    defer confFile.Close()
    scanner := bufio.NewScanner(confFile)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "NB_COMMANDS") {
            tree.numberOfSteps, err = strconv.Atoi(strings.Split(line, " ")[1])
            if err != nil {
                log.Fatal(err)
            }
        } else {
            dict := make(map[string]string)
            parts := strings.Split(line, " ")
            dict["Question"] = strings.Title(parts[0])
            order, err := strconv.Atoi(parts[1])
            if err != nil {
                log.Fatal(err)
            }
            dict["Order"] = strconv.Itoa(order)
            dict["Yes."] = parts[3]
            dict["No."] = strings.Title(strings.TrimRight(parts[5], "\r\n"))
            tree.listOfCommands = append(tree.listOfCommands, dict)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func main() {
	tree := DecisionalTree{}
	tree.handleEnv()
	tree.parseConf()

	fmt.Println("Question:")
	for _, command := range tree.listOfCommands {
		fmt.Println(command["Question"])
	}
	fmt.Println("\nOrdre:")
	for _, command := range tree.listOfCommands {
		fmt.Println(command["Order"])
	}
	fmt.Println("\nLink Yes:")
	for _, command := range tree.listOfCommands {
		fmt.Println(command["Yes."])
	}
	fmt.Println("\nLink No:")
	for _, command := range tree.listOfCommands {
		fmt.Println(command["No."])
	}
}


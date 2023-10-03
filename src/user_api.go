package main

import (
	"fmt"
	"os"
	"service/tree"
)

type UserAPI struct {
	decisionTree tree.DecisionTree
}

func NewUserAPI() *UserAPI {

	return &UserAPI{
		decisionTree: tree.DecisionTree{},
	}
}

func (api *UserAPI) LoadTree(treeConfFilepath string) error {

	var err error
	var treeConfFile *os.File

	// open and read file
	if treeConfFile, err = os.Open(treeConfFilepath); err != nil {
		return fmt.Errorf("Error opening tree config file(%v), err=%s", treeConfFilepath, err)
	}
	defer treeConfFile.Close()

	// decode json
	// if err = json.NewDecoder(treeConfFile).Decode(&api.decisionTree); err != nil {
	// 	log.Fatalf("Error decoding tree config file: %s", err)
	// 	return errors.New("Error decoding tree config file")
	// }

	api.decisionTree = tree.DecisionTree{}

	return nil
}

func (api *UserAPI) HandleUser(userMsg *UserMessage) error {

	LogInfo("function=UserAPI::HandleUser, message=Handling user: %s", userMsg)

	return nil
}

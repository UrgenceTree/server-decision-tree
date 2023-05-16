package user_api

import (
	"fmt"
	"os"
	"service/tree"
	"sync"
	"time"
)

type UserMessage struct {
	PhoneNumber string `json:"phone_number"`
	Text        string `json:"response"`
}

type UserAPI struct {
	decisionTree tree.DecisionTree
	wg           sync.WaitGroup
}

func NewUserAPI() *UserAPI {
	return &UserAPI{
		decisionTree: tree.DecisionTree{},
		wg:           sync.WaitGroup{},
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

func (api *UserAPI) Start() {

	api.wg.Add(1)

	go func(wg *sync.WaitGroup) {
		for {
			select {
			case <-time.After(10 * time.Second):
				fmt.Println("User API is running...")
			}
		}
	}(&api.wg)
}

func (api *UserAPI) HandleUser(userMsg *UserMessage) error {

	// do tree

	return nil
}

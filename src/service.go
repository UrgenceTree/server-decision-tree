package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"service/user_api"
)

type serviceConfig struct {
	TreeConfigFile string `json:"tree_config_file"`
	RabbitMQ       struct {
		URI       string `json:"uri"`
		QueueName string `json:"queueName"`
		Port      string `json:"port"`
	} `json:"rabbitmq"`
	UserAPI struct {
		BaseURL   string `json:"baseUrl"`
		Endpoints struct {
			GetUser    string `json:"getUser"`
			CreateUser string `json:"createUser"`
			UpdateUser string `json:"updateUser"`
			DeleteUser string `json:"deleteUser"`
		} `json:"endpoints"`
	} `json:"userAPI"`
}

type Service struct {
	uAPI   *user_api.UserAPI
	Config serviceConfig
}

func NewService() *Service {
	return &Service{
		uAPI:   user_api.NewUserAPI(),
		Config: serviceConfig{},
	}
}

func LoadConfig(configFilePath string) (*serviceConfig, error) {
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config serviceConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (s *Service) LoadConfig(confFilepath string) error {

	file, err := ioutil.ReadFile(confFilepath)
	if err != nil {
		return err
	}

	var config serviceConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	if err := s.uAPI.LoadTree(s.Config.TreeConfigFile); err != nil {
		return fmt.Errorf("Error loading tree: %s", err)
	}

	fmt.Printf("serviceConfig : %+v\n", s.Config)

	return nil
}
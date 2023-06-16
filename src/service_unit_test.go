package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {

	service := NewService()

	assert.NotNil(t, service)
	assert.IsType(t, &Service{}, service)
	assert.IsType(t, &UserAPI{}, service.uAPI)
	assert.IsType(t, serviceConfig{}, service.Config)
}

func createTempConfigFile(t *testing.T) string {

	config := serviceConfig{
		TreeConfigFile: "path/to/tree_config.json",
		RabbitMQ: struct {
			URI       string `json:"uri"`
			QueueName string `json:"queueName"`
			Port      string `json:"port"`
		}{
			URI:       "amqp://guest:guest@localhost",
			QueueName: "myQueue",
			Port:      "5672",
		},
	}

	// Create a temporary file to hold the configuration
	tempFile, err := ioutil.TempFile("", "config*.json")
	if err != nil {
		t.Fatalf("Error creating temporary file: %s", err)
	}
	defer tempFile.Close()

	// Write the configuration to the file
	configData, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Error marshaling config to JSON: %s", err)
	}

	_, err = tempFile.Write(configData)
	if err != nil {
		t.Fatalf("Error writing config to file: %s", err)
	}

	return tempFile.Name()
}

func removeTempFile(tempFile string) {

	if tempFile != "" {
		_ = os.Remove(tempFile)
	}
}

func TestService_LoadConfig(t *testing.T) {

	// Create a test instance of Service
	s := &Service{
		uAPI:   NewUserAPI(),
		Config: serviceConfig{},
	}

	// Load the configuration
	err := s.LoadConfig("service_conf.json")
	if err != nil {
		t.Fatalf("Error loading config: %s", err)
	}

	// Add assertions to verify the behavior of LoadConfig method
	// ...

	assert.Nil(t, err)
	assert.Equal(t, s.Config.TreeConfigFile, "tree/tree_conf.json")
	assert.Equal(t, s.Config.RabbitMQ.URI, "amqp://guest:guest@localhost:5672/")
	assert.Equal(t, s.Config.RabbitMQ.QueueName, "user_queue")
	assert.Equal(t, s.Config.RabbitMQ.Port, "5672")
	assert.Equal(t, s.Config.UserAPI.BaseURL, "http://localhost:8080/api")
	assert.Equal(t, s.Config.UserAPI.Endpoints.GetUser, "/users/{id}")
	assert.Equal(t, s.Config.UserAPI.Endpoints.CreateUser, "/users")
	assert.Equal(t, s.Config.UserAPI.Endpoints.UpdateUser, "/users/{id}")
	assert.Equal(t, s.Config.UserAPI.Endpoints.DeleteUser, "/users/{id}")

}

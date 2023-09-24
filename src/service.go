package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/streadway/amqp"
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
	uAPI *UserAPI
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
	msgs <-chan amqp.Delivery

	wg   sync.WaitGroup
	done chan struct{}

	Config serviceConfig
}

func NewService() *Service {

	return &Service{
		uAPI: NewUserAPI(),
		conn: nil,
		ch:   nil,
		q:    amqp.Queue{},
		msgs: nil,

		wg:   sync.WaitGroup{},
		done: make(chan struct{}),

		Config: serviceConfig{},
	}
}

func (s *Service) Start() {

	LogInfo("function=Service::Start, message=Starting service...")

	// ctrl+c handler
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		LogInfo("function=Service::Start, message=Received signal: %s. Stopping gently...", sig)
		s.Stop()
	}()

	if err := s.rabbitMQInit(); err != nil {
		LogFatal("function=Service::Start, error=%v", err.Error())
		return
	}

	s.wg.Add(1)
	go s.rolling()

	LogSuccess("function=Service::Start, message=Service started")

}

func (s *Service) Stop() {

	s.ch.Close()
	s.conn.Close()

	close(s.done)
	s.wg.Wait()
}

func (s *Service) Wait() {

	s.wg.Wait()
}

func (s *Service) LoadConfig(confFilepath string) error {

	LogInfo("function=Service::LoadConfig, message=Loading config file: %s", confFilepath)

	file, err := ioutil.ReadFile(confFilepath)
	if err != nil {
		return err
	}

	var config serviceConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	s.Config = config

	if err := s.uAPI.LoadTree(s.Config.TreeConfigFile); err != nil {
		return err
	}

	prettyPrint, err := json.MarshalIndent(s.Config, "", "    ")
	if err != nil {
		return err
	}

	LogSuccess("function=Service::LoadConfig, message=Config file loaded: %+v", string(prettyPrint))

	return nil
}

func (s *Service) rabbitMQInit() error {

	LogInfo("function=Service::rabbitMQInit, message=Initializing RabbitMQ...")

	var err error

	LogInfo("function=Service::rabbitMQInit, message=Connecting to RabbitMQ...")
	if s.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/"); err != nil {
		LogFatal("function=Service::connectToRabbitMQ, error=%v", err.Error())
		return err
	}
	LogSuccess("function=Service::rabbitMQInit, message=Connected to RabbitMQ")

	LogInfo("function=Service::rabbitMQInit, message=Creating channel...")
	if s.ch, err = s.conn.Channel(); err != nil {
		LogFatal("function=Service::createChannel, error=%v", err.Error())
		return err
	}
	LogSuccess("function=Service::rabbitMQInit, message=Channel created")

	LogInfo("function=Service::rabbitMQInit, message=Declaring queue...")
	s.q, err = s.ch.QueueDeclare(
		"decision_tree", // name
		false,           // durable?
		false,           // delete when unused?
		false,           // exclusive?
		false,           // no-wait?
		nil,             // arguments
	)
	if err != nil {
		LogError("function=Service::declareQueue, error=%v", err.Error())
		return err
	}
	LogSuccess("function=Service::rabbitMQInit, message=Queue declared")

	return nil
}

func (s *Service) rolling() error {

	defer s.wg.Done()

	var err error

	LogInfo("function=Service::rabbitMQInit, message=Binding queue...")
	s.msgs, err = s.ch.Consume(
		s.q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		LogError("function=Service::bindQueue, error=%v", err.Error())
		return err
	}

	LogSuccess("function=Service::rabbitMQInit, message=Queue bound")

	for {
		select {

		case <-s.done:
			LogInfo("function=Service::rolling, message=Shutting down...")
			return nil

		case msg := <-s.msgs:
			LogInfo("function=Service::rolling, message=Received message: %s", msg.Body)
			//s.uAPI.HandleUser(msg.Body)

		case <-time.After(5 * time.Second):
			LogInfo("function=Service::rolling, message=Service decision tree rolling...")

		}
	}
}

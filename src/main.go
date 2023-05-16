package main

import "log"

func main() {

	var s *Service = NewService()

	if err := s.LoadConfig("service_conf.json"); err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
}

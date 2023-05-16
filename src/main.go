package main

func main() {

	var s *Service = NewService()

	if err := s.LoadConfig("service_conf.json"); err != nil {
		LogFatal("function=main, message=Error loading config: %s", err)
	}
}

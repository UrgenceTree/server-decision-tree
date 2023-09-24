package main

func init() {

	InitLogger()
}

func main() {

	AddLogger(NewFileLogger("decision_tree.log"))

	var s *Service = NewService()

	if err := s.LoadConfig("service_conf.json"); err != nil {
		LogFatal("function=main, message=Error loading config: %s", err)
		return
	}

	s.Start()
	s.Wait()
}

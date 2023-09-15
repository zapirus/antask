package app

func (s *APIServer) confRouter() {
	s.router.HandleFunc("/analitycs", s.handler.MonitoringChangeHandler)
}

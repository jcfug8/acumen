package server

type Server struct {
	addr     string
	services []Service
}

func NewServer(addr string, services ...Service) *Server {
	return &Server{
		addr:     addr,
		services: services,
	}
}

func (s *Server) Addr() string {
	return ""
}

func (s *Server) Start() error {
	return nil
}

func (s *Server) Stop() error {
	return nil
}

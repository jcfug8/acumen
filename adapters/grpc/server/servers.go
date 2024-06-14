package server

type Servers []*Server

func NewServers(services ...Service) Servers {
	servers := make([]*Server, 0, len(services))
	serversByAddr := make(map[string]*Server)
	for _, s := range services {
		if server, ok := serversByAddr[s.Addr()]; ok {
			server.services = append(server.services, s)
		} else {
			server = &Server{
				addr:     s.Addr(),
				services: []Service{s},
			}
			servers = append(servers, server)
			serversByAddr[s.Addr()] = server
		}
	}
	return servers
}

func (s Servers) Start() error {
	for _, server := range s {
		if err := server.Start(); err != nil {
			return err
		}
	}
	return nil
}

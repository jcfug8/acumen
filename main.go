package main

import (
	"gpt_micro_test/adapters/grpc/server"
	"gpt_micro_test/domain"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	grpcResource "gpt_micro_test/adapters/grpc/resource"
	grpcUser "gpt_micro_test/adapters/grpc/user"
)

func main() {
	domain := domain.NewDomain()

	usersService := grpcUser.NewService(domain)
	resourceService := grpcResource.NewService(domain)

	servers := server.NewServers(usersService, resourceService)

	wg := &sync.WaitGroup{}
	for _, s := range servers {
		wg.Add(1)
		go func(s *server.Server) {
			if err := s.Start(); err != nil {
				log.Printf("server %s start return error: %v", s.Addr(), err)
			}
		}(s)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	for _, server := range servers {
		if err := server.Stop(); err != nil {
			log.Printf("server %s stop return error: %v", server.Addr(), err)
		}
	}
}
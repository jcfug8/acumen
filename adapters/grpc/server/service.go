package server

import "google.golang.org/grpc"

type Service interface {
	Name() string
	Stop() error
	Register(server *grpc.Server) error
	Start() error
	Addr() string
}

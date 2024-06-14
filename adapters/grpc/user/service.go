package user

import (
	"github.com/jcfug8/acumen/ports/user"

	"google.golang.org/grpc"
)

type Service struct {
	domain user.Domain
}

func NewService(domain user.Domain) *Service {
	return &Service{
		domain: domain,
	}
}

func (*Service) Name() string {
	return ""
}

func (*Service) Stop() error {
	return nil
}

func (*Service) Register(_ *grpc.Server) error {
	return nil
}

func (*Service) Start() error {
	return nil
}

func (*Service) Addr() string {
	return ""
}

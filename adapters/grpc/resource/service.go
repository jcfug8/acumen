package resource

import (
	"github.com/jcfug8/acumen/ports/resource"

	"google.golang.org/grpc"
)

type Service struct {
	domain resource.Domain
}

func NewService(domain resource.Domain) *Service {
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

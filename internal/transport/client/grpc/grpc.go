package grpc

import (
	v1 "github.com/teachme-group/user/pkg/api/grpc/v1"
	"google.golang.org/grpc"
)

type transport struct {
	service service
	v1.UnimplementedSignUpServiceServer
	v1.UnimplementedSignInServiceServer
}

func New(service service) *transport {
	return &transport{
		service: service,
	}
}

func (t transport) RegisterServer(server *grpc.Server) {
	v1.RegisterSignUpServiceServer(server, t)
	v1.RegisterSignInServiceServer(server, t)
}

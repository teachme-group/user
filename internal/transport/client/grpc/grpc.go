package grpc

import (
	"context"

	"github.com/Markuysa/pkg/tracer"
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

func (t transport) SignInInit(ctx context.Context, request *v1.SignInInitRequest) (*v1.SignInInitResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return nil, nil
}

func (t transport) SignInFinalize(ctx context.Context, request *v1.SignInFinalizeRequest) (*v1.SignInFinalizeResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return nil, nil
}

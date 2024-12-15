package grpc

import (
	"context"

	"github.com/Markuysa/pkg/tracer"
	v1 "github.com/teachme-group/user/pkg/api/grpc/v1"
)

func (t transport) SignInInit(ctx context.Context, request *v1.SignInInitRequest) (*v1.SignInInitResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return t.service.SignInInit(ctx, request)
}

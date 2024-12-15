package grpc

import (
	"context"

	"github.com/Markuysa/pkg/tracer"
	"github.com/google/uuid"
	v1 "github.com/teachme-group/user/pkg/api/grpc/v1"
)

func (t transport) GetOauthSignUpUrls(
	ctx context.Context,
	request *v1.GetOauthSignUpUrlRequest,
) (*v1.GetOauthSignUpUrlResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return t.service.GetOauthSignUpUrls(ctx, request)
}

func (t transport) GoogleSignUpCallback(
	ctx context.Context,
	request *v1.HandleOauthCallbackRequest,
) (*v1.HandleOauthCallbackResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return t.service.GoogleSignUpCallback(ctx, request)
}

func (t transport) SignUpInit(ctx context.Context, request *v1.SignUpInitRequest) (*v1.SignUpInitResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	request.SignUpToken = uuid.NewString()

	return t.service.SignUpRequest(ctx, request)
}

func (t transport) SignUpFinalize(ctx context.Context, request *v1.SignUpFinalizeRequest) (*v1.SignUpFinalizeResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return t.service.SignUpConfirm(ctx, request)
}

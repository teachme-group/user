package grpc

import (
	"context"

	v1 "github.com/teachme-group/user/pkg/api/grpc/v1"
)

type (
	service interface {
		signUpService
	}
	signUpService interface {
		GetOauthSignUpUrls(
			ctx context.Context,
			in *v1.GetOauthSignUpUrlRequest,
		) (*v1.GetOauthSignUpUrlResponse, error)
		GoogleSignUpCallback(
			ctx context.Context,
			request *v1.HandleOauthCallbackRequest,
		) (response *v1.HandleOauthCallbackResponse, err error)
		SignUpConfirm(
			ctx context.Context,
			request *v1.SignUpFinalizeRequest,
		) (response *v1.SignUpFinalizeResponse, err error)
		SignUpRequest(
			ctx context.Context,
			request *v1.SignUpInitRequest,
		) (response *v1.SignUpInitResponse, err error)
	}
)

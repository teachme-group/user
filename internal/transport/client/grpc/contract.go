package grpc

import (
	"context"

	v1 "github.com/teachme-group/user/pkg/api/grpc/v1"
)

type (
	service interface {
		signUpService
		signInService
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
		SignUpRequest(
			ctx context.Context,
			request *v1.SignUpInitRequest,
		) (response *v1.SignUpInitResponse, err error)
		SignUpConfirmEmail(
			ctx context.Context,
			request *v1.SignUpConfirmEmailRequest,
		) (response *v1.SignUpConfirmEmailResponse, err error)
		SignUpEnterPassword(
			ctx context.Context,
			request *v1.SignUpEnterPasswordRequest,
		) (response *v1.SignUpEnterPasswordResponse, err error)
	}

	signInService interface {
		SignInInit(ctx context.Context, request *v1.SignInInitRequest) (*v1.SignInInitResponse, error)
	}
)

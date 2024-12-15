package client

import (
	"context"

	"github.com/Markuysa/pkg/tracer"
	"github.com/google/uuid"
	sessionStorageV1 "github.com/teachme-group/session/pkg/api/grpc/v1"
	v1 "github.com/teachme-group/user/pkg/api/grpc/v1"
	oauthCli "github.com/teachme-group/user/pkg/oauth"
)

func (s *service) GetOauthSignUpUrls(
	ctx context.Context,
	in *v1.GetOauthSignUpUrlRequest,
) (*v1.GetOauthSignUpUrlResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	urls, err := s.oauth.AuthCodeURLs(uuid.NewString(), in.OauthProvider)
	if err != nil {
		return nil, err
	}

	return &v1.GetOauthSignUpUrlResponse{
		Urls: urls,
	}, nil
}

func (s *service) GoogleSignUpCallback(
	ctx context.Context,
	request *v1.HandleOauthCallbackRequest,
) (response *v1.HandleOauthCallbackResponse, err error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	userInfo, err := s.oauth.ProcessCallback(
		ctx,
		oauthCli.Google,
		request.GetBody(),
		request.GetCallbackUrl(),
	)
	if err != nil {
		return response, err
	}

	if err = s.repos.ValidateUserSignUp(ctx, userInfo.Email); err != nil {
		return response, err
	}

	user, err := s.repos.CreateUser(ctx, userInfo)
	if err != nil {
		return response, err
	}

	resp, err := s.sessionStorage.ClientSetSession(ctx, &sessionStorageV1.ClientSetSessionRequest{
		ClientId: user.ID.String(),
	})
	if err != nil {
		return response, err
	}

	return &v1.HandleOauthCallbackResponse{
		SessionId: resp.AccessToken,
	}, nil
}

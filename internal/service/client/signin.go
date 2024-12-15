package client

import (
	"context"

	"github.com/Markuysa/pkg/tracer"
	sessionStorageV1 "github.com/teachme-group/session/pkg/api/grpc/v1"
	v1 "github.com/teachme-group/user/pkg/api/grpc/v1"
	"github.com/teachme-group/user/pkg/errlist"
	"github.com/teachme-group/user/pkg/signer"
)

func (s *service) SignInInit(ctx context.Context, request *v1.SignInInitRequest) (*v1.SignInInitResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	user, err := s.repos.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if err = signer.ComparePasswords(user.Password, request.Password); err != nil {
		return nil, errlist.ErrInvalidLoginCredentials
	}

	resp, err := s.sessionStorage.ClientSetSession(ctx, &sessionStorageV1.ClientSetSessionRequest{
		ClientId: user.ID.String(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.SignInInitResponse{
		SessionId: resp.AccessToken,
	}, err
}

package client

import (
	"context"
	"time"

	v1 "github.com/teachme-group/session/pkg/api/grpc/v1"
	"github.com/teachme-group/user/internal/domain"
	oauthCli "github.com/teachme-group/user/pkg/oauth"
	"golang.org/x/oauth2"
)

type (
	repository interface {
		CreateUser(ctx context.Context, user domain.User) (domain.User, error)
		SaveSignUpStep(ctx context.Context, key string, step domain.SignUpStep, ttl time.Duration) error
		GetSignUpStep(ctx context.Context, key string) (domain.SignUpStep, error)
		ValidateUserSignUp(ctx context.Context, email string) error
	}
	mailer interface {
		Send(ctx context.Context, to, subject, body string) error
	}
	sessionStorage interface {
		ClientSetSession(ctx context.Context, in *v1.ClientSetSessionRequest) (*v1.ClientSetSessionResponse, error)
	}
	oauthClient interface {
		AuthCodeURLs(state string, provider *string, opts ...oauth2.AuthCodeOption) (map[string]string, error)
		ProcessCallback(ctx context.Context, provider oauthCli.Provider, body []byte, url string) (domain.User, error)
	}
)

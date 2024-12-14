package client

import (
	"context"
	v1 "session/pkg/api/grpc/v1"
	"time"

	"github.com/teachme-group/user/internal/domain"
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
		ClientSetSession(ctx context.Context, req *v1.ClientSetSessionRequest) (*v1.ClientSetSessionResponse, error)
	}
)

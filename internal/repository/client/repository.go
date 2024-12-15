package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	redisCli "github.com/redis/go-redis/v9"
	"github.com/teachme-group/user/internal/domain"
	"github.com/teachme-group/user/internal/storage/postgres"
	"github.com/teachme-group/user/internal/storage/redis"
	"github.com/teachme-group/user/pkg/errlist"
)

type repository struct {
	pgConn postgresStorage
	redis  redisStorage
}

func New(
	pool *pgxpool.Pool,
	redisClient *redisCli.Client,
) *repository {
	return &repository{
		pgConn: postgres.NewStorage(pool),
		redis:  redis.NewStorage(redisClient),
	}
}

func (r *repository) CreateUser(ctx context.Context, user domain.User) (created domain.User, err error) {
	pgUser, err := r.pgConn.Queries(ctx).CreateUser(ctx, postgres.CreateUserParams{
		Login: user.Login,
		Email: user.Email,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return created, fmt.Errorf("failed to create user: %w", err)
	}

	return userFromRepository(pgUser), nil
}

func (r *repository) SaveSignUpStep(ctx context.Context, key string, step domain.SignUpStep, ttl time.Duration) error {
	return r.redis.Save(ctx, key, step, ttl)
}

func (r *repository) GetSignUpStep(ctx context.Context, key string) (domain.SignUpStep, error) {
	result := domain.SignUpStep{}

	resp, err := r.redis.Get(ctx, key)
	if err != nil {
		return result, fmt.Errorf("failed to get sign up step: %w", err)
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal sign up step: %w", err)
	}

	return result, nil
}

func (r *repository) ValidateUserSignUp(ctx context.Context, email string) error {
	result, err := r.pgConn.Queries(ctx).ValidateUserSignUp(ctx, postgres.ValidateUserSignUpParams{
		Email: email,
	})
	if err != nil {
		return fmt.Errorf("failed to validate user sign up: %w", err)
	}

	if result {
		return errlist.ErrLoginOrEmailAlreadyExists
	}

	return nil
}

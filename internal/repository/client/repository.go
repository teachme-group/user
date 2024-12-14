package client

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	redisCli "github.com/redis/go-redis/v9"
	"gitlab.com/coinhubs/balance/internal/domain"
	"gitlab.com/coinhubs/balance/internal/storage/postgres"
	"gitlab.com/coinhubs/balance/internal/storage/redis"
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

func (r *repository) SaveEmailConfirmationCode(
	ctx context.Context,
	key string,
	code int,
) error {
	return r.redis.Save(ctx, key, code, time.Hour)
}

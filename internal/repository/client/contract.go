package client

import (
	"context"
	"time"

	"gitlab.com/coinhubs/balance/internal/storage/postgres"
)

type (
	redisStorage interface {
		Save(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	}
	postgresStorage interface {
		Queries(ctx context.Context) *postgres.Queries
	}
)

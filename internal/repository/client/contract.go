package client

import (
	"context"
	"time"

	"github.com/teachme-group/user/internal/storage/postgres"
)

type (
	redisStorage interface {
		Save(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	}
	postgresStorage interface {
		Queries(ctx context.Context) *postgres.Queries
	}
)

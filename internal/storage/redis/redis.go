package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisStorage struct {
	rd *redis.Client
}

func NewStorage(rd *redis.Client) *redisStorage {
	return &redisStorage{rd: rd}
}

func (r *redisStorage) Save(
	ctx context.Context,
	key string,
	value interface{},
	ttl time.Duration,
) error {
	return r.rd.Set(ctx, key, value, ttl).Err()
}

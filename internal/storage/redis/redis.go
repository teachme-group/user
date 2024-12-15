package redis

import (
	"context"
	"encoding/json"
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
	resp, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, resp, ttl).Err()
}

func (r *redisStorage) Get(
	ctx context.Context,
	key string,
) ([]byte, error) {
	resp, err := r.rd.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

package connection

import (
	"context"
	"time"

	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/redis/go-redis/v9"
)

// Cache is
var _ Cache = (*Redis)(nil)

// Redis is
type Redis struct {
	redis *redis.Client
}

// NewCache is
func NewCache(ctx context.Context, cfgs config.RedisDB) (*Redis, error) {
	options := &redis.Options{
		Addr:     cfgs.Address,
		Password: cfgs.Password,
		DB:       0, // use Default DB
	}

	rdb := redis.NewClient(options)
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, errlst.NewBadRequestError("error in rdb.Ping")
	}

	return &Redis{redis: rdb}, nil
}

// Cache interface is
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Del(ctx context.Context, key string) error
	Close() error
}

// Get is
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

// Set is
func (r *Redis) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.redis.Set(ctx, key, value, expiration).Err()
}

// Del is
func (r *Redis) Del(ctx context.Context, key string) error {
	return r.redis.Del(ctx, key).Err()
}

// Close is
func (r *Redis) Close() error {
	return r.redis.Close()
}

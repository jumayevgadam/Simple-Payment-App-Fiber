package connection

import (
	"context"
	"time"

	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/redis/go-redis/v9"
)

// Ensure Redis struct implements the Cache interface.
var _ Cache = (*Redis)(nil)

// Redis struct for using redis methods in application.
type Redis struct {
	redis *redis.Client
}

// NewCache creates and returns a new instance Redis struct.
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

// Cache interface for performing actions with redis.
type Cache interface {
	CacheRepository
}

// CacheRepository interface for managing redis repository.
type CacheRepository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Del(ctx context.Context, key string) error
	Close() error
}

// Get method fetches needed value using key.
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

// Set method insert a new key value pair in redisDB.
func (r *Redis) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.redis.Set(ctx, key, value, expiration).Err()
}

// Del method deletes the key value pair using identified key.
func (r *Redis) Del(ctx context.Context, key string) error {
	return r.redis.Del(ctx, key).Err()
}

// Close closes redisDB.
func (r *Redis) Close() error {
	return r.redis.Close()
}

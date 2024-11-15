package cache

import (
	"context"
	"strings"
	"time"

	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	"github.com/jumayevgadaym/tsu-toleg/internal/models/abstract"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// Ensure ClientRedisRepo implements the Store interface.
var _ Store = (*ClientRedisRepo)(nil)

// Store interface for using redisDB methods in app.
type Store interface {
	Get(ctx context.Context, argument abstract.CacheArgument) ([]byte, error)
	Set(ctx context.Context, argument abstract.CacheArgument, value []byte, duration time.Duration) error
	Del(ctx context.Context, argument abstract.CacheArgument) error
}

// ClientRedisRepo takes connection.Cache interface methods for performing decorator pattern in app.
type ClientRedisRepo struct {
	rdb connection.Cache
}

// NewClientRDRepository creates and returns a new instance of ClientRedisRepo.
func NewClientRDRepository(rdb connection.Cache) *ClientRedisRepo {
	return &ClientRedisRepo{rdb: rdb}
}

// getCacheKey from redisDB.
func (c *ClientRedisRepo) getCacheKey(objectType string, id string) string {
	return strings.Join([]string{
		objectType,
		id,
	}, ":")
}

// Get method takes needed value using cacheArgument.
func (c *ClientRedisRepo) Get(ctx context.Context, argument abstract.CacheArgument) ([]byte, error) {
	ctx, span := otel.Tracer("[ClientRedisRepo]").Start(ctx, "[Get]")
	defer span.End()

	key := argument.ToCacheStorage()
	cacheKey := c.getCacheKey(key.ObjectType, key.ID)
	valueString, err := c.rdb.Get(ctx, cacheKey)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return nil, errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "Successfully get data from redis repo")
	return []byte(valueString), nil
}

// Set method creates and insert a new key value pair into redisDB.
func (c *ClientRedisRepo) Set(ctx context.Context, argument abstract.CacheArgument, value []byte, duration time.Duration) error {
	ctx, span := otel.Tracer("[ClientRedisRepo]").Start(ctx, "[Set]")
	defer span.End()

	key := argument.ToCacheStorage()
	cacheKey := c.getCacheKey(key.ObjectType, key.ID)

	err := c.rdb.Set(ctx, cacheKey, string(value), duration)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "Successfully set data from redis repo")
	return nil
}

// Del method deletes cached value using cacheArgument from redisDB.
func (c *ClientRedisRepo) Del(ctx context.Context, argument abstract.CacheArgument) error {
	ctx, span := otel.Tracer("[ClientRedisRepo]").Start(ctx, "[Del]")
	defer span.End()

	key := argument.ToCacheStorage()
	cacheKey := c.getCacheKey(key.ObjectType, key.ID)

	err := c.rdb.Del(ctx, cacheKey)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "Successfully deleted data from redis repo")
	return nil
}

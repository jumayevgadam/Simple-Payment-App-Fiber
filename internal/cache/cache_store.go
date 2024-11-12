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

var _ Store = (*ClientRedisRepo)(nil)

// Store interface is
type Store interface {
	Get(ctx context.Context, argument abstract.CacheArgument) ([]byte, error)
	Set(ctx context.Context, argument abstract.CacheArgument, value []byte, duration time.Duration) error
	Del(ctx context.Context, argument abstract.CacheArgument) error
}

// ClientRedisRepo is
type ClientRedisRepo struct {
	rdb connection.Cache
}

// NewClientRDRepository is
func NewClientRDRepository(rdb connection.Cache) *ClientRedisRepo {
	return &ClientRedisRepo{rdb: rdb}
}

// Get is
func (c *ClientRedisRepo) getCacheKey(objectType string, id string) string {
	return strings.Join([]string{
		objectType,
		id,
	}, ":")
}

// Get is
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

// Set is
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

// Del is
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

package cache

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	"github.com/jumayevgadam/tsu-toleg/internal/models/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// Ensure ClientRedisRepo implements the Store interface.
var _ Store = (*ClientRedisRepo)(nil)

// Store interface for using redisDB methods in app.
type Store interface {
	CacheStore
	SessionStore
}

type CacheStore interface {
	Get(ctx context.Context, argument abstract.CacheArgument) ([]byte, error)
	Set(ctx context.Context, argument abstract.CacheArgument, value []byte, duration time.Duration) error
	Del(ctx context.Context, argument abstract.CacheArgument) error
}

// SessionStore interface handling operations related with refresh token.
type SessionStore interface {
	GetSession(ctx context.Context, params abstract.SessionArgument) ([]byte, error)
	PutSession(ctx context.Context, params abstract.SessionArgument) error
	DelSession(ctx context.Context, params abstract.SessionArgument) error
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
func (c *ClientRedisRepo) getCacheKey(objectType, id string) string {
	return strings.Join([]string{
		objectType,
		id,
	}, ":")
}

// getSessionKey from redisDB.
func (c *ClientRedisRepo) getSessionKey(sessionPrefix, userID string) string {
	return strings.Join([]string{
		sessionPrefix,
		userID,
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

// GetSession method receives refresh token using userId.
func (c *ClientRedisRepo) GetSession(ctx context.Context, params abstract.SessionArgument) ([]byte, error) {
	ctx, span := otel.Tracer("[ClientRedisRepo]").Start(ctx, "[GetSession]")
	defer span.End()

	key := params.ToSessionStorage()
	sessionKey := c.getSessionKey(key.SessionPrefix, key.UserID) // SessionKey is refresh token here.

	valueString, err := c.rdb.Get(ctx, sessionKey)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return nil, errlst.NewBadRequestError("error getting value string in this place") // error in this place....
	}

	span.SetStatus(codes.Ok, "successfully get session from redisDB")
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

// PutSession method insert and creates a session in redisDB.
func (c *ClientRedisRepo) PutSession(ctx context.Context, params abstract.SessionArgument) error {
	ctx, span := otel.Tracer("[ClientRedisRepo]").Start(ctx, "[PutSession]")
	defer span.End()

	key := params.ToSessionStorage()
	sessionKey := c.getSessionKey(key.SessionPrefix, key.UserID)

	sessBytes, err := json.Marshal(params)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	err = c.rdb.Set(ctx, sessionKey, string(sessBytes), params.ExpiresAt)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "Successfully created session")
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

// DelSession method removes session from redisDB.
func (c *ClientRedisRepo) DelSession(ctx context.Context, params abstract.SessionArgument) error {
	ctx, span := otel.Tracer("[ClientRedisRepo]").Start(ctx, "[DelSession]")
	defer span.End()

	key := params.ToSessionStorage()
	sessionKey := c.getSessionKey(key.SessionPrefix, key.UserID)

	err := c.rdb.Del(ctx, sessionKey)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "Successfully deleted session")
	return nil
}

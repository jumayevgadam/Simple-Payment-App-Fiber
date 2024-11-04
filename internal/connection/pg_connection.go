package connection

import (
	"context"
	"fmt"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
)

// DBOps is
var _ DB = (*Database)(nil)

// Querier is
type Querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

var (
	_ Querier = &pgxpool.Pool{}
	_ Querier = &pgxpool.Conn{}
)

// DB interface for general database operations
type DB interface {
	Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

// DBOps interface with Transaction
type DBOps interface {
	DB
	Begin(ctx context.Context, txOps pgx.TxOptions) (TxOps, error)
	Close() error
}

// Database struct is
type Database struct {
	Db *pgxpool.Pool
}

// GetDBConnection is
func GetDBConnection(ctx context.Context, cfg config.PostgresDB) (*Database, error) {
	db, err := pgxpool.New(ctx, fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SslMode,
	))
	if err != nil {
		return nil, fmt.Errorf("connection[pgxpool.New]: %w", err)
	}

	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("connection[Ping]: %w", err)
	}

	return &Database{Db: db}, nil
}

// Get method implements DB interface
func (d *Database) Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db, dest, query, args...)
}

// Select method implements DB interface
func (d *Database) Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db, dest, query, args...)
}

// QueryRow method implements DB interface
func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.Db.QueryRow(ctx, query, args...)
}

// Query method implements DB interface
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.Db.Query(ctx, query, args...)
}

// Exec method implements DB interface
func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.Db.Exec(ctx, query, args...)
}

// Begin starts a new transaction
func (d *Database) Begin(ctx context.Context, txOpts pgx.TxOptions) (TxOps, error) {
	if d == nil {
		return nil, fmt.Errorf("cannot start transaction")
	}

	c, err := d.Db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := d.Db.BeginTx(ctx, txOpts)
	if err != nil {
		c.Release()
		log.Printf("Failed to begin transaction: %v", err)
		return nil, fmt.Errorf("connection.Database.Begin: %w", err)
	}
	return &Transaction{Tx: tx}, nil
}

// Close closes the database connection pool
func (d *Database) Close() error {
	d.Db.Close()
	return nil
}

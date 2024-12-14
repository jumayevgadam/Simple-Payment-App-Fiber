package connection

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// Ensure Database struct implements the DB interface.
var _ DB = (*Database)(nil)

// Querier interface for using pgxscany.
type Querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

var (
	_ Querier = &pgxpool.Pool{}
	_ Querier = &pgxpool.Conn{}
)

// DB interface for general database operations.
type DB interface {
	Querier
	Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
}

// DBOps interface with Transaction.
type DBOps interface {
	DB
	Begin(ctx context.Context, txOps pgx.TxOptions) (TxOps, error)
	Close()
}

// Database struct performs database logic using pgxpool.Pool.
type Database struct {
	Db *pgxpool.Pool
}

// GetDBConnectionWithRetry attempts to connect to the database with retries and fails immediately after 3 attempts.
func GetDBConnection(ctx context.Context, cfg config.PostgresDB) (*Database, error) {
	const (
		retryAttempts = 3
		retryDelay    = 2 * time.Second
	)

	var db *pgxpool.Pool
	var err error

	// Retry connection attempts
	for i := 0; i < retryAttempts; i++ {
		db, err = connectToDB(ctx, cfg)
		if err == nil {
			// Successful connection.
			return &Database{Db: db}, nil
		}

		// Log and retry on failure
		fmt.Printf("Attempt %d/%d failed: %v. Retrying in %s...\n", i+1, retryAttempts, err, retryDelay)
		time.Sleep(retryDelay)
	}

	// After 3 failed attempts, log the error and immediately fatal the application.
	log.Fatalf("Failed to connect to database after %d attempts: %v. Exiting...\n", retryAttempts, err)

	// Return error after fatal
	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", retryAttempts, err)
}

// connectToDB handles a single attempt to connect to the database.
func connectToDB(ctx context.Context, cfg config.PostgresDB) (*pgxpool.Pool, error) {
	hostPort := net.JoinHostPort(cfg.Host, cfg.Port)
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		hostPort,
		cfg.Name,
		cfg.SslMode,
	)

	// Parse the connection string to create a pgxpool configuration
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("parsing connection config: %w", err)
	}

	// Configure the connection pool settings
	config.MaxConns = 200 // Max number of connections
	config.MinConns = 10  // Min number of connections

	// Create a new connection pool
	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	// Ping the database to verify the connection
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return db, nil
}

// Get method implements DB interface.
func (d *Database) Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db, dest, query, args...)
}

// Select method implements DB interface.
func (d *Database) Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db, dest, query, args...)
}

// QueryRow method implements DB interface.
func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.Db.QueryRow(ctx, query, args...)
}

// Query method implements DB interface.
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.Db.Query(ctx, query, args...)
}

// Exec method implements DB interface.
func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.Db.Exec(ctx, query, args...)
}

// Begin starts a new transaction.
func (d *Database) Begin(ctx context.Context, txOpts pgx.TxOptions) (TxOps, error) {
	if d == nil || d.Db == nil {
		return nil, errlst.ErrBeginTransaction
	}

	tx, err := d.Db.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, fmt.Errorf("connection.Database.Begin: %w", errlst.ErrBeginTransaction)
	}

	return &Transaction{Tx: tx}, nil
}

// Close closes the database connection pool.
func (d *Database) Close() {
	d.Db.Close()
}

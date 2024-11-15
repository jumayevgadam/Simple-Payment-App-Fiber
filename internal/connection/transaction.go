package connection

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TxOps interface for transaction operations.
type TxOps interface {
	Commit(ctx context.Context) error
	RollBack(ctx context.Context) error
	DB
}

// Transaction struct for handling transactions.
type Transaction struct {
	Tx   pgx.Tx
	Conn *pgxpool.Conn
}

// Get method implements TxOps interface.
func (tx *Transaction) Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, tx.Tx, dest, query, args...)
}

// Select method implements TxOps interface.
func (tx *Transaction) Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, tx.Tx, dest, query, args...)
}

// QueryRow method implements TxOps interface.
func (tx *Transaction) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return tx.Tx.QueryRow(ctx, query, args...)
}

// Query method implements TxOps interface.
func (tx *Transaction) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return tx.Tx.Query(ctx, query, args...)
}

// Exec executes a query that doesn't return rows.
func (tx *Transaction) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return tx.Tx.Exec(ctx, query, args...)
}

// Commit transaction.
func (tx *Transaction) Commit(ctx context.Context) error {
	return tx.Tx.Commit(ctx)
}

// RollBack Transaction.
func (tx *Transaction) RollBack(ctx context.Context) error {
	return tx.Tx.Rollback(ctx)
}

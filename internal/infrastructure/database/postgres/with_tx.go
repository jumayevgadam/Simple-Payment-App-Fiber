package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

// WithTransaction is
func (d *DataStoreImpl) WithTransaction(ctx context.Context, transactionFn func(db database.DataStore) error) error {
	db, ok := d.db.(connection.DBOps)
	if !ok {
		return fmt.Errorf("got error type assertion in WithTx")
	}

	//begin transaction in this place
	tx, err := db.Begin(ctx, pgx.TxOptions{})
	if err != nil {
		log.Println("error in db.Begin[WithTransaction]")
		return errlst.ParseErrors(err)
	}

	defer func() {
		if err != nil {
			// RollBack transaction if error occured
			if err = tx.RollBack(ctx); err != nil {
				log.Printf("postgres:[WithTransaction]: failed to rollback transaction: %v", err.Error())
			}
			log.Printf("postgres:[WithTransaction: failed in transaction]")
		}
	}()

	// transactionalDB is
	transactionalDB := &DataStoreImpl{db: tx}
	if err := transactionFn(transactionalDB); err != nil {
		return fmt.Errorf("postgres:[WithTransaction]: transaction function execution failed: %w", err)
	}

	// Commit the transaction if no error occurred during the transactionFn execution
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error in committing transaction: %w", err)
	}

	return nil
}

package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// WithTransaction method is a transaction method for performing multitasks, used in the service layer.
func (d *DataStoreImpl) WithTransaction(ctx context.Context, transactionFn func(db database.DataStore) error) error {
	// Assert that the db implements DBOps for transaction capabilities.
	db, ok := d.db.(connection.DBOps)
	if !ok {
		return errlst.ErrTypeAssertInTransaction
	}

	// Begin transaction.
	tx, err := db.Begin(ctx, pgx.TxOptions{})
	if err != nil {
		log.Println("error in db.Begin[WithTransaction]:", err)
		return errlst.ParseErrors(err)
	}

	// Ensure the transaction is rolled back if an error occurs.
	defer func() {
		if err != nil {
			if rbErr := tx.RollBack(ctx); rbErr != nil {
				log.Printf("postgres:[WithTransaction]: failed to rollback transaction: %v", rbErr)
			} else {
				log.Printf("postgres:[WithTransaction]: transaction rolled back due to error: %v", err)
			}
		}
	}()

	// Wrap the database in the transactional context.
	transactionalDB := &DataStoreImpl{db: tx}

	// Run the transaction function.
	err = transactionFn(transactionalDB)
	if err != nil {
		log.Printf("Error during transaction function execution: %v", err)

		return errlst.NewInternalServerError(err.Error())
	}

	// Commit the transaction if no error occurred during execution.
	if commitErr := tx.Commit(ctx); commitErr != nil {
		log.Printf("Error committing transaction: %v", commitErr)
		// Return a wrapped error if commit fails.
		return fmt.Errorf("error in committing transaction: %w", commitErr)
	}

	return nil
}

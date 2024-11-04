package postgres

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	"github.com/jumayevgadaym/tsu-toleg/internal/database"
	"github.com/jumayevgadaym/tsu-toleg/internal/roles"
	roleRepository "github.com/jumayevgadaym/tsu-toleg/internal/roles/repository"
)

var _ database.DataStore = (*DataStoreImpl)(nil)

// DataStore struct is
type DataStoreImpl struct {
	db       connection.DB
	role     roles.Repository
	roleInit sync.Once
}

// NewDataStore is
func NewDataStore(db connection.DBOps) database.DataStore {
	return &DataStoreImpl{
		db: db,
	}
}

// RolesRepo method is
func (d *DataStoreImpl) RolesRepo() roles.Repository {
	d.roleInit.Do(func() {
		d.role = roleRepository.NewRoleRepository(d.db)
	})

	return d.role
}

// WithTransaction is
func (d *DataStoreImpl) WithTransaction(ctx context.Context, transactionFn func(db database.DataStore) error) error {
	db, ok := d.db.(connection.DBOps)
	if !ok {
		return fmt.Errorf("got error type assertion in WithTx")
	}

	//begin transaction in this place
	tx, err := db.Begin(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("error in db.Begin[WithTransaction]: %w", err)
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

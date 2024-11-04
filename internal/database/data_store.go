package database

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/roles"
)

// DataStore is
type DataStore interface {
	WithTransaction(ctx context.Context, tx func(db DataStore) error) error
	RolesRepo() roles.Repository
}

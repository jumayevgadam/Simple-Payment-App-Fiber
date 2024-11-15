package database

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/app/faculties"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/groups"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/payment"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/roles"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/users"
)

// DataStore interface for performing all needed methods for repository layer of application.
type DataStore interface {
	WithTransaction(ctx context.Context, tx func(db DataStore) error) error
	RolesRepo() roles.Repository
	FacultiesRepo() faculties.Repository
	GroupsRepo() groups.Repository
	UsersRepo() users.Repository
	PaymentsRepo() payment.Repository
}

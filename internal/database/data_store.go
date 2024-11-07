package database

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/faculties"
	"github.com/jumayevgadaym/tsu-toleg/internal/groups"
	"github.com/jumayevgadaym/tsu-toleg/internal/payment"
	"github.com/jumayevgadaym/tsu-toleg/internal/roles"
	"github.com/jumayevgadaym/tsu-toleg/internal/users"
)

// DataStore is
type DataStore interface {
	WithTransaction(ctx context.Context, tx func(db DataStore) error) error
	RolesRepo() roles.Repository
	FacultiesRepo() faculties.Repository
	GroupsRepo() groups.Repository
	UsersRepo() users.Repository
	PaymentsRepo() payment.Repository
}

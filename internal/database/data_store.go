package database

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/common/faculties"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/groups"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/payment"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/roles"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/users"
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

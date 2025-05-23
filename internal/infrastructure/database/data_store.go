package database

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
)

// DataStore interface for performing all needed methods for repository layer of application.
type DataStore interface {
	WithTransaction(ctx context.Context, tx func(db DataStore) error) error
	RolesRepo() roles.Repository
	FacultiesRepo() faculties.Repository
	GroupsRepo() groups.Repository
	UsersRepo() users.Repository
	TimesRepo() times.Repository
	PaymentRepo() payments.Repository
}

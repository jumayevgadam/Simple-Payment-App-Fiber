package postgres

import (
	"sync"

	"github.com/jumayevgadaym/tsu-toleg/internal/app/faculties"
	facultyRepository "github.com/jumayevgadaym/tsu-toleg/internal/app/faculties/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/groups"
	groupRepository "github.com/jumayevgadaym/tsu-toleg/internal/app/groups/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/payment"
	paymentRepository "github.com/jumayevgadaym/tsu-toleg/internal/app/payment/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/roles"
	roleRepository "github.com/jumayevgadaym/tsu-toleg/internal/app/roles/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/users"
	userRepository "github.com/jumayevgadaym/tsu-toleg/internal/app/users/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/database"
)

// Ensure DataStoreImpl implements the database.DataStore interface.
var _ database.DataStore = (*DataStoreImpl)(nil)

// DataStore struct for performing all actions which needed in repository layer.
type DataStoreImpl struct {
	db          connection.DB
	role        roles.Repository
	roleInit    sync.Once
	faculty     faculties.Repository
	facultyInit sync.Once
	group       groups.Repository
	groupInit   sync.Once
	user        users.Repository
	userInit    sync.Once
	payment     payment.Repository
	paymentInit sync.Once
}

// NewDataStore creates and returns a new instance of DataStore.
func NewDataStore(db connection.DBOps) database.DataStore {
	return &DataStoreImpl{
		db: db,
	}
}

// RolesRepo method needs performing general repo methods for roles.
func (d *DataStoreImpl) RolesRepo() roles.Repository {
	d.roleInit.Do(func() {
		d.role = roleRepository.NewRoleRepository(d.db)
	})

	return d.role
}

// FacultiesRepo method needs performing general repo methods for faculties.
func (d *DataStoreImpl) FacultiesRepo() faculties.Repository {
	d.facultyInit.Do(func() {
		d.faculty = facultyRepository.NewFacultyRepository(d.db)
	})

	return d.faculty
}

// GroupsRepo method needs performing general repo methods for groups.
func (d *DataStoreImpl) GroupsRepo() groups.Repository {
	d.groupInit.Do(func() {
		d.group = groupRepository.NewGroupRepository(d.db)
	})

	return d.group
}

// UsersRepo method needs performing general repo methods for users.
func (d *DataStoreImpl) UsersRepo() users.Repository {
	d.userInit.Do(func() {
		d.user = userRepository.NewUserRepository(d.db)
	})

	return d.user
}

// PaymentsRepo method needs performing general repo methods for payments.
func (d *DataStoreImpl) PaymentsRepo() payment.Repository {
	d.paymentInit.Do(func() {
		d.payment = paymentRepository.NewPaymentRepository(d.db)
	})

	return d.payment
}

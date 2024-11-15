package postgres

import (
	"sync"

	"github.com/jumayevgadaym/tsu-toleg/internal/common/faculties"
	facultyRepository "github.com/jumayevgadaym/tsu-toleg/internal/common/faculties/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/groups"
	groupRepository "github.com/jumayevgadaym/tsu-toleg/internal/common/groups/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/payment"
	paymentRepository "github.com/jumayevgadaym/tsu-toleg/internal/common/payment/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/roles"
	roleRepository "github.com/jumayevgadaym/tsu-toleg/internal/common/roles/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/users"
	userRepository "github.com/jumayevgadaym/tsu-toleg/internal/common/users/repository"
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	"github.com/jumayevgadaym/tsu-toleg/internal/database"
)

var _ database.DataStore = (*DataStoreImpl)(nil)

// DataStore struct is
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

// FacultiesRepo method is
func (d *DataStoreImpl) FacultiesRepo() faculties.Repository {
	d.facultyInit.Do(func() {
		d.faculty = facultyRepository.NewFacultyRepository(d.db)
	})

	return d.faculty
}

// GroupsRepo method is
func (d *DataStoreImpl) GroupsRepo() groups.Repository {
	d.groupInit.Do(func() {
		d.group = groupRepository.NewGroupRepository(d.db)
	})

	return d.group
}

// UsersRepo method is
func (d *DataStoreImpl) UsersRepo() users.Repository {
	d.userInit.Do(func() {
		d.user = userRepository.NewUserRepository(d.db)
	})

	return d.user
}

// PaymentsRepo method is
func (d *DataStoreImpl) PaymentsRepo() payment.Repository {
	d.paymentInit.Do(func() {
		d.payment = paymentRepository.NewPaymentRepository(d.db)
	})

	return d.payment
}

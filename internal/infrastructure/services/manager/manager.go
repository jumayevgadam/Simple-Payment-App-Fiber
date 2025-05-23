package manager

import (
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	facultyService "github.com/jumayevgadam/tsu-toleg/internal/modules/faculties/service"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	groupService "github.com/jumayevgadam/tsu-toleg/internal/modules/groups/service"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
	paymentService "github.com/jumayevgadam/tsu-toleg/internal/modules/payments/service"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
	roleService "github.com/jumayevgadam/tsu-toleg/internal/modules/roles/service"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	timeService "github.com/jumayevgadam/tsu-toleg/internal/modules/times/service"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	userService "github.com/jumayevgadam/tsu-toleg/internal/modules/users/service"
)

// Ensuring services.DataService implements ServiceManager.
var _ services.DataService = (*ServiceManager)(nil)

type ServiceManager struct {
	role    roles.Service
	faculty faculties.Service
	group   groups.Service
	time    times.Service
	user    users.Service
	payment payments.Service
}

// NewServiceManager creates and returns a new instance of ServiceManager.
func NewServiceManager(dataStore database.DataStore, mw *middleware.Manager) services.DataService {
	return &ServiceManager{
		role:    roleService.NewRoleService(dataStore),
		faculty: facultyService.NewFacultyService(dataStore),
		group:   groupService.NewGroupService(dataStore),
		time:    timeService.NewTimeService(dataStore),
		user:    userService.NewUserService(mw, dataStore),
		payment: paymentService.NewPaymentService(dataStore),
	}
}

// IMPLEMENTING METHODS WITH SERVICEMANEGER.

func (sm *ServiceManager) RoleService() roles.Service {
	return sm.role
}

func (sm *ServiceManager) FacultyService() faculties.Service {
	return sm.faculty
}

func (sm *ServiceManager) GroupService() groups.Service {
	return sm.group
}

func (sm *ServiceManager) TimeService() times.Service {
	return sm.time
}

func (sm *ServiceManager) UserService() users.Service {
	return sm.user
}

func (sm *ServiceManager) PaymentService() payments.Service {
	return sm.payment
}

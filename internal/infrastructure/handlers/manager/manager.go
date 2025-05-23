package manager

import (
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/handlers"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	facultyHandler "github.com/jumayevgadam/tsu-toleg/internal/modules/faculties/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	groupHandler "github.com/jumayevgadam/tsu-toleg/internal/modules/groups/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
	paymentHandler "github.com/jumayevgadam/tsu-toleg/internal/modules/payments/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
	roleHandler "github.com/jumayevgadam/tsu-toleg/internal/modules/roles/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	timeHandler "github.com/jumayevgadam/tsu-toleg/internal/modules/times/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	userHandler "github.com/jumayevgadam/tsu-toleg/internal/modules/users/handler"
)

// Ensuring handlers.DataHandlers implements HandlerManager.
var _ handlers.DataHandlers = (*HandlerManager)(nil)

type HandlerManager struct {
	role    roles.Handlers
	faculty faculties.Handlers
	group   groups.Handlers
	time    times.Handlers
	user    users.Handlers
	payment payments.Handlers
}

// NewHandlerManager creates and returns a new instance of HandlerManager.
func NewHandlerManager(service services.DataService) handlers.DataHandlers {
	return &HandlerManager{
		role:    roleHandler.NewRoleHandler(service),
		faculty: facultyHandler.NewFacultyHandler(service),
		group:   groupHandler.NewGroupHandler(service),
		time:    timeHandler.NewTimeHandler(service),
		user:    userHandler.NewUserHandler(service),
		payment: paymentHandler.NewPaymentHandler(service),
	}
}

func (hm *HandlerManager) RoleHandler() roles.Handlers {
	return hm.role
}

func (hm *HandlerManager) FacultyHandler() faculties.Handlers {
	return hm.faculty
}

func (hm *HandlerManager) GroupHandler() groups.Handlers {
	return hm.group
}

func (hm *HandlerManager) TimeHandler() times.Handlers {
	return hm.time
}

func (hm *HandlerManager) UserHandler() users.Handlers {
	return hm.user
}

func (hm *HandlerManager) PaymentHandler() payments.Handlers {
	return hm.payment
}

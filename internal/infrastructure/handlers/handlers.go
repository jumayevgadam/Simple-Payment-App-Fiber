package handlers

import (
	"github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
)

// DataHandlers interface for general handling handlers in routes.
type DataHandlers interface {
	RoleHandler() roles.Handlers
	FacultyHandler() faculties.Handlers
	GroupHandler() groups.Handlers
	TimeHandler() times.Handlers
	UserHandler() users.Handlers
	PaymentHandler() payments.Handlers
}

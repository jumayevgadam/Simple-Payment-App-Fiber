package services

import (
	"github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payment"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
)

// DataService interface for using services generally in handlers.
type DataService interface {
	RoleService() roles.Service
	FacultyService() faculties.Service
	GroupService() groups.Service
	PaymentService() payment.Service
	UserService() users.Service
	TimeService() times.Service
}

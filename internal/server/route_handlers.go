// --- THIS GO FILE CONTAINS ALL NEEDED ENDPOINTS FOR THIS PROJECT --- //
// When request comes to server, then server -> infrastructure -> domain.

package server

import (
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	handlerManager "github.com/jumayevgadam/tsu-toleg/internal/infrastructure/handlers/manager"
	serviceManager "github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services/manager"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
)

const v1URL = "/api/v1"

// MapHandlers function contains all needed endpoints.
func (s *Server) MapHandlers(dataStore database.DataStore) {
	s.MapCustomMiddlewares()

	// Init MiddlewareManager.
	mdwManager := middleware.NewMiddlewareManager(s.Cfg, s.Logger)

	// Init v1 Path.
	v1 := s.Fiber.Group(v1URL)

	// Init Services.
	Services := serviceManager.NewServiceManager(dataStore, mdwManager)

	// Init Handlers.
	Handlers := handlerManager.NewHandlerManager(Services)

	authPath := v1.Group("/auth")
	{
		authPath.Post("/login", Handlers.UserHandler().Login())
	}

	superadminPath := v1.Group("/superadmin", mdwManager.RoleBasedMiddleware(constants.OnlySuperAdmin, constants.OnlySuperAdminID))
	{
		superadminPath.Post("/create-admin", Handlers.UserHandler().AddAdmin())
		superadminPath.Put("/update-admin/:admin_id", Handlers.UserHandler().UpdateAdmin())
		superadminPath.Delete("delete-admin/:admin_id", Handlers.UserHandler().DeleteAdmin())
	}

	// Init Users.
	adminPath := v1.Group("/admin", mdwManager.RoleBasedMiddleware(constants.AdminRoles, constants.AdminRoleIDs))
	{
		// ADMIN.
		adminPath.Post("/create-student", Handlers.UserHandler().AddStudent())
		adminPath.Get("/list-admins", Handlers.UserHandler().ListAdmins())
		adminPath.Get("/list-students", Handlers.UserHandler().ListStudents())
		adminPath.Get("/find-student", Handlers.UserHandler().AdminFindStudent())
		adminPath.Get("get-admin/:admin_id", Handlers.UserHandler().GetAdmin())
		adminPath.Get("get-student/:student_id", Handlers.UserHandler().GetStudent())
		adminPath.Delete("delete-student/:student_id", Handlers.UserHandler().DeleteStudent())
		adminPath.Put("/update-student/:student_id", Handlers.UserHandler().UpdateStudent())

		// Init Roles.
		roleGroup := adminPath.Group("/roles")
		{
			roleGroup.Post("/create", Handlers.RoleHandler().AddRole())
			roleGroup.Get("/", Handlers.RoleHandler().GetRoles())
			roleGroup.Get("/:id", Handlers.RoleHandler().GetRole())
			roleGroup.Delete("/:id", Handlers.RoleHandler().DeleteRole())
			roleGroup.Put("/:id", Handlers.RoleHandler().UpdateRole())
		}

		// Init Faculties.
		facultyGroup := adminPath.Group("/faculties")
		{
			facultyGroup.Post("/create", Handlers.FacultyHandler().AddFaculty())
			facultyGroup.Get("/", Handlers.FacultyHandler().ListFaculties())
			facultyGroup.Get("/list-groups", Handlers.FacultyHandler().ListGroupsByFacultyID())
			facultyGroup.Get("/:id", Handlers.FacultyHandler().GetFaculty())
			facultyGroup.Delete("/:id", Handlers.FacultyHandler().DeleteFaculty())
			facultyGroup.Put("/:id", Handlers.FacultyHandler().UpdateFaculty())
		}

		// Init Groups.
		groupPath := adminPath.Group("/groups")
		{
			groupPath.Post("/create", Handlers.GroupHandler().AddGroup())
			groupPath.Get("/", Handlers.GroupHandler().ListGroups())
			groupPath.Get("/students", Handlers.GroupHandler().ListStudentsByGroupID())
			groupPath.Get("/:id", Handlers.GroupHandler().GetGroup())
			groupPath.Delete("/:id", Handlers.GroupHandler().DeleteGroup())
			groupPath.Put("/:id", Handlers.GroupHandler().UpdateGroup())
		}

		// Init Times.
		timePath := adminPath.Group("/times")
		{
			timePath.Post("/create", Handlers.TimeHandler().AddTime())
			timePath.Get("/", Handlers.TimeHandler().ListTimes())
			timePath.Get("/active-year", Handlers.TimeHandler().SelectActiveYear())
			timePath.Get("/:time_id", Handlers.TimeHandler().GetTime())
			timePath.Delete("/:time_id", Handlers.TimeHandler().DeleteTime())
			timePath.Put("/:time_id", Handlers.TimeHandler().UpdateTime())
		}

		// Init Payments.
		paymentPath := adminPath.Group("/student-payments")
		{
			paymentPath.Get("/", Handlers.PaymentHandler().AdminListPaymentsByStudent())
			paymentPath.Put("/update/:student_id/:payment_id", Handlers.PaymentHandler().AdminUpdatePaymentOfStudent())
			paymentPath.Delete("/delete/:student_id/:payment_id", Handlers.PaymentHandler().AdminDeleteStudentPayment())
		}

		statisticsPath := adminPath.Group("/statistics")
		{
			statisticsPath.Get("/university", Handlers.PaymentHandler().AdminGetStatisticsAboutYear())
			statisticsPath.Get("/faculty/:faculty_id", Handlers.PaymentHandler().AdminGetStatisticsAboutFaculty())
		}
	}

	studentPath := v1.Group("/students", mdwManager.RoleBasedMiddleware(constants.StudentActionRoles, constants.StudentActionRoleIDs))
	{
		studentPath.Post("/add-payment", Handlers.PaymentHandler().AddPayment())
		studentPath.Get("/list-payments", Handlers.PaymentHandler().ListPaymentsByStudent())
		studentPath.Get("/get-payment", Handlers.PaymentHandler().GetPayment())
		studentPath.Put("/update-payment/:payment_id", Handlers.PaymentHandler().StudentUpdatePayment())
		studentPath.Delete("/delete/:payment_id", Handlers.PaymentHandler().StudentDeletePayment())
	}
}

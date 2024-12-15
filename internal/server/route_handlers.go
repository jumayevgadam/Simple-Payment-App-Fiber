// --- THIS GO FILE CONTAINS ALL NEEDED ENDPOINTS FOR THIS PROJECT --- //
// When request comes to server, then server -> infrastructure -> domain.

package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	handlerManager "github.com/jumayevgadam/tsu-toleg/internal/infrastructure/handlers/manager"
	serviceManager "github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services/manager"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
)

const v1URL = "/api/v1"

// MapHandlers function contains all needed endpoints.
func (s *Server) MapHandlers(dataStore database.DataStore) {
	s.Fiber.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":   "pong",
			"timestamp": time.Now(),
		})
	})

	s.Fiber.Use(func(c *fiber.Ctx) error {
		fmt.Println("Request Path:", c.Path())
		return c.Next()
	})

	s.Fiber.Static("/uploads", "./uploads", fiber.Static{
		Browse: true,
	})

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

	// Init Users.
	adminPath := v1.Group("/admin", mdwManager.RoleBasedMiddleware("admin", 2))
	{
		// ADMIN.
		adminPath.Post("/create-student", Handlers.UserHandler().AddStudent())
		adminPath.Post("/create-admin", Handlers.UserHandler().AddAdmin())
		adminPath.Get("/list-admins", Handlers.UserHandler().ListAdmins())
		adminPath.Get("/list-students", Handlers.UserHandler().ListStudents())
		adminPath.Get("get-admin/:admin_id", Handlers.UserHandler().GetAdmin())
		adminPath.Get("get-student/:student_id", Handlers.UserHandler().GetStudent())
		adminPath.Delete("delete-admin/:admin_id", Handlers.UserHandler().DeleteAdmin())
		adminPath.Delete("delete-student/:student_id", Handlers.UserHandler().DeleteStudent())
		adminPath.Put("/update-admin/:admin_id", Handlers.UserHandler().UpdateAdmin())
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
			facultyGroup.Get("/:id", Handlers.FacultyHandler().GetFaculty())
			facultyGroup.Delete("/:id", Handlers.FacultyHandler().DeleteFaculty())
			facultyGroup.Put("/:id", Handlers.FacultyHandler().UpdateFaculty())
			facultyGroup.Get("/:faculty_id/groups", Handlers.FacultyHandler().ListGroupsByFacultyID())
		}

		// Init Groups.
		groupPath := adminPath.Group("/groups")
		{
			groupPath.Post("/create", Handlers.GroupHandler().AddGroup())
			groupPath.Get("/", Handlers.GroupHandler().ListGroups())
			groupPath.Get("/:id", Handlers.GroupHandler().GetGroup())
			groupPath.Delete("/:id", Handlers.GroupHandler().DeleteGroup())
			groupPath.Put("/:id", Handlers.GroupHandler().UpdateGroup())
			groupPath.Get("/:group_id/students", Handlers.GroupHandler().ListStudentsByGroupID())
		}

		// Init Times.
		timePath := adminPath.Group("/times")
		{
			timePath.Post("/create", Handlers.TimeHandler().AddTime())
		}

		// Init Payments.
		paymentPath := adminPath.Group("/student-payments")
		{
			paymentPath.Get("/:student_id", Handlers.PaymentHandler().AdminListPaymentsByStudent())
		}
	}

	studentPath := v1.Group("/students", mdwManager.RoleBasedMiddleware(constants.Student, constants.DefaultRoleID))
	{
		studentPath.Post("/add-payment", Handlers.PaymentHandler().AddPayment())
		studentPath.Get("/list-payments", Handlers.PaymentHandler().ListPaymentsByStudent())
		studentPath.Get("/get-payment/:payment_id", Handlers.PaymentHandler().GetPayment())
		studentPath.Put("/update-payment/:payment_id", Handlers.PaymentHandler().StudentUpdatePayment())
	}
}

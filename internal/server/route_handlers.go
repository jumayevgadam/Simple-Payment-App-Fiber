// --- THIS GO FILE CONTAINS ALL NEEDED ENDPOINTS FOR THIS PROJECT --- //
// When request comes to server, then server -> infrastructure -> domain.

package server

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	handlerManager "github.com/jumayevgadam/tsu-toleg/internal/infrastructure/handlers/manager"
	serviceManager "github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services/manager"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware/permission"
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

	// Serve files dynamically based on the faculty and group structure.
	uploadsPath, err := filepath.Abs("./internal/uploads")
	if err != nil {
		s.Logger.Fatalf("Failed to resolve uploads path: %v", err)
	}

	// Init Static Route.
	s.Fiber.Static("/uploads", uploadsPath)

	// Init MiddlewareManager.
	mdwManager := middleware.NewMiddlewareManager(s.Cfg, s.Logger)

	// Init v1 Path.
	v1 := s.Fiber.Group(v1URL)

	// Init Services.
	Services := serviceManager.NewServiceManager(dataStore, mdwManager)

	// Init Handlers.
	Handlers := handlerManager.NewHandlerManager(Services)

	// Init Roles.
	roleGroup := v1.Group("/roles")
	{
		roleGroup.Post("/create",
			Handlers.RoleHandler().AddRole())

		roleGroup.Get("/",
			Handlers.RoleHandler().GetRoles())

		roleGroup.Get("/:id",
			Handlers.RoleHandler().GetRole())

		roleGroup.Delete("/:id",
			Handlers.RoleHandler().DeleteRole())

		roleGroup.Put("/:id",
			Handlers.RoleHandler().UpdateRole())
	}

	// Init Users.
	userGroup := v1.Group("/admin")
	{
		// ADMIN.
		userGroup.Post("/create-student",
			Handlers.UserHandler().AddStudent())

		userGroup.Post("/create-admin",
			Handlers.UserHandler().AddAdmin())

		userGroup.Get("/list-admins",
			Handlers.UserHandler().ListAdmins())

		userGroup.Get("/list-students",
			Handlers.UserHandler().ListStudents())

		userGroup.Get("get-admin/:admin_id",
			Handlers.UserHandler().GetAdmin())

		userGroup.Get("get-student/:student_id",
			Handlers.UserHandler().GetStudent())

		userGroup.Delete("delete-admin/:admin_id",
			Handlers.UserHandler().DeleteAdmin())

		userGroup.Delete("delete-student/:student_id",
			Handlers.UserHandler().DeleteStudent())
	}

	// Init Faculties.
	facultyGroup := v1.Group("/faculties")
	{
		facultyGroup.Post("/create",
			Handlers.FacultyHandler().AddFaculty())

		facultyGroup.Get("/",
			Handlers.FacultyHandler().ListFaculties())

		facultyGroup.Get("/:id",
			Handlers.FacultyHandler().GetFaculty())

		facultyGroup.Delete("/:id",
			Handlers.FacultyHandler().DeleteFaculty())

		facultyGroup.Put("/:id",
			Handlers.FacultyHandler().UpdateFaculty())

		facultyGroup.Get("/:faculty_id/groups",
			Handlers.FacultyHandler().ListGroupsByFacultyID())
	}

	// Init Groups.
	groupPath := v1.Group("/groups")
	{
		groupPath.Post("/create",
			Handlers.GroupHandler().AddGroup())

		groupPath.Get("/",
			Handlers.GroupHandler().ListGroups())

		groupPath.Get("/:id",
			Handlers.GroupHandler().GetGroup())

		groupPath.Delete("/:id",
			Handlers.GroupHandler().DeleteGroup())

		groupPath.Put("/:id",
			Handlers.GroupHandler().UpdateGroup())
	}

	// Init Payments.
	paymentGroup := v1.Group("/payments")
	{
		paymentGroup.Post("/add", mdwManager.RoleBasedMiddleware(permission.AddPayment),
			Handlers.PaymentHandler().AddPayment())

		paymentGroup.Get("/student", mdwManager.RoleBasedMiddleware(permission.StudentListPayments),
			Handlers.PaymentHandler().StudentListPaymentsByStudentID())

		paymentGroup.Get("/:payment_id",
			Handlers.PaymentHandler().GetPaymentByID())

		paymentGroup.Put("/:payment_id",
			Handlers.PaymentHandler().ChangePaymentStatus())
	}

	timeGroup := v1.Group("/time")
	{
		timeGroup.Post("/add", mdwManager.RoleBasedMiddleware(permission.AddTime),
			Handlers.TimeHandler().AddTime())
	}
}

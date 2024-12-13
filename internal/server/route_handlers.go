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

	// InitPermissions
	permissionGroup := v1.Group("/permissions")
	{
		permissionGroup.Post("/create",
			Handlers.RoleHandler().AddPermission())

		permissionGroup.Get("/",
			Handlers.RoleHandler().ListPermissions())

		permissionGroup.Get("/:id",
			Handlers.RoleHandler().GetPermission())

		permissionGroup.Delete("/:id",
			Handlers.RoleHandler().DeletePermission())

		permissionGroup.Put("/:id",
			Handlers.RoleHandler().UpdatePermission())
	}

	// Init RolePermissions.
	rolePermissionGroup := v1.Group("/role-permissions")
	{
		rolePermissionGroup.Post("/create",
			Handlers.RoleHandler().AddRolePermission())

		rolePermissionGroup.Get("/:role_id/permissions",
			Handlers.RoleHandler().GetPermissionsByRole())

		rolePermissionGroup.Get("/:permission_id/roles",
			Handlers.RoleHandler().GetRolesByPermission())

		rolePermissionGroup.Delete("/:role_id/and/:permission_id",
			Handlers.RoleHandler().DeleteRolePermission())
	}

	// Init Users.
	authGroup := v1.Group("/auth")
	{
		authGroup.Post("/register", Handlers.UserHandler().Register())
		authGroup.Post("/login", Handlers.UserHandler().Login())
	}

	usersGroup := v1.Group("/users")
	{
		usersGroup.Get("/",
			Handlers.UserHandler().ListUsers())

		usersGroup.Get("/:user_id",
			Handlers.UserHandler().GetUserByID())

		usersGroup.Delete("/:user_id",
			Handlers.UserHandler().DeleteUser())

		usersGroup.Put("/:user_id",
			Handlers.UserHandler().UpdateUser())
	}

	// Init Students.
	studentGroup := v1.Group("/students")
	{
		studentGroup.Get("/",
			Handlers.UserHandler().ListStudents())
	}

	// Init Faculties.
	facultyGroup := v1.Group("/faculties")
	{
		facultyGroup.Post("/create", mdwManager.RoleBasedMiddleware(permission.CreateFaculty),
			Handlers.FacultyHandler().AddFaculty())

		facultyGroup.Get("/", mdwManager.RoleBasedMiddleware(permission.ListFaculties),
			Handlers.FacultyHandler().ListFaculties())

		facultyGroup.Get("/:id", mdwManager.RoleBasedMiddleware(permission.GetFaculty),
			Handlers.FacultyHandler().GetFaculty())

		facultyGroup.Delete("/:id", mdwManager.RoleBasedMiddleware(permission.DeleteFaculty),
			Handlers.FacultyHandler().DeleteFaculty())

		facultyGroup.Put("/:id", mdwManager.RoleBasedMiddleware(permission.UpdateFaculty),
			Handlers.FacultyHandler().UpdateFaculty())

		facultyGroup.Get("/:faculty_id/groups",
			Handlers.FacultyHandler().ListGroupsByFacultyID())
	}

	// Init Groups.
	groupPath := v1.Group("/groups")
	{
		groupPath.Post("/create", mdwManager.RoleBasedMiddleware(permission.CreateGroup),
			Handlers.GroupHandler().AddGroup())

		groupPath.Get("/", mdwManager.RoleBasedMiddleware(permission.ListGroups),
			Handlers.GroupHandler().ListGroups())

		groupPath.Get("/:id", mdwManager.RoleBasedMiddleware(permission.GetGroup),
			Handlers.GroupHandler().GetGroup())

		groupPath.Delete("/:id", mdwManager.RoleBasedMiddleware(permission.DeleteGroup),
			Handlers.GroupHandler().DeleteGroup())

		groupPath.Put("/:id", mdwManager.RoleBasedMiddleware(permission.UpdateGroup),
			Handlers.GroupHandler().UpdateGroup())

		groupPath.Get("/:group_id/students", mdwManager.RoleBasedMiddleware(permission.ListStudentsByGroupID),
			Handlers.GroupHandler().ListStudentsByGroupID())
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

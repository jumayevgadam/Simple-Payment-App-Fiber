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
	"github.com/jumayevgadam/tsu-toleg/internal/middleware/permission"
)

const v1URL = "/api/v1"

// MapHandlers function contains all needed endpoints.
func (s *Server) MapHandlers(dataStore database.DataStore) {
	// s.Fiber.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowMethods: "GET,POST,PUT,DELETE",
	// 	AllowHeaders: "Content-Type, Authorization",
	// }))

	// Define a Ping Route.
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

	// Init Static Route.
	s.Fiber.Static("/uploads", "./uploads")

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
		usersGroup.Get("/", Handlers.UserHandler().ListUsers())
		usersGroup.Get("/:user_id", Handlers.UserHandler().GetUserByID())
		usersGroup.Delete("/:user_id", Handlers.UserHandler().DeleteUser())
		usersGroup.Put("/:user_id", Handlers.UserHandler().UpdateUser())
	}

	// Init Students.
	studentGroup := v1.Group("/students")
	{
		studentGroup.Get("/", Handlers.UserHandler().ListStudents())
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

		groupPath.Get("/:group_id/students",
			Handlers.GroupHandler().ListStudentsByGroupID())
	}

	// Init Payments.
	paymentGroup := v1.Group("/payments")
	{
		paymentGroup.Post("/add", mdwManager.RoleBasedMiddleware(permission.AddPayment),
			Handlers.PaymentHandler().AddPayment())
	}

	timeGroup := v1.Group("/time")
	{
		timeGroup.Post("/add", mdwManager.RoleBasedMiddleware(permission.AddTime),
			Handlers.TimeHandler().AddTime())
	}
}

// --- THIS GO FILE CONTAINS ALL NEEDED ENDPOINTS FOR THIS PROJECT --- //
// When request comes to server, then server -> infrastructure -> domain

package server

import (
	"github.com/gofiber/swagger"
	"github.com/jumayevgadam/tsu-toleg/docs"
	"github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	handlerManager "github.com/jumayevgadam/tsu-toleg/internal/infrastructure/handlers/manager"
	serviceManager "github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services/manager"
)

const v1URL = "/api/v1"

// MapHandlers function contains all needed endpoints.
func (s *Server) MapHandlers(dataStore database.DataStore) {
	// Init Swagger Doc details.
	docs.SwaggerInfo.Title = "API DOCUMENTATION OF TSU-TOLEG"
	s.Fiber.Get("/api-docs/tsu-toleg-api/*", swagger.HandlerDefault)

	// Init v1 Path.
	v1 := s.Fiber.Group(v1URL)

	// Init Services.
	Services := serviceManager.NewServiceManager(dataStore)

	// Init Handlers.
	Handlers := handlerManager.NewHandlerManager(Services)

	// Init Middlewares.
	mdwManager := middleware.NewMiddlewareManager(s.Cfg, s.Logger)

	// Init Roles.
	roleGroup := v1.Group("/role")
	{
		roleGroup.Post("/create", mdwManager.RoleBasedMiddleware("create:role"),
			Handlers.RoleHandler().AddRole())

		roleGroup.Get("/get-all", mdwManager.RoleBasedMiddleware("list:roles"),
			Handlers.RoleHandler().GetRoles())

		roleGroup.Get("/:id", mdwManager.RoleBasedMiddleware("get:role"),
			Handlers.RoleHandler().GetRole())

		roleGroup.Delete("/:id", mdwManager.RoleBasedMiddleware("delete:role"),
			Handlers.RoleHandler().DeleteRole())

		roleGroup.Put("/:id", mdwManager.RoleBasedMiddleware("update:role"),
			Handlers.RoleHandler().UpdateRole())
	}

	// InitPermissions
	permissionGroup := v1.Group("/permission")
	{
		permissionGroup.Post("/add", mdwManager.RoleBasedMiddleware("add:permission"),
			Handlers.RoleHandler().AddPermission())

		permissionGroup.Get("/list-all", mdwManager.RoleBasedMiddleware("list:permissions"),
			Handlers.RoleHandler().ListPermissions())

		permissionGroup.Get("/:id", mdwManager.RoleBasedMiddleware("get:permission"),
			Handlers.RoleHandler().GetPermission())

		permissionGroup.Delete("/:id", mdwManager.RoleBasedMiddleware("delete:permission"),
			Handlers.RoleHandler().DeletePermission())

		permissionGroup.Put("/:id", mdwManager.RoleBasedMiddleware("update:permission"),
			Handlers.RoleHandler().UpdatePermission())
	}

	// Init RolePermissions.
	rolePermissionGroup := v1.Group("/role-permission")
	{
		rolePermissionGroup.Post("/create", mdwManager.RoleBasedMiddleware("rolepermission:create"),
			Handlers.RoleHandler().AddRolePermission())

		rolePermissionGroup.Get("/:role_id", mdwManager.RoleBasedMiddleware("getpermissions:by:role"),
			Handlers.RoleHandler().GetPermissionsByRole())

		rolePermissionGroup.Get("/:permission_id", mdwManager.RoleBasedMiddleware("getroles:by:permission"),
			Handlers.RoleHandler().GetRolesByPermission())

		rolePermissionGroup.Delete("/:rol_id/and/:permission_id", mdwManager.RoleBasedMiddleware("delete:role:permission"),
			Handlers.RoleHandler().DeleteRolePermission())
	}

	// Init Users
	authGroup := v1.Group("/auth")
	{
		authGroup.Post("/register", Handlers.UserHandler().Register())
		authGroup.Post("/login", Handlers.UserHandler().Login())
	}

	// Init Faculties.
	facultyGroup := v1.Group("/faculty")
	{
		facultyGroup.Post("/create", mdwManager.RoleBasedMiddleware("create:faculty"),
			Handlers.FacultyHandler().AddFaculty())

		facultyGroup.Get("/get-all", mdwManager.RoleBasedMiddleware("list:faculties"),
			Handlers.FacultyHandler().ListFaculties())

		facultyGroup.Get("/:id", mdwManager.RoleBasedMiddleware("get:faculty"),
			Handlers.FacultyHandler().GetFaculty())

		facultyGroup.Delete("/:id", mdwManager.RoleBasedMiddleware("delete:faculty"),
			Handlers.FacultyHandler().DeleteFaculty())

		facultyGroup.Put("/:id", mdwManager.RoleBasedMiddleware("update:faculty"),
			Handlers.FacultyHandler().UpdateFaculty())
	}

	// Init Groups.
	groupPath := v1.Group("/group")
	{
		groupPath.Post("/add", mdwManager.RoleBasedMiddleware("add:group"),
			Handlers.GroupHandler().AddGroup())

		groupPath.Get("/get-all", mdwManager.RoleBasedMiddleware("list:groups"),
			Handlers.GroupHandler().ListGroups())

		groupPath.Get("/:id", mdwManager.RoleBasedMiddleware("get:group"),
			Handlers.GroupHandler().GetGroup())

		groupPath.Delete("/:id", mdwManager.RoleBasedMiddleware("delete:group"),
			Handlers.GroupHandler().DeleteGroup())

		groupPath.Put("/:id", mdwManager.RoleBasedMiddleware("update:group"),
			Handlers.GroupHandler().UpdateGroup())
	}

	// Init Payments.
	paymentGroup := v1.Group("/payment")
	{
		paymentGroup.Post("/add", mdwManager.RoleBasedMiddleware("add:payment"),
			Handlers.PaymentHandler().AddPayment())
	}
}

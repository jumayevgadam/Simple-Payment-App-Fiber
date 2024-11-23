package server

import (
	"github.com/gofiber/swagger"
	_ "github.com/jumayevgadam/tsu-toleg/docs"
	mwMngr "github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	facultyHTTP "github.com/jumayevgadam/tsu-toleg/internal/modules/faculties/routes"
	groupHTTP "github.com/jumayevgadam/tsu-toleg/internal/modules/groups/routes"
	roleHTTP "github.com/jumayevgadam/tsu-toleg/internal/modules/roles/routes"
	userHTTP "github.com/jumayevgadam/tsu-toleg/internal/modules/users/routes"
)

const (
	v1URL = "/api/v1"
)

// MapHandlers function takes all http routes.
func (s *Server) MapHandlers() error {
	// docs.SwaggerInfo.Title = "TSU-TOLEG API"
	s.Fiber.Get("/swagger/*", swagger.HandlerDefault)
	s.Fiber.Get("/swagger/*", swagger.New(swagger.Config{}))

	v1 := s.Fiber.Group(v1URL)

	mwOps := mwMngr.NewMiddlewareManager(s.Cfg, s.CacheStore)

	// roleHTTP is for app/role part of project.
	roleHTTP.Routes(v1, s.DataStore, s.CacheStore)
	// facultyHTTP  is for app/faculty part of project.
	facultyHTTP.Routes(v1, mwOps, s.DataStore)
	// groupHTTP route is for app/group part of project.
	groupHTTP.Routes(v1, mwOps, s.DataStore)
	// userHTTP route is for app/user part of project.
	userHTTP.Routes(v1, mwOps, s.DataStore, s.CacheStore)

	return nil
}

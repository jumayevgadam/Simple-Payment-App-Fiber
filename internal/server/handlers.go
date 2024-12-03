package server

import (
	"github.com/gofiber/swagger"
	"github.com/jumayevgadam/tsu-toleg/docs"
	mwMngr "github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	facultyHTTP "github.com/jumayevgadam/tsu-toleg/internal/modules/faculties/routes"
	groupHTTP "github.com/jumayevgadam/tsu-toleg/internal/modules/groups/routes"
	paymentHTTP "github.com/jumayevgadam/tsu-toleg/internal/modules/payment/routes"
	roleHTTP "github.com/jumayevgadam/tsu-toleg/internal/modules/roles/routes"
	userHTTP "github.com/jumayevgadam/tsu-toleg/internal/modules/users/routes"
)

const (
	v1URL = "/api/v1"
)

// MapHandlers function takes all http routes.
func (s *Server) MapHandlers() error {
	docs.SwaggerInfo.Title = "TALYP-TOLEG API"
	s.Fiber.Get("/api-docs/talyp-toleg-api/*", swagger.HandlerDefault)
	v1 := s.Fiber.Group(v1URL)

	mwOps := mwMngr.NewMiddlewareManager(s.Cfg, s.Logger)

	// roleHTTP is for app/role part of project.
	roleHTTP.Routes(v1, s.DataStore)
	// facultyHTTP  is for app/faculty part of project.
	facultyHTTP.Routes(v1, mwOps, s.DataStore)
	// groupHTTP route is for app/group part of project.
	groupHTTP.Routes(v1, mwOps, s.DataStore)
	// userHTTP route is for app/user part of project.
	userHTTP.Routes(v1, mwOps, s.DataStore)
	// paymentHTTP route is for app/payment part of project.
	paymentHTTP.Routes(v1, mwOps, s.DataStore)

	return nil
}

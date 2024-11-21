package server

import (
	"github.com/gofiber/fiber/v2"
	mwMngr "github.com/jumayevgadaym/tsu-toleg/internal/common/middleware"
	facultyHTTP "github.com/jumayevgadaym/tsu-toleg/internal/features/faculties/routes"
	groupHTTP "github.com/jumayevgadaym/tsu-toleg/internal/features/groups/routes"
	roleHTTP "github.com/jumayevgadaym/tsu-toleg/internal/features/roles/routes"
	userHTTP "github.com/jumayevgadaym/tsu-toleg/internal/features/users/routes"
)

const (
	v1URL = "/api/v1"
)

// MapHandlers function takes all http routes.
func (s *Server) MapHandlers(f *fiber.App) error {
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

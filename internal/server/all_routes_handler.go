package server

import (
	"github.com/gofiber/fiber/v2"
	facultyHTTP "github.com/jumayevgadaym/tsu-toleg/internal/app/faculties/routes"
	groupHTTP "github.com/jumayevgadaym/tsu-toleg/internal/app/groups/routes"
	roleHTTP "github.com/jumayevgadaym/tsu-toleg/internal/app/roles/routes"
	userHTTP "github.com/jumayevgadaym/tsu-toleg/internal/app/users/routes"
)

const (
	v1URL = "/api/v1"
)

func (s *Server) MapHandlers(f *fiber.App) error {
	v1 := s.Fiber.Group(v1URL)

	// roleHTTP route is
	roleHTTP.Routes(v1, s.DataStore, s.CacheStore)
	// facultyHTTP route is
	facultyHTTP.Routes(v1, s.DataStore)
	// groupHTTP route is
	groupHTTP.Routes(v1, s.DataStore)
	// userHTTP route is
	userHTTP.Routes(v1, s.DataStore)

	return nil
}

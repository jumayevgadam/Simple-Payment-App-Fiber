package server

import (
	"github.com/gofiber/fiber/v2"
	// userHTTP "github.com/jumayevgadam/tsu-toleg/internal/users/routes"
	facultyHTTP "github.com/jumayevgadaym/tsu-toleg/internal/faculties/routes"
	groupHTTP "github.com/jumayevgadaym/tsu-toleg/internal/groups/routes"
	roleHTTP "github.com/jumayevgadaym/tsu-toleg/internal/roles/routes"
	userHTTP "github.com/jumayevgadaym/tsu-toleg/internal/users/routes"
)

const (
	v1URL = "/api/v1"
)

func (s *Server) MapHandlers(f *fiber.App) error {
	v1 := s.Fiber.Group(v1URL)

	// roleHTTP route is
	roleHTTP.Routes(v1, s.DataStore)
	// facultyHTTP route is
	facultyHTTP.Routes(v1, s.DataStore)
	// groupHTTP route is
	groupHTTP.Routes(v1, s.DataStore)
	// userHTTP route is
	userHTTP.Routes(v1, s.DataStore)

	return nil
}

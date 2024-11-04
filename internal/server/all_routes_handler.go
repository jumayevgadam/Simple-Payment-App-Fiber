package server

import (
	"github.com/gofiber/fiber/v2"
	roleHTTP "github.com/jumayevgadaym/tsu-toleg/internal/roles/routes"
)

const (
	v1URL = "/api/v1"
)

func (s *Server) MapHandlers(f *fiber.App) error {
	v1 := s.Fiber.Group(v1URL)

	// roleHTTP route is
	roleHTTP.Routes(v1, s.DataStore)

	return nil
}

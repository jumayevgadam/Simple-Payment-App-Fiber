package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/internal/database"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

// Server struct is
type Server struct {
	Fiber     *fiber.App
	Cfg       *config.Config
	DataStore database.DataStore
}

// NewServer is
func NewServer(
	cfg *config.Config,
	dataStore database.DataStore,
) *Server {
	server := &Server{
		Fiber:     fiber.New(),
		Cfg:       cfg,
		DataStore: dataStore,
	}

	return server
}

// Run application
func (s *Server) Run() error {
	if err := s.MapHandlers(s.Fiber); err != nil {
		return errlst.ParseErrors(err)
	}

	return s.Fiber.Listen(":" + s.Cfg.Server.HTTPPort)
}

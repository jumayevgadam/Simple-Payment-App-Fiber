package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
)

// Server struct keeps all configurations needed for application.
type Server struct {
	Fiber     *fiber.App
	Cfg       *config.Config
	DataStore database.DataStore
	Logger    logger.Logger
}

// NewServer creates and returns a new instance of Server.
func NewServer(
	cfg *config.Config,
	dataStore database.DataStore,
	logger logger.Logger,
) *Server {
	httpServer := fiber.Config{
		ReadTimeout:  cfg.Server.ReadTimeOut,
		WriteTimeout: cfg.Server.WriteTimeOut,
	}

	server := &Server{
		Fiber:     fiber.New(httpServer),
		Cfg:       cfg,
		DataStore: dataStore,
		Logger:    logger,
	}

	return server
}

// Run method for running application.
func (s *Server) Run() error {
	if err := s.Fiber.Listen(":" + s.Cfg.Server.HTTPPort); err != nil {
		s.Logger.Errorf("error listening port: %s", s.Cfg.Server.HTTPPort)
	}

	return nil
}

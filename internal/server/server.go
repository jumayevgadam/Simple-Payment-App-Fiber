package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jumayevgadam/tsu-toleg/docs"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
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
	server := &Server{
		Fiber:     fiber.New(),
		Cfg:       cfg,
		DataStore: dataStore,
		Logger:    logger,
	}

	return server
}

// Run method for running application.
func (s *Server) Run() error {
	if err := s.MapHandlers(); err != nil {
		return errlst.ParseErrors(err)
	}

	err := s.Fiber.Listen(":" + s.Cfg.Server.HTTPPort)
	if err != nil {
		return fmt.Errorf("failed to listen app: %w", err)
	}

	return nil
}

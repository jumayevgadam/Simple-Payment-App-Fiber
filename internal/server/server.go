package server

import (
	"os"
	"os/signal"
	"syscall"

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

	server.Logger.InitLogger()

	return server
}

func (s *Server) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		s.Logger.Infof("server started on port %s", s.Cfg.Server.HTTPPort)

		if err := s.Fiber.Listen(":" + s.Cfg.Server.HTTPPort); err != nil {
			s.Logger.Errorf("error occured when running http port: %s", s.Cfg.Server.HTTPPort)
		}
	}()

	s.MapHandlers(s.DataStore)

	<-quit
	s.Logger.Info("got interruption signal")

	if err := s.Fiber.Shutdown(); err != nil {
		s.Logger.Errorf("error occured when shutting down application")
	}

	s.Logger.Info("application shutdown successfully.")

	return nil
}

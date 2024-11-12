package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/cache"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/internal/database"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

// Server struct is
type Server struct {
	Fiber      *fiber.App
	Cfg        *config.Config
	DataStore  database.DataStore
	CacheStore cache.Store
}

// NewServer is
func NewServer(
	cfg *config.Config,
	dataStore database.DataStore,
	cacheStore cache.Store,
) *Server {
	server := &Server{
		Fiber:      fiber.New(),
		Cfg:        cfg,
		DataStore:  dataStore,
		CacheStore: cacheStore,
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

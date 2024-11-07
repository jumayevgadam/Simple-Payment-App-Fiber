package handler

import userOps "github.com/jumayevgadaym/tsu-toleg/internal/users"

var (
	_ userOps.Handler = (*UserHandler)(nil)
)

// UserHandler is
type UserHandler struct {
	service userOps.Service
}

// NewUserHandler is
func NewUserHandler(service userOps.Service) *UserHandler {
	return &UserHandler{service: service}
}

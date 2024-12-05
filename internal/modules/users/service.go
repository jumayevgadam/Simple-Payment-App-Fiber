package users

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
)

// Service interface for performing actions in this layer.
type Service interface {
	Register(ctx context.Context, req userModel.SignUpReq) (int, error)
	Login(ctx context.Context, loginReq userModel.LoginReq) (string, error)
}

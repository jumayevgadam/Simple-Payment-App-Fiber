package users

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
)

// Service interface for performing actions in this layer.
type Service interface {
	CreateUser(ctx context.Context, req userModel.SignUpReq, role string) (int, error)
	Login(ctx context.Context, loginReq userModel.LoginReq, role string) (userModel.UserWithTokens, error)
}

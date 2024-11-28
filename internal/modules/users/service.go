package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
)

// Service interface for performing actions in this layer.
type Service interface {
	CreateUser(ctx context.Context, req userModel.SignUpReq, role string) (int, error)
	Login(ctx context.Context, loginReq userModel.LoginReq, role string) (userModel.UserWithTokens, error)
	RenewAccessToken(ctx *fiber.Ctx, refreshToken string) (string, string, error)
}

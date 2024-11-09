package users

import (
	"context"

	userModel "github.com/jumayevgadaym/tsu-toleg/internal/models/user"
)

// Repository interface for performing actions in this layer
type Repository interface {
	CreateUser(ctx context.Context, res *userModel.SignUpRes) (int, error)
	GetUserByUsername(ctx context.Context, username string) (*userModel.AllUserDAO, error)
}

package users

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
)

// Repository interface for performing actions in this layer.
type Repository interface {
	CreateUser(ctx context.Context, res userModel.SignUpRes) (int, error)
	GetUserByUsername(ctx context.Context, username string) (*userModel.Details, error)
}

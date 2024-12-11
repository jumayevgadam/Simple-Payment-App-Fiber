package users

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Service interface for performing actions in this layer.
type Service interface {
	Register(ctx context.Context, req userModel.SignUpReq) (int, error)
	Login(ctx context.Context, loginReq userModel.LoginReq) (string, error)
	ListAllUsers(ctx context.Context, pagination abstract.PaginationQuery) (
		abstract.PaginatedResponse[*userModel.AllUserDTO], error,
	)
	UpdateUser(ctx context.Context, userID int) error
}

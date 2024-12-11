package users

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Repository interface for performing actions in this layer.
type Repository interface {
	CreateUser(ctx context.Context, res userModel.SignUpRes) (int, error)
	GetUserByID(ctx context.Context, userID int) (*userModel.AllUserDAO, error)
	GetUserByUsername(ctx context.Context, username string) (*userModel.Details, error)
	GetStudentDetailsForPayment(ctx context.Context, studentID int) (*userModel.StudentInfoData, error)
	ListAllUsers(ctx context.Context, paginationData abstract.PaginationData) ([]*userModel.AllUserDAO, error)
	CountAllUsers(ctx context.Context) (int, error)
	UpdateUser(ctx context.Context, userID int, updateRes *userModel.UpdateUserDetailsData) error
	DeleteUser(ctx context.Context, userID int) error
}

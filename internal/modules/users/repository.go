package users

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

type Repository interface {
	Login(ctx context.Context, username string) (userModel.LoginResponseData, error)
	// ADMIN.
	AddAdmin(ctx context.Context, res userModel.AdminResponse) (int, error)
	GetAdmin(ctx context.Context, adminID int) (*userModel.AdminData, error)
	DeleteAdmin(ctx context.Context, adminID int) error
	CountAdmins(ctx context.Context) (int, error)
	UpdateAdmin(ctx context.Context, updateData userModel.AdminUpdateData) (string, error)
	ListAdmins(ctx context.Context, paginationData abstract.PaginationData) ([]*userModel.AdminData, error)

	AddStudent(ctx context.Context, res userModel.Response) (int, error)
	GetStudent(ctx context.Context, studentID int) (*userModel.StudentData, error)
	DeleteStudent(ctx context.Context, studentID int) error
	UpdateStudent(ctx context.Context, updateData userModel.StudentUpdateData) (string, error)
	CountStudents(ctx context.Context) (int, error)
	CountStudentsByGroupID(ctx context.Context, groupID int) (int, error)
	ListStudents(ctx context.Context, paginationData abstract.PaginationData) ([]*userModel.StudentData, error)
	ListStudentsByGroupID(ctx context.Context, groupID int, paginationData abstract.PaginationData) (
		[]*userModel.StudentDataByGroupID, error,
	)

	// STUDENT.
}

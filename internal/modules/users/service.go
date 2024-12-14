package users

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

type Service interface {
	// ADMIN.
	AddAdmin(ctx context.Context, request userModel.AdminRequest) (int, error)
	GetAdmin(ctx context.Context, adminID int) (*userModel.Admin, error)
	DeleteAdmin(ctx context.Context, adminID int) error
	ListAdmins(ctx context.Context, paginationQuery abstract.PaginationQuery) (
		abstract.PaginatedResponse[*userModel.Admin], error,
	)

	AddStudent(ctx context.Context, request userModel.Request) (int, error)
	GetStudent(ctx context.Context, studentID int) (*userModel.Student, error)
	DeleteStudent(ctx context.Context, studentID int) error
	ListStudents(ctx context.Context, paginationQuery abstract.PaginationQuery) (
		abstract.PaginatedResponse[*userModel.Student], error,
	)
	ListStudentsByGroupID(ctx context.Context, groupID int, paginationData abstract.PaginationQuery) (
		abstract.PaginatedResponse[*userModel.StudentResGroupID], error,
	)

	// STUDENT.

}

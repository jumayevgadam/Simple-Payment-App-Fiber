package users

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

type Service interface {
	Login(ctx context.Context, loginRequest userModel.LoginRequest) (userModel.LoginResponseWithToken, error)
	// ADMIN.
	AddAdmin(ctx context.Context, request userModel.AdminRequest) (int, error)
	GetAdmin(ctx context.Context, adminID int) (*userModel.Admin, error)
	DeleteAdmin(ctx context.Context, adminID int) error
	UpdateAdmin(ctx context.Context, adminID int, updateReq userModel.AdminUpdateRequest) (string, error)
	ListAdmins(ctx context.Context, paginationQuery abstract.PaginationQuery) (
		abstract.PaginatedResponse[*userModel.Admin], error,
	)
	AdminFindStudent(ctx context.Context, filter userModel.FilterStudent, paginationQuery abstract.PaginationQuery) (
		abstract.PaginatedResponse[*userModel.AllStudentDTO], error,
	)

	AddStudent(ctx context.Context, request userModel.Request) (int, error)
	GetStudent(ctx context.Context, studentID int) (*userModel.Student, error)
	DeleteStudent(ctx context.Context, studentID int) error
	UpdateStudent(ctx context.Context, studentID int, updateRequest userModel.StudentUpdateRequest) (string, error)
	ListStudents(ctx context.Context, paginationQuery abstract.PaginationQuery) (
		abstract.PaginatedResponse[*userModel.Student], error,
	)
	ListStudentsByGroupID(ctx context.Context, groupID int, paginationData abstract.PaginationQuery) (
		abstract.PaginatedResponse[*userModel.StudentResGroupID], error,
	)

	// STUDENT.

}

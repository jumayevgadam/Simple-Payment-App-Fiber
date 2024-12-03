package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// Ensure UserRepository implements the users.Repository interface.
var (
	_ users.Repository = (*UserRepository)(nil)
)

// UserRepository manages database methods for users.
type UserRepository struct {
	psqlDB connection.DB
}

// NewUserRepository creates and returns a new instance of UserRepository.
func NewUserRepository(psqlDB connection.DB) *UserRepository {
	return &UserRepository{psqlDB: psqlDB}
}

// CreateUser repo insert user data into db and returns id.
func (r *UserRepository) CreateUser(ctx context.Context, user userModel.SignUpRes) (int, error) {
	var userID int

	if err := r.psqlDB.QueryRow(
		ctx,
		createUserQuery,
		user.RoleID,
		user.GroupID,
		user.Name,
		user.Surname,
		user.UserName,
		user.Password,
	).Scan(&userID); err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return userID, nil
}

// GetUserByUsername fetches user by using identified username.
func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*userModel.Details, error) {
	var details userModel.Details

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&details,
		getDetailsByUsernameQuery,
		username,
	)
	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &details, nil
}

// GetStudentDetailsForPayment repository method for Payment.
func (r *UserRepository) GetStudentDetailsForPayment(ctx context.Context, studentID int) (*userModel.StudentInfoData, error) {
	var studentInfo userModel.StudentInfoData

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&studentInfo,
		getStudentInfoDetailsQuery,
		studentID,
	)
	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &studentInfo, nil
}

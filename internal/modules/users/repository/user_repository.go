package repository

import (
	"context"
	"log"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
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

	log.Println("Before query, checking pool state:", r.psqlDB)

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
		log.Println("[UserRepository][CreateUser]", err)
		return -1, errlst.ParseSQLErrors(err)
	}

	return userID, nil
}

// GetUserByID repo method fetches user by its id.
func (r *UserRepository) GetUserByID(ctx context.Context, userID int) (*userModel.AllUserDAO, error) {
	var userAllDAO userModel.AllUserDAO

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&userAllDAO,
		getUserByIDQuery,
		userID,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &userAllDAO, nil
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
func (r *UserRepository) GetStudentDetailsForPayment(ctx context.Context, studentID int) (
	*userModel.StudentInfoData, error,
) {
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

// CountAllUsers repo.
func (r *UserRepository) CountAllUsers(ctx context.Context) (int, error) {
	var totalCountOfAllUser int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalCountOfAllUser,
		totalCountOfAllUserQuery,
	)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalCountOfAllUser, nil
}

// UpdateUser repo.
func (r *UserRepository) UpdateUser(ctx context.Context, userID int, updateRes *userModel.UpdateUserDetailsData) error {
	err := r.psqlDB.QueryRow(
		ctx,
		updateUserDetailsQuery,
		updateRes.RoleID,
		updateRes.GroupID,
		updateRes.Name,
		updateRes.Surname,
		updateRes.Username,
		userID,
	).Scan(nil)

	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	return nil
}

// CountAllUsers repo.
func (r *UserRepository) ListAllUsers(ctx context.Context, paginationData abstract.PaginationData) (
	[]*userModel.AllUserDAO, error,
) {
	var allUsers []*userModel.AllUserDAO

	offset := (paginationData.Page - 1) * paginationData.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&allUsers,
		listAllUsersQuery,
		offset,
		paginationData.Limit,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return allUsers, nil
}

// DeleteUser repo.
func (r *UserRepository) DeleteUser(ctx context.Context, userID int) error {
	_, err := r.psqlDB.Exec(
		ctx,
		deleteUserQuery,
		userID,
	)

	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	return nil
}

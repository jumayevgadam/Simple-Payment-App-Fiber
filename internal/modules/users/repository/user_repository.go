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
func (r *UserRepository) CreateUser(ctx context.Context, roleID int, user userModel.SignUpRes) (int, error) {
	var userID int

	if err := r.psqlDB.QueryRow(
		ctx,
		createUserQuery,
		roleID,
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
func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (userModel.AllUserDAO, error) {
	var userDAO userModel.AllUserDAO

	if err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&userDAO,
		getUserByUsernameQuery,
		username,
	); err != nil {
		return userModel.AllUserDAO{}, errlst.ParseSQLErrors(err)
	}

	return userDAO, nil
}

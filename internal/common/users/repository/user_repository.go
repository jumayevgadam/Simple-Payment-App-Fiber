package repository

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/common/users"
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	userModel "github.com/jumayevgadaym/tsu-toleg/internal/models/user"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

var (
	_ users.Repository = (*UserRepository)(nil)
)

// UserRepository is
type UserRepository struct {
	psqlDB connection.DB
}

// NewUserRepository is
func NewUserRepository(psqlDB connection.DB) *UserRepository {
	return &UserRepository{psqlDB: psqlDB}
}

// CreateUser repo is
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

// GetUserByUsername repo is
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

package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

var _ users.Repository = (*UserRepository)(nil)

type UserRepository struct {
	psqlDB connection.DB
}

func NewUserRepository(psqlDB connection.DB) *UserRepository {
	return &UserRepository{psqlDB: psqlDB}
}

func (r *UserRepository) Login(ctx context.Context, username string) (userModel.LoginResponseData, error) {
	var loginResponseData userModel.LoginResponseData

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&loginResponseData,
		loginUserCheckWithQuery,
		username,
	)

	if err != nil {
		return userModel.LoginResponseData{}, errlst.ParseSQLErrors(err)
	}

	return loginResponseData, nil
}

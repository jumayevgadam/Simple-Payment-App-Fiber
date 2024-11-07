package repository

import "github.com/jumayevgadaym/tsu-toleg/internal/connection"

// UserRepository is
type UserRepository struct {
	psqlDB connection.DB
}

// NewUserRepository is
func NewUserRepository(psqlDB connection.DB) *UserRepository {
	return &UserRepository{psqlDB: psqlDB}
}

package repository

import "github.com/jumayevgadaym/tsu-toleg/internal/connection"

// GroupRepository is
type GroupRepository struct {
	psqlDB connection.DB
}

// NewGroupRepository is
func NewGroupRepository(psqlDB connection.DB) *GroupRepository {
	return &GroupRepository{psqlDB: psqlDB}
}

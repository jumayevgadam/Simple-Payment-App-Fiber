package repository

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/common/groups"
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	groupModel "github.com/jumayevgadaym/tsu-toleg/internal/models/group"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

var (
	_ groups.Repository = (*GroupRepository)(nil)
)

// GroupRepository is
type GroupRepository struct {
	psqlDB connection.DB
}

// NewGroupRepository is
func NewGroupRepository(psqlDB connection.DB) *GroupRepository {
	return &GroupRepository{psqlDB: psqlDB}
}

// AddGroup repo is
func (r *GroupRepository) AddGroup(ctx context.Context, groupDAO *groupModel.GroupRes) (int, error) {
	var groupID int

	if err := r.psqlDB.QueryRow(
		ctx,
		addGroupQuery,
		groupDAO.FacultyID,
		groupDAO.ClassCode,
	).Scan(&groupID); err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return groupID, nil
}

// GetGroup repo is
func (r *GroupRepository) GetGroup(ctx context.Context, groupID int) (groupModel.GroupDAO, error) {
	var groupDAO groupModel.GroupDAO

	if err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&groupDAO,
		getGroupQuery,
		groupID,
	); err != nil {
		return groupDAO, errlst.ParseSQLErrors(err)
	}

	return groupDAO, nil
}

// ListGroups repo is
func (r *GroupRepository) ListGroups(ctx context.Context) ([]groupModel.GroupDAO, error) {
	var groupDAOs []groupModel.GroupDAO

	if err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&groupDAOs,
		listGroupsQuery,
	); err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return groupDAOs, nil
}

// DeleteGroup repo is
func (r *GroupRepository) DeleteGroup(ctx context.Context, groupID int) error {
	_, err := r.psqlDB.Exec(
		ctx,
		deleteGroupQuery,
		groupID,
	)
	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	return nil
}

// UpdateGroup repo is
func (r *GroupRepository) UpdateGroup(ctx context.Context, groupDAO groupModel.GroupDAO) (string, error) {
	var res string

	if err := r.psqlDB.QueryRow(
		ctx,
		updateGroupQuery,
		groupDAO.FacultyID,
		groupDAO.ClassCode,
		groupDAO.ID,
	).Scan(&res); err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return res, nil
}

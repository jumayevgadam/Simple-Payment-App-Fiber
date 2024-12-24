package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	groupModel "github.com/jumayevgadam/tsu-toleg/internal/models/group"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// Ensure GroupRepository implements the groups.Repository interface.
var (
	_ groups.Repository = (*GroupRepository)(nil)
)

// GroupRepository performs database actions for groups.
type GroupRepository struct {
	psqlDB connection.DB
}

// NewGroupRepository creates and returns a new instance of GroupRepository.
func NewGroupRepository(psqlDB connection.DB) *GroupRepository {
	return &GroupRepository{psqlDB: psqlDB}
}

// AddGroup repo insert group data into db and returns id.
func (r *GroupRepository) AddGroup(ctx context.Context, groupDAO *groupModel.Res) (int, error) {
	var groupID int

	if err := r.psqlDB.QueryRow(
		ctx,
		addGroupQuery,
		groupDAO.FacultyID,
		groupDAO.GroupCode,
		groupDAO.CourseYear,
		groupDAO.Index,
	).Scan(&groupID); err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return groupID, nil
}

// GetGroup repo fetches a group using identified id.
func (r *GroupRepository) GetGroup(ctx context.Context, groupID int) (*groupModel.DAO, error) {
	var groupDAO groupModel.DAO

	if err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&groupDAO,
		getGroupQuery,
		groupID,
	); err != nil {
		return &groupDAO, errlst.ParseSQLErrors(err)
	}

	return &groupDAO, nil
}

// CountGroups repo method gives totalCount of groups.
func (r *GroupRepository) CountGroups(ctx context.Context) (int, error) {
	var totalCount int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalCount,
		countGroupsQuery,
	)
	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalCount, nil
}

// ListGroups repo fetches a list of groups.
func (r *GroupRepository) ListGroups(ctx context.Context, pagination abstract.PaginationData) (
	[]*groupModel.DAO, error,
) {
	var groupDAOs []*groupModel.DAO
	offset := (pagination.CurrentPage - 1) * pagination.Limit

	if err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&groupDAOs,
		listGroupsQuery,
		offset,
		pagination.Limit,
	); err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return groupDAOs, nil
}

// DeleteGroup repo deletes a group from db using identified id.
func (r *GroupRepository) DeleteGroup(ctx context.Context, groupID int) error {
	result, err := r.psqlDB.Exec(
		ctx,
		deleteGroupQuery,
		groupID,
	)

	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	if result.RowsAffected() == 0 {
		return errlst.ErrGroupNotFound
	}

	return nil
}

// UpdateGroup repo updates group data with a new group data and identified id.
func (r *GroupRepository) UpdateGroup(ctx context.Context, groupDAO *groupModel.DAO) (string, error) {
	var res string

	if err := r.psqlDB.QueryRow(
		ctx,
		updateGroupQuery,
		groupDAO.FacultyID,
		groupDAO.GroupCode,
		groupDAO.CourseYear,
		groupDAO.Index,
		groupDAO.ID,
	).Scan(&res); err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return res, nil
}

// CountGroupsByFacultyID repo.
func (r *GroupRepository) CountGroupsByFacultyID(ctx context.Context, facultyID int) (int, error) {
	var totalGroupCount int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalGroupCount,
		countGroupsByFacultyIDQuery,
		facultyID,
	)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalGroupCount, nil
}

func (r *GroupRepository) ListGroupsByFacultyID(ctx context.Context, facultyID int, paginationData abstract.PaginationData) (
	[]*groupModel.DAO, error,
) {
	var groups []*groupModel.DAO
	offset := (paginationData.CurrentPage - 1) * paginationData.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&groups,
		listGroupsByFacultyIDQuery,
		facultyID,
		offset,
		paginationData.Limit,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return groups, nil
}

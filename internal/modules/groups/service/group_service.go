package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	groupModel "github.com/jumayevgadam/tsu-toleg/internal/models/group"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/samber/lo"
)

// Ensure GroupService implements the groups.Service interface.
var (
	_ groups.Service = (*GroupService)(nil)
)

// GroupService performs buisiness logic for app/group part of application.
type GroupService struct {
	repo database.DataStore
}

// NewGroupService creates and returns a new instance of GroupService.
func NewGroupService(repo database.DataStore) *GroupService {
	return &GroupService{repo: repo}
}

// AddGroup service insert a group to db and returns id.
func (s *GroupService) AddGroup(ctx context.Context, groupRequest *groupModel.Req) (int, error) {
	groupID, err := s.repo.GroupsRepo().AddGroup(ctx, groupRequest.ToStorage())
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return groupID, nil
}

// GetGroup service fetches a group from db using identified id.
func (s *GroupService) GetGroup(ctx context.Context, groupID int) (*groupModel.DTO, error) {
	var groupDTO *groupModel.DTO

	groupDAO, err := s.repo.GroupsRepo().GetGroup(ctx, groupID)
	if err != nil {
		return groupDTO, errlst.ParseErrors(err)
	}

	groupDTO = groupDAO.ToServer()

	return groupDTO, nil
}

// ListGroups service fetches a list of groups from db and returns it.
func (s *GroupService) ListGroups(ctx context.Context, pagination abstract.PaginationQuery) (
	abstract.PaginatedResponse[*groupModel.DTO], error,
) {
	var (
		groupAllDatas     []*groupModel.DAO
		err               error
		groupListResponse abstract.PaginatedResponse[*groupModel.DTO]
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		var count int

		count, err = db.GroupsRepo().CountGroups(ctx)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		groupListResponse.TotalItems = count

		groupAllDatas, err = db.GroupsRepo().ListGroups(ctx, pagination.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})
	if err != nil {
		return abstract.PaginatedResponse[*groupModel.DTO]{}, errlst.ParseErrors(err)
	}

	groupList := lo.Map(
		groupAllDatas,
		func(item *groupModel.DAO, _ int) *groupModel.DTO {
			return item.ToServer()
		},
	)

	groupListResponse.Items = groupList
	groupListResponse.Page = pagination.Page
	groupListResponse.Limit = len(groupList)

	return groupListResponse, nil
}

// DeleteGroup service deletes a group from db using identified id.
func (s *GroupService) DeleteGroup(ctx context.Context, groupID int) error {
	if err := s.repo.GroupsRepo().DeleteGroup(ctx, groupID); err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}

// UpdateGroup service edits group data with a new group data using identified group id.
func (s *GroupService) UpdateGroup(ctx context.Context, groupID int, inputValue *groupModel.UpdateGroupReq) (
	string, error,
) {
	var (
		resFromDB string
		err       error
	)
	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// check group exist in this id
		_, err = db.GroupsRepo().GetGroup(ctx, groupID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		resFromDB, err = db.GroupsRepo().UpdateGroup(ctx, inputValue.ToStorage(groupID))
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return "", errlst.ParseErrors(err)
	}

	return resFromDB, nil
}

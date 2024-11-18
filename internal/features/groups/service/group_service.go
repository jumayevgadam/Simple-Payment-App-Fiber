package service

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/features/groups"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/database"
	groupModel "github.com/jumayevgadaym/tsu-toleg/internal/models/group"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"go.opentelemetry.io/otel"
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
func (s *GroupService) AddGroup(ctx context.Context, groupRequest *groupModel.GroupReq) (int, error) {
	ctx, span := otel.Tracer("[GroupService]").Start(ctx, "[AddGroup]")
	defer span.End()

	groupID, err := s.repo.GroupsRepo().AddGroup(ctx, groupRequest.ToStorage())
	if err != nil {
		tracing.ErrorTracer(span, err)
		return -1, errlst.ParseErrors(err)
	}

	return groupID, nil
}

// GetGroup service fetches a group from db using identified id.
func (s *GroupService) GetGroup(ctx context.Context, groupID int) (groupModel.GroupDTO, error) {
	ctx, span := otel.Tracer("[GroupService]").Start(ctx, "[GetGroup]")
	defer span.End()
	var groupDTO groupModel.GroupDTO

	groupDAO, err := s.repo.GroupsRepo().GetGroup(ctx, groupID)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return groupDTO, errlst.ParseErrors(err)
	}

	groupDTO = groupDAO.ToServer()

	return groupDTO, nil
}

// ListGroups service fetches a list of groups from db and returns it.
func (s *GroupService) ListGroups(ctx context.Context) ([]groupModel.GroupDTO, error) {
	ctx, span := otel.Tracer("[GroupService]").Start(ctx, "[ListGroups]")
	defer span.End()

	var groupDTOs []groupModel.GroupDTO
	if err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		groupDAOs, err := db.GroupsRepo().ListGroups(ctx)
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		for _, res := range groupDAOs {
			groupDTOs = append(groupDTOs, res.ToServer())
		}

		return nil
	}); err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return groupDTOs, nil
}

// DeleteGroup service deletes a group from db using identified id.
func (s *GroupService) DeleteGroup(ctx context.Context, groupID int) error {
	ctx, span := otel.Tracer("[GroupService]").Start(ctx, "[DeleteGroup]")
	defer span.End()

	if err := s.repo.GroupsRepo().DeleteGroup(ctx, groupID); err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	return nil
}

// UpdateGroup service edits group data with a new group data using identified group id.
func (s *GroupService) UpdateGroup(ctx context.Context, groupDTO groupModel.GroupDTO) (string, error) {
	ctx, span := otel.Tracer("[GroupService]").Start(ctx, "[UpdateGroup]")
	defer span.End()
	var (
		resFromDB string
		err       error
	)
	if err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// check group exist in this id
		_, err = db.GroupsRepo().GetGroup(ctx, groupDTO.ID)
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		resFromDB, err = db.GroupsRepo().UpdateGroup(ctx, groupDTO.ToStorage())
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		return nil
	}); err != nil {
		tracing.ErrorTracer(span, err)
		return "", errlst.ParseErrors(err)
	}

	return resFromDB, nil
}

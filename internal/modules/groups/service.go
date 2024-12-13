package groups

import (
	"context"

	groupModel "github.com/jumayevgadam/tsu-toleg/internal/models/group"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Service interface for performing actions in this layer (for groups).
type Service interface {
	AddGroup(ctx context.Context, groupDTO *groupModel.Req) (int, error)
	GetGroup(ctx context.Context, groupID int) (*groupModel.DTO, error)
	ListGroups(ctx context.Context, pagination abstract.PaginationQuery) (
		abstract.PaginatedResponse[*groupModel.DTO], error)
	DeleteGroup(ctx context.Context, groupID int) error
	UpdateGroup(ctx context.Context, groupID int, updateInput *groupModel.UpdateGroupReq) (string, error)
	ListGroupsByFacultyID(ctx context.Context, facultyID int, pagination abstract.PaginationQuery) (
		abstract.PaginatedResponse[*groupModel.DTO], error)
}

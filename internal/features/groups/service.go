package groups

import (
	"context"

	groupModel "github.com/jumayevgadaym/tsu-toleg/internal/models/group"
	"github.com/jumayevgadaym/tsu-toleg/pkg/abstract"
)

// Service interface for performing actions in this layer (for groups).
type Service interface {
	AddGroup(ctx context.Context, groupDTO *groupModel.GroupReq) (int, error)
	GetGroup(ctx context.Context, groupID int) (*groupModel.GroupDTO, error)
	ListGroups(ctx context.Context, pagination abstract.PaginationQuery) ([]*groupModel.GroupDTO, error)
	DeleteGroup(ctx context.Context, groupID int) error
	UpdateGroup(ctx context.Context, groupDTO *groupModel.GroupDTO) (string, error)
}

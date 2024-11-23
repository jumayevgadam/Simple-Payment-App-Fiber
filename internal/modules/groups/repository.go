package groups

import (
	"context"

	groupModel "github.com/jumayevgadam/tsu-toleg/internal/models/group"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Repository interface for performing actions in groups repo (layer).
type Repository interface {
	AddGroup(ctx context.Context, groupDAO *groupModel.GroupRes) (int, error)
	GetGroup(ctx context.Context, groupID int) (*groupModel.GroupDAO, error)
	ListGroups(ctx context.Context, pagination abstract.PaginationData) ([]*groupModel.GroupDAO, error)
	DeleteGroup(ctx context.Context, groupID int) error
	UpdateGroup(ctx context.Context, groupModel *groupModel.GroupDAO) (string, error)
}

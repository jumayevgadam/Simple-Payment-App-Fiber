package times

import (
	"context"

	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

type Repository interface {
	AddTime(ctx context.Context, response timeModel.Response) (int, error)
	GetTime(ctx context.Context, timeID int) (timeModel.DAO, error)
	CountOfTimes(ctx context.Context) (int, error)
	ListTimes(ctx context.Context, paginationData abstract.PaginationData) ([]timeModel.DAO, error)
	DeleteTime(ctx context.Context, timeID int) error
	UpdateTime(ctx context.Context, timeDAO timeModel.DAO) (string, error)
	SelectActiveYear(ctx context.Context) (timeModel.DAO, error)
}

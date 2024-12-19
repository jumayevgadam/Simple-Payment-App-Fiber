package times

import (
	"context"

	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

type Service interface {
	AddTime(ctx context.Context, request timeModel.Request) (int, error)
	GetTime(ctx context.Context, timeID int) (timeModel.DTO, error)
	ListTimes(ctx context.Context, pagination abstract.PaginationQuery) (abstract.PaginatedResponse[timeModel.DTO], error)
	DeleteTime(ctx context.Context, timeID int) error
	UpdateTime(ctx context.Context, timeID int, updateRequest *timeModel.UpdateRequest) (string, error)
	SelectActiveYear(ctx context.Context) (timeModel.DTO, error)
}

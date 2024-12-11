package faculties

import (
	"context"

	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Service interface for performing methods in Service layer.
type Service interface {
	AddFaculty(ctx context.Context, request *facultyModel.Req) (int, error)
	GetFaculty(ctx context.Context, facultyID int) (*facultyModel.DTO, error)
	ListFaculties(ctx context.Context, pagination abstract.PaginationQuery) (
		abstract.PaginatedResponse[*facultyModel.DTO], error)
	DeleteFaculty(ctx context.Context, facultyID int) error
	UpdateFaculty(ctx context.Context, facultyID int, facultyModel *facultyModel.UpdateInputReq) (string, error)
}

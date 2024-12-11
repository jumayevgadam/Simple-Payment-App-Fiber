package faculties

import (
	"context"

	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Repository interface for faculties.
type Repository interface {
	AddFaculty(ctx context.Context, facultyDAO *facultyModel.Res) (int, error)
	GetFaculty(ctx context.Context, facultyID int) (*facultyModel.DAO, error)
	CountFaculties(ctx context.Context) (int, error)
	ListFaculties(ctx context.Context, paginationData abstract.PaginationData) ([]*facultyModel.DAO, error)
	DeleteFaculty(ctx context.Context, facultyID int) error
	UpdateFaculty(ctx context.Context, facultyDAO *facultyModel.DAO) (string, error)
}

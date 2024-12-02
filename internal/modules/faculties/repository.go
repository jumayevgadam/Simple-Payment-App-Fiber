package faculties

import (
	"context"

	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Repository interface for faculties.
type Repository interface {
	AddFaculty(ctx context.Context, facultyDAO *facultyModel.DAO) (int, error)
	GetFaculty(ctx context.Context, facultyID int) (*facultyModel.FacultyData, error)
	CountFaculties(ctx context.Context) (int, error)
	ListFaculties(ctx context.Context, paginationData abstract.PaginationData) ([]*facultyModel.FacultyData, error)
	DeleteFaculty(ctx context.Context, facultyID int) error
	UpdateFaculty(ctx context.Context, facultyDAO *facultyModel.FacultyData) (string, error)
}

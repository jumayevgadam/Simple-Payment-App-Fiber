package faculties

import (
	"context"

	facultyModel "github.com/jumayevgadaym/tsu-toleg/internal/models/faculty"
)

// Repository interface for faculties
type Repository interface {
	AddFaculty(ctx context.Context, facultyDAO facultyModel.DAO) (int, error)
	GetFaculty(ctx context.Context, facultyID int) (facultyModel.DAO, error)
	ListFaculties(ctx context.Context) ([]facultyModel.DAO, error)
	DeleteFaculty(ctx context.Context, facultyID int) error
	UpdateFaculty(ctx context.Context, facultyDAO facultyModel.DAO) (string, error)
}

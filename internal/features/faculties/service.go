package faculties

import (
	"context"

	facultyModel "github.com/jumayevgadaym/tsu-toleg/internal/models/faculty"
)

// Service interface for performing methods in Service layer.
type Service interface {
	AddFaculty(ctx context.Context, facultyDTO facultyModel.DTO) (int, error)
	GetFaculty(ctx context.Context, facultyID int) (facultyModel.DTO, error)
	ListFaculties(ctx context.Context) ([]facultyModel.DTO, error)
	DeleteFaculty(ctx context.Context, facultyID int) error
	UpdateFaculty(ctx context.Context, facultyModel facultyModel.DTO) (string, error)
}

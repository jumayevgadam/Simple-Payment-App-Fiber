package faculties

import (
	"context"

	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
)

// Service interface for performing methods in Service layer.
type Service interface {
	AddFaculty(ctx context.Context, facultyDTO *facultyModel.DTO) (int, error)
	GetFaculty(ctx context.Context, facultyID int) (*facultyModel.Faculty, error)
	ListFaculties(ctx context.Context) ([]*facultyModel.Faculty, error)
	DeleteFaculty(ctx context.Context, facultyID int) error
	UpdateFaculty(ctx context.Context, facultyID int, facultyModel *facultyModel.UpdateInputReq) (string, error)
}

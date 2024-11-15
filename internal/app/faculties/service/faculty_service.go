package service

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/app/faculties"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/database"
	facultyModel "github.com/jumayevgadaym/tsu-toleg/internal/models/faculty"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// Ensure FacultyService implements the faculties.Service interface.
var (
	_ faculties.Service = (*FacultyService)(nil)
)

// FacultyService struct  buisiness logic part for app/faculty part of application.
type FacultyService struct {
	repo database.DataStore
}

// NewFacultyService creates and returns a new instance of FacultyService.
func NewFacultyService(repo database.DataStore) *FacultyService {
	return &FacultyService{repo: repo}
}

// AddFaculty service insert faculty datas into database.
func (s *FacultyService) AddFaculty(ctx context.Context, facultyDTO facultyModel.DTO) (int, error) {
	ctx, span := otel.Tracer("[FacultyService]").Start(ctx, "[AddFaculty]")
	defer span.End()

	roleID, err := s.repo.FacultiesRepo().AddFaculty(ctx, facultyDTO.ToStorage())
	if err != nil {
		tracing.ErrorTracer(span, err)
		return -1, errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "successfully added faculty")
	return roleID, nil
}

// GetFaculty service fetches faculty from DB using identified id.
func (s *FacultyService) GetFaculty(ctx context.Context, facultyID int) (facultyModel.DTO, error) {
	ctx, span := otel.Tracer("[FacultyService]").Start(ctx, "[GetFaculty]")
	defer span.End()

	facultyDAO, err := s.repo.FacultiesRepo().GetFaculty(ctx, facultyID)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return facultyModel.DTO{}, errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "successfully got faculty")
	return facultyDAO.ToServer(), nil
}

// ListFaculties service fetches a list of faculties from DB and returns it.
func (s *FacultyService) ListFaculties(ctx context.Context) ([]facultyModel.DTO, error) {
	ctx, span := otel.Tracer("[FacultyService]").Start(ctx, "[ListFaculties]")
	defer span.End()
	var facultyDTOs []facultyModel.DTO

	faculties, err := s.repo.FacultiesRepo().ListFaculties(ctx)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return nil, errlst.ParseErrors(err)
	}

	for _, facultyRes := range faculties {
		facultyDTOs = append(facultyDTOs, facultyRes.ToServer())
	}

	return facultyDTOs, nil
}

// DeleteFaculty service deletes faculty from DB using identified id.
func (s *FacultyService) DeleteFaculty(ctx context.Context, facultyID int) error {
	ctx, span := otel.Tracer("[FacultyService]").Start(ctx, "[DeleteFaculty]")
	defer span.End()

	if err := s.repo.FacultiesRepo().DeleteFaculty(ctx, facultyID); err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	return nil
}

// UpdateFaculty service updates faculty data using a new faculty data and id.
func (s *FacultyService) UpdateFaculty(ctx context.Context, facultyDTO facultyModel.DTO) (string, error) {
	ctx, span := otel.Tracer("[FacultyService]").Start(ctx, "[UpdateFaculty]")
	defer span.End()

	var updateRes string
	if err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// Check faculty exist in that id
		_, err := db.FacultiesRepo().GetFaculty(ctx, facultyDTO.ID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		updateRes, err = db.FacultiesRepo().UpdateFaculty(ctx, facultyDTO.ToStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	}); err != nil {
		tracing.ErrorTracer(span, err)
		return "", errlst.ParseErrors(err)
	}

	return updateRes, nil
}

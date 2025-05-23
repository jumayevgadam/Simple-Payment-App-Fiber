package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/samber/lo"
)

// Ensure FacultyService implements the faculties.Service interface.
var (
	_ faculties.Service = (*FacultyService)(nil)
)

// FacultyService struct  buisiness logic part for modules/faculty part of moduleslication.
type FacultyService struct {
	repo database.DataStore
}

// NewFacultyService creates and returns a new instance of FacultyService.
func NewFacultyService(repo database.DataStore) *FacultyService {
	return &FacultyService{repo: repo}
}

// AddFaculty service insert faculty datas into database.
func (s *FacultyService) AddFaculty(ctx context.Context, request *facultyModel.Req) (int, error) {
	roleID, err := s.repo.FacultiesRepo().AddFaculty(ctx, request.ToStorage())
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return roleID, nil
}

// GetFaculty service fetches faculty from DB using identified id.
func (s *FacultyService) GetFaculty(ctx context.Context, facultyID int) (*facultyModel.DTO, error) {
	facultyDAO, err := s.repo.FacultiesRepo().GetFaculty(ctx, facultyID)
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return facultyDAO.ToServer(), nil
}

// ListFaculties service fetches a list of faculties from DB and returns it.
func (s *FacultyService) ListFaculties(ctx context.Context, pagination abstract.PaginationQuery) (
	abstract.PaginatedResponse[*facultyModel.DTO], error,
) {
	var (
		facultyAllData      []*facultyModel.DAO
		err                 error
		facultyListResponse abstract.PaginatedResponse[*facultyModel.DTO]
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		var count int

		count, err = db.FacultiesRepo().CountFaculties(ctx)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		facultyListResponse.TotalItems = count

		facultyAllData, err = db.FacultiesRepo().ListFaculties(ctx, pagination.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})
	if err != nil {
		return abstract.PaginatedResponse[*facultyModel.DTO]{}, errlst.ParseErrors(err)
	}

	facultyList := lo.Map(
		facultyAllData,
		func(item *facultyModel.DAO, _ int) *facultyModel.DTO {
			return item.ToServer()
		},
	)

	facultyListResponse.Items = facultyList
	facultyListResponse.CurrentPage = pagination.CurrentPage
	facultyListResponse.Limit = pagination.Limit
	facultyListResponse.ItemsInCurrentPage = len(facultyList)

	return facultyListResponse, nil
}

// DeleteFaculty service deletes faculty from DB using identified id.
func (s *FacultyService) DeleteFaculty(ctx context.Context, facultyID int) error {
	if err := s.repo.FacultiesRepo().DeleteFaculty(ctx, facultyID); err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}

// UpdateFaculty service updates faculty data using a new faculty data and id.
func (s *FacultyService) UpdateFaculty(ctx context.Context, facultyID int, updateReq *facultyModel.UpdateInputReq) (
	string, error,
) {
	var updateRes string
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// Check faculty exist in that id.
		_, err := db.FacultiesRepo().GetFaculty(ctx, facultyID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		updateRes, err = db.FacultiesRepo().UpdateFaculty(ctx, updateReq.ToStorage(facultyID))
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return "", errlst.ParseErrors(err)
	}

	return updateRes, nil
}

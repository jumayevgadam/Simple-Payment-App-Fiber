package time

import "github.com/jumayevgadam/tsu-toleg/pkg/constants"

// Request model.
type Request struct {
	StartYear int `form:"start-year" validate:"required"`
	EndYear   int `form:"end-year" validate:"required"`
}

// Response model.
type Response struct {
	StartYear int `db:"start_year"`
	EndYear   int `db:"end_year"`
}

// ToPsqlDBStorage sends Request model to storage.
func (r *Request) ToPsqlDBStorage() Response {
	return Response{
		StartYear: r.StartYear,
		EndYear:   r.EndYear,
	}
}

// DTO model.
type DTO struct {
	ID        int `json:"id"`
	StartYear int `json:"startYear"`
	EndYear   int `json:"endYear"`
}

// DAO model.
type DAO struct {
	ID        int `db:"id"`
	StartYear int `db:"start_year"`
	EndYear   int `db:"end_year"`
}

// ToServer sends DAO model to server.
func (d *DAO) ToServer() DTO {
	return DTO{
		ID:        d.ID,
		StartYear: d.StartYear,
		EndYear:   d.EndYear,
	}
}

// UpdateRequest model and validate that.
type UpdateRequest struct {
	StartYear int `form:"start-year"`
	EndYear   int `form:"end-year"`
}

func (u *UpdateRequest) Validate() (string, error) {
	if u.StartYear == 0 && u.EndYear == 0 {
		return constants.NoUpdateResponse, nil
	}

	return "", nil
}

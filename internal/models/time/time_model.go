package time

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
	ID        int  `json:"id"`
	StartYear int  `json:"startYear"`
	EndYear   int  `json:"endYear"`
	IsActive  bool `json:"isActiveYear"`
}

// DAO model.
type DAO struct {
	ID        int  `db:"id"`
	StartYear int  `db:"start_year"`
	EndYear   int  `db:"end_year"`
	IsActive  bool `db:"is_active"`
}

// ToServer sends DAO model to server.
func (d *DAO) ToServer() DTO {
	return DTO{
		ID:        d.ID,
		StartYear: d.StartYear,
		EndYear:   d.EndYear,
		IsActive:  d.IsActive,
	}
}

// UpdateRequest model and validate that.
type UpdateRequest struct {
	StartYear int  `form:"start-year,omitempty"`
	EndYear   int  `form:"end-year,omitempty"`
	IsActive  bool `form:"is-active,omitempty"`
}

func (u *UpdateRequest) ToPsqlDBStorage(timeID int) DAO {
	return DAO{
		ID:        timeID,
		StartYear: u.StartYear,
		EndYear:   u.EndYear,
		IsActive:  u.IsActive,
	}
}



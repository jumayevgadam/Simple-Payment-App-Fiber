package time

type AcademicYearRequest struct {
	StartYear int `json:"startYear" validate:"required"`
	EndYear   int `json:"endYear"`
}

type AcademicYearData struct {
	StartYear int `db:"start_year"`
	EndYear   int `db:"end_year"`
}

func (a *AcademicYearRequest) ToPsqlDBStorage() AcademicYearData {
	return AcademicYearData{
		StartYear: a.StartYear,
		EndYear:   a.EndYear,
	}
}

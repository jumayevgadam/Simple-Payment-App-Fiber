package statistics

type AcademicYear struct {
	StartYear int `form:"start-year" validate:"required,gte=2024"`
	EndYear   int `form:"end-year"`
}

type AcademicYearData struct {
	StartYear int `db:"start_year"`
	EndYear   int `db:"end_year"`
}

func (a *AcademicYear) ToStorage() AcademicYearData {
	return AcademicYearData{
		StartYear: a.StartYear,
		EndYear:   a.EndYear,
	}
}

// -------------------------- FIRST SEMESTER STATISTICS ----------------------.
type FirstSemesterPaymentData struct {
	OnlyFirstSemester       int `db:"only_first_semester"`
	BothSemester            int `db:"both_semester"`
	NotPerformFirstSemester int `db:"not_first_semester"`
}

type FirstSemesterPaymentRes struct {
	OnlyFirstSemester       int `json:"onlyFirstSemesterPaid"`
	BothSemester            int `json:"bothSemesterPaid"`
	NotPerformFirstSemester int `json:"notPaidForFirstSemester"`
}

type FirstSemester struct {
}

func (f *FirstSemesterPaymentData) ToServer() FirstSemesterPaymentRes {
	return FirstSemesterPaymentRes{
		OnlyFirstSemester:       f.OnlyFirstSemester,
		BothSemester:            f.BothSemester,
		NotPerformFirstSemester: f.NotPerformFirstSemester,
	}
}

// -------------------------- SECOND SEMESTER STATISTICS ---------------------.

type SecondSemesterPaymentData struct {
	OnlySecondSemester       int `db:"only_second_semester"`
	BothSemester             int `db:"both_semester"`
	NotPerformSecondSemester int `db:"not_second_semester"`
}

type SecondSemesterPaymentRes struct {
	OnlySecondSemester       int `json:"onlySecondSemester"`
	BothSemester             int `json:"bothSemesterPaid"`
	NotPerformSecondSemester int `json:"notPaidSecondSemester"`
}

func (s *SecondSemesterPaymentData) ToServer() SecondSemesterPaymentRes {
	return SecondSemesterPaymentRes{
		OnlySecondSemester:       s.OnlySecondSemester,
		BothSemester:             s.BothSemester,
		NotPerformSecondSemester: s.NotPerformSecondSemester,
	}
}

// -------------------------- ENTIRE YEAR STATISTICS -------------------------.
type EntireYearPaymentData struct {
	FirstSemester       int `db:"first_semester"`
	SecondSemester      int `db:"second_semester"`
	NotPerformedPayment int `db:"not_performed_payment"`
}

type EntireYearPaymentRes struct {
	FirstSemester       int `json:"firstSemesterPaymentsEntireYear"`
	SecondSemester      int `json:"secondSemesterPaymentsEntireYear"`
	NotPerformedPayment int `json:"notPerformedPaymentsEntireYear"`
}

func (e *EntireYearPaymentData) ToServer() EntireYearPaymentRes {
	return EntireYearPaymentRes{
		FirstSemester:       e.FirstSemester,
		SecondSemester:      e.SecondSemester,
		NotPerformedPayment: e.NotPerformedPayment,
	}
}

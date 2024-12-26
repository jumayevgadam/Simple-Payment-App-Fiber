package statistics

// ----------------------- STATISTICS MODELS -------------.

// ----------------------- STATISTICS ABOUT UNIVERSITY ---.

// ----------------------- DAO MODELS --------------------.
type StatisticsAboutUniversityData struct {
	FirstSemesterData  FirstSemesterData  `db:"first_semester_data"`
	SecondSemesterData SecondSemesterData `db:"second_semester_data"`
	BothSemesterData   BothSemesterData   `db:"both_semester_data"`
	TotalStudents      int                `db:"count_total_students"`
}

type FirstSemesterData struct {
	Paid    int `db:"count_first_semester_paid"`
	NotPaid int `db:"count_first_semester_not_paid"`
}

type SecondSemesterData struct {
	Paid    int `db:"count_second_semester_paid"`
	NotPaid int `db:"count_second_semester_not_paid"`
}

type BothSemesterData struct {
	Paid    int `db:"count_both_semester_paid"`
	NotPaid int `db:"count_both_semester_not_paid"`
}

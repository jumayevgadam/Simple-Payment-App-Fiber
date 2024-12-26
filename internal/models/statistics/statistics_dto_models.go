package statistics

// ----------------------- STATISTICS MODELS -------------

// ----------------------- STATISTICS ABOUT UNIVERSITY ---

// ----------------------- DTO MODELS --------------------
type StatisticsAboutUniversity struct {
	FirstSemester  FirstSemester  `json:"firstSemester"`
	SecondSemester SecondSemester `json:"secondSemester"`
	BothSemester   BothSemester   `json:"bothSemester"`
	TotalStudents  int            `json:"totalStudentCount"`
}

type FirstSemester struct {
	Paid    int `json:"studentsFirstSemesterPaid"`
	NotPaid int `json:"studentsFirstSemesterNotPaid"`
}

type SecondSemester struct {
	Paid    int `json:"studentsSecondSemesterPaid"`
	NotPaid int `json:"studentsSecondSemesterNotPaid"`
}

type BothSemester struct {
	Paid    int `json:"studentsBothSemesterPaid"`
	NotPaid int `json:"studentsBothSemesterNotPaid"`
}

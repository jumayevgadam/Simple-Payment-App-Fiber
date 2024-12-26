package statistics

func (f *FirstSemesterData) ToServer() FirstSemester {
	return FirstSemester{
		Paid:    f.Paid,
		NotPaid: f.NotPaid,
	}
}

func (s *SecondSemesterData) ToServer() SecondSemester {
	return SecondSemester{
		Paid:    s.Paid,
		NotPaid: s.NotPaid,
	}
}

func (e *BothSemesterData) ToServer() BothSemester {
	return BothSemester{
		Paid:    e.Paid,
		NotPaid: e.NotPaid,
	}
}

func (d *StatisticsAboutUniversityData) ToServer() StatisticsAboutUniversity {
	return StatisticsAboutUniversity{
		FirstSemester:  d.FirstSemesterData.ToServer(),
		SecondSemester: d.SecondSemesterData.ToServer(),
		BothSemester:   d.BothSemesterData.ToServer(),
		TotalStudents:  d.TotalStudents,
	}
}

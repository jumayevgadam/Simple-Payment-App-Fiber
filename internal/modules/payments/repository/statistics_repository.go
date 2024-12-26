package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/models/statistics"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (r *PaymentRepository) AdminGetStatisticsAboutYear(ctx context.Context, academicYear timeModel.AcademicYearData) (
	statistics.StatisticsAboutUniversityData, error,
) {
	var (
		firstSemesterData  statistics.FirstSemesterData
		secondSemesterData statistics.SecondSemesterData
		bothSemesterData   statistics.BothSemesterData
		totalStudentCount  int
		yearID             int
	)

	// select year id given start year and year.
	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&yearID,
		getAcademicYearQuery,
		academicYear.StartYear,
		academicYear.EndYear,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	// select first semester payments.
	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&firstSemesterData,
		firstSemesterStatisticsQuery,
		yearID,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	// select second semester payments.
	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&secondSemesterData,
		secondSemesterStatisticsQuery,
		yearID,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	// select both semester payments.
	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&bothSemesterData,
		bothSemesterStatisticsQuery,
		yearID,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	// select total student count by university.
	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalStudentCount,
		totalStudentQuery,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	return statistics.StatisticsAboutUniversityData{
		FirstSemesterData:  firstSemesterData,
		SecondSemesterData: secondSemesterData,
		BothSemesterData:   bothSemesterData,
		TotalStudents:      totalStudentCount,
	}, nil
}

func (r *PaymentRepository) AdminGetStatisticsAboutFaculty(ctx context.Context, facultyID int, academicYear timeModel.AcademicYearData) (
	statistics.StatisticsAboutUniversityData, error,
) {
	var (
		firstSemesterData  statistics.FirstSemesterData
		secondSemesterData statistics.SecondSemesterData
		bothSemesterData   statistics.BothSemesterData
		totalStudentCount  int
		yearID             int
	)

	// select year id given start year and year.
	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&yearID,
		getAcademicYearQuery,
		academicYear.StartYear,
		academicYear.EndYear,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	// get statistics about first semester.
	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&firstSemesterData,
		firstSemesterStatisticsByFacultyQuery,
		facultyID,
		yearID,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	// get statistics about second semester.
	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&secondSemesterData,
		secondSemesterStatisticsByFacultyQuery,
		facultyID,
		yearID,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	// get statistics about both semester.
	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&bothSemesterData,
		bothSemesterStatisticsByFacultyQuery,
		facultyID,
		yearID,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	// get total student count by faculty.
	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalStudentCount,
		totalStudentsCountByFacultyQuery,
		facultyID,
	)
	if err != nil {
		return statistics.StatisticsAboutUniversityData{}, errlst.ParseSQLErrors(err)
	}

	return statistics.StatisticsAboutUniversityData{
		FirstSemesterData:  firstSemesterData,
		SecondSemesterData: secondSemesterData,
		BothSemesterData:   bothSemesterData,
		TotalStudents:      totalStudentCount,
	}, nil
}

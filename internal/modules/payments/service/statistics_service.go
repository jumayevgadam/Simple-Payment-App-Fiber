package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/models/statistics"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (s *PaymentService) AdminGetStatisticsAboutYear(ctx context.Context, academicYearReq timeModel.AcademicYearRequest) (
	statistics.StatisticsAboutUniversity, error,
) {
	var (
		statisticsData statistics.StatisticsAboutUniversityData
		err            error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		statisticsData, err = db.PaymentRepo().AdminGetStatisticsAboutYear(ctx, academicYearReq.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return statistics.StatisticsAboutUniversity{}, errlst.ParseErrors(err)
	}

	return statisticsData.ToServer(), nil
}

func (s *PaymentService) AdminGetStatisticsAboutFaculty(ctx context.Context, facultyID int, academicYearReq timeModel.AcademicYearRequest) (
	statistics.StatisticsAboutUniversity, error,
) {
	var (
		statisticsData statistics.StatisticsAboutUniversityData
		err            error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		statisticsData, err = db.PaymentRepo().AdminGetStatisticsAboutFaculty(ctx, facultyID, academicYearReq.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return statistics.StatisticsAboutUniversity{}, errlst.ParseErrors(err)
	}

	return statisticsData.ToServer(), nil
}

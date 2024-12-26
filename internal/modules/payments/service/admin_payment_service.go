package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (s *PaymentService) GetPayment(ctx context.Context, studentID, paymentID int) (*paymentModel.AllPaymentDTO, error) {
	paymentData, err := s.repo.PaymentRepo().GetPaymentByID(ctx, paymentID)
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	if paymentData.StudentID != studentID {
		return nil, errlst.ErrMisMatchedStudentID
	}

	return paymentData.ToServer(), nil
}

func (s *PaymentService) AdminUpdatePaymentOfStudent(ctx context.Context, studentID, paymentID int, paymentStatus string) (
	string, error,
) {
	var (
		updateRes string
		err       error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		paymentData, err := db.PaymentRepo().GetPaymentByID(ctx, paymentID)
		if err != nil {
			return errlst.ErrPaymentNotFound
		}

		if paymentData.StudentID != studentID {
			return errlst.ErrMisMatchedStudentID
		}

		updateRes, err = db.PaymentRepo().AdminUpdatePaymentOfStudent(ctx, studentID, paymentID, paymentStatus)
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

func (s *PaymentService) AdminDeleteStudentPayment(ctx context.Context, studentID, paymentID int) error {
	activeYear, err := s.repo.TimesRepo().SelectActiveYear(ctx)
	if err != nil {
		return errlst.ErrActiveYearNotFound
	}

	err = s.repo.PaymentRepo().AdminDeleteStudentPayment(ctx, studentID, paymentID, activeYear.ID)
	if err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}

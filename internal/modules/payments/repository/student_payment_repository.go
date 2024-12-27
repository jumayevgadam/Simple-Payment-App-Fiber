package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (r *PaymentRepository) AddPayment(ctx context.Context, paymentData paymentModel.Response) (int, error) {
	var paymentID int

	err := r.psqlDB.QueryRow(
		ctx,
		addPaymentQuery,
		paymentData.StudentID,
		paymentData.TimeID,
		paymentData.PaymentType,
		paymentData.PaymentStatus,
		paymentData.CurrentPaidSum,
		paymentData.CheckPhoto,
	).Scan(&paymentID)

	if err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return paymentID, nil
}

func (r *PaymentRepository) CheckType3Payment(ctx context.Context, studentID, timeID int) (bool, int, error) {
	var (
		count      int
		paymentSum int
	)

	err := r.psqlDB.QueryRow(
		ctx,
		checkType3PaymentQuery,
		studentID,
		timeID,
	).Scan(&count)

	if err != nil {
		return false, 0, errlst.ParseSQLErrors(err)
	}

	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&paymentSum,
		totalPaymentSumQuery,
		studentID,
		timeID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, 0, nil
		}

		return false, 0, errlst.ParseSQLErrors(err)
	}

	return count > 0, paymentSum, nil
}

func (r *PaymentRepository) IsPerformedPaymentCheck(ctx context.Context, studentID, timeID int) (bool, int, error) {
	var (
		exists                     bool
		firstSemesterPaymentAmount int
	)

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&exists,
		isPerformedPaymentCheckQuery,
		studentID,
		timeID,
	)

	if err != nil {
		return false, 0, errlst.ParseSQLErrors(err)
	}

	if !exists {
		return false, 0, nil
	}

	if exists {
		err = r.psqlDB.Get(
			ctx,
			r.psqlDB,
			&firstSemesterPaymentAmount,
			firstSemesterPaymentAmountQuery,
			studentID,
			timeID,
		)

		if err != nil {
			// error occured in this place
			return false, 0, errlst.ParseSQLErrors(err)
		}
	}

	return exists, firstSemesterPaymentAmount, nil
}

func (r *PaymentRepository) StudentUpdatePayment(ctx context.Context, paymentData paymentModel.UpdatePaymentData) (string, error) {
	var (
		res string
	)

	err := r.psqlDB.QueryRow(
		ctx,
		studentUpdatePaymentQuery,
		paymentData.PaymentType,
		paymentData.CurrentPaidSum,
		paymentData.CheckPhoto,
		paymentData.StudentID,
		paymentData.PaymentID,
	).Scan(&res)

	if err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return res, nil
}

func (r *PaymentRepository) ListPaymentsByStudent(ctx context.Context, studentID int, paginationData abstract.PaginationData) (
	[]*paymentModel.AllPaymentDAO, userModel.StudentNameAndSurnameData, error,
) {
	var (
		paymentListByStudent []*paymentModel.AllPaymentDAO
		studentData          userModel.StudentNameAndSurnameData
	)
	offset := (paginationData.CurrentPage - 1) * paginationData.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&paymentListByStudent,
		listPaymentsByStudentQuery,
		studentID,
		offset,
		paginationData.Limit,
	)

	if err != nil {
		return nil, userModel.StudentNameAndSurnameData{}, errlst.ParseSQLErrors(err)
	}

	err = r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&studentData,
		studentNameAndSurnameQuery,
		studentID,
	)

	if err != nil {
		return nil, userModel.StudentNameAndSurnameData{}, errlst.ErrStudentNotFound
	}

	return paymentListByStudent, studentData, nil
}

func (r *PaymentRepository) StudentDeletePayment(ctx context.Context, studentID, paymentID, timeID int) error {
	result, err := r.psqlDB.Exec(
		ctx,
		studentDeletePaymentQuery,
		paymentID,
		studentID,
		timeID,
	)

	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	if result.RowsAffected() == 0 {
		return errlst.ErrPaymentNotFound
	}

	return nil
}

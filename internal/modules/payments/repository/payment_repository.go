package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

var _ payments.Repository = (*PaymentRepository)(nil)

type PaymentRepository struct {
	psqlDB connection.DB
}

func NewPaymentRepository(psqlDB connection.DB) *PaymentRepository {
	return &PaymentRepository{psqlDB: psqlDB}
}

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

func (r *PaymentRepository) GetStudentInfoForPayment(ctx context.Context, studentID int) (*paymentModel.StudentInfoForPayment, error) {
	var studentDataForPayment paymentModel.StudentInfoForPayment

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&studentDataForPayment,
		studentInfoForPaymentQuery,
		studentID,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &studentDataForPayment, nil
}

func (r *PaymentRepository) GetPaymentByID(ctx context.Context, paymentID int) (*paymentModel.AllPaymentDAO, error) {
	var paymentData paymentModel.AllPaymentDAO

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&paymentData,
		getPaymentByIDQuery,
		paymentID,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &paymentData, nil
}

func (r *PaymentRepository) CountPaymentByStudent(ctx context.Context, studentID int) (int, error) {
	var totalCount int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalCount,
		countPaymentsByStudentQuery,
		studentID,
	)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalCount, nil
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

func (r *PaymentRepository) AdminGetPaymentStatusOfStudent(ctx context.Context, studentID, paymentID int) (string, error) {
	var currentStatus string

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&currentStatus,
		adminGetPaymentStatusQuery,
		studentID,
		paymentID,
	)

	if err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return currentStatus, nil
}

func (r *PaymentRepository) AdminUpdatePaymentOfStudent(ctx context.Context, studentID, paymentID int, paymentStatus string) (
	string, error,
) {
	var updatedRes string

	err := r.psqlDB.QueryRow(
		ctx,
		adminUpdatePaymentStatusQuery,
		paymentStatus,
		studentID,
		paymentID,
	).Scan(&updatedRes)

	if err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return updatedRes, nil
}

func (r *PaymentRepository) ListPaymentsByStudentID(ctx context.Context, studentID int) ([]*paymentModel.PaymentsByStudentID, error) {
	var paymentsByStudentID []*paymentModel.PaymentsByStudentID

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&paymentsByStudentID,
		paymentsByStudentIDQuery,
		studentID,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return paymentsByStudentID, nil
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
			return false, 0, errlst.ParseSQLErrors(err)
		}
	}

	return exists, firstSemesterPaymentAmount, nil
}

func (r *PaymentRepository) CurrentPaymentAmount(ctx context.Context, studentID, timeID int) (int, error) {
	var totalPaymentAmount int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		totalPaymentAmount,
		`SELECT
			 SUM(payment_amount) FROM payments 
		WHERE student_id = $1 AND time_id = $2
		GROUP BY student_id, time_id;`,
		studentID, timeID,
	)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalPaymentAmount, nil
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

func (r *PaymentRepository) AdminDeleteStudentPayment(ctx context.Context, studentID, paymentID, timeID int) error {
	result, err := r.psqlDB.Exec(
		ctx,
		adminDeleteStudentPayment,
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

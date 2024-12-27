package repository

import (
	"context"

	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

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

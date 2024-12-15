package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
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

func (r *PaymentRepository) AddPayment(ctx context.Context, paymentData *paymentModel.Response) (int, error) {
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
	[]*paymentModel.AllPaymentDAO, error,
) {
	var paymentListByStudent []*paymentModel.AllPaymentDAO
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
		return nil, errlst.ParseSQLErrors(err)
	}

	return paymentListByStudent, nil
}

func (r *PaymentRepository) StudentUpdatePayment(ctx context.Context, paymentData paymentModel.UpdatePaymentData) (string, error) {
	var res string

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

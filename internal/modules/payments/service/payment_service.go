package service

import (
	"context"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
	"github.com/samber/lo"
)

var _ payments.Service = (*PaymentService)(nil)

type PaymentService struct {
	repo database.DataStore
}

func NewPaymentService(repo database.DataStore) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) AddPayment(ctx *fiber.Ctx, checkPhoto *multipart.FileHeader, studentID int, paymentReq *paymentModel.Request) (
	int, error,
) {
	var (
		paymentID int
		err       error
	)

	err = s.repo.WithTransaction(ctx.Context(), func(db database.DataStore) error {
		studentDataForPayment, err := db.PaymentRepo().GetStudentInfoForPayment(ctx.Context(), studentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		checkPhotoURL, err := utils.SaveImage(
			ctx, checkPhoto, studentDataForPayment.FacultyName, studentDataForPayment.GroupCode,
			studentDataForPayment.FullName, studentDataForPayment.Username, paymentReq.PaymentType,
		)

		if err != nil {
			return errlst.ParseErrors(err)
		}

		paymentID, err = db.PaymentRepo().AddPayment(ctx.Context(), paymentReq.ToPsqlDBStorage(studentID, checkPhotoURL))
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return paymentID, nil
}

func (s *PaymentService) GetPayment(ctx context.Context, studentID, paymentID int) (*paymentModel.AllPaymentDTO, error) {
	paymentData, err := s.repo.PaymentRepo().GetPaymentByID(ctx, paymentID)
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	if paymentData.StudentID != studentID {
		return nil, errlst.NewBadRequestError(
			"studentID's mismatched: [paymentService][GetPaymentByID]",
		)
	}

	return paymentData.ToServer(), nil
}

func (s *PaymentService) ListPaymentsByStudent(ctx context.Context, studentID int, paginationQuery abstract.PaginationQuery) (
	abstract.PaginatedResponse[*paymentModel.AllPaymentDTO], error,
) {
	var (
		allPaymentListDataByStudent   []*paymentModel.AllPaymentDAO
		paymentListResponseByStudent  abstract.PaginatedResponse[*paymentModel.AllPaymentDTO]
		totalCountOfPaymentsByStudent int
		err                           error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err = db.UsersRepo().GetStudent(ctx, studentID)
		if err != nil {
			return errlst.NewNotFoundError("student can not found: [paymentService][ListPaymentsByStudent]")
		}

		totalCountOfPaymentsByStudent, err = db.PaymentRepo().CountPaymentByStudent(ctx, studentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		paymentListResponseByStudent.TotalItems = totalCountOfPaymentsByStudent

		allPaymentListDataByStudent, err = db.PaymentRepo().ListPaymentsByStudent(ctx, studentID, paginationQuery.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return abstract.PaginatedResponse[*paymentModel.AllPaymentDTO]{}, errlst.ParseErrors(err)
	}

	paymentListByStudent := lo.Map(
		allPaymentListDataByStudent,
		func(item *paymentModel.AllPaymentDAO, _ int) *paymentModel.AllPaymentDTO {
			return item.ToServer()
		},
	)

	paymentListResponseByStudent.Items = paymentListByStudent
	paymentListResponseByStudent.CurrentPage = paginationQuery.CurrentPage
	paymentListResponseByStudent.Limit = paginationQuery.Limit
	paymentListResponseByStudent.ItemsInCurrentPage = len(paymentListByStudent)

	return paymentListResponseByStudent, nil
}

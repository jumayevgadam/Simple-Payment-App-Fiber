package service

import (
	"context"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/samber/lo"
)

// Ensuring PaymentService implements payment.Service interface.
var (
	_ payment.Service = (*PaymentService)(nil)
)

// PaymentService for logic in Payment.
type PaymentService struct {
	repo database.DataStore
}

// NewPaymentService creates and returns a new instance of PaymentService.
func NewPaymentService(repo database.DataStore) *PaymentService {
	return &PaymentService{repo: repo}
}

// AddPayment service method for adding payment.
func (p *PaymentService) AddPayment(c *fiber.Ctx, studentID int, checkPhoto *multipart.FileHeader, request *paymentModel.Request) (
	int, error,
) {
	var (
		paymentID int
		err       error
	)

	// perform adding payment with transaction.
	err = p.repo.WithTransaction(c.Context(), func(db database.DataStore) error {
		return nil
	})
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return paymentID, nil
}

// GetPaymentByID service.
func (p *PaymentService) GetPaymentByID(ctx context.Context, paymentID int) (*paymentModel.AllPaymentDTO, error) {
	paymentDAO, err := p.repo.PaymentsRepo().GetPaymentByID(ctx, paymentID)
	if err != nil {
		return nil, errlst.NewNotFoundError("payment not found" + err.Error())
	}

	return paymentDAO.ToServer(), nil
}

func (p *PaymentService) StudentListPaymentsByStudentID(ctx context.Context, studentID int, paginationQuery abstract.PaginationQuery) (
	abstract.PaginatedResponse[*paymentModel.AllPaymentDTO], error,
) {
	var (
		studentPaymentDatas         []*paymentModel.AllPaymentDAO
		studentPaymentsListResponse abstract.PaginatedResponse[*paymentModel.AllPaymentDTO]
		totalStudentPaymentCount    int
		err                         error
	)

	err = p.repo.WithTransaction(ctx, func(db database.DataStore) error {
		totalStudentPaymentCount, err = db.PaymentsRepo().CountPaymentsByStudentID(ctx, studentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		studentPaymentsListResponse.TotalItems = totalStudentPaymentCount

		studentPaymentDatas, err = db.PaymentsRepo().StudentListPaymentsByStudentID(ctx, studentID, paginationQuery.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return abstract.PaginatedResponse[*paymentModel.AllPaymentDTO]{}, errlst.ParseErrors(err)
	}

	studentPaymentList := lo.Map(
		studentPaymentDatas,
		func(item *paymentModel.AllPaymentDAO, _ int) *paymentModel.AllPaymentDTO {
			return item.ToServer()
		},
	)

	studentPaymentsListResponse.Items = studentPaymentList
	studentPaymentsListResponse.CurrentPage = paginationQuery.CurrentPage
	studentPaymentsListResponse.Limit = paginationQuery.Limit
	studentPaymentsListResponse.ItemsInCurrentPage = len(studentPaymentList)

	return studentPaymentsListResponse, nil
}

package service

import (
	"context"
	"mime/multipart"
	"os"

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
		// Select active year timeID.
		activeYear, err := db.TimesRepo().SelectActiveYear(ctx.Context())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		existingPayment, err := db.PaymentRepo().CheckType3Payment(ctx.Context(), studentID, activeYear.ID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		if existingPayment {
			return errlst.NewBadRequestError("You cannot perform payment for type 3, because this action has already been performed.")
		}

		// implement logic here for student can do action with payment_type 1 and 2.
		paymentsByStudentID, err := db.PaymentRepo().ListPaymentsByStudentID(ctx.Context(), studentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		for _, payment := range paymentsByStudentID {
			if payment.PaymentType == paymentReq.PaymentType {
				return errlst.NewBadRequestError("already u performed some payments for this year")
			}
		}

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

		paymentID, err = db.PaymentRepo().AddPayment(ctx.Context(), paymentReq.ToPsqlDBStorage(studentID, activeYear.ID, checkPhotoURL))
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
	var (
		paymentData *paymentModel.AllPaymentDAO
		err         error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		paymentData, err = db.PaymentRepo().GetPaymentByID(ctx, paymentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		if paymentData.StudentID != studentID {
			return errlst.NewBadRequestError(
				"studentID's mismatched: [paymentService][GetPaymentByID]",
			)
		}

		return nil
	})

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

func (s *PaymentService) StudentUpdatePayment(c *fiber.Ctx, studentID, paymentID int, updateReq paymentModel.UpdatePaymentRequest) (
	string, error,
) {
	var (
		updateResponse string
		checkPhotoURL  string
		err            error
	)

	err = s.repo.WithTransaction(c.Context(), func(db database.DataStore) error {
		paymentData, err := db.PaymentRepo().GetPaymentByID(c.Context(), paymentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		if paymentData.StudentID != studentID {
			return errlst.NewBadRequestError("studentID mismatch with payment record")
		}

		studentInfo, err := db.PaymentRepo().GetStudentInfoForPayment(c.Context(), studentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		photo, _ := utils.ReadImage(c, "check-photo")
		if photo == nil {
			checkPhotoURL = paymentData.CheckPhoto
		}

		if photo != nil {
			// Remove the old photo.
			if paymentData.CheckPhoto != "" {
				os.Remove("./uploads/" + paymentData.CheckPhoto)
			}

			checkPhotoURL, err = utils.SaveImage(
				c, photo, studentInfo.FacultyName, studentInfo.GroupCode,
				studentInfo.FullName, studentInfo.Username, updateReq.PaymentType,
			)

			if err != nil {
				return errlst.ParseErrors(err)
			}

			updateReq.PhotoURL = checkPhotoURL
		}

		updateResponse, err = db.PaymentRepo().StudentUpdatePayment(c.Context(), updateReq.ToPsqlDBStorage(
			studentID, paymentID))

		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return "", errlst.ParseErrors(err)
	}

	return updateResponse, nil
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
			return errlst.NewBadRequestError(err)
		}

		if paymentData.StudentID != studentID {
			return errlst.NewBadRequestError("mismatched studentID in [paymentService][AdminUpdatePaymentOfStudent]")
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

package service

import (
	"context"
	"errors"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/helpers"
	"github.com/samber/lo"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
)

func (s *PaymentService) AddPayment(ctx *fiber.Ctx, checkPhoto *multipart.FileHeader, studentID int, paymentReq paymentModel.Request) (
	int, error,
) {
	var (
		paymentID int
		// currentPaymentAmount int
		err error
	)

	err = s.repo.WithTransaction(ctx.Context(), func(db database.DataStore) error {
		// Select active year.
		activeYear, err := db.TimesRepo().SelectActiveYear(ctx.Context())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		existingPayment, totalPaymentSum, err := db.PaymentRepo().CheckType3Payment(
			ctx.Context(), studentID, activeYear.ID)

		if err != nil {
			return errlst.ParseErrors(err)
		}

		if existingPayment || totalPaymentSum >= constants.FullPaymentPrice {
			return errlst.ErrPaymentPerform
		}

		isPerformedPaymentForFirstSemester, firstSemesterPaymentAmount, err := db.PaymentRepo().IsPerformedPaymentCheck(
			ctx.Context(), studentID, activeYear.ID)

		if err != nil {
			return errlst.ParseErrors(err)
		}

		if err := helpers.CheckPayment(paymentReq, firstSemesterPaymentAmount, isPerformedPaymentForFirstSemester); err != nil {
			return err
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

func (s *PaymentService) ListPaymentsByStudent(ctx context.Context, studentID int, paginationQuery abstract.PaginationQuery) (
	abstract.PaginatedResponse[*paymentModel.AllPaymentDTO], userModel.StudentNameAndSurname, error,
) {
	var (
		allPaymentListDataByStudent   []*paymentModel.AllPaymentDAO
		paymentListResponseByStudent  abstract.PaginatedResponse[*paymentModel.AllPaymentDTO]
		totalCountOfPaymentsByStudent int
		studentData                   userModel.StudentNameAndSurnameData
		err                           error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err = db.UsersRepo().GetStudent(ctx, studentID)
		if err != nil {
			return errlst.ErrStudentNotFound
		}

		totalCountOfPaymentsByStudent, err = db.PaymentRepo().CountPaymentByStudent(ctx, studentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		paymentListResponseByStudent.TotalItems = totalCountOfPaymentsByStudent

		allPaymentListDataByStudent, studentData, err = db.PaymentRepo().ListPaymentsByStudent(ctx, studentID, paginationQuery.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return abstract.PaginatedResponse[*paymentModel.AllPaymentDTO]{}, userModel.StudentNameAndSurname{}, errlst.ParseErrors(err)
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

	return paymentListResponseByStudent, studentData.ToServer(), nil
}

func (s *PaymentService) StudentUpdatePayment(c *fiber.Ctx, studentID, paymentID int, updateReq paymentModel.UpdatePaymentRequest) (
	string, error,
) {
	var (
		updateResponse string
		checkPhotoURL  string
		// paymentCount   int
		err error
	)

	err = s.repo.WithTransaction(c.Context(), func(db database.DataStore) error {
		paymentData, err := db.PaymentRepo().GetPaymentByID(c.Context(), paymentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		if paymentData.StudentID != studentID {
			return errlst.NewBadRequestError("studentID mismatch with payment record")
		}

		_, err = db.PaymentRepo().CountPaymentByStudent(c.Context(), studentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		studentInfo, err := db.PaymentRepo().GetStudentInfoForPayment(c.Context(), studentID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		photo, err := utils.ReadImage(c, "check-photo")
		if err != nil && !errors.Is(err, errlst.ErrNoUploadedFile) {
			return errlst.ParseErrors(err)
		}

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

		// if err := helpers.UpdatePaymentChecker(paymentData, paymentCount, updateReq); err != nil {
		// 	return err
		// }

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

func (s *PaymentService) StudentDeletePayment(ctx context.Context, studentID, paymentID int) error {
	activeYear, err := s.repo.TimesRepo().SelectActiveYear(ctx)
	if err != nil {
		return errlst.ParseErrors(err)
	}

	err = s.repo.PaymentRepo().StudentDeletePayment(ctx, studentID, paymentID, activeYear.ID)
	if err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}

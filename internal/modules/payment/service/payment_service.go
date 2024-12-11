package service

import (
	"log"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
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
		// get name, surname, courseYear and groupCode using studentID.
		studentDataForPayment, err := db.UsersRepo().GetStudentDetailsForPayment(c.Context(), studentID)
		if err != nil {
			log.Println(err)
			return errlst.ParseErrors(err)
		}

		// save image dynamic folder using groupCode, name_surname
		checkPhotoURL, err := utils.SaveImage(
			c, checkPhoto,
			studentDataForPayment.GroupCode,
			studentDataForPayment.FullName,
			studentDataForPayment.Username,
		)
		if err != nil {
			return errlst.NewBadRequestError("can not save check photo that directory")
		}

		// save into payments table using studentID, courseYear, and return response
		paymentID, err = db.PaymentsRepo().AddPayment(c.Context(), request.ToPsqlDBStorage(studentID, checkPhotoURL))
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

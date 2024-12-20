package helpers

import (
	"github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func CheckPayment(request payment.Request, firstSemesterPaymentAmount int, firstSemesterPayment bool) error {
	fullPrice := constants.FullPaymentPrice
	minimum := constants.AtLeastPaymentPrice

	switch firstSemesterPayment {
	case true:
		if request.PaymentType != "2" || request.PaymentType == "3" {
			return errlst.NewBadRequestError(
				"u performed first semester payment, please perform 2nd semester payment",
			)
		}

		if request.PaymentType == "2" && request.CurrentPaidSum != fullPrice-firstSemesterPaymentAmount {
			return errlst.NewBadRequestError(
				"u performed first semester payment, unnecessary payment implemented for second semester",
			)
		}

	case false:
		if request.PaymentType == "1" && request.CurrentPaidSum < minimum || request.CurrentPaidSum > fullPrice {
			return errlst.NewBadRequestError(
				"can not perform first semester payment, unnecessary payment balance for first semester payment",
			)
		}

		if request.PaymentType == "2" {
			return errlst.NewBadRequestError(
				"u did not perform first semester payment,u can not perform second semester payment, please do for first semester or perform full payment",
			)
		}

		if request.PaymentType == "3" && request.CurrentPaidSum != fullPrice {
			return errlst.NewBadRequestError(
				"please implement true payment balance for full payment",
			)
		}

	default:
		errlst.NewBadRequestError("u implement wrong payment_type")
	}

	return nil
}

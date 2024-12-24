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
			return errlst.ErrDidNotPerformFullPayment
		}

		if request.PaymentType == "2" && request.CurrentPaidSum < fullPrice-firstSemesterPaymentAmount {
			return errlst.ErrSecondSemesterPayment
		}

	case false:
		if request.PaymentType == "1" && request.CurrentPaidSum < minimum || request.CurrentPaidSum > fullPrice {
			return errlst.ErrFirstSemesterPayment
		}

		if request.PaymentType == "2" {
			return errlst.ErrDidNotPerformPayment
		}

		if request.PaymentType == "3" && request.CurrentPaidSum < fullPrice {
			return errlst.ErrFullPayment
		}

	default:
		return errlst.ErrInPaymentType
	}

	return nil
}

package helpers

import (
	"github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func CheckPayment(request payment.Request) error {
	fullPrice := constants.FullPaymentPrice
	minimum := constants.AtLeastPaymentPrice

	if request.PaymentType == "1" && request.CurrentPaidSum >= fullPrice && request.CurrentPaidSum < minimum {
		return errlst.NewBadRequestError(
			"error: can not perform payment for 1 type, incorrect paid sum",
		)
	}

	if request.PaymentType == "2" && request.CurrentPaidSum >= fullPrice {
		return errlst.NewBadRequestError(
			"error: can not perform payment for 2 type, incorrect paid sum",
		)
	}

	if request.PaymentType == "3" && request.CurrentPaidSum != fullPrice {
		return errlst.NewBadRequestError(
			"error: can not perform payment for 3 type, incorrect paid sum",
		)
	}

	return nil
}

package helpers

import (
	"errors"

	"github.com/jumayevgadam/tsu-toleg/internal/models/payment"
)

func CheckerFuncForUpdate(paymentData *payment.AllPaymentDAO, updateReq payment.UpdatePaymentRequest) error {
	switch paymentData.PaymentType {
	case "1":
		if updateReq.PaymentType == "2" {
			if updateReq.CurrentPaidSum == 0 && paymentData.CurrentPaidSum < someNumber {
				return errors.New("err1")
			}
		}

		// if updateReq.PaymentType == "3" {
		// 	if updateReq.CurrentPaidSum < fullPrice
		// }
	}

	return nil
}

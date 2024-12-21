package helpers

import (
	"errors"

	"github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
)

// UpdatePaymentCheck method helps to us updating fields of payments table.
func UpdatePayment(request *payment.UpdatePaymentRequest, currentPaymentType string, paymentCount int, currentBalance int) error {
	fullPrice := constants.FullPaymentPrice

	switch paymentCount {
	case 2:
		if request.PaymentType != "" && request.PaymentType == "3" {
			return errors.New("you can not update payment type to full type")
		}
	case 1:
		switch currentPaymentType {
		case "1":
			if request.PaymentType == "2" {
				return errors.New("you can not update payment")
			}
		}

		if currentPaymentType == "3" && request.PaymentType != "3" {
			return errors.New("")
		}

		if currentPaymentType == "1" && request.PaymentType == "3" && request.CurrentPaidSum != fullPrice {
			return errors.New("current payment type 1, for updating to full payment, current paid sum is not equal to full price")
		}

		if request.PaymentType == "2" && currentPaymentType == "3" {
			return errors.New("can not update payment ")
		}
	}

	return nil
}

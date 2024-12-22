package helpers

import (
	"errors"

	"github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
)

func UpdatePaymentChecker(paymentData payment.UpdatePaymentData, paymentTypes []string) error {
	for _, pType := range paymentTypes {
		switch len(pType) {
		case 2:
			if paymentData.PaymentType == "3" {
				return errors.New("can not update payment for full update, two times performed payment")
			}

			switch pType {
			case "1":
				if paymentData.CurrentPaidSum >= constants.FullPaymentPrice || paymentData.CurrentPaidSum < constants.AtLeastPaymentPrice {
					return errors.New("wrong payment amount implemented for first semester payment update")
				}

			case "2":
				if paymentData.PaymentType == "3" {
					return errors.New("can not update for full payment, current payment type for second semester")
				}
			}

		case 1:
			switch pType {
			case "1":
				if paymentData.PaymentType == "3" && paymentData.CurrentPaidSum != constants.FullPaymentPrice {
					return errors.New("wrong payment amount implemented for full payment in update payment")
				}

				if paymentData.PaymentType == "" && paymentData.CurrentPaidSum < constants.AtLeastPaymentPrice || paymentData.CurrentPaidSum >= constants.FullPaymentPrice {
					return errors.New("current payment type 1, wrong payment balance implemented for first semester in update")
				}

			case "3":
				if paymentData.PaymentType == "2" {
					return errors.New("current payment type 3, can not update payment for second semester")
				}

				if paymentData.PaymentStatus == "1" {
					if paymentData.CurrentPaidSum == 0 {
						return errors.New("current payment type 3, for updating payment to first semester, please implement payment balance")
					} else {
						if paymentData.CurrentPaidSum < constants.AtLeastPaymentPrice || paymentData.CurrentPaidSum >= constants.FullPaymentPrice {
							return errors.New("current payment type 3, wrong payment amount implemented for first semester in update")
						}
					}
				}
			}

		}
	}

	return nil
}

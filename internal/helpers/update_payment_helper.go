package helpers

import (
	"errors"

	"github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
)

// constants.
const (
	fullPrice                  = constants.FullPaymentPrice
	atLeastPrice               = constants.AtLeastPaymentPrice
	secondSemesterAtLeastPrice = fullPrice - atLeastPrice
)

var someNumber int = 3000

// errors.
var (
	ErrInvalidBalance = errors.New("birinji semestr tölegi üçin nädogry balans")
)

// UpdatePaymentChecker helper function need for us validating request before calling database.
func UpdatePaymentChecker(paymentData *payment.AllPaymentDAO, paymentCount int, updateRequest payment.UpdatePaymentRequest) error {
	switch paymentCount {
	case 2:
		if updateRequest.PaymentType == "3" {
			// return error when updating payment for full payment.
			return errors.New("can not update to payment full payment")
		}

		switch paymentData.PaymentType {
		case "1":
			if updateRequest.PaymentType == "" && updateRequest.CurrentPaidSum < atLeastPrice {
				return errors.New("invalid balance implemented for payment type 1")
			}

			if updateRequest.PaymentType == "1" && updateRequest.CurrentPaidSum == 0 {
				if paymentData.CurrentPaidSum < atLeastPrice {
					return errors.New("invalid balance for first semester payment")
				}
			}

		case "2":
			if updateRequest.PaymentType == "" && updateRequest.CurrentPaidSum < secondSemesterAtLeastPrice {
				return errors.New("invalid balance for second semester")
			}

			if updateRequest.PaymentType == "2" && updateRequest.CurrentPaidSum == 0 {
				if paymentData.CurrentPaidSum < fullPrice-atLeastPrice {
					return errors.New("invalid balance implemented for second semester")
				}
			}

			if updateRequest.PaymentType == "2" && updateRequest.CurrentPaidSum < secondSemesterAtLeastPrice {
				return errors.New("error occured when performing second semester payment")
			}
		default:
			return errors.New("error occured unknown type")
		}
	case 1:
		switch paymentData.PaymentType {
		case "3":
			if updateRequest.PaymentType == "" && updateRequest.CurrentPaidSum < fullPrice {
				return errors.New("error occured when performing full payment")
			}

			if updateRequest.PaymentType == "3" && updateRequest.CurrentPaidSum == 0 {
				if paymentData.CurrentPaidSum < fullPrice {
					return errors.New("invalid balance implemented in database")
				}
			}

			if updateRequest.PaymentType == "3" && updateRequest.CurrentPaidSum < fullPrice {
				return errors.New("error occured when performing full payment")
			}

			if updateRequest.PaymentType == "2" && updateRequest.CurrentPaidSum < secondSemesterAtLeastPrice {
				return errors.New("error occured when implementing second semester payment when current payment full")
			}

			if updateRequest.PaymentType == "2" && updateRequest.CurrentPaidSum == 0 {
				if paymentData.CurrentPaidSum < secondSemesterAtLeastPrice {
					return errors.New("error occured when performing second semester payment")
				}
			}

			if updateRequest.PaymentType == "1" && updateRequest.CurrentPaidSum < atLeastPrice {
				return errors.New("error occured when performing update for first semester when performed full payment")
			}

			if updateRequest.PaymentType == "1" && updateRequest.CurrentPaidSum == 0 {
				if paymentData.CurrentPaidSum < atLeastPrice {
					return errors.New("error occured when performing action hereeeeeeee")
				}
			}
		case "1":
			if updateRequest.PaymentType == "3" {
				if updateRequest.CurrentPaidSum == 0 && paymentData.CurrentPaidSum < fullPrice {
					return errors.New("err12")
				}

				if updateRequest.CurrentPaidSum < fullPrice {
					return errors.New("error occured when performing payment when case 1")
				}
			}

			if updateRequest.PaymentType == "2" {
				if updateRequest.CurrentPaidSum == 0 && paymentData.CurrentPaidSum < someNumber {
					return errors.New("error occured when performing changing payment to second semester type")
				}

				if updateRequest.CurrentPaidSum < someNumber {
					return errors.New("error occured when performing aaaaaaaaaaaaaaaaaaa")
				}
			}

			if updateRequest.PaymentType == "" && updateRequest.CurrentPaidSum < atLeastPrice {
				return errors.New("err134")
			}
		}

	}

	return nil
}

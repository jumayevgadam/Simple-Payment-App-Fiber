package payment

// Request model for payment.
type Request struct {
	PaymentType    string `form:"payment-type" json:"paymentType" validate:"required"`
	CheckPhoto     string `form:"check-photo" json:"checkPhoto"`
	CurrentPaidSum int    `form:"current-balance" json:"currentPayedBalance" validate:"required"`
}

// Response model for taking request into Storage.
type Response struct {
	StudentID      int    `db:"student_id"`
	CheckPhoto     string `db:"check_photo"`
	PaymentType    string `db:"payment_type"`
	CurrentPaidSum int    `db:"payment_amount"`
	PaymentStatus  string `db:"payment_status"`
}

// ToDBStorage sends request to psqlDB in our case.
func (r *Request) ToPsqlDBStorage(studentID int, photoURL string) *Response {
	return &Response{
		StudentID:      studentID,
		CheckPhoto:     photoURL,
		PaymentType:    r.PaymentType,
		CurrentPaidSum: r.CurrentPaidSum,
	}
}

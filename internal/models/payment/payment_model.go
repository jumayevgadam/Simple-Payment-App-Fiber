package payment

import "time"

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
	TimeID         int    `db:"time_id"`
}

// ToPsqlDBStorage sends request to psqlDB in our case.
func (r *Request) ToPsqlDBStorage(studentID int, photoURL string) *Response {
	return &Response{
		StudentID:      studentID,
		CheckPhoto:     photoURL,
		PaymentType:    r.PaymentType,
		CurrentPaidSum: r.CurrentPaidSum,
	}
}

// AllPaymentDAO model.
type AllPaymentDAO struct {
	ID             int       `db:"payment_id"`
	StudentID      int       `db:"student_id"`
	StartYear      int       `db:"start_year"`
	EndYear        int       `db:"end_year"`
	PaymentType    string    `db:"payment_type"`
	CheckPhoto     string    `db:"check_photo"`
	CurrentPaidSum int       `db:"payment_amount"`
	PaymentStatus  string    `db:"payment_status"`
	UploadedAt     time.Time `db:"uploaded_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// AllPaymentDTO model.
type AllPaymentDTO struct {
	ID             int       `json:"paymentID"`
	StudentID      int       `json:"studentID"`
	StartYear      int       `json:"startYear"`
	EndYear        int       `json:"endYear"`
	PaymentType    string    `json:"paymentType"`
	CheckPhoto     string    `json:"checkPhotoURL"`
	CurrentPaidSum int       `json:"paymentAmount"`
	PaymentStatus  string    `json:"paymentStatus"`
	UploadedAt     time.Time `json:"uploadedAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (a *AllPaymentDAO) ToServer() *AllPaymentDTO {
	return &AllPaymentDTO{
		ID:             a.ID,
		StudentID:      a.StudentID,
		StartYear:      a.StartYear,
		EndYear:        a.EndYear,
		PaymentType:    a.PaymentType,
		CheckPhoto:     a.CheckPhoto,
		CurrentPaidSum: a.CurrentPaidSum,
		PaymentStatus:  a.PaymentStatus,
		UploadedAt:     a.UploadedAt,
		UpdatedAt:      a.UpdatedAt,
	}
}

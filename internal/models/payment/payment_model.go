package payment

import (
	"time"

	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
)

// Request model for payment.
type Request struct {
	PaymentType    string `form:"payment-type" json:"paymentType" validate:"required"`
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
func (r *Request) ToPsqlDBStorage(studentID, timeID int, photoURL string) *Response {
	return &Response{
		StudentID:      studentID,
		TimeID:         timeID,
		CheckPhoto:     photoURL,
		PaymentType:    r.PaymentType,
		CurrentPaidSum: r.CurrentPaidSum,
	}
}

// AllPaymentDAO model.
type AllPaymentDAO struct {
	ID             int       `db:"id"`
	StudentID      int       `db:"student_id"`
	PaymentType    string    `db:"payment_type"`
	CourseYear     int       `db:"course_year"`
	CheckPhoto     string    `db:"check_photo"`
	CurrentPaidSum int       `db:"payment_amount"`
	PaymentStatus  string    `db:"payment_status"`
	StartYear      int       `db:"start_year"`
	EndYear        int       `db:"end_year"`
	UploadedAt     time.Time `db:"uploaded_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// AllPaymentDTO model.
type AllPaymentDTO struct {
	ID             int       `json:"paymentID"`
	StudentID      int       `json:"studentID"`
	PaymentType    string    `json:"paymentType"`
	CourseYear     int       `json:"courseYear"`
	CheckPhoto     string    `json:"checkPhotoURL"`
	CurrentPaidSum int       `json:"paymentAmount"`
	PaymentStatus  string    `json:"paymentStatus"`
	StartYear      int       `json:"startYear"`
	EndYear        int       `json:"endYear"`
	UploadedAt     time.Time `json:"uploadedAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (a *AllPaymentDAO) ToServer() *AllPaymentDTO {
	return &AllPaymentDTO{
		ID:             a.ID,
		StudentID:      a.StudentID,
		PaymentType:    a.PaymentType,
		CourseYear:     a.CourseYear,
		CheckPhoto:     a.CheckPhoto,
		CurrentPaidSum: a.CurrentPaidSum,
		PaymentStatus:  a.PaymentStatus,
		StartYear:      a.StartYear,
		EndYear:        a.EndYear,
		UploadedAt:     a.UploadedAt,
		UpdatedAt:      a.UpdatedAt,
	}
}

type StudentInfoForPayment struct {
	StudentID   int    `db:"id"`
	RoleID      int    `db:"role_id"`
	FacultyName string `db:"faculty_name"`
	GroupCode   string `db:"group_code"`
	FullName    string `db:"full_name"`
	Username    string `db:"username"`
}

type UpdatePaymentRequest struct {
	CurrentPaidSum int    `form:"current-paid-sum"`
	PaymentType    string `form:"payment-type"`
	PhotoURL       string
}

type UpdatePaymentData struct {
	PaymentID      int    `db:"id"`
	StudentID      int    `db:"student_id"`
	TimeID         int    `db:"time_id"`
	CheckPhoto     string `db:"check_photo"`
	CurrentPaidSum int    `db:"current_balance"`
	PaymentType    string `db:"payment_type"`
	PaymentStatus  string `db:"payment_status"`
}

func (u *UpdatePaymentRequest) ToPsqlDBStorage(studentID, paymentID int) UpdatePaymentData {
	return UpdatePaymentData{
		PaymentID:      paymentID,
		StudentID:      studentID,
		CurrentPaidSum: u.CurrentPaidSum,
		CheckPhoto:     u.PhotoURL,
		PaymentType:    u.PaymentType,
	}
}

func (u *UpdatePaymentRequest) Validate() (string, error) {
	if u.CurrentPaidSum == 0 && u.PaymentType == "" {
		return constants.NoUpdateResponse, nil
	}

	return "", nil
}

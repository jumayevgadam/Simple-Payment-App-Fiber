package payment

// Request model for payment.
type Request struct {
	CourseYear     int    `form:"course-year" json:"courseYear" validate:"required"`
	PaymentType    string `form:"payment-type" json:"paymentType" validate:"required"`
	CheckPhoto     string `form:"check-photo" json:"checkPhoto" validate:"required"`
	CurrentPaidSum int    `form:"current-balance" json:"currentPayedBalance" validate:"required"`
}

// Response model for taking request into Storage.
type Response struct {
	StudentID      int    `db:"student_id"`
	CourseYear     int    `db:"course_year"`
	CheckPhoto     string `db:"check_photo"`
	PaymentType    string `db:"payment_type"`
	CurrentPaidSum int    `db:"-"`
}

package payment

type PaymentsByStudentID struct {
	StudentID     int    `db:"student_id"`
	PaymentType   string `db:"payment_type"`
	PaymentStatus string `db:"payment_status"`
}

package repository

// SQL QUERIES FOR PAYMENT
const (
	// addPaymentQuery is.
	addPaymentQuery = `
		INSERT INTO payments (student_id, payment_type, payment_status, payment_amount, check_photo)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;`
)

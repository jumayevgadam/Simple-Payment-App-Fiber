package repository

// SQL QUERIES FOR PAYMENT.
const (
	// addPaymentQuery is.
	addPaymentQuery = `
		INSERT INTO payments (
			student_id, payment_type, payment_status, payment_amount, check_photo, time_id)
		VALUES (
			$1, 
			$2, 
			COALESCE(NULLIF($3, '')::payment_status_enum, 'In Progress'),
			$4, 
			$5, 
			COALESCE(NULLIF($6, 0), 1)
		)
		RETURNING 
			id;`

	// getPaymentByIDQuery is.
	getPaymentByIDQuery = `
		SELECT 
			id, 
			student_id,
			payment_type, 
			payment_status, 
			payment_amount, 
			check_photo, 
			uploaded_at, 
			updated_at
		FROM 
			payments
		WHERE 
			id = $1;`
)

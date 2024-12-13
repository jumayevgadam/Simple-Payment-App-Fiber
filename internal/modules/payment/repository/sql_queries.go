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

	// countPaymentsByStudentIDQuery is.
	countPaymentsByStudentIDQuery = `
		SELECT COUNT(student_id) 
		FROM payments
		WHERE student_id = $1;`

	// listPaymentsByStudentIDQuery is.
	listPaymentsByStudentIDQuery = `
		SELECT 
			p.payment_id,
			p.student_id,
			p.payment_type,
			p.payment_status,
			p.payment_amount,
			p.check_photo,
			p.uploaded_at,
			p.updated_at,
			p.start_year,
			p.end_year
		FROM 
			payment_details_mv AS p
		WHERE student_id = $1
		ORDER BY payment_id DESC OFFSET $2 LIMIT $3;`
)

package repository

const (
	studentInfoForPaymentQuery = `
		SELECT 
			u.id,
			u.role_id,
			f.faculty_name,
			g.group_code,
			CONCAT(u.name, '_', u.surname) full_name,
			u.username
		FROM 
			users u 
			INNER JOIN groups g ON g.id = u.group_id
			INNER JOIN faculties f ON f.id = g.faculty_id
		WHERE 
			u.id = $1 AND u.role_id = 3;`

	addPaymentQuery = `
			INSERT INTO payments (
				student_id,
				time_id,
				payment_type,
				payment_status,
				payment_amount,
				check_photo
			) VALUES (
				$1,
				$2,
				$3,
				COALESCE(NULLIF($4, '')::payment_status_enum, 'In Progress'), 
				$5,
				$6
			)
			RETURNING id;`

	getPaymentByIDQuery = `
		SELECT 
			p.id,
			p.student_id,
			p.payment_type,
			t.start_year,
			t.end_year,
			g.course_year,
			p.payment_status,
			p.payment_amount,
			p.check_photo,
			p.uploaded_at,
			p.updated_at
		FROM payments p
		INNER JOIN times t ON t.id = p.time_id
		INNER JOIN users u ON u.id = p.student_id
		INNER JOIN groups g ON g.id = u.group_id
		WHERE p.id = $1 AND u.role_id = 3;`

	countPaymentsByStudentQuery = `
		SELECT COUNT(student_id)
		FROM payments
		WHERE student_id = $1;`

	listPaymentsByStudentQuery = `
		SELECT 
			p.id,
			p.student_id, 
			p.payment_type,
			g.course_year,
			p.payment_status,
			p.payment_amount,
			p.check_photo,
			t.start_year,
			t.end_year,
			p.uploaded_at,
			p.updated_at
		FROM payments p
		INNER JOIN times t ON t.id = p.time_id
		INNER JOIN users u ON u.id = p.student_id
		INNER JOIN groups g ON g.id = u.group_id
		WHERE p.student_id = $1 AND u.role_id = 3
		ORDER BY p.id ASC OFFSET $2 LIMIT $3;`

	studentNameAndSurnameQuery = `
		SELECT 
			name,
			surname
		FROM users
		WHERE id = $1 AND role_id = 3;`

	studentUpdatePaymentQuery = `
		UPDATE payments
		SET
			payment_type = COALESCE(NULLIF($1, '')::payment_type_enum, payment_type),
			payment_amount = COALESCE(NULLIF($2, 0), payment_amount),
			check_photo = COALESCE(NULLIF($3, ''), check_photo)
		WHERE student_id = $4 AND id = $5 AND (payment_status = 'In Progress' OR payment_status = 'Rejected')
		RETURNING 'payment ops successfully updated';`

	adminGetPaymentStatusQuery = `
		SELECT payment_status
		FROM payments
		WHERE student_id = $1 AND id = $2;`

	adminUpdatePaymentStatusQuery = `
		UPDATE payments
		SET payment_status = COALESCE(NULLIF($1, '')::payment_status_enum, payment_status)
		WHERE student_id = $2 AND id = $3
		RETURNING 'payment status changed';`

	checkType3PaymentQuery = `
		SELECT COUNT(*)
		FROM payments
		WHERE student_id = $1 AND time_id = $2 AND payment_type = '3' 
		AND (payment_status = 'Accepted' OR payment_status = 'In Progress');`

	totalPaymentSumQuery = `
		SELECT SUM(payment_amount)
		FROM payments 
		WHERE student_id = $1 AND time_id = $2
		GROUP BY student_id;`

	paymentsByStudentIDQuery = `
		SELECT student_id, payment_type, payment_status
		FROM payments 
		WHERE student_id = $1 AND time_id = 1;`

	isPerformedPaymentCheckQuery = `
		SELECT EXISTS (SELECT 1 FROM payments WHERE student_id = $1 AND payment_type = '1' AND time_id = $2);`

	firstSemesterPaymentAmountQuery = `
		SELECT payment_amount FROM payments
		WHERE student_id = $1 AND time_id = $2;`

	studentDeletePaymentQuery = `
		DELETE FROM payments
		WHERE id = $1 AND student_id = $2 AND time_id = $3 AND payment_type IS NOT 'Accepted';`

	adminDeleteStudentPayment = `
		DELETE FROM payments
		WHERE id = $1 AND student_id = $2 AND time_id = $3;`
)

// --------------------- STATISTICS QUERIES ---------------------//.
const ()

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
				COALESCE(NULLIF($2, 0), 1),
				$3,
				COALESCE(NULLIF($4, '')::payment_status_enum, 'In Progress'), 
				$5,
				$6
			)
			RETURNING id;`

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
		FROM payments
		WHERE id = $1;`

	countPaymentsByStudentQuery = `
		SELECT COUNT(student_id)
		FROM payments
		WHERE student_id = $1;`

	listPaymentsByStudentQuery = `
		SELECT 
			id,
			student_id, 
			payment_type,
			payment_status,
			payment_amount,
			check_photo,
			uploaded_at,
			updated_at
		FROM payments
		WHERE student_id = $1
		ORDER BY id DESC OFFSET $2 LIMIT $3;`

	studentUpdatePaymentQuery = `
		UPDATE payments
		SET
			payment_type = COALESCE(NULLIF($1, '')::payment_type_enum, payment_type),
			payment_amount = COALESCE(NULLIF($2, 0), payment_amount),
			check_photo = COALESCE(NULLIF($3, ''), check_photo)
		WHERE student_id = $4 AND id = $5
		RETURNING 'payment ops successfully updated';`
)

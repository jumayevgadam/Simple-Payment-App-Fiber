package repository

// ------------------------ STATISTICS QUERY -------------------------------
const (
	getAcademicYearQuery = `
		SELECT id FROM times
		WHERE start_year = $1 AND end_year = $2;`

	firstSemesterStatisticsQuery = `
		SELECT 
		  (SELECT COUNT(*)
			FROM (
				SELECT student_id
				FROM payments
				WHERE payment_type IN ('1', '3') 
				AND time_id = $1
				GROUP BY student_id
				HAVING SUM(CASE WHEN payment_status = 'Accepted' THEN 1 ELSE 0 END) = COUNT(*)
			) AS paid_students
			) AS count_first_semester_paid,
		  (SELECT COUNT(*)
			FROM users
			WHERE role_id = 3
			AND id NOT IN (
				SELECT student_id
				FROM payments
				WHERE payment_type IN ('1', '3') 
					AND time_id = $1
				GROUP BY student_id
				HAVING SUM(CASE WHEN payment_status = 'Accepted' THEN 1 ELSE 0 END) = COUNT(*)
			)
			) AS count_first_semester_not_paid;`

	secondSemesterStatisticsQuery = `
		SELECT 
			(SELECT COUNT(*)
			FROM (
				SELECT student_id
				FROM payments
				WHERE payment_type IN ('2', '3')
				AND time_id = $1
				GROUP BY student_id
				HAVING SUM(CASE WHEN payment_status = 'Accepted' THEN 1 ELSE 0 END) = COUNT(*)
			) AS paid_students
			) AS count_second_semester_paid,
			(SELECT COUNT(*)
			FROM users
			WHERE role_id = 3
			AND id NOT IN (
				SELECT student_id
				FROM payments
				WHERE payment_type IN ('2', '3')
					AND time_id = $1
				GROUP BY student_id
				HAVING SUM(CASE WHEN payment_status = 'Accepted' THEN 1 ELSE 0 END) = COUNT(*)
			)
			) AS count_second_semester_not_paid;`

	bothSemesterStatisticsQuery = `
		WITH first_semester_payers AS (
			SELECT student_id
			FROM payments
			WHERE payment_status = 'Accepted'
				AND time_id = $1
				AND payment_type = '1'  -- First semester payments
			GROUP BY student_id
			HAVING COUNT(*) = (
				SELECT COUNT(*)
				FROM payments
				WHERE student_id = payments.student_id
					AND time_id = $1
					AND payment_type = '1'
					AND payment_status = 'Accepted'
			)  -- Ensure all first semester payments are accepted
		),
		second_semester_payers AS (
			SELECT student_id
			FROM payments
			WHERE payment_status = 'Accepted'
				AND time_id = $1
				AND payment_type = '2'  
			GROUP BY student_id
			HAVING COUNT(*) = (
				SELECT COUNT(*)
				FROM payments
				WHERE student_id = payments.student_id
					AND time_id = $1
					AND payment_type = '2'
					AND payment_status = 'Accepted'
			)  -- Ensure all second semester payments are accepted
		),
		both_semester_payers AS (
			SELECT student_id
			FROM payments
			WHERE payment_status = 'Accepted'
				AND time_id = $1
				AND payment_type IN ('1', '2')  
			GROUP BY student_id
			HAVING 
				-- Ensure all payments for both semesters are accepted
				student_id IN (SELECT student_id FROM first_semester_payers)
				AND student_id IN (SELECT student_id FROM second_semester_payers)
		)
		SELECT 
			(SELECT COUNT(*) FROM both_semester_payers) AS count_both_semester_paid,
			(SELECT COUNT(*) 
				FROM users
				WHERE role_id = 3 
					AND id NOT IN (
						SELECT student_id
						FROM both_semester_payers
					)
			) AS count_both_semester_not_paid;`

	totalStudentQuery = `
		SELECT COUNT(role_id) 
		FROM users WHERE role_id = 3;`

	firstSemesterStatisticsByFacultyQuery = `
		SELECT 
		  (SELECT COUNT(*)
			FROM (
				SELECT p.student_id
				FROM payments p
				INNER JOIN users u ON u.id = p.student_id
				INNER JOIN "groups" g ON g.id = u.group_id
				WHERE p.payment_type IN ('1', '3') 
				AND g.faculty_id = $1
				AND p.time_id = $2
				GROUP BY p.student_id
				HAVING SUM(CASE WHEN p.payment_status = 'Accepted' THEN 1 ELSE 0 END) = COUNT(*)
			) AS paid_students
			) AS count_first_semester_paid,
		  (SELECT COUNT(*)
			FROM users u INNER JOIN "groups" g2 ON g2.id = u.group_id
			WHERE u.role_id = 3 AND g2.faculty_id = $1
			AND u.id NOT IN (
				SELECT p.student_id
				FROM payments p INNER JOIN users u ON u.id = p.student_id
				INNER JOIN "groups" g3 ON g3.id = u.group_id
				WHERE p.payment_type IN ('1', '3') 
					AND g3.faculty_id = $1
					AND p.time_id = $2
				GROUP BY p.student_id
				HAVING SUM(CASE WHEN p.payment_status = 'Accepted' THEN 1 ELSE 0 END) = COUNT(*)
			)
			) AS count_first_semester_not_paid;`

	secondSemesterStatisticsByFacultyQuery = `
		SELECT 
			(SELECT COUNT(DISTINCT p.student_id)
				FROM payments p
				INNER JOIN users u ON u.id = p.student_id
				INNER JOIN groups g ON u.group_id = g.id
				WHERE g.faculty_id = $1 
					AND p.payment_type IN ('2', '3') 
					AND p.payment_status = 'Accepted' 
					AND p.time_id = $2) AS count_second_semester_paid,
			(SELECT COUNT(*) 
				FROM users u
				INNER JOIN groups g ON u.group_id = g.id
				WHERE g.faculty_id = $1 
				AND u.role_id = 3 
				AND u.id NOT IN (
					SELECT DISTINCT p.student_id
					FROM payments p
					WHERE p.payment_type IN ('2', '3') 
						AND p.payment_status = 'Accepted' 
						AND p.time_id = $2
			)) AS count_second_semester_not_paid;`

	bothSemesterStatisticsByFacultyQuery = `
		WITH both_semester_payers AS (
			SELECT p.student_id
			FROM payments p
			INNER JOIN users u ON u.id = p.student_id
			INNER JOIN groups g ON u.group_id = g.id
			WHERE g.faculty_id = $1 
			  AND p.payment_status = 'Accepted'
			  AND p.time_id = $2
			  AND p.payment_type IN ('1', '2', '3')
			GROUP BY p.student_id
			HAVING 
				COUNT(DISTINCT p.payment_type) = 2 OR 
				MAX(p.payment_type) = '3'
		)
		SELECT 
			(SELECT COUNT(*) FROM both_semester_payers) AS count_both_semester_paid,
			(SELECT COUNT(*) 
			 FROM users u
			 INNER JOIN groups g ON u.group_id = g.id
			 WHERE g.faculty_id = $1 
			   AND u.role_id = 3 
			   AND u.id NOT IN (
				   SELECT student_id
				   FROM both_semester_payers
			   )) AS count_both_semester_not_paid;
		`

	totalStudentsCountByFacultyQuery = `
		SELECT COUNT(u.id) 
		FROM users u
		INNER JOIN groups g ON g.id = u.group_id
		WHERE g.faculty_id = $1 AND u.role_id = 3;`
)

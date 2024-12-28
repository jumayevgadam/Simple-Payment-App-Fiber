package repository

// ------------------- FILTER STUDENT QUERY ------------------
var (
	adminFindStudentBaseQuery = `
        SELECT 
            u.id AS student_id,
            u.name AS student_name,
            u.surname AS student_surname,
            u.username AS student_username,
            COALESCE(r.role, '') AS role_name,
            COALESCE(f.faculty_name, '') AS faculty_name,
            COALESCE(g.group_code, '') AS group_code,
            COALESCE(g.course_year, 0) AS course_year
        FROM 
            users u
        LEFT JOIN 
            roles r ON u.role_id = r.id
        LEFT JOIN 
            groups g ON u.group_id = g.id
        LEFT JOIN 
            faculties f ON g.faculty_id = f.id
        WHERE 
            u.role_id = 3`

	firstSemesterPaidQuery = `
        AND u.id IN (
            SELECT p.student_id
            FROM payments p
            WHERE p.payment_type IN ('1', '3')
            GROUP BY p.student_id
            HAVING COUNT(*) = SUM(CASE WHEN p.payment_status = 'Accepted' THEN 1 ELSE 0 END)
        )
    `

	firstSemesterNotPaidQuery = `
        AND u.id NOT IN (
            SELECT p.student_id
            FROM payments p
            WHERE p.payment_type IN ('1', '3')
            GROUP BY p.student_id
            HAVING COUNT(*) = SUM(CASE WHEN p.payment_status = 'Accepted' THEN 1 ELSE 0 END)
        )
    `

	secondSemesterPaidQuery = `
        AND u.id IN (
            SELECT p.student_id
            FROM payments p
            WHERE p.payment_type IN ('2', '3')
            GROUP BY p.student_id
            HAVING COUNT(*) = SUM(CASE WHEN p.payment_status = 'Accepted' THEN 1 ELSE 0 END)
        )
    `

	secondSemesterNotPaidQuery = `
        AND u.id NOT IN (
            SELECT p.student_id
            FROM payments p
            WHERE p.payment_type IN ('2', '3')
            GROUP BY p.student_id
            HAVING COUNT(*) = SUM(CASE WHEN p.payment_status = 'Accepted' THEN 1 ELSE 0 END)
        )
    `

	bothSemesterPaidQuery = `
        AND u.id IN (
            WITH first_semester_payers AS (
                SELECT student_id
                FROM payments 
                WHERE payment_type IN ('1', '3')
                AND payment_status = 'Accepted'
                GROUP BY student_id
				HAVING COUNT(*) = SUM(CASE WHEN payment_status = 'Accepted' THEN 1 ELSE 0 END)
            ),
            second_semester_payers AS (
                SELECT student_id 
                FROM payments 
                WHERE payment_type IN ('2', '3')
                AND payment_status = 'Accepted'
                GROUP BY student_id
				HAVING COUNT(*) = SUM(CASE WHEN payment_status = 'Accepted' THEN 1 ELSE 0 END)
            )
            SELECT student_id
            FROM first_semester_payers
            INTERSECT 
            SELECT student_id
            FROM second_semester_payers
        )
		AND u.id NOT IN (
            SELECT student_id
            FROM payments 
            WHERE payment_type IN ('1', '2', '3')
            GROUP BY student_id
            HAVING COUNT(*) != SUM(CASE WHEN payment_status = 'Accepted' THEN 1 ELSE 0 END)
        )`

	bothSemesterNotPaidQuery = `
        AND u.id NOT IN (
            WITH first_semester_payers AS (
                SELECT student_id
                FROM payments
                WHERE payment_type IN ('1', '3')
                AND payment_status = 'Accepted'
                GROUP BY student_id
            ),
            second_semester_payers AS (
                SELECT student_id
                FROM payments
                WHERE payment_type IN ('2', '3')
                AND payment_status = 'Accepted'
                GROUP BY student_id
            )
            SELECT student_id
            FROM first_semester_payers
            INTERSECT
            SELECT student_id
            FROM second_semester_payers
        )
		OR u.id IN (
            SELECT student_id
            FROM payments 
            WHERE payment_type IN ('1', '2', '3')
            GROUP BY student_id
            HAVING COUNT(*) != SUM(CASE WHEN payment_status = 'Accepted' THEN 1 ELSE 0 END)
        )`

	limitOffSetQuery = `
        ORDER BY 
            u.surname ASC
        OFFSET $%d LIMIT $%d;
    `
)

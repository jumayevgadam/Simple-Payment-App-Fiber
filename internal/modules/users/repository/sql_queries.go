package repository

const limitOffSet = ` ORDER BY id DESC OFFSET $1 LIMIT $2;`

const limitOffSetTwo = ` ORDER BY id DESC OFFSET $2 LIMIT $3;`

// AUTH.
const (
	loginUserCheckWithQuery = `
		SELECT 
			u.id AS id,
			u.role_id AS role_id,
			r.role AS role_type,
			u.username AS username,
			u.name AS name,
			u.surname AS surname,
			u.password AS password
		FROM users AS u 
		INNER JOIN roles AS r ON r.id = u.role_id
		WHERE u.username = $1;`
)

// ADMIN.
const (
	addStudentQuery = `
		INSERT INTO users (
			role_id,
			group_id,
			name, 
			surname,
			username,
			password
		) VALUES (
			COALESCE(NULLIF($1, 0), 3),
			$2,
			$3,
			$4,
			$5,
			$6
		)
		RETURNING id;`

	addAdminQuery = `
		INSERT INTO users (
			role_id,
			group_id,
			name,
			surname,
			username,
			password
		) VALUES (
		 	2,
			NULL,
			$1,
			$2,
			$3,
			$4
		) RETURNING id;`

	listAdminsQuery = `
		SELECT 
			id,
			role_id,
			name,
			surname,
			username,
			password,
			created_at,
			updated_at
		FROM 
			users
		WHERE 
			group_id IS NULL` + limitOffSet

	totalAdminCountQuery = `
		SELECT COUNT(id) 
		FROM users name
		WHERE role_id = 2 AND group_id IS NULL;`

	listStudentsQuery = `SELECT 
			id,
			role_id,
			group_id,
			name,
			surname,
			username,
			password,
			created_at,
			updated_at
		FROM 
			users
		WHERE 
			role_id = 3` + limitOffSet

	totalStudentCountQuery = `
		SELECT COUNT(id)
		FROM users
		WHERE role_id = 3;`

	getAdminQuery = `
		SELECT
			id,
			role_id,
			name,
			surname,
			username,
			password,
			created_at,
			updated_at
		FROM users
		WHERE id = $1 AND group_id IS NULL;`

	getStudentQuery = `
		SELECT
			id, 
			role_id,
			group_id,
			name,
			surname,
			username,
			password,
			created_at,
			updated_at
		FROM users
		WHERE id = $1 AND group_id IS NOT NULL;`

	deleteAdminQuery = `
		DELETE FROM users
		WHERE id = $1 AND role_id = 2 AND group_id IS NULL;`

	deleteStudentQuery = `
		DELETE FROM users
		WHERE id = $1 AND role_id = 3 AND group_id IS NOT NULL;`

	listStudentsByGroupIDQuery = `
		SELECT 
			id,
			role_id,
			group_id,
			CONCAT(name, ' ', surname) AS full_name,
			username,
			password,
			created_at,
			updated_at
		FROM users
		WHERE group_id = $1 AND role_id = 3` + limitOffSetTwo

	countStudentsByGroupIDQuery = `
		SELECT COUNT(group_id)
		FROM users
		WHERE group_id = $1 AND role_id = 3;`

	updateAdminQuery = `
		UPDATE users
		SET name = COALESCE(NULLIF($1, ''), name),
			surname = COALESCE(NULLIF($2, ''), surname), 
			username = COALESCE(NULLIF($3, ''), username),
			password = COALESCE(NULLIF($4, ''), password),
			updated_at = NOW()
		WHERE id = $5 AND role_id = 2
		RETURNING 'admin updated successfully';`

	updateStudentQuery = `
		UPDATE users
		SET group_id = COALESCE(NULLIF($1, 0), group_id),
			name = COALESCE(NULLIF($2, ''), name),
			surname = COALESCE(NULLIF($3, ''), surname),
			username = COALESCE(NULLIF($4, ''), username),
			password = COALESCE(NULLIF($5, ''), password),
			updated_at = NOW()
		WHERE id = $6 AND role_id = 3
		RETURNING 'student successfully updated';`

	adminFindStudentQuery = `
		SELECT 
			u.id AS student_id,
			u.name AS student_name,
			u.surname AS student_surname,
			u.username AS student_username,
			r.role AS role_name,
			f.faculty_name,
			f.faculty_code,
			g.group_code,
			g.course_year,
			COUNT(*) OVER() AS total_count
		FROM 
			users u
		LEFT JOIN 
			roles r ON u.role_id = r.id
		LEFT JOIN 
			groups g ON u.group_id = g.id
		LEFT JOIN 
			faculties f ON g.faculty_id = f.id
		WHERE 
			(u.name ILIKE '%' || COALESCE($1, '') || '%' OR $1 IS NULL)
			AND (u.surname ILIKE '%' || COALESCE($2, '') || '%' OR $2 IS NULL)
			AND (u.username ILIKE '%' || COALESCE($3, '') || '%' OR $3 IS NULL)
			AND (g.group_code ILIKE '%' || COALESCE($4, '') || '%' OR $4 IS NULL)
			AND (f.faculty_name ILIKE '%' || COALESCE($5, '') || '%' OR $5 IS NULL)
		ORDER BY 
			u.id
		OFFSET $6 LIMIT $7;
	`
)

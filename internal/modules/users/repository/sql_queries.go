package repository

const limitOffSet = ` ORDER BY surname ASC OFFSET $1 LIMIT $2;`

const limitOffSetTwo = ` ORDER BY surname ASC OFFSET $2 LIMIT $3;`

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
			name,
			surname,
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
)

package repository

const limitOffSet = ` ORDER BY id DESC OFFSET $1 LIMIT $2;`

const limitOffSetTwo = ` ORDER BY id DESC OFFSET $2 LIMIT $3;`

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
)

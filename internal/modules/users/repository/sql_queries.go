package repository

const (
	// createUserQuery is.
	createUserQuery = `
		INSERT INTO users (role_id, group_id, name, surname, username, password)
     	VALUES (COALESCE(NULLIF($1, 0), 3), $2, $3, $4, $5, $6)
		RETURNING id;`

	// getUserByIDQuery is.
	getUserByIDQuery = `
		SELECT 
			id,
			role_id,
			COALESCE(group_id, 0) AS group_id,
			name,
			surname,
			username,
			password,
			created_at,
			updated_at
		FROM
			users
		WHERE 
			id = $1;`

	// getUserByUsernameQuery is.
	getDetailsByUsernameQuery = `
		SELECT 
			id, 
			role_id,
			username,
			password
		FROM users
		WHERE username = $1;`

	// getStudentInfoDetailsQuery is.
	getStudentInfoDetailsQuery = `
		SELECT
			COALESCE(g.course_year, 0) AS course_year,
			COALESCE(g.group_code, '') AS group_code,
			CONCAT(u.name, '-', u.surname) AS full_name,
			u.username AS username,
			f.name AS faculty_name
		FROM 
			users AS u 
		INNER JOIN
			groups AS g ON u.group_id = g.id
		INNER JOIN 
			faculties AS f on f.id = g.faculty_id	
		WHERE 
			u.id = $1 AND u.group_id IS NOT NULL;`

	// listAllUsersQuery is.
	listAllUsersQuery = `
		SELECT 
			id,
			role_id, 
			COALESCE(group_id, 0) AS group_id,
			name,
			surname,
			username,
			password,
			created_at,
			updated_at
		FROM 
			users
		ORDER BY 
			role_id ASC OFFSET $1 LIMIT $2;	
		`

	// totalCountOfAllUserQuery is.
	totalCountOfAllUserQuery = `
		SELECT COUNT(id) FROM users;`

	// updateUserDetailsQuery is.
	updateUserDetailsQuery = `
		UPDATE users
		SET 
			role_id = COALESCE(NULLIF($1, 0), role_id),
			group_id = COALESCE(NULLIF($2, 0), group_id),
			name = COALESCE(NULLIF($3, ''), name),
			surname = COALESCE(NULLIF($4, ''), surname),
			username = COALESCE(NULLIF($5, ''), username)
		WHERE
			id = $6;`

	// deleteUserQuery is.
	deleteUserQuery = `
		DELETE FROM users WHERE id = $1;`

	// listAllStudentsQuery is.
	listAllStudentsQuery = `
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
		FROM
			users
		WHERE
			role_id = 3
		ORDER BY id DESC OFFSET $1 LIMIT $2;`

	// countAllStudentsQuery is.
	countAllStudentsQuery = `
		SELECT COUNT(id) FROM users WHERE role_id = 3	;`
)

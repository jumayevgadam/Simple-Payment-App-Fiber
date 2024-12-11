package repository

const (
	// createUserQuery is.
	createUserQuery = `
		INSERT INTO users (role_id, group_id, name, surname, username, password)
     	VALUES (COALESCE(NULLIF($1, 0), 3), $2, $3, $4, $5, $6)
		RETURNING id;`

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
			u.username AS username
		FROM 
			users AS u 
		INNER JOIN
			groups AS g ON u.group_id = g.id
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
)

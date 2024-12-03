package repository

const (
	// createUserQuery is.
	createUserQuery = `
		INSERT INTO users (role_id, group_id, name, surname, username, password)
     	VALUES (NULLIF($1, 0), $2, $3, $4, $5, $6)
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
			g.course_year AS course_year,
			g.group_code AS group_code,
			CONCAT(u.name, '-', u.surname) AS full_name,
			u.username AS username
		FROM 
			users AS u 
		INNER JOIN
			groups AS g ON u.group_id = g.id
		WHERE 
			u.id = $1 AND u.group_id IS NOT NULL;`
)

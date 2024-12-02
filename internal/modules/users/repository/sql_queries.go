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
)

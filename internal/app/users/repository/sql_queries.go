package repository

const (
	// createUserQuery is.
	createUserQuery = `
		INSERT INTO users (role_id, group_id, name, surname, username, password)
     	VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;`

	// getUserByUsernameQuery is.
	getUserByUsernameQuery = `
		SELECT *
		FROM users
		WHERE username = $1;`
)

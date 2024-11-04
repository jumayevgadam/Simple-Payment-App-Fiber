package repository

// WE WRITE QUERIES IN THIS PLACE
const (
	// addRoleQuery is
	addRoleQuery = `
		INSERT INTO roles (name)
		VALUES ($1) 
		RETURNING id;`

	// getRoleQuery is
	getRoleQuery = `
	SELECT id, name
	FROM roles
	WHERE id = $1;`
)

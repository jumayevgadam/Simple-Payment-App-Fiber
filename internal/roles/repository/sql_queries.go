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

	// getRolesQuery is
	getRolesQuery = `
		SELECT *
		FROM roles;`

	// deleteRoleQuery is
	deleteRoleQuery = `
		DELETE 
		FROM roles 
		WHERE id = $1;`

	// fetchCurrentRoleQuery is
	fetchCurrentRoleQuery = `
		SELECT name
		FROM roles
		WHERE id = $1;`

	// updateRoleQuery is
	updateRoleQuery = `
		UPDATE roles
		SET 
			name = COALESCE(NULLIF($1, ''), name)
		WHERE id = $2	
		RETURNING 'role changed successfully';`
)

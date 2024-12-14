package repository

// WE WRITE QUERIES IN THIS PLACE FOR ROLES.
const (
	// addRoleQuery is.
	addRoleQuery = `
		INSERT INTO roles (role)
		VALUES ($1) 
		RETURNING id;`

	// getRoleQuery is.
	getRoleQuery = `
		SELECT id, role
		FROM roles
		WHERE id = $1;`

	// getRoleIDByRoleName is.
	getRoleByRoleName = `
		SELECT id, role FROM roles WHERE name = $1;`

	// getRolesQuery is.
	getRolesQuery = `
		SELECT *
		FROM roles;`

	// deleteRoleQuery is.
	deleteRoleQuery = `
		DELETE 
		FROM roles 
		WHERE id = $1;`

	// fetchCurrentRoleQuery is.
	fetchCurrentRoleQuery = `
		SELECT role
		FROM roles
		WHERE id = $1;`

	// updateRoleQuery is.
	updateRoleQuery = `
		UPDATE roles
		SET 
			role = COALESCE(NULLIF($1, ''), role)
		WHERE id = $2	
		RETURNING 'role changed successfully';`
)

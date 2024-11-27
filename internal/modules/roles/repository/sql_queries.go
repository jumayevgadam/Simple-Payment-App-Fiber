package repository

// WE WRITE QUERIES IN THIS PLACE FOR ROLES
const (
	// addRoleQuery is.
	addRoleQuery = `
		INSERT INTO roles (name)
		VALUES ($1) 
		RETURNING id;`

	// getRoleQuery is.
	getRoleQuery = `
		SELECT id, name
		FROM roles
		WHERE id = $1;`

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
		SELECT name
		FROM roles
		WHERE id = $1;`

	// updateRoleQuery is.
	updateRoleQuery = `
		UPDATE roles
		SET 
			name = COALESCE(NULLIF($1, ''), name)
		WHERE id = $2	
		RETURNING 'role changed successfully';`
)

// WE WRITE QUERIES FOR PERMISSIONS IN THIS CONST.
const (
	// addPermissionQuery is.
	addPermissionQuery = `
		INSERT INTO permissions (permission_type)
		VALUES ($1) 
		RETURNING id;`

	// getPermissionQuery is.
	getPermissionQuery = `
		SELECT
			id,
			permission_type
		FROM permissions
		WHERE id = $1;`

	// listPermissionsQuery is.
	listPermissionsQuery = `
		SELECT
			id,
			permission_type
		FROM permissions
		ORDER BY COALESCE(NULLIF($1, ''), permission_type) DESC OFFSET $2 LIMIT $3;`

	// deletePermissionQuery is.
	deletePermissionQuery = `
		DELETE 
		FROM permissions
		WHERE id = $1;`

	// updatePermissionQuery is.
	updatePermissionQuery = `
		UPDATE permissions 
		SET permission_type = COALESCE(NULLIF($1, ''), permission_type)
		WHERE id = $2
		RETURNING 'permission successfully edited'`
)

// WE WRITE QUERIES FOR ROLE PERMISSIONS IN THIS CONST.
const (
	// addRolePermissionQuery is.
	addRolePermissionQuery = `
		INSERT INTO role_permissions (role_id, permission_id)
		VALUES ($1, $2)
		RETURNING 'role-permission successfully created.'`

	// getPermissionsByRoleQuery is.
	getPermissionsByRoleQuery = `
		SELECT 
			role_id, 
			permission_id
		FROM role_permissions
		WHERE role_id = $1;`

	// getRolesByPermissionQuery is.
	getRolesByPermissionIDQuery = `
		SELECT 
			role_id,
			permission_id
		FROM role_permissions
		WHERE permission_id = $1;`

	getRolesByPermissionQuery = `
		SELECT rp.role_id
		FROM role_permissions rp
		JOIN permissions p ON rp.permission_id = p.id 
		WHERE p.permission_type = $1;`

	// deleteRolePermissionQuery is.
	deleteRolePermissionQuery = `
		DELETE FROM role_permissions
		WHERE 
			role_id = $1 AND permission_id = $2;`
)

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
		ORDER BY id DESC OFFSET $1 LIMIT $2;`

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
		SELECT 
			rp.role_id AS id, r.name AS name
		FROM roles AS r
		INNER JOIN 
			role_permissions AS rp ON rp.role_id = r.id 
		INNER JOIN 
			permissions AS p ON p.id = rp.permission_id
		WHERE p.permission_type = $1;`

	// deleteRolePermissionQuery is.
	deleteRolePermissionQuery = `
		DELETE FROM role_permissions
		WHERE 
			role_id = $1 AND permission_id = $2;`

	// getPermissionsByRoleQuery is.
	getPermissionsByRoleIDQuery = `
		SELECT p.permission_type
		FROM permissions p 
		INNER JOIN role_permissions rp ON rp.permission_id = p.id
		INNER JOIN roles r ON r.id = rp.role_id
		WHERE r.id = $1;`
)

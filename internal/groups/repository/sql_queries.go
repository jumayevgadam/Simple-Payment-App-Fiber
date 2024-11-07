package repository

// SQL QUERIES
const (
	// addGroupQuery is
	addGroupQuery = `
		INSERT INTO groups (faculty_id, class_code)
		VALUES ($1, $2)
		RETURNING id;`

	// getGroupQuery is
	getGroupQuery = `
		SELECT id, faculty_id, class_code
		FROM groups
		WHERE id = $1;`

	// listGroupsQuery is
	listGroupsQuery = `
		SELECT id, faculty_id, class_code
		FROM groups;`

	// deleteGroupQuery is
	deleteGroupQuery = `
		DELETE
		FROM groups
		WHERE id = $1;`

	// updateGroupQuery is
	updateGroupQuery = `
		UPDATE groups
		SET faculty_id = COALESCE(NULLIF($1, 0), faculty_id),
			class_code = COALESCE(NULLIF($2, ''), class_code)
		WHERE id = $3
		RETURNING 'group ops successfully edited';`
)

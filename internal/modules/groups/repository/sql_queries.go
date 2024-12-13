package repository

// SQL QUERIES.
const (
	// addGroupQuery is.
	addGroupQuery = `
		INSERT INTO groups (faculty_id, group_code, course_year)
		VALUES ($1, $2, $3)
		RETURNING id;`

	// getGroupQuery is.
	getGroupQuery = `
		SELECT 
			id, faculty_id, group_code, course_year
		FROM groups
		WHERE id = $1;`

	// countGroupsQuery is.
	countGroupsQuery = `
		SELECT COUNT(*)
		FROM groups;`

	// listGroupsQuery is.
	listGroupsQuery = `
		SELECT
			id, faculty_id, group_code, course_year
		FROM groups
		ORDER BY id DESC OFFSET $1 LIMIT $2;`

	// deleteGroupQuery is.
	deleteGroupQuery = `
		DELETE
		FROM groups
		WHERE id = $1;`

	// updateGroupQuery is.
	updateGroupQuery = `
		UPDATE groups
		SET faculty_id = COALESCE(NULLIF($1, 0), faculty_id),
			group_code = COALESCE(NULLIF($2, ''), group_code),
			course_year = COALESCE(NULLIF($3, 0), course_year)
		WHERE id = $4
		RETURNING 'group ops successfully edited';`
)

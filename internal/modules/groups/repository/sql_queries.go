package repository

// SQL QUERIES.
const (
	// addGroupQuery is.
	addGroupQuery = `
		INSERT INTO groups (faculty_id, group_code, course_year, group_index)
		VALUES (
			$1, 
			$2, 
			$3, 
			$4)
		RETURNING id;`

	// getGroupQuery is.
	getGroupQuery = `
		SELECT 
			id, 
			faculty_id, 
			group_code, 
			course_year, 
			COALESCE(group_index, 0) AS group_index
		FROM groups
		WHERE id = $1;`

	// countGroupsQuery is.
	countGroupsQuery = `
		SELECT COUNT(*)
		FROM groups;`

	// listGroupsQuery is.
	listGroupsQuery = `
		SELECT
			id, 
			faculty_id, 
			group_code, 
			course_year, 
			COALESCE(group_index, 0) AS group_index
		FROM groups
		ORDER BY group_index ASC OFFSET $1 LIMIT $2;`

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
			course_year = COALESCE(NULLIF($3, 0), course_year),
			group_index = COALESCE(NULLIF($4, 0), group_index)
		WHERE id = $5
		RETURNING 'group ops successfully edited';`

	// countGroupsByFacultyIDQuery is.
	countGroupsByFacultyIDQuery = `
		SELECT COUNT(faculty_id)
		FROM groups
		WHERE faculty_id = $1;`

	// listGroupsByFacultyIDQuery is.
	listGroupsByFacultyIDQuery = `
		SELECT 
			id,
			faculty_id,
			group_code,
			course_year,
			COALESCE(group_index, 0) AS group_index
		FROM 
			groups
		WHERE 
			faculty_id = $1
		ORDER BY group_index ASC OFFSET $2 LIMIT $3;`
)

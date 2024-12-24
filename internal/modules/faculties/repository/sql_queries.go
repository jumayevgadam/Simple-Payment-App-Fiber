package repository

const (
	// addFacultyQuery is.
	addFacultyQuery = `
		INSERT INTO faculties (faculty_name, faculty_code, faculty_index)
		VALUES ($1, $2, $3)
		RETURNING id;`

	// getFacultyQuery is.
	getFacultyQuery = `
		SELECT 
			id,
			faculty_name,
			faculty_code,
			COALESCE(faculty_index, 0) AS faculty_index
		FROM faculties
		WHERE id = $1;`

	// listFacultiesQuery is.
	listFacultiesQuery = `
		SELECT
			id,
			faculty_name,
			faculty_code,
			COALESCE(faculty_index, 0) AS faculty_index
		FROM faculties
		ORDER BY faculty_index ASC OFFSET $1 LIMIT $2;`

	// countFacultiesQuery is.
	countFacultiesQuery = `
		SELECT COUNT(id) 
		FROM faculties;`

	// deleteFacultyQuery is.
	deleteFacultyQuery = `
		DELETE 
		FROM faculties 
		WHERE id = $1;`

	// updateFacultyQuery is.
	updateFacultyQuery = `
		UPDATE faculties 
		SET faculty_name = COALESCE(NULLIF($1, ''), faculty_name),
			faculty_code = COALESCE(NULLIF($2, ''), faculty_code),
			faculty_index = COALESCE(NULLIF($3, 0), faculty_index)
		WHERE id = $4
		RETURNING 'successfully updated faculty'`
)

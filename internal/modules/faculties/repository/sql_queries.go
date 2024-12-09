package repository

const (
	// addFacultyQuery is.
	addFacultyQuery = `
		INSERT INTO faculties (name, faculty_code)
		VALUES ($1, $2)
		RETURNING id;`

	// getFacultyQuery is.
	getFacultyQuery = `
		SELECT 
			id,
			name,
			faculty_code
		FROM faculties
		WHERE id = $1;`

	// listFacultiesQuery is.
	listFacultiesQuery = `
		SELECT
			id,
			name,
			faculty_code
		FROM faculties
		ORDER BY id DESC OFFSET $2 LIMIT $3;`

	// countFacultiesQuery is.
	countFacultiesQuery = `
		SELECT COUNT(*) 
		FROM faculties;`

	// deleteFacultyQuery is.
	deleteFacultyQuery = `
		DELETE 
		FROM faculties 
		WHERE id = $1;`

	// updateFacultyQuery is.
	updateFacultyQuery = `
		UPDATE faculties 
		SET name = COALESCE(NULLIF($1, ''), name),
			faculty_code = COALESCE(NULLIF($2, ''), faculty_code)
		WHERE id = $3
		RETURNING 'successfully updated faculty'`
)

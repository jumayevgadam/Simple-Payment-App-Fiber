package repository

const (
	// addFacultyQuery is.
	addFacultyQuery = `
		INSERT INTO faculties (name, code)
		VALUES ($1, $2)
		RETURNING id;`

	// getFacultyQuery is.
	getFacultyQuery = `
		SELECT 
			id,
			name,
			code
		FROM faculties
		WHERE id = $1;`

	// listFacultiesQuery is.
	listFacultiesQuery = `
		SELECT
			id,
			name,
			code
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
			code = COALESCE(NULLIF($2, ''), code)
		WHERE id = $3
		RETURNING 'successfully updated faculty'`
)

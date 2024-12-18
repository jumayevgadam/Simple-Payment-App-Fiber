package errlst

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

// ParseSQLErrors returns corresponding RestErr.
func ParseSQLErrors(err error) RestErr {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// CLASS 22
		case "22001": // numeric value out of range
			return NewBadRequestError("Numeric" + ErrRange.Error() + pgErr.Message + "\n")

		// CLASS 23
		case "23505": // Unique violation
			return NewConflictError("Unique constraint violation: " + ErrConflict.Error() + pgErr.Message + "\n")
		case "23503": // Foreign key violation
			return NewBadRequestError("Foreign key violation: " + pgErr.Message + "\n")
		case "23502": // Not null violation
			return NewBadRequestError("Not null violation: " + pgErr.Message + "\n")

		// CLASS 40
		case "40001": // serialization failure
			return NewConflictError("Serialization error: " + pgErr.Message + "\n")
		// CLASS 42
		case "42601": // syntax error
			return NewBadRequestError("Syntax error in sql statements: " + pgErr.Message + "\n")
		}
	}

	if strings.Contains(err.Error(), "scany") {
		return NewBadRequestError(err.Error() + "\n")
	}

	if strings.Contains(err.Error(), "no corresponding field found") {
		return NewBadRequestError(err.Error() + "\n")
	}

	return NewBadRequestError(err.Error() + "\n")
}

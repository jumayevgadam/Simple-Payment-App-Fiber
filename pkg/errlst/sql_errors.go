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
			return NewBadRequestError("Numeric" + ErrRange.Error() + pgErr.Message)

		// CLASS 23
		case "23505": // Unique violation
			return NewConflictError("Unique constraint violation: " + ErrConflict.Error() + pgErr.Message)
		case "23503": // Foreign key violation
			return NewBadRequestError("Foreign key violation: " + pgErr.Message)
		case "23502": // Not null violation
			return NewBadRequestError("Not null violation: " + pgErr.Message)

		// CLASS 40
		case "40001": // serialization failure
			return NewConflictError("Serialization error: " + pgErr.Message)
		// CLASS 42
		case "42601": // syntax error
			return NewBadRequestError("Syntax error in sql statements: " + pgErr.Message)
		}
	}

	if strings.Contains(err.Error(), "scany") {
		return NewBadRequestError(err.Error())
	}

	if strings.Contains(err.Error(), "no corresponding field found") {
		return NewBadRequestError(err.Error())
	}

	return NewBadRequestError(err.Error())
}

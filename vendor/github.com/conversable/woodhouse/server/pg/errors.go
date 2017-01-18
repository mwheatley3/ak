package pg

import (
	"database/sql"
	"github.com/jackc/pgx"
)

// A helpful set of postgres error codes
const (
	ErrCodeNotNullViolation    = "23502"
	ErrCodeForeignKeyViolation = "23503"
	ErrCodeUniqueViolation     = "23505"
	ErrCodeCheckViolation      = "23514"
)

// IsPgErr checks that the the error is a postgres error
// and that the postgres error code matches the input code
func IsPgErr(err error, code string) bool {
	e, ok := err.(pgx.PgError)

	if !ok {
		return false
	}

	return e.Code == code
}

// IsConstraintErr checks if the error is an integrity restraint violation
// and matches the provided code and constraint
func IsConstraintErr(err error, code string, constraint string) bool {
	e, ok := err.(pgx.PgError)

	if !ok {
		return false
	}

	return e.Code == code && e.ConstraintName == constraint
}

// IsFKErr checks that the error is a foreign key
// constraint error against a specific constraint
func IsFKErr(err error, constraint string) bool {
	return IsConstraintErr(err, ErrCodeForeignKeyViolation, constraint)
}

// IsCheckErr checks that the error is a check violation
// against a specific constraint
func IsCheckErr(err error, constraint string) bool {
	return IsConstraintErr(err, ErrCodeCheckViolation, constraint)
}

// IsUniqErr checks that the error is a unique constraint violation
// against a specific constraint
func IsUniqErr(err error, constraint string) bool {
	return IsConstraintErr(err, ErrCodeUniqueViolation, constraint)
}

// IsNoRowsErr is a convenience check against sql.ErrNoRows
func IsNoRowsErr(err error) bool {
	return err == sql.ErrNoRows
}

// IgnoreNoRows returns the input error
// if the error does NOT match sql.ErrNoRows
//
// useful in situations where you don't want to
// leak ErrNoRows past your function
func IgnoreNoRows(err error) error {
	if IsNoRowsErr(err) {
		return nil
	}

	return err
}

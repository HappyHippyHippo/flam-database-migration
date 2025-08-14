package migration

import (
	"errors"

	flam "github.com/happyhippyhippo/flam"
)

var (
	ErrMigrationNotFound = errors.New("migration not found")
)

func newErrNilReference(
	arg string,
) error {
	return flam.NewErrorFrom(
		flam.ErrNilReference,
		arg)
}

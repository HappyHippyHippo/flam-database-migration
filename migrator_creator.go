package migration

import (
	flam "github.com/happyhippyhippo/flam"
)

type MigratorCreator interface {
	Accept(config flam.Bag) bool
	Create(config flam.Bag) (Migrator, error)
}

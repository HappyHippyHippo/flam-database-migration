package migration

import (
	database "github.com/happyhippyhippo/flam-database"
)

type Migration interface {
	Group() string
	Version() string
	Description() string
	Up(connection database.Connection) error
	Down(connection database.Connection) error
}

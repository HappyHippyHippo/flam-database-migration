package migration

import (
	flam "github.com/happyhippyhippo/flam"
)

type LoggerCreator interface {
	Accept(config flam.Bag) bool
	Create(config flam.Bag) (Logger, error)
}

package migration

import (
	log "github.com/happyhippyhippo/flam-log"
)

var (
	DefaultConnection    = ""
	DefaultLogger        = ""
	DefaultLogChannel    = "flam"
	DefaultLogStartLevel = log.Info
	DefaultLogErrorLevel = log.Error
	DefaultLogDoneLevel  = log.Info
)

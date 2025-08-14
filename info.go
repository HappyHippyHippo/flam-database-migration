package migration

import (
	"time"
)

type Info struct {
	Version     string
	Description string
	InstalledAt *time.Time
}

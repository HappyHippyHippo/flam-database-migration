package migration

import (
	"time"
)

type record struct {
	ID uint `gorm:"primaryKey"`

	Version     string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (record) TableName() string {
	return "__migrations"
}

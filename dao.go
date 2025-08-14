package migration

import (
	"errors"

	"gorm.io/gorm"

	database "github.com/happyhippyhippo/flam-database"
)

type dao struct{}

func newDao(
	connection database.Connection,
) (*dao, error) {
	d := &dao{}

	if e := connection.AutoMigrate(&record{}); e != nil {
		return nil, e
	}

	return d, nil
}

func (dao dao) List(
	connection database.Connection,
) ([]record, error) {
	var models []record
	result := connection.Order("created_at desc").Find(&models)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return models, nil
}

func (dao dao) Last(
	connection database.Connection,
) (*record, error) {
	model := &record{}
	result := connection.Order("created_at desc").FirstOrInit(model, record{Version: ""})
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return model, nil
}

func (dao dao) Up(
	connection database.Connection,
	version string,
	description string,
) (*record, error) {
	model := &record{Version: version, Description: description}
	result := connection.Create(model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (dao dao) Down(
	connection database.Connection,
	last *record,
) error {
	if last.Version != "" {
		result := connection.Delete(last)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

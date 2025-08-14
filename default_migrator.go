package migration

import (
	"slices"
	"time"

	"gorm.io/gorm"

	database "github.com/happyhippyhippo/flam-database"
)

type defaultMigrator struct {
	connection database.Connection
	logger     Logger
	migrations []Migration
	dao        *dao
}

func newDefaultMigrator(
	connection database.Connection,
	logger Logger,
	migrations []Migration,
) (Migrator, error) {
	dao, e := newDao(connection)
	if e != nil {
		return nil, e
	}

	return &defaultMigrator{
		connection: connection,
		logger:     logger,
		migrations: migrations,
		dao:        dao,
	}, nil
}

func (migrator *defaultMigrator) List() ([]Info, error) {
	installed, e := migrator.dao.List(migrator.connection)
	if e != nil {
		return nil, e
	}

	var migrations []Info
	for _, migration := range migrator.migrations {
		var createdAt *time.Time
		for _, m := range installed {
			if m.Version == migration.Version() {
				createdAt = &m.CreatedAt
				break
			}
		}

		migrations = append(migrations, Info{
			Version:     migration.Version(),
			Description: migration.Description(),
			InstalledAt: createdAt,
		})
	}

	return migrations, nil
}

func (migrator *defaultMigrator) Current() (*Info, error) {
	last, e := migrator.dao.Last(migrator.connection)
	switch {
	case e != nil:
		return nil, e
	case last.ID == 0:
		return nil, nil
	}

	return &Info{
		Version:     last.Version,
		Description: last.Description,
		InstalledAt: &last.CreatedAt,
	}, nil
}

func (migrator *defaultMigrator) CanUp() bool {
	list, e := migrator.List()
	if e != nil {
		return false
	}

	for _, migration := range list {
		if migration.InstalledAt == nil {
			return true
		}
	}

	return false
}

func (migrator *defaultMigrator) CanDown() bool {
	list, e := migrator.List()
	if e != nil {
		return false
	}

	for _, migration := range list {
		if migration.InstalledAt != nil {
			return true
		}
	}

	return false
}

func (migrator *defaultMigrator) Up() error {
	last, e := migrator.dao.Last(migrator.connection)
	if e != nil {
		return e
	}

	if last.ID == 0 && len(migrator.migrations) > 0 {
		return migrator.up(migrator.migrations[0])
	}

	if migrator.migrations[len(migrator.migrations)-1].Version() != last.Version {
		for i, migration := range migrator.migrations {
			if migration.Version() == last.Version {
				return migrator.up(migrator.migrations[i+1])
			}
		}
	}

	return ErrMigrationNotFound
}

func (migrator *defaultMigrator) UpAll() error {
	last, e := migrator.dao.Last(migrator.connection)
	if e != nil {
		return e
	}

	execute := false
	for _, migration := range migrator.migrations {
		if last.ID == 0 || execute {
			if e := migrator.up(migration); e != nil {
				return e
			}
		} else if migration.Version() == last.Version {
			execute = true
		}
	}

	return nil
}

func (migrator *defaultMigrator) Down() error {
	last, e := migrator.dao.Last(migrator.connection)
	if e != nil {
		return e
	}

	if last.ID != 0 {
		for i, migration := range migrator.migrations {
			if migration.Version() == last.Version {
				return migrator.down(migrator.migrations[i], last)
			}
		}
	}

	return ErrMigrationNotFound
}

func (migrator *defaultMigrator) DownAll() error {
	slices.Reverse(migrator.migrations)
	for i, migration := range migrator.migrations {
		last, e := migrator.dao.Last(migrator.connection)
		if e != nil {
			return e
		}

		execute := false
		if migration.Version() == last.Version {
			execute = true
		}

		if execute {
			if e := migrator.down(migrator.migrations[i], last); e != nil {
				return e
			}
		}
	}

	return nil
}

func (migrator *defaultMigrator) up(
	migration Migration,
) error {
	_ = migrator.logUpStart(migration)

	if e := migrator.connection.Transaction(func(tx *gorm.DB) error {
		if e := migration.Up(tx); e != nil {
			_ = migrator.logUpError(migration, e)
			return e
		}

		if _, e := migrator.dao.Up(tx, migration.Version(), migration.Description()); e != nil {
			_ = migrator.logUpError(migration, e)
			return e
		}

		return nil
	}); e != nil {
		return e
	}

	_ = migrator.logUpDone(migration)

	return nil
}

func (migrator *defaultMigrator) down(
	migration Migration,
	record *record,
) error {
	_ = migrator.logDownStart(migration)

	if e := migrator.connection.Transaction(func(tx *gorm.DB) error {
		if e := migration.Down(tx); e != nil {
			_ = migrator.logDownError(migration, e)
			return e
		}

		if e := migrator.dao.Down(tx, record); e != nil {
			_ = migrator.logDownError(migration, e)
			return e
		}

		return nil
	}); e != nil {
		return e
	}

	_ = migrator.logDownDone(migration)

	return nil
}

func (migrator *defaultMigrator) logUpStart(
	migration Migration,
) error {
	if migrator.logger == nil {
		return nil
	}

	return migrator.logger.LogUpStart(
		Info{
			Version:     migration.Version(),
			Description: migration.Description(),
		},
	)
}

func (migrator *defaultMigrator) logUpError(
	migration Migration,
	e error,
) error {
	if migrator.logger == nil {
		return nil
	}

	return migrator.logger.LogUpError(
		Info{
			Version:     migration.Version(),
			Description: migration.Description(),
		},
		e,
	)
}

func (migrator *defaultMigrator) logUpDone(
	migration Migration,
) error {
	if migrator.logger == nil {
		return nil
	}

	return migrator.logger.LogUpDone(
		Info{
			Version:     migration.Version(),
			Description: migration.Description(),
		},
	)
}

func (migrator *defaultMigrator) logDownStart(
	migration Migration,
) error {
	if migrator.logger == nil {
		return nil
	}

	return migrator.logger.LogDownStart(
		Info{
			Version:     migration.Version(),
			Description: migration.Description(),
		},
	)
}

func (migrator *defaultMigrator) logDownError(
	migration Migration,
	e error,
) error {
	if migrator.logger == nil {
		return nil
	}

	return migrator.logger.LogDownError(
		Info{
			Version:     migration.Version(),
			Description: migration.Description(),
		},
		e,
	)
}

func (migrator *defaultMigrator) logDownDone(
	migration Migration,
) error {
	if migrator.logger == nil {
		return nil
	}

	return migrator.logger.LogDownDone(
		Info{
			Version:     migration.Version(),
			Description: migration.Description(),
		},
	)
}

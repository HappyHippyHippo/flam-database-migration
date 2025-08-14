package migration

type Facade interface {
	HasLogger(id string) bool
	ListLoggers() []string
	GetLogger(id string) (Logger, error)
	AddLogger(id string, logger Logger) error

	HasMigrator(id string) bool
	ListMigrators() []string
	GetMigrator(id string) (Migrator, error)
	AddMigrator(id string, migrator Migrator) error
}

type facade struct {
	loggerFactory   loggerFactory
	migratorFactory migratorFactory
}

func newFacade(
	loggerFactory loggerFactory,
	migratorFactory migratorFactory,
) Facade {
	return &facade{
		loggerFactory:   loggerFactory,
		migratorFactory: migratorFactory,
	}
}

func (facade facade) HasLogger(
	id string,
) bool {
	return facade.loggerFactory.Has(id)
}

func (facade facade) ListLoggers() []string {
	return facade.loggerFactory.List()
}

func (facade facade) GetLogger(
	id string,
) (Logger, error) {
	return facade.loggerFactory.Get(id)
}

func (facade facade) AddLogger(
	id string,
	logger Logger,
) error {
	return facade.loggerFactory.Add(id, logger)
}

func (facade facade) HasMigrator(
	id string,
) bool {
	return facade.migratorFactory.Has(id)
}

func (facade facade) ListMigrators() []string {
	return facade.migratorFactory.List()
}

func (facade facade) GetMigrator(
	id string,
) (Migrator, error) {
	return facade.migratorFactory.Get(id)
}

func (facade facade) AddMigrator(
	id string,
	migrator Migrator,
) error {
	return facade.migratorFactory.Add(id, migrator)
}

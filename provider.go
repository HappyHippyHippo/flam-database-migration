package migration

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
	config "github.com/happyhippyhippo/flam-config"
	log "github.com/happyhippyhippo/flam-log"
)

type provider struct{}

func NewProvider() flam.Provider {
	return &provider{}
}

func (provider) Id() string {
	return providerId
}

func (provider) Register(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	var e error
	provide := func(constructor any, opts ...dig.ProvideOption) bool {
		e = container.Provide(constructor, opts...)
		return e == nil
	}

	_ = provide(newDao) &&
		provide(newDefaultLoggerCreator, dig.Group(LoggerCreatorGroup)) &&
		provide(newLoggerFactory) &&
		provide(newDefaultMigratorCreator, dig.Group(MigratorCreatorGroup)) &&
		provide(newMigratorFactory) &&
		provide(newMigrationPool) &&
		provide(newFacade)

	return e
}

func (provider) Boot(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	return container.Invoke(func(
		configFacade config.Facade,
		migrationFacade Facade,
	) error {
		DefaultConnection = configFacade.String(PathDefaultConnection, DefaultConnection)
		DefaultLogger = configFacade.String(PathDefaultLogger, DefaultLogger)
		DefaultLogChannel = configFacade.String(PathDefaultLogChannel, DefaultLogChannel)
		DefaultLogStartLevel = log.LevelFrom(configFacade.Get(PathDefaultLogStartLevel), DefaultLogStartLevel)
		DefaultLogErrorLevel = log.LevelFrom(configFacade.Get(PathDefaultLogErrorLevel), DefaultLogErrorLevel)
		DefaultLogDoneLevel = log.LevelFrom(configFacade.Get(PathDefaultLogDoneLevel), DefaultLogDoneLevel)

		if configFacade.Bool(PathBoot) {
			for _, id := range migrationFacade.ListMigrators() {
				migrator, e := migrationFacade.GetMigrator(id)
				if e != nil {
					return e
				}

				if e := migrator.UpAll(); e != nil {
					return e
				}
			}
		}

		return nil
	})
}

func (provider) Close(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	return container.Invoke(func(
		migratorFactory migratorFactory,
		loggerFactory loggerFactory,
	) error {
		if e := migratorFactory.Close(); e != nil {
			return e
		}

		if e := loggerFactory.Close(); e != nil {
			return e
		}

		return nil
	})
}

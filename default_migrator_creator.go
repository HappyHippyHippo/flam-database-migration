package migration

import (
	flam "github.com/happyhippyhippo/flam"
	database "github.com/happyhippyhippo/flam-database"
)

type defaultMigratorCreator struct {
	databaseFacade database.Facade
	loggerFactory  loggerFactory
	pool           *migrationPool
}

func newDefaultMigratorCreator(
	databaseFacade database.Facade,
	loggerFactory loggerFactory,
	pool *migrationPool,
) MigratorCreator {
	return &defaultMigratorCreator{
		databaseFacade: databaseFacade,
		loggerFactory:  loggerFactory,
		pool:           pool,
	}
}

func (creator defaultMigratorCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == MigratorDriverDefault &&
		config.String("group") != ""
}

func (creator defaultMigratorCreator) Create(
	config flam.Bag,
) (Migrator, error) {
	connectionId := config.String("connection", DefaultConnection)
	connection, e := creator.databaseFacade.GetConnection(connectionId)
	if e != nil {
		return nil, e
	}

	var logger Logger
	if loggerId := config.String("logger", DefaultLogger); loggerId != "" {
		logger, e = creator.loggerFactory.Get(loggerId)
		if e != nil {
			return nil, e
		}
	}

	return newDefaultMigrator(
		connection,
		logger,
		creator.pool.Group(config.String("group")),
	)
}

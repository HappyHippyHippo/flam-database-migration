package migration

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
)

type migratorFactory flam.Factory[Migrator]

type migratorFactoryArgs struct {
	dig.In

	Creators      []MigratorCreator `group:"flam.database_migration.migrators.creator"`
	FactoryConfig flam.FactoryConfig
}

func newMigratorFactory(
	args migratorFactoryArgs,
) (migratorFactory, error) {
	var creators []flam.ResourceCreator[Migrator]
	for _, creator := range args.Creators {
		creators = append(creators, creator)
	}

	return flam.NewFactory(
		creators,
		PathMigrators,
		args.FactoryConfig,
		nil)
}

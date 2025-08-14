package migration

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
)

type loggerFactory flam.Factory[Logger]

type loggerFactoryArgs struct {
	dig.In

	Creators      []LoggerCreator `group:"flam.database_migration.loggers.creator"`
	FactoryConfig flam.FactoryConfig
}

func newLoggerFactory(
	args loggerFactoryArgs,
) (loggerFactory, error) {
	var creators []flam.ResourceCreator[Logger]
	for _, creator := range args.Creators {
		creators = append(creators, creator)
	}

	return flam.NewFactory(
		creators,
		PathLoggers,
		args.FactoryConfig,
		nil)
}

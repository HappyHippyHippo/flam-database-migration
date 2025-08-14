package migration

const (
	providerId = "flam.database_migration.provider"

	LoggerCreatorGroup    = "flam.database_migration.loggers.creator"
	LoggerDriverDefault   = "flam.database_migration.loggers.driver.default"
	MigratorCreatorGroup  = "flam.database_migration.migrators.creator"
	MigratorDriverDefault = "flam.database_migration.migrators.driver.default"
	MigrationGroup        = "flam.database_migration.migrations"

	PathDefaultConnection    = "flam.database_migration.defaults.connection"
	PathDefaultLogger        = "flam.database_migration.defaults.logger"
	PathDefaultLogChannel    = "flam.database_migration.defaults.log.channel"
	PathDefaultLogStartLevel = "flam.database_migration.defaults.log.start.level"
	PathDefaultLogErrorLevel = "flam.database_migration.defaults.log.error.level"
	PathDefaultLogDoneLevel  = "flam.database_migration.defaults.log.done.level"
	PathBoot                 = "flam.database_migration.boot"
	PathLoggers              = "flam.database_migration.loggers"
	PathMigrators            = "flam.database_migration.migrators"
)

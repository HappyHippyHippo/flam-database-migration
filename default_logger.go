package migration

import (
	"fmt"

	log "github.com/happyhippyhippo/flam-log"
)

type defaultLoggerLevels struct {
	start log.Level
	error log.Level
	done  log.Level
}

type defaultLogger struct {
	logFacade log.Facade
	channel   string
	level     defaultLoggerLevels
}

func newDefaultLogger(
	logFacade log.Facade,
	channel string,
	startLevel log.Level,
	errorLevel log.Level,
	doneLevel log.Level,
) Logger {
	return &defaultLogger{
		logFacade: logFacade,
		channel:   channel,
		level: defaultLoggerLevels{
			start: startLevel,
			error: errorLevel,
			done:  doneLevel,
		},
	}
}

func (logger *defaultLogger) LogUpStart(
	migration Info,
) error {
	return logger.logFacade.Signal(
		logger.level.start,
		logger.channel,
		fmt.Sprintf("Migration '%s' up action started", migration.Description))
}

func (logger *defaultLogger) LogUpError(
	migration Info,
	e error,
) error {
	return logger.logFacade.Signal(
		logger.level.error,
		logger.channel,
		fmt.Sprintf("Migration '%s' up action error: %s", migration.Version, e.Error()))
}

func (logger *defaultLogger) LogUpDone(
	migration Info,
) error {
	return logger.logFacade.Signal(
		logger.level.done,
		logger.channel,
		fmt.Sprintf("Migration '%s' up action terminated", migration.Version))
}

func (logger *defaultLogger) LogDownStart(
	migration Info,
) error {
	return logger.logFacade.Signal(
		logger.level.start,
		logger.channel,
		fmt.Sprintf("Migration '%s' down action started", migration.Version))
}

func (logger *defaultLogger) LogDownError(
	migration Info,
	e error,
) error {
	return logger.logFacade.Signal(
		logger.level.error,
		logger.channel,
		fmt.Sprintf("Migration '%s' down action error: %s", migration.Version, e.Error()))
}

func (logger *defaultLogger) LogDownDone(
	migration Info,
) error {
	return logger.logFacade.Signal(
		logger.level.done,
		logger.channel,
		fmt.Sprintf("Migration '%s' down action terminated", migration.Version))
}

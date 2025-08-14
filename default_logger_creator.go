package migration

import (
	flam "github.com/happyhippyhippo/flam"
	log "github.com/happyhippyhippo/flam-log"
)

type defaultLoggerCreator struct {
	logFacade log.Facade
}

func newDefaultLoggerCreator(
	logFacade log.Facade,
) LoggerCreator {
	return &defaultLoggerCreator{
		logFacade: logFacade,
	}
}

func (defaultLoggerCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == LoggerDriverDefault
}

func (creator defaultLoggerCreator) Create(
	config flam.Bag,
) (Logger, error) {
	return newDefaultLogger(
		creator.logFacade,
		config.String("channel", DefaultLogChannel),
		log.LevelFrom(config.Get("levels.start"), DefaultLogStartLevel),
		log.LevelFrom(config.Get("levels.error"), DefaultLogErrorLevel),
		log.LevelFrom(config.Get("levels.done"), DefaultLogDoneLevel),
	), nil
}

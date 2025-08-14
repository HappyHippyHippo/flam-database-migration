package migration

type Logger interface {
	LogUpStart(migration Info) error
	LogUpError(migration Info, e error) error
	LogUpDone(migration Info) error

	LogDownStart(migration Info) error
	LogDownError(migration Info, e error) error
	LogDownDone(migration Info) error
}

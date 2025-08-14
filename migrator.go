package migration

type Migrator interface {
	List() ([]Info, error)
	Current() (*Info, error)
	CanUp() bool
	CanDown() bool
	Up() error
	UpAll() error
	Down() error
	DownAll() error
}

package core

type Migration struct{}

func NewMigration() *Migration {
	return &Migration{}
}

func (m *Migration) Exec() error {
	return nil
}

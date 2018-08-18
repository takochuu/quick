package core

type Migration struct{}

func NewMigration() *Migration {
	return &Migration{}
}

func (m *Migration) Exec() error {
	// TODO DBに接続する
	// TODO 当たっていないSchemaを調べる
	// TODO 当たっていないSchemaを当てる
	return nil
}

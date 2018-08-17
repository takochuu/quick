package factory

import "github.com/takochuu/quick/core"

const (
	migration = "migrate"
	create    = "create"
	rollback  = "rollback"
	up        = "up"
	down      = "down"
)

type Executor interface {
	Exec() error
}

func New(cmd string) Executor {
	var instance Executor
	switch cmd {
	case migration:
		instance = core.NewMigration()
	case create:
	case rollback:
	case up:
	case down:
	default:
	}

	return instance
}

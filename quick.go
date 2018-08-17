package quick

import (
	"os"

	"github.com/takochuu/quick/factory"
)

const (
	migration = "migrate"
	create    = "create"
	rollback  = "rollback"
	up        = "up"
	down      = "down"
)

type Quick struct {
	Executor factory.Executor
	Stderr   *os.File
	Stdin    *os.File
	Stdout   *os.File
}

func New(cmd string) *Quick {
	return &Quick{
		Executor: factory.New(cmd),
		Stderr:   os.Stderr,
		Stdin:    os.Stdin,
		Stdout:   os.Stdout,
	}
}

func (q *Quick) Run(dryRun bool) error {
	// TODO Begin

	if err := q.Executor.Exec(); err != nil {
		return err
	}

	if dryRun {
		// TODO Rollback
		return nil
	}
	// TODO Commit

	return nil
}

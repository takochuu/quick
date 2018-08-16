package quick

import "os"

type Quick struct {
	Stderr *os.File
	Stdin  *os.File
	Stdout *os.File
}

func New() *Quick {
	return &Quick{
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
	}
}

func (q *Quick) Run(cmd string, dryRun bool) error {

	if dryRun {
		// TODO Rollback
		return nil
	}
	// TODO Commit

	return nil
}

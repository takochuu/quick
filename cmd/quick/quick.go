package main

import (
	"os"

	"flag"

	"github.com/golang/glog"
	"github.com/takochuu/quick"
)

var (
	dryRun = flag.Bool("dryRun", false, "")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	os.Exit(run(flag.Args()))
}

func run(args []string) int {

	if len(args) == 0 {
		flag.Usage()
		return 1
	}

	cmd := args[0]
	cli := quick.New()
	if err := cli.Run(cmd, *dryRun); err != nil {
		glog.Errorln(err)
		return 1
	}
	return 0
}

func usage() {
	glog.Info("Usage")
}

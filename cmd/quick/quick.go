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

	// TODO options
	cmd := args[0]
	conf := args[1]
	cli := quick.New(cmd, conf)
	if err := cli.Run(*dryRun); err != nil {
		glog.Errorln(err)
		return 1
	}
	return 0
}

func usage() {
	glog.Info("Usage")
}

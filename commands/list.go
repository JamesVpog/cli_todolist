package commands

import (
	"flag"
	"fmt"
	"os"
)

type ListCommand struct {
	fs flag.FlagSet
}

func NewListCommand() *ListCommand {
	lc := &ListCommand{
		fs: *flag.NewFlagSet("list", flag.ContinueOnError),
	}
	lc.fs.Usage = func() {fmt.Fprintf(os.Stderr, `Usage: %s 
Lists the task list from file storage
`,
		lc.fs.Name())
}
	return lc
}

func (lc *ListCommand) Init(args []string) error {
	return lc.fs.Parse(args)
}

func (lc *ListCommand) Name() string {
	return lc.fs.Name()
}

func (lc *ListCommand) Run() error {
	list()
	return nil 
}

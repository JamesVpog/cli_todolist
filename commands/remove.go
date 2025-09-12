package commands

import (
	"flag"
	"fmt"
	"os"
)

type RemoveCommand struct {
	fs flag.FlagSet
}

func NewRemoveCommand() *RemoveCommand {
	rc := &RemoveCommand{
		fs : *flag.NewFlagSet("rm", flag.ContinueOnError),
	}

	rc.fs.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage: %s 
Removes the task list from file storage

`,
			rc.fs.Name())
	}
	return rc
}

func (rc *RemoveCommand) Init(args []string) error {
	return rc.fs.Parse(args)
}
func (rc *RemoveCommand) Name() string {
	return rc.fs.Name()
}
func (rc *RemoveCommand) Run() error {
	dataPath, err := getDataPath()
	if err != nil {
		return(err)
	}
	err = os.Remove(dataPath)
	if err != nil {
		return(err)
	}

	fmt.Println("Removed the todo list")
	return nil
}

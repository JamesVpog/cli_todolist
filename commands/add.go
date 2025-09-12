// add sub-command implementation
package commands

import (
	"flag"
	"fmt"
	"os"
)

type AddCommand struct {
	fs    *flag.FlagSet
	tasks []string
}

func NewAddCommand() *AddCommand {
	ac := &AddCommand{
		fs: flag.NewFlagSet("add", flag.ContinueOnError),
	}
	ac.fs.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage: %s [TASK...]
Add one or more tasks

Positional arguments:
TASK    names of tasks to add to the task list 
`,
			ac.fs.Name())
	}
	return ac
}

func (a *AddCommand) Init(args []string) error {
	return a.fs.Parse(args)
}

func (a *AddCommand) Name() string {
	return a.fs.Name()
}

// add tasks with a status of pending into tasks.json
func (a *AddCommand) Run() error {

	a.tasks = a.fs.Args()
	if len(a.tasks) < 1 {
		fmt.Fprintf(os.Stderr, "Nothing specified, nothing added.\n")
		fmt.Fprintf(os.Stderr, "hint: Maybe you wanted to say `todo add 'example task 1`?\n")
		return nil
	}
	// start_id will be 0 if current_tasks is empty
	var start_id int
	current_tasks, err := loadTasks()
	if err != nil {
		return (err)
	}

	if len(current_tasks) != 0 {
		// tasks.json exists
		start_id = current_tasks[len(current_tasks)-1].ID + 1

	}

	var new_tasks []Task

	for _, v := range a.tasks {
		var new_task Task
		new_task.Description = v
		new_task.ID = start_id
		new_task.Status = "[ ]"

		start_id += 1
		new_tasks = append(new_tasks, new_task)
	}

	// current tasks and new tasks combine
	current_tasks = append(current_tasks, new_tasks...)

	err = saveTasks(current_tasks)

	if err != nil {
		panic(err)
	}
	fmt.Println("Added tasks to tasks.json!")

	list()
	
	return nil //no errors!
}

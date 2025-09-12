package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type DeleteCommand struct {
	fs       *flag.FlagSet
	taskNums []int
}

func NewDeleteCommand() *DeleteCommand {
	dc := &DeleteCommand{
		fs: flag.NewFlagSet("del", flag.ContinueOnError),
	}
	dc.fs.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage: %s [TASKNUM...]
Deletes one or more tasks starting at 0 for task 1, 1 for task 2 etc.

Positional arguments:
TASKNUM    task numbers of the tasks to be deleted from the tasks.json
`,
			dc.fs.Name())
	}
	return dc
}

func (dc *DeleteCommand) Init(args []string) error {
	return dc.fs.Parse(args)
}

func (dc *DeleteCommand) Name() string {
	return dc.fs.Name()
}

// Given a slice of task numbers, delete each task from the tasks.json
func (dc *DeleteCommand) Run() error {

	dc.taskNums = stringSliceToIntSlice(dc.fs.Args())

	if len(dc.taskNums) < 1 {
		fmt.Fprintf(os.Stderr, "Nothing specified, nothing deleted.\n")
		fmt.Fprintf(os.Stderr, "hint: Maybe you wanted to say `todo del 0 1'?\n")
		return nil
	}
	current_tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}


	// make a map of task.id to index which looks up the tasks
	taskMap := make(map[int]int)
	for i, task := range current_tasks {
		taskMap[task.ID] = i
	}
	// for each task_num, grab the task it refers to, and remove it
	for i := range dc.taskNums {
		taskNum := dc.taskNums[i]

		// lookup the taskNum in taskMap
		// using the index works because the tasks.json is ordered
		index, exists := taskMap[taskNum]

		// TODO: will be weird when deleting lots of tasks
		// if it exists, then remove the index at that slice wtih append
		// this appends eleements until the index, and elements after the index, effectively skipping the index
		if exists {
			current_tasks = append(current_tasks[:index], current_tasks[index+1:]...)
		} else {
			// trying to delete a task with invalid id
			error_msg := fmt.Sprintf("task %d does not exist in tasks.json\n", taskNum)
			return errors.New(error_msg)
		}
	}
	// go through all the tasks and re-id them since they are out of order
	for i := range current_tasks {
		current_tasks[i].ID = i
	}
	err = saveTasks(current_tasks)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted task(s)!")
	list()
	return nil
}

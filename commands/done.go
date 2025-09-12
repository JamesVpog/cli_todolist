package commands

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type DoneCommand struct {
	fs       *flag.FlagSet
	taskNums []int
}

func NewDoneCommand() *DoneCommand {
	dc := &DoneCommand{
		fs: flag.NewFlagSet("done", flag.ContinueOnError),
	}
	dc.fs.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage: %s [TASKNUM...]
Completes one or more tasks starting at 0 for task 1, 1 for task 2 etc.

Positional arguments:
TASKNUM    task numbers of the tasks to be marked as complete in the tasks.json
`,
			dc.fs.Name())
	}
	return dc
}

func (dc *DoneCommand) Init(args []string) error {
	return dc.fs.Parse(args)
}
func (dc *DoneCommand) Name() string {
	return dc.fs.Name()
}

// https://stackoverflow.com/questions/24972950/go-convert-strings-in-array-to-integer
func stringSliceToIntSlice(s []string) (intSlice []int) {

	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, j)
	}
	return intSlice
}

// change the status of each task in dc.taskNums to done
func (dc *DoneCommand) Run() error {

	if len(dc.fs.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "Nothing specified, nothing completed.\n")
		fmt.Fprintf(os.Stderr, "hint: Maybe you wanted to say `todo done 0 1'`?\n")
		return nil
	}
	dc.taskNums = stringSliceToIntSlice(dc.fs.Args())

	current_tasks, err := loadTasks()
	if err != nil {
		return(err)
	}

	// make a map of task.id to index which looks up the tasks
	taskMap := make(map[int]int)
	for i, task := range current_tasks {
		taskMap[task.ID] = i
	}
	// for each task_num, grab the task it refers to and set its status to [x]
	for i := range dc.taskNums {
		taskNum := dc.taskNums[i]

		// lookup the taskNum in taskMap
		// using the index works because the tasks.json is ordered
		index, exists := taskMap[taskNum]

		// if it exists, then change its status to x
		if exists {
			current_tasks[index].Status = "[x]"
		} else {
			// trying to delete a task with invalid id
			fmt.Printf("task %d does not exist in tasks.json\n", taskNum)
			return(err)
		}
	}

	err = saveTasks(current_tasks)
	if err != nil {
		return(err)
	}

	list()

	fmt.Println("Completed task(s)!")
	
	return nil
}




// holds exeuction of commands and structure of commands, also helper functions
package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

// all commands follow this Runner interface
type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

// have to use uppercase First letter to make the Task fields visible to the encoding/json package
// In Go, uppercase first letter means exported variable that is visible to other packages
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

const helpText = 
`
todo is a todolist manager cli written in Go.

Usage:

        todo <command> [arguments]

The commands are:
		add		create a new task 
		done		complete a task 
		del		delete a task
		list		view all tasks 
		rm		delete the todo list

Use "todo <command> -h" for more information about a command.
`
func Root(args []string) error {

	if len(args) < 1 {
        fmt.Fprintf(os.Stderr, "todo: no command specified\nRun 'todo -help' for the full list of commands.\n")
        os.Exit(1)
        return nil
	}
	// all the commands
	cmds := []Runner{
		NewAddCommand(),
		NewDoneCommand(),
		NewRemoveCommand(),
		NewDeleteCommand(),
		NewListCommand(),
	}

	subcmd := os.Args[1]

	// check if help message 
	if subcmd == "-help" {
		fmt.Print(helpText)
		return nil
	}
	for _, cmd := range cmds {
		// if our subcmd matches any of the commands
		if cmd.Name() == subcmd {
			// init the command
			err := cmd.Init(os.Args[2:])
			if err != nil {
				return err
			}
			// run the command with side effects
			return cmd.Run()
		}
	}
	return fmt.Errorf("todo: %s is not a todo command. See 'todo -help'. ", subcmd)
}

func getDataPath() (filePath string, err error) {
	// Stores in
	// ~/.local/share/todo/tasks.json for Linux
	// /Library/Application\ Support/todo/tasks.json for macOS
	//  %AppData%\todo\tasks.json for Windows

	return xdg.DataFile("todo/tasks.json")
}

// If tasks.json doesn't exist, it is an empty slice of Tasks and not an error
func loadTasks() (tasks []Task, err error) {

	dataPath, err := getDataPath()
	if err != nil {
		return nil, err
	}

	// Create directory if it doesn't exist
	os.MkdirAll(filepath.Dir(dataPath), 0755)

	data, err := os.ReadFile(dataPath)

	if errors.Is(err, os.ErrNotExist) {
		return tasks, nil
	}

	if err != nil {
		//file-read error
		return nil, err
	}
	//else there is tasks.json, unmarshall it
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		//error unmarshalling data
		return nil, err
	}

	return tasks, nil
}

// given tasks, saves tasks to tasks.json
func saveTasks(tasks []Task) (err error) {
	dataPath, err := getDataPath()
	if err != nil {
		return err
	}
	// marshal data to json format
	b, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(dataPath, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// shared between add, done, list, and del
// Output the tasks.json list to stdout in pretty format
func list() {
	current_tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}

	// should not call list with empty tasks
	if len(current_tasks) == 0 {
		fmt.Println("Please create a task first with ./todo add <task_1> <task_2> ... ")
		os.Exit(1)
	}

	fmt.Println("Current list of tasks:")
	for _, t := range current_tasks {
		fmt.Printf("%s  %s\n", t.Status, t.Description)
	}
}

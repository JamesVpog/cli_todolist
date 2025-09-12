package main

import (
	"fmt"
	"os"

	"github.com/JamesVpog/todo/commands"
)

// handle errors returned from commands
func main() {
	if err := commands.Root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// const default_help_text string = `./todo is a todolist manager cli written in Go.

// Usage:

//         todo <command> [arguments]

// The commands are:
// 		add		create a new task
// 		done		complete a task
// 		del		delete a task
// 		list		view all tasks
// 		rm		delete the todo list

// Use "todo help <command>" for more information about a command.

// `

// // have to use uppercase First letter to make the Task fields visible to the encoding/json package
// // In Go, uppercase first letter means exported variable that is visible to other packages
// type Task struct {
// 	ID          int    `json:"id"`
// 	Description string `json:"description"`
// 	Status      string `json:"status"`
// }

// func main() {
// 	if len(os.Args) == 1 {
// 		fmt.Print(default_help_text)
// 		os.Exit(1)
// 	}

// 	command := os.Args[1]

// 	switch command {
// 	case "add":
// 		task_names := os.Args[2:]
// 		if len(task_names) == 0 {
// 			fmt.Println("Not enough arguments. At least one task name is required to create a task after the ./todo add command.")
// 			return
// 		}
// 		add(task_names)
// 	case "done":
// 		task_numbers := os.Args[2:]
// 		if len(task_numbers) == 0 {
// 			fmt.Println("Not enough arguments. At least one task number is required to create a task after the ./todo done command.")
// 			return
// 		}
// 		complete(task_numbers)
// 	case "del":
// 		task_numbers := os.Args[2:]
// 		del(task_numbers)
// 	case "list":
// 		list()
// 	case "help":
// 		help_cmd := os.Args[2]
// 		switch help_cmd {
// 		case "add":
// 			fmt.Println("usage: add")
// 		}
// 		fmt.Println("TODO: implement help text for all the commands")
// 	case "rm":
// 		rm()
// 	default:
// 		fmt.Print(default_help_text)

// 		return
// 	}

// }
// func getDataPath() (filePath string, err error) {
// 	// Stores in
// 	// ~/.local/share/todo/tasks.json for Linux
// 	// /Library/Application\ Support/todo/tasks.json for macOS
// 	//  %AppData%\todo\tasks.json for Windows

// 	return xdg.DataFile("todo/tasks.json")
// }

// // returns a slice of Tasks from tasks.json and err.
// //
// //	If tasks.json doesn't exist, it is an empty slice of Tasks and not an error
// func loadTasks() (tasks []Task, err error) {

// 	dataPath, err := getDataPath()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create directory if it doesn't exist
// 	os.MkdirAll(filepath.Dir(dataPath), 0755)

// 	data, err := os.ReadFile(dataPath)

// 	if errors.Is(err, os.ErrNotExist) {
// 		return tasks, nil
// 	}

// 	if err != nil {
// 		//file-read error
// 		return nil, err
// 	}
// 	//else there is tasks.json, unmarshall it
// 	err = json.Unmarshal(data, &tasks)
// 	if err != nil {
// 		//error unmarshalling data
// 		return nil, err
// 	}

// 	return tasks, nil
// }

// // given tasks, saves tasks to tasks.json
// func saveTasks(tasks []Task) (err error) {
// 	dataPath, err := getDataPath()
// 	if err != nil {
// 		return err
// 	}
// 	// marshal data to json format
// 	b, err := json.MarshalIndent(tasks, "", "	")
// 	if err != nil {
// 		return err
// 	}

// 	err = os.WriteFile(dataPath, b, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Given a slice of task numbers, change the status of each task to done
// func complete(task_numbers []string) {

// 	current_tasks, err := loadTasks()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// make a map of task.id to index which looks up the tasks
// 	taskMap := make(map[int]int)
// 	for i, task := range current_tasks {
// 		taskMap[task.ID] = i
// 	}
// 	// for each task_num, grab the task it refers to and set its status to [x]
// 	for i := range task_numbers {
// 		taskNum, err := strconv.Atoi(task_numbers[i])
// 		if err != nil {
// 			panic(err)
// 		}

// 		// lookup the taskNum in taskMap
// 		// using the index works because the tasks.json is ordered
// 		index, exists := taskMap[taskNum]

// 		// if it exists, then change its status to x
// 		if exists {
// 			current_tasks[index].Status = "[x]"
// 		} else {
// 			// trying to delete a task with invalid id
// 			fmt.Printf("task %d does not exist in tasks.json\n", taskNum)
// 			return
// 		}
// 	}

// 	err = saveTasks(current_tasks)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Completed task(s)!")
// 	list()

// }

// // Given a slice of task numbers, delete each task from the tasks.json
// func del(task_numbers []string) {
// 	current_tasks, err := loadTasks()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// make a map of task.id to index which looks up the tasks
// 	taskMap := make(map[int]int)
// 	for i, task := range current_tasks {
// 		taskMap[task.ID] = i
// 	}
// 	// for each task_num, grab the task it refers to, and remove it
// 	for i := range task_numbers {
// 		taskNum, err := strconv.Atoi(task_numbers[i])
// 		if err != nil {
// 			panic(err)
// 		}

// 		// lookup the taskNum in taskMap
// 		// using the index works because the tasks.json is ordered
// 		index, exists := taskMap[taskNum]

// 		// if it exists, then remove the index at that slice wtih append
// 		// this appends eleements until the index, and elements after the index, effectively skipping the index
// 		if exists {
// 			current_tasks = append(current_tasks[:index], current_tasks[index+1:]...)
// 		} else {
// 			// trying to delete a task with invalid id
// 			fmt.Printf("task %d does not exist in tasks.json\n", taskNum)
// 			return
// 		}
// 	}
// 	// go through all the tasks and re-id them since they are out of order
// 	for i := range current_tasks {
// 		current_tasks[i].ID = i
// 	}
// 	err = saveTasks(current_tasks)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Deleted task(s)!")
// 	list()
// }

// // Output the tasks.json list to stdout in pretty format
// func list() {
// 	current_tasks, err := loadTasks()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// should not call list with empty tasks
// 	if len(current_tasks) == 0 {
// 		fmt.Println("Please create a task first with ./todo add <task_1> <task_2> ... ")
// 		os.Exit(1)
// 	}

// 	fmt.Println("Current list of tasks:")
// 	for _, t := range current_tasks {
// 		fmt.Printf("%s  %s\n", t.Status, t.Description)
// 	}
// }

// // removes the todo list from file storage
// func rm() {
// 	dataPath, err := getDataPath()
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = os.Remove(dataPath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Removed the todo list")
// }

package main

// TODO: make use of JSON for storage
import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)


const default_help_text string = `./todo is a todolist manager cli written in Go.

Usage:

        ./todo <command> [arguments]

The commands are:
		add			create a new task 
		done		complete a task 
		del			delete a task
		list		view all tasks 

Use "./todo help <command>" for more information about a command.

`

// have to use uppercase First letter to make the Task fields visible to the encoding/json package
// In Go, uppercase first letter means exported variable that is visible to other packages 
type Task struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
    Status      string `json:"status"`
}

func main() {
	// use os.args to get the commands
	// TODO: list of commands (add task_1 task_2 ... , list, done [task number], del [task number])

	if len(os.Args) == 1 {
		fmt.Print(default_help_text)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		task_names := os.Args[2:]
		if len(task_names) == 0 {
			fmt.Println("Not enough arguments. At least one task name is required to create a task after the ./todo command.")
			return
		}
		add(task_names)
	case "done":
		task_numbers := os.Args[2:]
		complete(task_numbers)
	case "del":
		task_numbers := os.Args[2:]
		del(task_numbers)
	case "list":
		list()
	default:
		fmt.Print(default_help_text)
		return
	}

}

// Given a slice of strings, add tasks with status of pending into tasks.json
func add(tasks []string) {

	start_id := 0

	_ , err := os.OpenFile("tasks.json", os.O_RDONLY, 0644)

	// tasks.json does not exist, write to file
	if errors.Is(err, os.ErrNotExist) { 

		file, err := os.Create("tasks.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file.Close()

		start_id = 1

		var new_tasks []Task

		
		for _ , v := range tasks {
			var new_task Task
			new_task.Description = v
			new_task.ID = start_id 
			new_task.Status = "[ ]"

			start_id += 1
			new_tasks = append(new_tasks, new_task)
		}
		
		// marshal data to json format
		b, err := json.MarshalIndent(new_tasks, "", "	")
		if err != nil {
			panic(err)
		}

		err = os.WriteFile("tasks.json", b, 0644)
		if err != nil {
			panic(err)
		}
		
	} else { // tasks.json does exist 
		//
		data, err := os.ReadFile("tasks.json")
		if err != nil {
			panic(err)
		}
		var current_tasks []Task 

		err = json.Unmarshal(data, &current_tasks)
		if err != nil {
			panic(err)
		}

		// get the latest ID plus one to start the new tasks
		start_id = current_tasks[len(current_tasks) - 1].ID + 1

		var new_tasks []Task
		
		for _ , v := range tasks {
			var new_task Task
			new_task.Description = v
			new_task.ID = start_id 
			new_task.Status = "[ ]"

			start_id += 1
			new_tasks = append(new_tasks, new_task)
		}
		
		// combine old tasks with new tasks
		current_tasks = append(current_tasks, new_tasks...)

		// marshal data to json format
		b, err := json.MarshalIndent(current_tasks, "", "	")
		if err != nil {
			panic(err)
		}

		err = os.WriteFile("tasks.json", b, 0644)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Added tasks to tasks.json!")
}

// Given a slice of task numbers, change the status of each task to done
func complete(task_numbers []string) {

}

// Given a slice of task numbers, delete each task from the tasks.json
func del(task_numbers []string) {

}

// Output the tasks.json list to stdout in pretty format 
func list() {

	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Printf("%s. Please create a task first with ./todo add <task_1> <task_2> ... \n", err)
		return
	}

	//unmarshal json 
	var tasks []Task 

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}	

	fmt.Println("Current list of tasks:")
	for _, t := range tasks {
		fmt.Printf("%s  %s\n",t.Status, t.Description)
	}
}

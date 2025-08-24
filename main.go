package main

// TODO: make use of JSON for storage
import (
	"bufio"
	"fmt"
	"os"
)


type Task struct {
	num int
	name string
	status string
}
func main() {
	// use os.args to get the commands 
	// TODO: list of commands (add "task, ... ," , list, done [task number], del [task number])
	
	if len(os.Args) == 1 {
		//print the default help thingy
		fmt.Println("TODO: default help thingy")
		return 
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
		fmt.Println("Added task to tasks.txt!")
	case "done":
		task_numbers := os.Args[2:]
		complete(task_numbers)
	case "del":
		task_numbers := os.Args[2:]
		del(task_numbers)
	case "list":
		list()
	default:
		fmt.Println("Command not recognized. Call ./todo --help for more information.")
		return 
	}

}

// Given a slice of strings, add tasks with status of pending into tasks.json
func add(tasks []string) {

	// TODO: make new IDs from the tasks.json 
	
	var new_tasks []Task

	for i, v := range tasks {
		var new_task Task
		new_task.name = v 
		new_task.num = i + 1
		new_task.status = "pending"

		new_tasks = append(new_tasks, new_task)
	}

	// create a new file
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	// write each new task to txt file 
	for _, v := range new_tasks {
		var status_box string
		if v.status == "pending" {
			status_box = "[ ]"
		} else {
			status_box = "[x]"
		}
		task_string :=  status_box + " " + v.name  + "\n"
		file.WriteString(task_string)
	}

	file.Close()
}

// Given a slice of task numbers, change the status of each task to done 
func complete(task_numbers []string) {

}

// Given a slice of task numbers, delete each task from the tasks.json
func del(task_numbers []string) {

}

// Output the current task list to terminal 
func list() {

	file, err := os.Open("tasks.txt") 
	if err != nil {
		fmt.Printf("%s. Please create a task first with ./todo add <task_1> <task_2> ... \n", err)
		return 
	}
	
	scanner := bufio.NewScanner(file)

	fmt.Println("Current list of tasks:")
	for (scanner.Scan()) {
		fmt.Println(scanner.Text())
	}
	file.Close()
	
}
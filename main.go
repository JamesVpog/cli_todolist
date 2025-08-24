package main

import (
	"fmt"
	"os"
)

func main() {
	// use os.args to get the commands 
	// TODO: list of commands (add "task desc", list, done [task number], del [task number])
	
	if len(os.Args) == 1 {
		//print the default help thingy
		fmt.Println("TODO: default help thingy")
		return 
	}
	
	command := os.Args[1]

	
	switch command {
	case "list":
		list()
	default:
		fmt.Println("Command not recognized. Call ./todo --help for more information.")
		return 
	}

}

// Output the current task list to terminal 
func list() {
	// look for json file in current directory 
	file, err := os.Open("tasks.json") 
	if err != nil {
		fmt.Printf("%s. Please create a task first with ./todo create", err)
		return 
	}
	
	fmt.Println("Found file, tasks.json!")
	file.Close()
	
}
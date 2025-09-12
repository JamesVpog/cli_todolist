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
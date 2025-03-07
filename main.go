package main

import (
	"fmt"
	"os"
)

func main() {
	// Get cmd line args
	var args []string = os.Args[1:]

	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "add":
		fmt.Println(args[0])
	case "list":
		fmt.Println(args[0])
	case "complete":
		fmt.Println(args[0])
	case "delete":
		fmt.Println(args[0])
	default:
		fmt.Println("Not a valid command:", args[0])
	}

	return
}
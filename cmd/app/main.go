package main

import (
	"fmt"
	"os"
	"todoapp/internal/cmd"
	"todoapp/internal/model"
	"todoapp/internal/utils"
)

func main() {
	// Default task file path
	var filePath string = "./list.csv"

	// Get cmd line args
	var args []string = os.Args[1:]

	// Check if no args passed
	if err := hasNoArgs(args); err != nil {
		fmt.Println(err)
		return
	}

	// Try to read file and check if error reading
	f, err := utils.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read data from csv
	data, err := utils.ReadCsv(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close file
	utils.CloseFile(f)

	// Parse data rows into todo models
	var s []model.Todo = utils.ParseData(data)

	// Run command that user input
	switch args[0] {
	case "add":
		cmd.AddTodo(args, s, filePath)
		
	case "list":
		cmd.DisplayTodos(s)
		
	case "complete":
		cmd.Complete(args, s, filePath)

	case "delete":
		cmd.Delete(args, s, filePath)

	default:
		fmt.Println("Not a valid command:", args[0])
	}
}


func hasNoArgs(args []string) error {
	if n := len(args); n == 0 {
		return fmt.Errorf("error: received no command")
	}
	return nil
}
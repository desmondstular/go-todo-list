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
		panic(err)
	}

	// Try to read file and check if error reading
	f, err := utils.OpenFile(filePath)
	if err != nil {
		panic(err)
	}

	// Read data from csv
	data, err := utils.ReadCsv(f)
	if err != nil {
		panic(err)
	}

	// Close file
	utils.CloseFile(f)

	// Parse data rows into model.Todo structs
	var s []model.Todo = utils.ParseData(data)

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
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	// "strconv"
	"syscall"
)

type todo struct {
	id string
	description string
	createdAt string
	isComplete string
}

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
	f, err := openFile(filePath)
	if err != nil {
		panic(err)
	}
	
	// Defer file close
	defer closeFile(f)

	// Read data from csv
	data, err := readCsv(f)
	if err != nil {
		panic(err)
	}

	// Parse data rows into todo structs
	s := parseData(data)
	
	fmt.Println(s[0].id)

	switch args[0] {
	case "add":
		fmt.Println(args[0])
		if err := addTask("test", data); err != nil {
			panic(err)
		}
	case "list":
		fmt.Println(args[0])
	case "complete":
		fmt.Println(args[0])
	case "delete":
		fmt.Println(args[0])
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


func openFile(filePath string) (*os.File, error) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	// Exclusive lock obtained on file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}


func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}


func readCsv(f *os.File) ([][]string, error) {
	reader := csv.NewReader(f)
	return reader.ReadAll()
}

func parseData(data [][]string) []todo {
	var s []todo = make([]todo, 0, 3)

	for _, row := range data[1:] {
		s = append(s, todo{row[0], row[1], row[2], row[3]})
	}

	return s
}


func addTask(description string, data [][]string) error {
	
	return nil
}


func getNewId(s []todo) int {
	
	return 1
}
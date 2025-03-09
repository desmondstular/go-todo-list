package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

type Todo struct {
	Id string
	Description string
	CreatedAt time.Time
	IsComplete bool
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

	// Parse data rows into Todo structs
	s := parseData(data)
	

	switch args[0] {
	case "add":
		fmt.Println(args[0])

		if len(args) < 2 {
			panic("no description was passed")
		}
		addTask(args[1], s)
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

func parseData(data [][]string) []Todo {
	var s []Todo = make([]Todo, 0, 3)

	for _, row := range data[1:] {
		if timeStr, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", row[2]); err != nil {
			panic(err)
		} else {
			if comp, err := strconv.ParseBool(row[3]); err != nil {
				panic(err)
			} else {
				s = append(s, Todo{row[0], row[1], timeStr, comp})
			}
		}
	}

	return s
}

func addTask(description string, list []Todo) {
	var id string = getNewId(list)
	now := time.Now()
	newTd := Todo{id, description, now, false}
	fmt.Println(newTd)
}

func getNewId(s []Todo) string {
	// Find current max Id
	var max int = 0

	for _, i := range s {
		temp, e := strconv.Atoi(i.Id)

		if e == nil && temp > max {
			max = temp
		}
	}

	return strconv.Itoa(max+1)
}
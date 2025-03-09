package cmd

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
	"todoapp/internal/model"
	"todoapp/internal/utils"
)


func AddTodo(args []string, sl []model.Todo, filePath string) {
	if len(args) > 1 {
		var td model.Todo = CreateTask(args[1], sl)
		AddTaskToCsv(td, filePath)
	} else {
		panic("no description was passed")
	}
}


func AddTaskToCsv(td model.Todo, filePath string) {
	// Try to read file and check if error reading
	f, err := utils.OpenFile(filePath)
	if err != nil {
		panic(err)
	}

	defer utils.CloseFile(f)

	// Seek to end
	f.Seek(0, io.SeekEnd)

	// Write to file
	w := csv.NewWriter(f)
	sl := td.ToSlice()
	w.Write(sl)
	w.Flush()
}


func CreateTask(description string, list []model.Todo) model.Todo {
	var id string = GetNewTaskId(list)
	now := time.Now()
	newTd := model.Todo{id, description, now, false}
	return newTd
}


func GetNewTaskId(s []model.Todo) string {
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
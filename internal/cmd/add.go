package cmd

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"todoapp/internal/model"
)


func AddTodo(args []string, sl []model.Todo, f *os.File) {
	if len(args) > 1 {
		var td model.Todo = CreateTask(args[1], sl)
		AddTaskToCsv(f, td)
	} else {
		panic("no description was passed")
	}
}


func AddTaskToCsv(f *os.File, td model.Todo) {
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
package cmd

import (
	"encoding/csv"
	"fmt"
	// "io"
	"os"
	"todoapp/internal/model"
	"todoapp/internal/utils"
)

func Delete(args []string, s []model.Todo, filePath string) {
	if len(args) < 2 {
		panic("no id was passed for deletion")
	}

	// Grab id from args to delete
	id := args[1]

	if s, err := deleteTodoFromSlice(s, id); err != nil {
		fmt.Println(err)
		return
	} else {
		overwriteCsv(s, filePath)
	}
}

func deleteTodoFromSlice(s []model.Todo, id string) ([]model.Todo, error) {
	for idx, td := range s {
		if td.Id == id {
			s = append(s[:idx], s[idx+1:]...)
			return s, nil
		}
	}
	return s, fmt.Errorf("Could not find row with that id")
}

func overwriteCsv(s []model.Todo, filePath string) {
	// Truncate file to clear contents
	if err := os.Truncate(filePath, 0); err != nil {
		panic("overwriteCsv(): could not truncate file")
	}

	// Attempt to open file
	f, err := utils.OpenFile(filePath)
	if err != nil {
		panic("overwriteCsv: unable to open list file")
	}

	defer utils.CloseFile(f)
	
	// Create new writer and write headers
	w := csv.NewWriter(f)
	w.Write(model.Header())

	// Write each row
	for _, td := range s {
		w.Write(td.ToSlice())
	}

	w.Flush()
}
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"todoapp/internal/model"
	"todoapp/internal/utils"
)


func Complete(args []string, s []model.Todo, filePath string) {
	if len(args) < 2 {
		panic("no id was passed")
	}

	// Get id from args list
	id := args[1]

	if err := updateTodoDone(&s, id); err != nil {
		fmt.Println(err)
		return
	}

	// Update done column as complete
	if err := updateCsvDone(s, filePath); err != nil {
		fmt.Println(err)
	}
}


func updateTodoDone(s *[]model.Todo, id string) error {
	for idx, _ := range *s {
		if (*s)[idx].Id == id {
			if (*s)[idx].IsComplete == false {
				(*s)[idx].IsComplete = true
				return nil
			} else {
				return fmt.Errorf("task is already complete")
			}
			
		}
	}

	return fmt.Errorf("unable to mark task complete")
}


func updateCsvDone(s []model.Todo, filePath string) error {
	// Truncate file to clear contents
	if err := os.Truncate(filePath, 0); err != nil {
		return fmt.Errorf("unable to truncate file")
	}

	// Attempt to open file
	f, err := utils.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("overwriteCsv: unable to open list file")
	}

	// Create new writer and write headers
	w := csv.NewWriter(f)
	w.Write(model.Header())

	// Write each row
	for _, td := range s {
		w.Write(td.ToSlice())
	}

	w.Flush()

	return nil
}
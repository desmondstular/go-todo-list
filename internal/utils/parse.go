package utils

import (
	"strconv"
	"time"
	"todoapp/internal/model"
)

func ParseData(data [][]string) []model.Todo {
	var s []model.Todo = make([]model.Todo, 0, 3)

	if len(data) > 1 {
		for _, row := range data[1:] {
			if timeStr, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", row[2]); err != nil {
				panic(err)
			} else {
				if comp, err := strconv.ParseBool(row[3]); err != nil {
					panic(err)
				} else {
					s = append(s, model.Todo{row[0], row[1], timeStr, comp})
				}
			}
		}	
	}
	
	return s
}
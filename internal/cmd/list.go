package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"todoapp/internal/model"
)

func DisplayTodos(s []model.Todo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
	for _, td := range s {
		var tdStr string = td.FormatString()
		fmt.Fprintln(w, tdStr)
	}

	w.Flush()
}
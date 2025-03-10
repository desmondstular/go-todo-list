package model

import (
	"fmt"
	"github.com/mergestat/timediff"
	"strconv"
	"time"
)

type Todo struct {
	Id string
	Description string
	CreatedAt time.Time
	IsComplete bool
}

func (t Todo) ToSlice() []string {
	sl := make([]string, 4)
	timeStr := t.CreatedAt.Format("2006-01-02 15:04:05.999999 -0700 MST")
	sl[0] = t.Id
	sl[1] = t.Description
	sl[2] = timeStr
	sl[3] = strconv.FormatBool(t.IsComplete)
	return sl
}

func (t Todo) FormatString() string {
	timeStr := timediff.TimeDiff(t.CreatedAt)
	boolStr := strconv.FormatBool(t.IsComplete)
	return fmt.Sprintf("%v\t%v\t%v\t%v", t.Id, t.Description, timeStr, boolStr)
}

func Header() []string {
	return []string {"Id", "Description", "CreatedAt", "Done"}
}

func HeaderString() string {
	return fmt.Sprintf("ID,Description,CreatedAt,Done\n")
}
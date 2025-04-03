package models

import (
	"fmt"
	"time"
)

// Task TODO: add child tasks
// Task TODO: add priority with an enum
type Task struct {
	Id       uint `gorm:"primary_key"`
	Title    string
	Desc     string
	Deadline time.Time
	Finished bool `gorm:"default:false"`
}

func NewTask(title string, desc string, deadline time.Time, finished bool) *Task {
	return &Task{
		Title:    title,
		Desc:     desc,
		Deadline: deadline,
		Finished: finished,
	}
}

func (t Task) String() string {
	return fmt.Sprintf(
		"[#%d] %s\n  Desc: %s\n  Deadline: %s\n  Finished: %t\n",
		t.Id,
		t.Title,
		t.Desc,
		t.Deadline.Format("2006-01-02 15:04"),
		t.Finished,
	)
}

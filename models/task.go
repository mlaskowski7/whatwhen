package models

import "time"

// Task TODO: add child tasks
type Task struct {
	Id       uint `gorm:"primary_key"`
	Title    string
	Desc     string
	Deadline time.Time
	Finished bool `gorm:"default:false"`
}

package task

import (
	"time"
)

// type 'Task' represents the complete task model ready for consume or insert
type Task struct {
	ID     string     `json:"id"`
	Emoji  string     `json:"emoji"`
	Name   string     `json:"name"`
	Due    *time.Time `json:"due"`
	Done   bool       `json:"done"`
	Active bool       `json:"active"`
}

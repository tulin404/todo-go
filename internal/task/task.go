package task

import (
	"time"

	"github.com/google/uuid"
)

// type 'Task' represents the complete task model ready for consume or insert
type Task struct {
	ID     uuid.UUID  `json:"id"`
	Name   string     `json:"name"`
	Due    *time.Time `json:"due"`
	Done   bool       `json:"done"`
	Active bool       `json:"active"`
}

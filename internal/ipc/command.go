package ipc

import (
	"github.com/tulin404/todo-go/internal/task"
)

// type 'Command' is a DTO struct for IPC between todo daemon <-> todo program
type Command struct {
	Command  string 	`json:"command"`
	Task     task.Task	`json:"task"`
}

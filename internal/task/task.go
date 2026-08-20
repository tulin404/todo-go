package task

import (
	"time"

	"github.com/google/uuid"
	"github.com/tulin404/todo-go/internal/storage"
)

type InputTask struct {
	Name   string     `json:"name"`
	Due    *time.Time `json:"due"`
	Done   bool       `json:"done"`
	Active bool       `json:"active"`
}

type Task struct {
	ID     uuid.UUID  `json:"id"`
	Name   string     `json:"name"`
	Due    *time.Time `json:"due"`
	Done   bool       `json:"done"`
	Active bool       `json:"active"`
}

func Add(inputTask InputTask, due *time.Time) {
	file := storage.VerifyStorageFile()
	defer file.Close()

	newTask := Task{
		ID:     uuid.New(),
		Name:   inputTask.Name,
		Due:    due,
		Done:   false,
		Active: true,
	}

	storage.SaveTask(file, newTask)
}

package task

import (
	"time"

	"github.com/google/uuid"
	"github.com/tulin404/todo-go/internal/storage"
)

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

	storage.Save(file, newTask)
}

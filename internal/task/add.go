package task

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tulin404/todo-go/internal/storage"
)

type AddInput struct {
	Name string     `json:"name"`
	Due  *time.Time `json:"due"`
}

func Add(input AddInput) error {
	file, error := storage.VerifyStorageFile()
	if error != nil {
		return fmt.Errorf("failed to add task:")
	}
	defer file.Close()

	newTask := Task{
		ID:     uuid.New(),
		Name:   input.Name,
		Due:    input.Due,
		Done:   false,
		Active: true,
	}

	return storage.Save(file, newTask)
}

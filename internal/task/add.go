package task

import (
	"fmt"
	"time"

	"github.com/tulin404/todo-go/internal/storage"
)

var currentID uint64 = 0

// type 'AddInput' is a simple DTO for standardization and type safety
type AddInput struct {
	Emoji string
	Name  string
	Due   *time.Time
}

// 'Add' opens the task file (for storage) and adds a task
func Add(input AddInput) error {
	newTask := Task{
		ID:     currentID + 1,
		Emoji:  input.Emoji,
		Name:   input.Name,
		Due:    input.Due,
		Done:   false,
		Active: true,
	}

	if err := storage.Save(newTask); err != nil {
		return fmt.Errorf("failed to save tasks file: %v", err)
	}

	return nil
}

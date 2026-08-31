package task

import (
	"fmt"
	"time"

	"github.com/tulin404/todo-go/internal/helpers"
	"github.com/tulin404/todo-go/internal/storage"
)

// type 'AddInput' is a simple DTO for standardization and type safety
type AddInput struct {
	Emoji string
	Name  string
	Due   *time.Time
}

// 'Add' opens the task file (for storage) and adds a task
func Add(input AddInput) error {
	id, err := helpers.GenerateID()
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	newTask := Task{
		ID:     id,
		Emoji:  input.Emoji,
		Name:   input.Name,
		Due:    input.Due,
		Done:   false,
		Active: true,
	}

	if err := storage.Save(newTask); err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	return nil
}

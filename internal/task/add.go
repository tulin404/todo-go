package task

import (
	"fmt"
	"time"

	"github.com/tulin404/todo-go/internal/helpers"
	"github.com/tulin404/todo-go/internal/storage"
	"github.com/tulin404/todo-go/internal/timeutil"
)

// type 'AddInput' is a simple DTO for standardization and type safety
type AddInput struct {
	Icon string
	Name string
	Due  string
}

// 'Add' opens the task file (for storage) and adds a task
func Add(input AddInput) error {
	id, err := helpers.GenerateID()
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	var canParse bool
	var parsedDue *time.Time

	if input.Due != "" {
		canParse = true
	}

	if canParse {
		parsedDue, err = timeutil.FromNow(input.Due)
		if err != nil {
			return fmt.Errorf("failed to add task: %w", err)
		}
	}
	newTask := Task{
		ID:     id,
		Icon:   input.Icon,
		Name:   input.Name,
		Due:    parsedDue,
		Done:   false,
		Active: true,
	}

	if err := storage.Save(newTask); err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	return nil
}

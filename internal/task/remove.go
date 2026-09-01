package task

import (
	"fmt"

	"github.com/tulin404/todo-go/internal/storage"
)

func Remove(id string) error {
	if err := storage.Remove(id); err != nil {
		return fmt.Errorf("failed to remove task: %w", err)
	}

	return nil
}

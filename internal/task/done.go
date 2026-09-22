package task

import (
	"fmt"

	"github.com/tulin404/todo-go/internal/storage"
)

// Good redundancy for SoC
// 'Done' is responsible for calling the Done method from the storage package and treating its possible errors
func Done(id string) error {
	if err := storage.Done(id); err != nil {
		return fmt.Errorf("failed to mark task as done: %w", err)
	}

	return nil
}

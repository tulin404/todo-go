package task

import (
	"fmt"

	"github.com/tulin404/todo-go/internal/storage"
)

// Good redundancy for SoC
// 'Reset' is responsible for calling the Reset method from the storage package and treating its possible errors
func Reset() error {
	if err := storage.Reset(); err != nil {
		return fmt.Errorf("failed to reset tasks file: %w", err)
	}

	return nil
}

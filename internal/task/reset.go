package task

import (
	"fmt"

	"github.com/tulin404/todo-go/internal/storage"
)

func Reset() error {
	if err := storage.Reset(); err != nil {
		return fmt.Errorf("failed to reset tasks file: %v", err)
	}

	return nil
}

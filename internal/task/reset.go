package task

import (
	"fmt"

	"github.com/tulin404/todo-go/internal/storage"
)

func Reset() error {
	file, err := storage.VerifyStorageFile()
	if err != nil {
		return fmt.Errorf("failed to reset tasks: %v", err)
	}
	defer file.Close()

	return storage.Reset()
}

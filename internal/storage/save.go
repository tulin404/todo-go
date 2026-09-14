package storage

import (
	"fmt"
	"encoding/json"
)

// 'Save' edits the tasks file and saves the received task
func Save(task any) error {
	file, err := VerifyStorageFile()
	if err != nil {
		return fmt.Errorf("failed to verify tasks file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(task); err != nil {
		return fmt.Errorf("failed to write new task on tasks.json\n %v", err)
	}

	return nil
}

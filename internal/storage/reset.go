package storage

import "fmt"

// 'Reset' resets the tasks file, wiping out every task
func Reset() error {
	file, err := VerifyStorageFile()
	if err != nil {
		return fmt.Errorf("failed to verify tasks file: %w", err)
	}
	defer file.Close()

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to reset file pointer: %w", err)
	}

	return nil
}

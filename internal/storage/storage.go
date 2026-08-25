package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// 'NewScanner' is an abstraction layer for no direct coupling between package task and direct data readers
func NewScanner(file *os.File) (*bufio.Scanner, error) {
	// ABSTRACTION LAYER
	return bufio.NewScanner(file), nil
}

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

// 'Next' scans the next line of the scanner selected file on demand
func Next(scanner *bufio.Scanner) ([]byte, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("failed to read tasks file:\n %v", err)
		}

		return nil, nil
	}

	line := scanner.Bytes()

	if len(line) <= 0 {
		return nil, nil
	}

	return line, nil
}

func Reset() error {
	file, err := VerifyStorageFile()
	if err != nil {
		return fmt.Errorf("failed to verify tasks file: %v", err)
	}
	defer file.Close()

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %v", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to reset file pointer: %v", err)
	}

	return nil
}

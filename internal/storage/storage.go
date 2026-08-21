package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// 'NewScanner' is an abstraction layer for no direct coupling between package task and direct data readers
func NewScanner(file *os.File) *bufio.Scanner {
	// ABSTRACTION LAYER
	return bufio.NewScanner(file)
}

// 'Save' edits the tasks file and saves the received task
func Save(file *os.File, task any) error {
	encoder := json.NewEncoder(file)

	if err := encoder.Encode(task); err != nil {
		return fmt.Errorf("failed to write new task on tasks.json\n %v", err)
	}

	return nil
}

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

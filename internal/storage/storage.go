package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 'NewScanner' is an abstraction layer for no direct coupling between package task and direct data readers
func NewScanner(file *os.File) (*bufio.Scanner, error) {
	// ABSTRACTION LAYER
	return bufio.NewScanner(file), nil
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

// 'Remove' identifies and removes certain task by its ID
func Remove(id string) error {
	file, err := VerifyStorageFile()
	if err != nil {
		return fmt.Errorf("failed to verify tasks file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		var task struct {
			ID string `json:"id"`
		}

		if err := json.Unmarshal([]byte(line), &task); err != nil {
			return fmt.Errorf("failed to parse json: %w", err)
		}

		if task.ID == id {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
        return fmt.Errorf("failed to scan tasks file: %w", err)
    }

    file.Close()

    // EVEN THOUGHT THIS APPROACH IS LESS MEMORY EFFICIENT, IT REDUCES SYSTEM CALLS AND MAKES SURE THAT THE FILE WILL BE OVERWRITTEN ONLY IF SCAN SUCCEDS
    return os.WriteFile(
    	file.Name(),
     	[]byte(strings.Join(lines, "\n")+"\n"),
      	0644,
    )
}

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

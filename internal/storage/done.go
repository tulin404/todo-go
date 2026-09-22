package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 'Done' identifies and marks certain task as done by its id
func Done(id string) error {
	file, err := VerifyStorageFile()
	if err != nil {
		return fmt.Errorf("failed to verify tasks file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		// NIL MAP
		// USING MAP FOR MARSHALING EVERY STRUCT FIELD
		var taskData map[string]any

		if err := json.Unmarshal([]byte(line), &taskData); err != nil {
			return fmt.Errorf("failed to parse json: %w", err)
		}

		if taskData["id"] == id {
			taskData["done"] = true
			updated, err := json.Marshal(taskData)
			if err != nil {
				return fmt.Errorf("failed to marshal task: %w", err)
			}

			line = string(updated)
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan tasks file: %w", err)
	}

	// EVEN THOUGHT THIS APPROACH IS LESS MEMORY EFFICIENT, IT REDUCES SYSTEM CALLS AND MAKES SURE THAT THE FILE WILL BE OVERWRITTEN ONLY IF SCAN SUCCEDS
	return os.WriteFile(
		file.Name(),
		[]byte(strings.Join(lines, "\n")+"\n"),
		0644,
	)
}

package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Notified(id string) error {
	file, err := VerifyStorageFile()
	if err != nil {
		return fmt.Errorf("failed to verify tasks file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		var task struct {
			ID       string `json:"id"`
			Notified bool   `json:"notified"`
		}

		if err := json.Unmarshal([]byte(line), &task); err != nil {
			return fmt.Errorf("failed to parse json: %w", err)
		}

		if task.ID == id {
			task.Notified = true
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

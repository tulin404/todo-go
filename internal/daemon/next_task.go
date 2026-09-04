package daemon

import (
	"encoding/json"
	"fmt"

	"github.com/tulin404/todo-go/internal/storage"
	"github.com/tulin404/todo-go/internal/task"
)

// NextTask returns the unnotified task with the earliest due date, including overdue tasks.
func NextTask() (*task.Task, error) {
	file, err := storage.VerifyStorageFile()
	if err != nil {
		return nil, fmt.Errorf("failed to verify tasks file: %w", err)
	}

	scanner := storage.NewScanner(file)
	var pastTask task.Task

	for {
		line, err := storage.Next(scanner)

		if err != nil {
			return nil, fmt.Errorf("failed to refresh: %w", err)
		}

		if line == nil {
			break
		}

		var currentTask task.Task

		if err := json.Unmarshal(line, &currentTask); err != nil {
			fmt.Printf("\033[33mWarning:\033[0m invalid json line, skipping\n %v", err)
			continue
		}

		if currentTask.Due == nil || currentTask.Notified {
			continue
		}

		// NO PAST TASK, FIRST ITERATION
		if pastTask.Name == "" {
			pastTask = currentTask
			continue
		}

		if currentTask.Due.Before(*pastTask.Due) {
			pastTask = currentTask
		}
	}

	return &pastTask, nil
}

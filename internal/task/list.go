package task

import (
	"encoding/json"
	"fmt"

	"github.com/tulin404/todo-go/internal/storage"
)

// 'List' lists all the tasks and can filter them
func List(filters []string) error {
	file, err := storage.VerifyStorageFile()
	if err != nil {
		return fmt.Errorf("failed to list tasks: %v", err)
	}
	defer file.Close()

	// ABSTRACTION WRAPPER FOR NOT DIRECTLY USING BUFIO
	scanner := storage.NewScanner(file)

	taskCount := 0

	// NOTHING STAYS IN MEMORY FOR PRINTING AFTER, EVERYTHING IS PRINTED AT RUNTIME
	fmt.Printf("%-6s %-30s\n", "EMOJI", "TASK")
	fmt.Printf("%-6s %-30s\n", "-----", "------------------------------")
	for {
		line, err := storage.Next(scanner)

		if err != nil {
			return fmt.Errorf("\033[1;31mFatal error:\033[0m failed to read tasks file\n %v", err)
		}

		if line == nil {
			break
		}

		var task Task

		if err := json.Unmarshal(line, &task); err != nil {
			fmt.Printf("\033[33mWarning:\033[0m invalid json line, skipping\n %v", err)
			continue
		}

		fmt.Printf("%-6s %-30s\n", task.Emoji, task.Name)
		taskCount++
	}

	if taskCount <= 0 {
		fmt.Println("YAY! You don't have nothing to do! Wait... would that be great?")
	}

	return nil
}

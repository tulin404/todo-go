package task

import (
	"encoding/json"
	"fmt"

	"github.com/tulin404/todo-go/internal/storage"
)

func ListRaw() ([]Task, error) {
	file, err := storage.VerifyStorageFile()
	if err != nil {
		return nil, fmt.Errorf("failed to verify tasks file: %v", err)
	}
	defer file.Close()

	scanner, err := storage.NewScanner(file)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %v", err)
	}

	var tasks []Task

	for {
		line, err := storage.Next(scanner)

		if err != nil {
			return nil, fmt.Errorf("failed to list task: %v", err)
		}

		if line == nil {
			break
		}

		var task Task

		if err := json.Unmarshal(line, &task); err != nil {
			fmt.Printf("\033[33mWarning:\033[0m invalid json line, skipping\n %v", err)
			continue
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// 'List' lists all the tasks and can filter them
func List(filters []string) error {
	file, err := storage.VerifyStorageFile()
	if err != nil {
		return fmt.Errorf("failed to verify tasks file: %v", err)
	}
	defer file.Close()

	// ABSTRACTION WRAPPER FOR NOT DIRECTLY USING BUFIO
	scanner, err := storage.NewScanner(file)
	if err != nil {
		return fmt.Errorf("failed to list tasks: %v", err)
	}

	taskCount := 0

	// NOTHING STAYS IN MEMORY FOR PRINTING AFTER, EVERYTHING IS PRINTED AT RUNTIME
	for {
		line, err := storage.Next(scanner)

		if err != nil {
			return fmt.Errorf("failed to list task: %v", err)
		}

		if line == nil {
			break
		}

		// JUST PRINTS IN THE FIRST ITERATION AND DOESNT PRINT IF THE USER DOESNT HAVE TASKS (after line == nil)
		if taskCount == 0 {
			fmt.Printf("%-6s %-30s\n", "EMOJI", "TASK")
			fmt.Printf("%-6s %-30s\n", "-----", "------------------------------")
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
	} else {
		fmt.Printf("\nYou have %d tasks\n", taskCount)
	}

	return nil
}

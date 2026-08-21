package task

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tulin404/todo-go/internal/storage"
)

// List LISTS ALL THE TASKS AND CAN FILTER THEM
func List(filters []string) {
	file := storage.VerifyStorageFile()

	// ABSTRACTION WRAPPER FOR NOT DIRECTLY USING BUFIO
	scanner := storage.NewScanner(file)

	taskCount := 0

	for {
		line, err := storage.Next(scanner)

		if err != nil {
			log.Fatalf("\033[1;31mFatal error:\033[0m failed to read tasks file\n %v", err)
		}

		if line == nil {
			break
		}

		var task Task

		if err := json.Unmarshal(line, &task); err != nil {
			fmt.Printf("\033[33mWarning:\033[0m invalid json line, skipping\n %v", err)
			continue
		}

		fmt.Println(task)
		taskCount++
	}

	if taskCount <= 0 {
		fmt.Println("YAY! You don't have nothing to do! Wait... would that be great?")
	}
}

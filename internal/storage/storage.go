package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tulin404/todo-go/internal/task"
)

func SaveTask(file *os.File, task task.Task) {
	encoder := json.NewEncoder(file)

	if err := encoder.Encode(task); err != nil {
		log.Fatalf("Fatal error: failed to write new task on tasks.json\n %v", err)
	}
}

func ListTasks(file *os.File, filters []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) <= 0 {
			continue
		}

		var task task.Task
		taskCount := 0

		if err := json.Unmarshal(line, &task); err != nil {
			fmt.Println(task.Name)
			taskCount++
		}

		if taskCount <= 0 {
			fmt.Println("YAY! You don't have nothing to do! Wait... would that be great?")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Fatal error: failed to read tasks file\n %v", err)
	}
}

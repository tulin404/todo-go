package storage

import (
	"encoding/json"
	"os"
)

func SaveTask(file *os.File, task any) {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(task); err != nil {
		panic("Fatal error: Failed to write new task on tasks.json.\n" + err.Error())
	}
}

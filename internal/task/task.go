package task

import (
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/tulin404/todo-go/internal/storage"
)

type InputTask struct {
	Name   string     `json:"name"`
	Due    *time.Time `json:"due"`
	Done   bool       `json:"done"`
	Active bool       `json:"active"`
}

type InsertTask struct {
	ID     uuid.UUID  `json:"id"`
	Name   string     `json:"name"`
	Due    *time.Time `json:"due"`
	Done   bool       `json:"done"`
	Active bool       `json:"active"`
}

func verifyStorageFile() *os.File {
	dirPath, err := os.UserCacheDir()
	if err != nil {
		panic("Fatal error: No cache directory on system.\n" + err.Error())
	}

	dirPath = filepath.Join(dirPath, "todo-go")

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		panic("Fatal error: Failed to create application's folder.\n" + err.Error())
	}

	filePath := filepath.Join(dirPath, "tasks.jsonl")

	file, err := os.OpenFile(
		filePath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic("Fatal error: Failed to open file " + filePath + ".\n" + err.Error())
	}

	return file
}

func Add(inputTask InputTask, due *time.Time) {
	file := verifyStorageFile()
	defer file.Close()

	newTask := InsertTask{
		ID:     uuid.New(),
		Name:   inputTask.Name,
		Due:    due,
		Done:   false,
		Active: true,
	}

	storage.SaveTask(file, newTask)
}

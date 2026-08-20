package task

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
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
	dirPath, dirPathErr := os.UserCacheDir()
	if dirPathErr != nil {
		panic("Fatal error: No cache directory on system.\n" + dirPathErr.Error())
	}

	dirPath = filepath.Join(dirPath, "todo-go")
	if err := os.Mkdir(dirPath, 0755); err != nil {
		panic("Fatal error: Failed to create applications's folder.\n" + err.Error())
	}

	filePath := filepath.Join(dirPath, "tasks.json")

	_, statErr := os.Stat(filePath)

	var (
		file    *os.File
		fileErr error
	)

	if errors.Is(statErr, os.ErrNotExist) {
		file, fileErr = os.Create(filePath)
	} else if statErr != nil {
		panic("Fatal error: 'tasks.json' located in '%s' is corrupted or obstructed.\n" + filePath)
	}

	if fileErr != nil {
		panic("Fatal error: Failed to create file" + filePath + ".\n" + fileErr.Error())
	}

	return file
}

func Add(inputTask InputTask, due *time.Time) {
	file := verifyStorageFile()

	newTask := InsertTask{
		ID:     uuid.New(),
		Name:   inputTask.Name,
		Due:    due,
		Done:   false,
		Active: true,
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(newTask); err != nil {
		panic("Fatal error: Failed to write new task on tasks.json.\n" + err.Error())
	}

	defer file.Close()
}

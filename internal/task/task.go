package task

import (
	"errors"
	"os"
	"path/filepath"
)

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func verifyFile() {
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
		file *os.File
		fileErr error
	)
	
	if errors.Is(statErr, os.ErrNotExist) {
		file, fileErr = os.Create(filePath)
	} else {
		panic("Fatal error: 'tasks.json' located in '%s' is corrupted or obstructed.\n" + filePath)
	}

	if fileErr != nil  {
		panic("Fatal error: Failed to create file" + filePath + ".\n" + fileErr.Error())
	}

	defer file.Close()
}

func Add() {
	
}

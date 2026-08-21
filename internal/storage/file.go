package storage

import (
	"log"
	"os"
	"path/filepath"
)

func VerifyStorageFile() *os.File {
	dirPath, err := os.UserCacheDir()
	if err != nil {
		log.Fatalf("Fatal error: no cache directory on system\n %v", err.Error())
	}

	dirPath = filepath.Join(dirPath, "todo-go")

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		log.Fatalf("Fatal error: failed to create application's folder\n %v", err)
	}

	filePath := filepath.Join(dirPath, "tasks.jsonl")

	file, err := os.OpenFile(
		filePath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatalf("Fatal error: failed to open file %s\n %v", filePath, err)
	}

	return file
}

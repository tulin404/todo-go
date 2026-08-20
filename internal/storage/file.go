package storage

import (
	"os"
	"path/filepath"
)

func VerifyStorageFile() *os.File {
	dirPath, err := os.UserCacheDir()
	if err != nil {
		panic("Fatal error: no cache directory on system.\n" + err.Error())
	}

	dirPath = filepath.Join(dirPath, "todo-go")

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		panic("Fatal error: failed to create application's folder\n" + err.Error())
	}

	filePath := filepath.Join(dirPath, "tasks.jsonl")

	file, err := os.OpenFile(
		filePath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic("Fatal error: failed to open file " + filePath + "\n" + err.Error())
	}

	return file
}

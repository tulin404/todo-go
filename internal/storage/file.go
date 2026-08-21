package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// 'userDataDir' returns the data directory of the user based on it's OS
func userDataDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		dir := os.Getenv("APPDATA")
		if dir == "" {
			return "", fmt.Errorf("APPDATA is not defined\n")
		}

		return dir, nil

	case "linux":
		dir := os.Getenv("XDG_DATA_HOME")
		if dir != "" {
			return dir, nil
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("\033[1;31mFatal error:\033[0m failed to get user's home directory: %v", err)
		}

		return filepath.Join(home, ".local", "share"), nil

	default:
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user's home directory: %v", err)
		}

		return home, nil
	}
}

/*
'VerifyStorageFile' verifies the storage file before any type of action that depends on it, returning a file descriptor for access
OBS: IT DOES NOT CLOSE THE FILE BY ITSELF
*/
func VerifyStorageFile() (*os.File, error) {
	dir, err := userDataDir()
	if err != nil {
		return nil, fmt.Errorf("failed to verify storage file\n %v", err)
	}

	dir = filepath.Join(dir, "todo-go")

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create application's folder\n %v", err)
	}

	filePath := filepath.Join(dir, "tasks.jsonl")

	file, err := os.OpenFile(
		filePath,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0644,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s\n %v", filePath, err)
	}

	return file, nil
}

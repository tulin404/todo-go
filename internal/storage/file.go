package storage

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func userDataDir() string {
	switch runtime.GOOS {
	case "windows":
		dir := os.Getenv("APPDATA")
		if dir == "" {
			log.Fatal("\033[1;31mFatal error:\033[0m APPDATA is not defined\n")
		}

		return dir

	case "linux":
		dir := os.Getenv("XDG_DATA_HOME")
		if dir != "" {
			return dir
		}

		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("\033[1;31mFatal error:\033[0m failed to get user's home directory: %v", err)
		}

		return filepath.Join(home, ".local", "share")

	default:
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("\033[1;31mFatal error:\033[0m failed to get user's home directory: %v", err)
		}

		return home
	}
}

func VerifyStorageFile() *os.File {
	dir := userDataDir()

	dir = filepath.Join(dir, "todo-go")

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("\033[1;31mFatal error:\033[0m failed to create application's folder\n %v", err)
	}

	filePath := filepath.Join(dir, "tasks.jsonl")

	file, err := os.OpenFile(
		filePath,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0644,
	)
	if err != nil {
		log.Fatalf("\033[1;31mFatal error:\033[0m failed to open file %s\n %v", filePath, err)
	}

	return file
}

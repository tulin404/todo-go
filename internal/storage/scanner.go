package storage

import (
	"bufio"
	"fmt"
	"os"
)

// 'NewScanner' is an abstraction layer for no direct coupling between package task and direct data readers
func NewScanner(file *os.File) (*bufio.Scanner) {
	// ABSTRACTION LAYER
	return bufio.NewScanner(file)
}

// 'Next' scans the next line of the scanner selected file on demand
func Next(scanner *bufio.Scanner) ([]byte, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("failed to read tasks file:\n %v", err)
		}

		return nil, nil
	}

	return scanner.Bytes(), nil
}

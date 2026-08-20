package main

import (
	"os"

	"github.com/tulin404/todo-go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

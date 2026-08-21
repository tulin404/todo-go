package main

import (
	"os"

	"github.com/tulin404/todo-go/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}

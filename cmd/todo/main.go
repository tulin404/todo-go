package main

import (
	"os"

	"github.com/tulin404/todo-go/internal/cli"
)

// Program's entry point
func main() {
	// GRANTS COBRA'S CONTROL OVER THE CLI EXECUTION FLOW
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"log"

	"github.com/tulin404/todo-go/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatalf("\033[1;31mFatal error:\033[0m %v", err)
	}
}

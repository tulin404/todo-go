package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo <command>")
		return
	}

	args := os.Args[1:]

	switch args {
		case "add":
		case "remove":
		case "done":
		case "list":
	}
}

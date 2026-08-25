package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

var due string

/*
var 'addCmd' represents the "add" subcommand and is directly linked to the rootCmd
Cobra's tree: todo -> add
*/
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a task",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		scanner := bufio.NewScanner(os.Stdin)

		var emoji string = "📝"

		fmt.Print("Type an emoji to this task (empty for default 📝): ")

		if scanner.Scan() {
			input := scanner.Text()

			if input != "" {
				emoji = input
			}
		}

		addDto := task.AddInput{
			Emoji: strings.TrimSpace(emoji),
			Name:  args[0],
		}

		task.Add(addDto)
		fmt.Println("\033[32m✔ \033[0m Task added")

		return nil
	},
}

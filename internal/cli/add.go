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
		var emoji string = "📝"
		fmt.Print("Type an emoji to this task (empty for default 📝): ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			input := scanner.Text()

			if input != "" {
				emoji = input
			}
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("failed to read input: %v", err)
		}

		addDto := task.AddInput{
			Emoji: strings.TrimSpace(emoji),
			Name:  args[0],
		}

		if err := task.Add(addDto); err != nil {
			return err
		}
		fmt.Println("\033[32m✔ \033[0m Task added")

		return nil
	},
}

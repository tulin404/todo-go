package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

/*
var 'resetCmd' represents the "reset" subcommand and is directly linked to the rootCmd
Cobra's tree: todo -> reset
*/
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Cleans all tasks",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("This will delete ALL your tasks, are you sure? [y/n] (n) ")

		var input string
		_, err := fmt.Scanln(&input)

		if err != nil && err.Error() != "unexpected newline" {
			return fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.ToLower(strings.TrimSpace(input))

		if input != "y" {
			fmt.Println("Operation canceled")
			return nil
		}

		if err := task.Reset(); err != nil {
			return err
		}

		fmt.Println("All the tasks were cleaned.")
		return nil
	},
}

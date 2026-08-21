package cli

import (
	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple todo cli.",

	RunE: func(cmd *cobra.Command, args []string) error {
		// THIS IMPLEMENTS A BASIC VERSION OF todo list
		if len(args) < 1 {
			task.List(nil)
		}

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand()
}

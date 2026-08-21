package cli

import (
	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

// var 'rootCmd' represents the root cli command ('todo') and serves as the entry point of the command tree.
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple todo cli.",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		// THIS IMPLEMENTS A BASIC VERSION OF todo list
		if len(args) < 1 {
			return task.List(nil)
		}

		return nil
	},
}

// 'Execute' is a wrapper that abstracts Cobra's execution workflow and initialize the CLI command tree using Cobra's .Execute()
func Execute() error {
	return rootCmd.Execute()
}

// 'init' appends the subcomands and flags to the its parent command
func init() {
	rootCmd.AddCommand(addCmd)
}

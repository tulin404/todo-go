package cli

import (
	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

var due string

var addCmd = &cobra.Command{
	Use:  "add [task]",
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		addDto := task.AddInput{
			Name: args[0],
		}

		task.Add(addDto)
		return nil
	},
}

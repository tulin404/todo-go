package cli

import (
	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

var removeCmd = &cobra.Command{
	Use:   "remove [id]",
	Short: "Completely removes a task",
	Args:  cobra.ArbitraryArgs,

	RunE: func(cmd *cobra.Command, args []string) error {

		task.Remove(args[0])
	},
}

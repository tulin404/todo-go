package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Cleans all tasks",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := task.Reset(); err != nil {
			return err
		}

		fmt.Println("All the tasks were cleaned.")
		return nil
	},
}

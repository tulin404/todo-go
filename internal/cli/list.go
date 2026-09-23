package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

// flags
var (
	done bool
	all  bool
)

/*
var 'listCmd' represents the "list" subcommand and is directly linked to the rootCmd
Cobra's tree: todo -> list
*/
var listCmd = &cobra.Command{
	Use:   "list [filter]",
	Short: "List tasks with filter",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		if done && all {
			return fmt.Errorf("--done and --all cannot be used together")
		}
		if done {
			return task.List("done")
		}
		if all {
			return task.List("all")
		}

		return task.List("")
	},
}

func init() {
	listCmd.Flags().BoolVar(&done, "done", false, "List completed tasks")
	listCmd.Flags().BoolVar(&done, "all", false, "List all tasks")
}

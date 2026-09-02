package cli

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

var removeCmd = &cobra.Command{
	Use:   "remove [id]",
	Short: "Completely removes a task",
	Args:  cobra.ArbitraryArgs,

	ValidArgsFunction: func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		tasks, err := task.ListRaw()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		toCompleteLower := strings.ToLower(toComplete)

		var completions []string
		for _, task := range tasks {
			hasPrefix := strings.HasPrefix(strings.ToLower(task.Name), toCompleteLower)
			alreadyTyped := slices.Contains(args, task.ID)

			if hasPrefix && !alreadyTyped {
				completions = append(
					completions,
					fmt.Sprintf("%s\t%s", task.ID, task.Name),
				)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := task.Remove(args[0]); err != nil {
			return err
		}
		fmt.Println("🗑️ Task removed")
		return nil
	},
}

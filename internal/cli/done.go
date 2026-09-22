package cli

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/task"
)

/*
var 'doneCmd' represents the "done" subcommand and is directly linked to the rootCmd
Cobra's tree: todo -> done
*/
var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Marks a task as done",
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
					fmt.Sprintf("%s\t%s %s", task.ID, task.Icon, task.Name),
				)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			if err := task.Done(id); err != nil {
				return err
			}
		}
		fmt.Println("✓ Task done. Congratulations!")
		return nil
	},
}

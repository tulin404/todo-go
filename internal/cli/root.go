package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple todo cli.",

	RunE: func(cmd *cobra.Command, args []string) error {
		// THIS IMPLEMENTS A BASIC VERSION OF todo list
		if len(args) < 1 {

		}

		fmt.Println(args[0])
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand()
}

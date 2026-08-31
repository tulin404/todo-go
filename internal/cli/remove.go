package cli

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use: "remove [task]",
	Short: "Completely removes a task",
	Args: cobra.ArbitraryArgs,

	RunE: func(cmd *cobra.Command, args []string) error {

	},
}

package cli

import (
	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/daemon"
	"github.com/tulin404/todo-go/internal/ipc"
)

/*
var 'removeCmd' represents the "daemon" subcommand and is directly linked to the rootCmd. It isn't suposed to be ran directly
Cobra's tree: todo -> daemon
*/
var daemonCmd = &cobra.Command{
	Use: "daemon",
	Short: "Background daemon for todo (start with systemctl)",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		listener, err := ipc.Listen()
		if err != nil {
			return err
		}
		defer listener.Close()

		task, err := daemon.NextTask()
		if err != nil {
			return err
		}

		timer := daemon.NewTimer(*task.Due)

		commands := make(chan ipc.Command)

		go ipc.AcceptCommands(listener, commands)

		for {
			select {
				case command := <- commands:
					newTask, err := daemon.HandleCommand(*task, command)
					if err != nil {
						return err
					}
					// NO NEW TASK FOR SUB
					if newTask == nil {
						break
					}

					timer = daemon.NewTimer(*newTask.Due)
			}
		}
	},
}

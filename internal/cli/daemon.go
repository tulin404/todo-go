package cli

import (
	"github.com/spf13/cobra"
	"github.com/tulin404/todo-go/internal/daemon"
	"github.com/tulin404/todo-go/internal/ipc"
)

/*
var 'removeCmd' represents the "daemon" subcommand and is directly linked to the rootCmd. It isn't supposed to be run directly
Cobra's tree: todo -> daemon
*/
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Background daemon for todo (start with systemctl)",
	Args:  cobra.NoArgs,

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
			// NEW COMMAND INCOME
			case command := <-commands:
				newTask, recalc := daemon.HandleCommand(*task, command)

				// add case
				// NO NEW TASK FOR SUB
				if newTask != nil {
					task = newTask
					timer = daemon.NewTimer(*newTask.Due)
					break
				}

				// remove case
				// TASK REMOVED, NEED TO RECALC
				if recalc {
					task, err = daemon.NextTask()
					if err != nil {
						return err
					}
				}

			// TIME TO NOTIFY
			case <-timer.C:
				err := daemon.Notify(*task)
				if err != nil {
					return err
				}

				newTask, err := daemon.NextTask()
				if err != nil {
					return err
				}

				if newTask == nil {
					// Não há mais tarefas para agendar
					timer = nil
					continue
				}

				task = newTask
				timer = daemon.NewTimer(*task.Due)
			}
		}
	},
}

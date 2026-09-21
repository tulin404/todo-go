package daemon

import (
	"github.com/tulin404/todo-go/internal/ipc"
	"github.com/tulin404/todo-go/internal/task"
)

func HandleCommand(
	currentTask task.Task,
	command ipc.Command,
) (
	*task.Task,
	bool,
) {
	switch command.Command {
	case "task_added":
		if command.Task.Due.Before(*currentTask.Due) {
			return &command.Task, false
		}
		return nil, false

	case "task_removed":
		if command.Task == currentTask {
			return nil, true
		}
		return nil, false

	default:
		return nil, false
	}
}

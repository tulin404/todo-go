package daemon

import (
	"github.com/tulin404/todo-go/internal/ipc"
	"github.com/tulin404/todo-go/internal/task"
)

func HandleCommand(currentTask task.Task, command ipc.Command) (*task.Task, error) {
	switch command.Command {
	case "task_added":
		if command.Task.Due.Before(*currentTask.Due) {
			return &command.Task, nil
		}
		return nil, nil

	default:
		return nil, nil
	}
}

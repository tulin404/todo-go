package daemon

import (
	"github.com/gen2brain/beeep"
	"github.com/tulin404/todo-go/internal/storage"
	"github.com/tulin404/todo-go/internal/task"
)

func Notify(task *task.Task) error {
	err := beeep.Notify(task.Name, "Task", task.Icon)
	if err != nil {
		return err
	}

	err = storage.Notified(task.ID)
	if err != nil {
		return err
	}

	return nil
}

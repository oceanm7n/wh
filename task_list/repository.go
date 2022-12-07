package task_list

import (
	"context"

	"github.com/oceanm7n/wh/models"
)

type Repository interface {
	GetTaskList(ctx context.Context) (models.TaskList, error)
	AddTask(ctx context.Context, task models.Task) error
}

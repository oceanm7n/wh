package usecase

import (
	"context"

	"github.com/oceanm7n/wh/models"
	"github.com/oceanm7n/wh/task_list"
)

type TaskListUsecase struct {
	userRepo task_list.Repository
}

func NewTaskListUsecase(userRepo task_list.Repository) *TaskListUsecase {
	return &TaskListUsecase{
		userRepo: userRepo,
	}
}

func (u *TaskListUsecase) GetTaskList(ctx context.Context) (models.TaskList, error) {
	return u.userRepo.GetTaskList(ctx)
}

func (u *TaskListUsecase) AddTask(ctx context.Context, task models.Task) error {
	return u.userRepo.AddTask(ctx, task)
}

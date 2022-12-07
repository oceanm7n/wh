package usecase

import (
	"context"

	"github.com/oceanm7n/wh/models"
	"github.com/stretchr/testify/mock"
)

type TaskListMock struct {
	mock.Mock
}

func (m *TaskListMock) GetTaskList(ctx context.Context) (models.TaskList, error) {
	m.Called()

	return models.TaskList{}, nil
}

func (m *TaskListMock) AddTask(ctx context.Context, task models.Task) error {
	args := m.Called(task)

	return args.Error(0)
}

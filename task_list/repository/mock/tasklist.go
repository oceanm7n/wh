package mock

import (
	"context"

	"github.com/oceanm7n/wh/models"
	"github.com/stretchr/testify/mock"
)

type TaskListStorageMock struct {
	mock.Mock
}

func (s *TaskListStorageMock) GetTaskList(ctx context.Context) (models.TaskList, error) {
	args := s.Called()
	return models.TaskList{}, args.Error(1)
}

func (m *TaskListStorageMock) AddTask(ctx context.Context, task models.Task) error {
	args := m.Called(task)

	return args.Error(0)
}

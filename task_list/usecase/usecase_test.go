package usecase

import (
	"context"
	"testing"

	"github.com/oceanm7n/wh/models"
	"github.com/oceanm7n/wh/task_list/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestTaskListBL(t *testing.T) {
	// create mock for repository
	mockRepo := new(mock.TaskListStorageMock)

	uc := NewTaskListUsecase(mockRepo)

	mockRepo.On("GetTaskList").Return(models.TaskList{}, nil)
	_, err := uc.GetTaskList(context.Background())
	assert.NoError(t, err)

	mockRepo.On("AddTask", models.Task{}).Return(nil)
	err = uc.AddTask(context.Background(), models.Task{})
	assert.NoError(t, err)
}

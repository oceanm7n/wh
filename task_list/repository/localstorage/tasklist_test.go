package localstorage

import (
	"context"
	"testing"

	"github.com/oceanm7n/wh/models"
)

func TestLocalStorage(t *testing.T) {
	s := NewTasklistLocalStorage()

	newTask := models.Task{
		ID:   "1",
		Name: "test",
	}

	// try to add task and check that no error
	err := s.AddTask(context.Background(), newTask)
	if err != nil {
		t.Error("Error adding task")
	}

	// try to get task list and check that no error
	taskList, err := s.GetTaskList(context.Background())
	if err != nil {
		t.Error("Error getting task list")
	}

	// check that newTask is same as taskList.Tasks[0]
	if taskList.Tasks[0] != newTask {
		t.Error("Error: taskList.Tasks[0] != newTask")
	}
}

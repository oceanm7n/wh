package localstorage

import (
	"context"
	"sync"

	"github.com/oceanm7n/wh/models"
)

type TasklistLocalStorage struct {
	mutex    *sync.Mutex
	tasklist *models.TaskList
}

func NewTasklistLocalStorage() *TasklistLocalStorage {
	return &TasklistLocalStorage{
		mutex:    &sync.Mutex{},
		tasklist: &models.TaskList{},
	}
}

func (t *TasklistLocalStorage) GetTaskList(ctx context.Context) (models.TaskList, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return *t.tasklist, nil
}

func (t *TasklistLocalStorage) AddTask(ctx context.Context, task models.Task) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.tasklist.Tasks = append(t.tasklist.Tasks, task)

	return nil
}

package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oceanm7n/wh/models"
	"github.com/oceanm7n/wh/task_list"
)

type TaskListHandler struct {
	tasklistUC task_list.UseCase
}

func NewTaskListHandler(tasklistUC task_list.UseCase) *TaskListHandler {
	return &TaskListHandler{
		tasklistUC: tasklistUC,
	}
}

func (h *TaskListHandler) GetTaskList(c *gin.Context) {
	tasklist, err := h.tasklistUC.GetTaskList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tasklist)
}

func (h *TaskListHandler) AddTask(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.tasklistUC.AddTask(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, task)
}

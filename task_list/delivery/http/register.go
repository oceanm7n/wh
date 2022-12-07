package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oceanm7n/wh/task_list"
)

func RegisterHTTPEndpoints(router *gin.Engine, tasklistUC task_list.UseCase) {
	tasklistHandler := NewTaskListHandler(tasklistUC)
	router.GET("/tasklist", tasklistHandler.GetTaskList)
	router.POST("/tasklist", tasklistHandler.AddTask)
}

package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/oceanm7n/wh/task_list"
	task_http "github.com/oceanm7n/wh/task_list/delivery/http"
	task_local_storage "github.com/oceanm7n/wh/task_list/repository/localstorage"
	task_uc "github.com/oceanm7n/wh/task_list/usecase"
)

type App struct {
	httpServer *http.Server

	tasklistUC task_list.UseCase
}

func NewApp() *App {
	return &App{
		tasklistUC: task_uc.NewTaskListUsecase(task_local_storage.NewTasklistLocalStorage()),
	}
}

func hello_world(c *gin.Context) {
	// reply with hello world
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})

}

func (a *App) Run(port string) error {

	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/", hello_world)

	task_http.RegisterHTTPEndpoints(router, a.tasklistUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if os.Getenv("DOCKER_BUILD") == "true" {
		go func() {
			if err := a.httpServer.ListenAndServe(); err != nil {
				log.Fatalf("Failed to listen and serve: %+v", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Interrupt)

		<-quit

		ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdown()

		return a.httpServer.Shutdown(ctx)
	} else {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}

		return a.httpServer.Shutdown(context.Background())
	}
}

package router

import (
	"task_manager/controllers"
	"task_manager/data"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskService data.TaskService) *gin.Engine {
	r := gin.Default()
	taskController := controllers.NewTaskController(taskService)

	
	tasks := r.Group("/tasks")
	{
		tasks.POST("", taskController.CreateTask)
		tasks.GET("", taskController.GetAllTasks)
		tasks.GET("/:id", taskController.GetTask)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id", taskController.DeleteTask)
	}

	return r
}
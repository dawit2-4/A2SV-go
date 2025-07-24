package routers

import (
	"log"
	"os"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskService data.TaskService, userService data.UserManager) *gin.Engine {
	r := gin.Default()
	taskController := controllers.NewTaskController(taskService)
	authController := controllers.NewAuthController(userService)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}

	r.POST("/register", authController.RegisterUser)
	r.POST("/login", authController.LogIn)

	tasks := r.Group("/tasks").Use(middlewares.AuthMiddleware(jwtSecret))
	{
		tasks.POST("", taskController.CreateTask)
		tasks.GET("", taskController.GetAllTasks)
		tasks.GET("/:id", taskController.GetTask)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id", middlewares.AdminOnly(), taskController.DeleteTask)
	}

	return r
}
// Package routers configures HTTP routes for the Task Manager API.
package routers

import (
	"task_manger/Delivery/controllers"
	"task_manger/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController, jwtService Infrastructure.JWTService) *gin.Engine {
	r := gin.Default()

	//Public routes
	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.LogIn)

	//Protected routes
	tasks := r.Group("/tasks").Use(Infrastructure.AuthMiddleware(jwtService)) 
	{
		tasks.POST("", taskController.CreateTask)
		tasks.GET("", taskController.GetAllTasks)
		tasks.GET("/:id", taskController.GetTask)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id",Infrastructure.AdminOnlyMiddleware(),taskController.DeleteTask)
	}

	return r
}
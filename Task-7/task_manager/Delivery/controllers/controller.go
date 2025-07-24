package controllers

import (
	"net/http"
	"task_manager/Domain"
	"task_manager/Usecase"

	"github.com/gin-gonic/gin"
)

// TaskController handles task-related HTTP requests
type TaskController struct {
	taskUsecase Usecase.TaskUsecase
}

// NewTaskController creates a new TaskController
func NewTaskController(taskUsecase Usecase.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase: taskUsecase}
}

// UserController handles user-related HTTP requests
type UserController struct {
	userUsecase Usecase.UserUsecase
}

// NewUserController creates a new UserController
func NewUserController(userUsecase Usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

// CreateTask handles POST /tasks to create a new task
func (tc *TaskController) CreateTask(c *gin.Context) {
	var task Domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	createdTask, err := tc.taskUsecase.CreateTask(ctx, task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    createdTask,
	})
}

// GetTask handles GET /tasks/:id to retrieve a task
func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	task, err := tc.taskUsecase.GetTaskByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

// GetAllTasks handles GET /tasks to retrieve all tasks
func (tc *TaskController) GetAllTasks(c *gin.Context) {
	ctx := c.Request.Context()
	tasks, err := tc.taskUsecase.GetAllTasks(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// UpdateTask handles PUT /tasks/:id to update a task
func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task Domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()
	updatedTask, err := tc.taskUsecase.UpdateTask(ctx, id, task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task":    updatedTask,
	})
}

// DeleteTask handles DELETE /tasks/:id to delete a task
func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	if err := tc.taskUsecase.DeleteTask(ctx, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// RegisterUser handles POST /register to create a new user
func (uc *UserController) RegisterUser(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()
	createdUser, err := uc.userUsecase.RegisterUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    createdUser,
	})
}

// LogIn handles POST /login to authenticate a user
func (uc *UserController) LogIn(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()
	token, err := uc.userUsecase.LogIn(ctx, loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

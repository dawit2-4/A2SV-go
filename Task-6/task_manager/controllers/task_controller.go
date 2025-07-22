package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidInput   = fmt.Errorf("invalid input data")
	ErrInternalServer = fmt.Errorf("internal server error")
	ErrUnauthorized   = fmt.Errorf("unauthorized access")
)

type AuthControllerInterface interface {
	RegisterUser(c *gin.Context)
	LogIn(c *gin.Context)
}

type AuthController struct {
	UserService data.UserManager
}

func NewAuthController(userService data.UserManager) AuthControllerInterface {
	return &AuthController{UserService: userService}
}

type TaskController struct {
	taskService data.TaskService
}

func NewTaskController(taskService data.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (ac *AuthController) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("%w: %v", ErrInvalidInput, err).Error()})
		return
	}

	ctx := c.Request.Context()
	createdUser, err := ac.UserService.RegisterUser(ctx, user)
	if err != nil {
		if errors.Is(err, data.ErrUsernameTaken) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("%w: %v", ErrInternalServer, err).Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "user": createdUser})
}

func (ac *AuthController) LogIn(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("%w: %v", ErrInvalidInput, err).Error()})
		return
	}

	ctx := c.Request.Context()
	token, err := ac.UserService.LogIn(ctx, loginData.Username, loginData.Password)
	if err != nil {
		if errors.Is(err, data.ErrInvalidCredential) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("%w: %v", ErrInternalServer, err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("%w: invalid request body", ErrInvalidInput).Error()})
		return
	}

	createdTask, err := tc.taskService.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, createdTask)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.taskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("%w: %v", ErrInternalServer, err).Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("%w: invalid request body", ErrInvalidInput).Error()})
		return
	}

	updatedTask, err := tc.taskService.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := tc.taskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
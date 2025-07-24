package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService data.TaskService
}

func NewTaskController (taskService data.TaskService) *TaskController{
	return &TaskController{taskService: taskService}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task

	if err:= c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	createdTask, err:= tc.taskService.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.IndentedJSON(http.StatusCreated, createdTask)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := tc.taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func(tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.taskService.GetAllTasks() 
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func(tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invaild request body."})
		return
	}
	updatedTask, err := tc.taskService.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err:= tc.taskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

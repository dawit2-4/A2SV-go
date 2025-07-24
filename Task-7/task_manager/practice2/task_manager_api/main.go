package main

import (
	"net/http"
	"task_manager_api/data"
	"task_manager_api/models"

	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()
	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"Tasks:": data.Tasks})
	})

router.GET("/tasks/:id", func(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range data.Tasks {
		if task.ID == id {
			ctx.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
})


	router.PUT("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var updatedTaks models.Task

		err := ctx.ShouldBindJSON(&updatedTaks); if err != nil{
			ctx.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

		for i, task := range data.Tasks{
			if task.ID == id {
				if updatedTaks.Title != "" {
					data.Tasks[i].Title = updatedTaks.Title
				}
				if updatedTaks.Description != "" {
					data.Tasks[i].Description = updatedTaks.Description
				}
				ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated"})
				return
			}
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error:":"Task not found"})
			
		}

		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error:":"Task not found"})
		

	})

	router.DELETE("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for i, val := range data.Tasks {
			if val.ID == id {
				data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
				ctx.IndentedJSON(http.StatusOK, gin.H{"message:":"Taks deleted succesfully."})
				return
			}
		}

		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error:":"Task not found"})
	})

	router.POST("/tasks", func (ctx *gin.Context)  {
		var newTask models.Task

		err := ctx.ShouldBindJSON(&newTask); if err != nil{
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error:":err})
		}

		data.Tasks = append(data.Tasks, newTask)
		ctx.IndentedJSON(http.StatusOK, gin.H{"message:":"Task added successfully."})
	})
	router.Run("localhost:8080")
}
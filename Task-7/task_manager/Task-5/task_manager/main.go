package main

import (
	"log"
	"task_manager/data"
	"task_manager/models"
	"task_manager/router"
)

func main () {
	models.ConnectDatabase()
	defer models.DisconnectDatabase()

	collection := models.GetTaskCollection()

	taskService := data.NewTaskService(collection)

	r := router.SetupRouter(taskService)
	log.Println("Server is running on http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
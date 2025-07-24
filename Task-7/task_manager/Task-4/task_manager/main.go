package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main () {
	taskService := data.NewTaskService()

	r := router.SetupRouter(taskService)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
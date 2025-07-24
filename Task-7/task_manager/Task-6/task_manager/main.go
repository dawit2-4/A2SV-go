package main

import (
	"log"
	"os"
	"task_manager/data"
	"task_manager/models"
	"task_manager/routers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	connectionString := os.Getenv("MONGODB_URI")
	if connectionString == "" {
		connectionString = "mongodb://localhost:27017"
	}
	models.ConnectDatabase(connectionString)
	defer models.DisconnectDatabase()

	taskCollection := models.GetTaskCollection()
	taskService := data.NewTaskService(taskCollection)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}
	userService := data.NewUserService(models.GetMongoClient(), jwtSecret)

	r := routers.SetupRouter(taskService, userService)
	log.Println("Server is running on http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server error:", err)
	}
}
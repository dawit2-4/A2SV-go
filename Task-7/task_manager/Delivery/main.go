// Package main initializes and runs the Task Manager API server.
package main

import (
	"context"
	"log"
	"os"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/routers"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
	"task_manager/Usecase"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initMongoClient initializes the MongoDB client.
func initMongoClient(connectionString string) *mongo.Client {
	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetMaxPoolSize(100).
		SetMinPoolSize(10)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}
	return client
}

// main starts the Task Manager API server.
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables or defaults")
	}

	// Load environment variables
	connectionString := os.Getenv("MONGODB_URI")
	if connectionString == "" {
		connectionString = "mongodb://localhost:27017"
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "H7k9pQzX2mW3vL8rT4sY6uN9jF2aB5cC7dE8="
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "tasks"
	}
	tasksCollection := os.Getenv("TASKS_COLLECTION")
	if tasksCollection == "" {
		tasksCollection = "tasks"
	}
	usersCollection := os.Getenv("USERS_COLLECTION")
	if usersCollection == "" {
		usersCollection = "users"
	}

	// Initialize MongoDB client
	client := initMongoClient(connectionString)
	defer client.Disconnect(context.Background())

	// Initialize repositories
	taskRepo := Repositories.NewMongoTaskRepository(client, dbName, tasksCollection)
	userRepo := Repositories.NewMongoUserRepository(client, dbName, usersCollection)

	// Initialize services
	jwtService := Infrastructure.NewJWTService(jwtSecret)
	passwordService := Infrastructure.NewPasswordService()

	// Initialize use cases
	taskUsecase := Usecase.NewTaskUsecase(taskRepo)
	userUsecase := Usecase.NewUserUsecase(userRepo, jwtService, passwordService)

	// Initialize controllers and router
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)
	router := routers.SetupRouter(taskController, userController, jwtService)

	// Start server
	log.Println("Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server error:", err)
	}
}
package main

import (
	"go_mongo/model"
	"go_mongo/controller" // Import the new controller package
	"log"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	r.Use(gin.Logger())

	// Define API endpoints
	r.POST("/movies", controller.CreateMovieHandler)
	r.POST("/movies/bulk", controller.CreateMoviesHandler)
	r.PUT("/movies/:id", controller.UpdateMovieHandler)
	r.DELETE("/movies/:id", controller.DeleteMovieHandler)
	r.GET("/movies/name/:name", controller.FindMovieHandler)
	r.GET("/movies/names/:name", controller.FindAllMoviesHandler)
	r.GET("/movies", controller.ListAllMoviesHandler)
	r.DELETE("/movies", controller.DeleteAllMoviesHandler)

	model.ConnectDatabase()
	log.Println("Server started")
	r.Run()
}
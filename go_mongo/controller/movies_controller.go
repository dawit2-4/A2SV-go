package controller

import (
	"net/http"

	"go_mongo/model" // Adjust import path based on your module structure

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateMovieHandler handles creating a single movie
func CreateMovieHandler(c *gin.Context) {
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.InsertOne(movie)
	c.JSON(http.StatusCreated, gin.H{"message": "Movie created successfully"})
}

// CreateMoviesHandler handles creating multiple movies
func CreateMoviesHandler(c *gin.Context) {
	var movies []model.Movie
	if err := c.ShouldBindJSON(&movies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.InsertMany(movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Movies created successfully"})
}

// UpdateMovieHandler handles updating a movie
func UpdateMovieHandler(c *gin.Context) {
	movieID := c.Param("id")

	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.UpdateMovie(movieID, movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

// DeleteMovieHandler handles deleting a movie
func DeleteMovieHandler(c *gin.Context) {
	movieID := c.Param("id")

	if err := model.DeleteMovie(movieID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}

// FindMovieHandler handles finding a movie by name
func FindMovieHandler(c *gin.Context) {
	movieName := c.Param("name")

	movie := model.Find(movieName)
	if movie.ID == primitive.NilObjectID {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// FindAllMoviesHandler handles finding all movies with a specific name
func FindAllMoviesHandler(c *gin.Context) {
	movieName := c.Param("name")

	movies := model.FindAll(movieName)
	c.JSON(http.StatusOK, movies)
}

// ListAllMoviesHandler handles listing all movies
func ListAllMoviesHandler(c *gin.Context) {
	movies := model.ListAll("")
	c.JSON(http.StatusOK, movies)
}

// DeleteAllMoviesHandler handles deleting all movies
func DeleteAllMoviesHandler(c *gin.Context) {
	if err := model.DeleteAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All movies deleted successfully"})
}
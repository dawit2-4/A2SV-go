package main

import (
	// "fmt"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, Quantity: 3},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99, Quantity: 0},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, Quantity: 2},
}

func getAlbum (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func albumById(c *gin.Context) {
	id := c.Param("id")
	album, err := getAlbumById(id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func getAlbumById(id string) (*album, error){
	for i,a := range albums {
		if a.ID == id {
			return &albums[i], nil
		}
	}
	return nil, errors.New("album not found")
}

func createAlbum (c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func returnAlbum (c *gin.Context) {
	id, exists := c.GetQuery("id")

	if !exists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "missing id parameter."})
		return
	}
	album, err := getAlbumById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found."})
		return
	}

	album.Quantity += 1
	c.IndentedJSON(http.StatusOK, album)
}

func checkOutAlbum (c *gin.Context) {
	id, exists := c.GetQuery("id")

	if !exists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "missing id parameter."})
		return
	}

	album, err := getAlbumById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found."})
		return
	}

	if album.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Album not available."})
		return
	}

	album.Quantity -= 1
	c.IndentedJSON(http.StatusOK, album)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbum)
	router.GET("/albums/:id", albumById)
	router.PATCH("/checkout", checkOutAlbum)
	router.PATCH("/return", returnAlbum)
	router.POST("/albums", createAlbum)
	router.Run("localhost: 8080")
}
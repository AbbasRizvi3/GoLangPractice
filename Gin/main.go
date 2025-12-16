package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"albums": albums,
		"query":  c.Query("name"),
	})

}
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func showDashboard(c *gin.Context) {
	fmt.Fprintf(c.Writer, "The request type is: %v\n", c.Request.Method)
	fmt.Fprintf(c.Writer, "Welcome")
}

func samplePath(c *gin.Context) {
	fmt.Fprintf(c.Writer, "The parameter is: %v\n", c.Param("id"))

}

func main() {
	router := gin.Default()
	router.GET("/", showDashboard)
	router.GET("/albums", getAlbums)
	router.GET("/samplePath/:id", samplePath)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

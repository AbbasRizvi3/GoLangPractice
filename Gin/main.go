package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type User struct {
	Name  string
	Email string
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	var results []album
	id := c.Param("id")
	for _, val := range albums {
		if val.ID == id || id == "" {
			results = append(results, val)
		}
	}
	c.JSON(200, gin.H{
		"result": results,
	})
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	albums = append(albums, newAlbum)
	c.JSON(201, gin.H{
		"result": albums,
	})
}
func showDashboard(c *gin.Context) {
	fmt.Print("Inside Dashboard")
	fmt.Fprintf(c.Writer, "The request type is: %v\n", c.Request.Method)
	fmt.Fprintf(c.Writer, "Welcome")
}

func samplePath(c *gin.Context) {
	fmt.Fprintf(c.Writer, "The parameter is: %v\n", c.Param("id"))

}

func middleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("Inside Middleware 1")
		c.Next()
		fmt.Print("exiting middle ware 1")
	}
}

func middleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("Inside Middleware 2")
		c.Next()
		fmt.Print("exiting middle ware 2")
	}
}

func middleWare3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("Inside Middleware 3")
		c.Next()
		fmt.Print("exiting middle ware 3")
	}
}

func getUser(c *gin.Context) {
	user := User{
		Name:  "feehu",
		Email: "snake.com",
	}
	c.HTML(200, "page.html", user)
}

func main() {
	router := gin.Default()
	router.GET("/", middleWare1(), middleWare2(), showDashboard)
	custom := router.Group("/albums")
	custom.Use(middleWare3())

	custom.GET("", getAlbums)
	custom.GET("/samplePath/:id", samplePath)
	custom.POST("", postAlbums)
	router.StaticFile("/image", "../Public/download.jpeg")

	router.LoadHTMLGlob("Templates/*")
	router.GET("/getUser", getUser)

	router.Run("localhost:8080")
}

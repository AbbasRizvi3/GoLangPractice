package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

type Foo1 struct {
	Type     string `json:"type" binding:"required,eq=username"`
	Username string `json:"username" binding:"required"`
}

type Foo2 struct {
	Type     string `json:"type" binding:"required,eq=password"`
	Password string `json:"password" binding:"required"`
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

type myForm struct {
	Colors []string `form:"colors[]"`
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

func main() {
	router := gin.Default()
	// router.GET("/", middleWare1(), middleWare2(), showDashboard)
	// custom := router.Group("/albums")
	// custom.Use(middleWare3())

	// custom.GET("", getAlbums)
	// custom.GET("/samplePath/:id", samplePath)
	// custom.POST("", postAlbums)
	// router.StaticFile("/image", "../Public/download.jpeg")

	// router.LoadHTMLGlob("Templates/*")
	// router.GET("/getUser", getUser)
	router.LoadHTMLFiles("Templates/form.html")
	router.GET("/", indexHandler)
	router.POST("/", formHandler)

	router.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, gin.H{
			"success": true,
			"result":  data,
		})
	})

	router.POST("/update", func(ctx *gin.Context) {
		var fo1 Foo1
		var fo2 Foo2

		if err := ctx.ShouldBindBodyWith(&fo1, binding.JSON); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Username updated", "data": fo1.Username})
			return
		}

		if err := ctx.ShouldBindBodyWith(&fo2, binding.JSON); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Password updated", "data": fo2.Password})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Body did not match any valid structure"})
	})

	router.Run("localhost:8080")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

// when we need to read the body multiple times
func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// This reads c.Request.Body and stores the result into the context.
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	}
}

// func SomeHandler(c *gin.Context) {
// 	objA := formA{}
// 	objB := formB{}

// 	if errA := c.ShouldBind(&objA); errA == nil {
// 		c.String(http.StatusOK, `the body should be formA`)
// 		// Always an error is occurred by this because c.Request.Body is EOF now.
// 	} else if errB := c.ShouldBind(&objB); errB == nil {
// 		c.String(http.StatusOK, `the body should be formB`)
// 	}
// }

func dashboardHandler(c *gin.Context) {
	query := c.Params.ByName("page")
	c.HTML(200, "Dashboard.html", gin.H{
		"title": query,
	})
}

// templates (used mainly for ssr)
func showDashboard(c *gin.Context) {
	c.HTML(200, "Dashboard.html", gin.H{
		"title": "Home",
	})
}

func checkboxhandler(c *gin.Context) {
	c.HTML(200, "Checkbox.html", nil)
}

type myForm struct {
	Colors []string `form:"colors[]"`
}

func handler2(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../Template/*")
	router.GET("/", showDashboard)
	router.GET("/:page", dashboardHandler)
	router.GET("/somehandler", SomeHandler)
	router.GET("/handlecheckbox", checkboxhandler)
	router.POST("/handlecheckbox", handler2)

	router.Run()
}

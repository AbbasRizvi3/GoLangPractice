package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func midddlewareExample() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("This is a middleware example")
		c.Next()
	}
}

func main() {

	router := gin.Default()

	router.GET("/cookie", midddlewareExample(), func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}

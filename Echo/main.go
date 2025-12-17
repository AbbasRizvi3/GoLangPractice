package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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

func getAlbums(c echo.Context) error {
	var results []album
	id := c.Param("id")
	for _, val := range albums {
		if val.ID == id || id == "" {
			results = append(results, val)
		}
	}
	c.JSON(200, results)
	return nil
}

func samplePath(c echo.Context) error {
	parameter := c.Param("id")
	query := c.QueryParam("name")
	fmt.Fprintf(c.Response().Writer, "The paramter is %v and query is %v", parameter, query)
	return nil
}

func postAlbums(c echo.Context) {
	var newAlbum album
	err := c.Bind(&newAlbum)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	albums = append(albums, newAlbum)
	c.JSON(201, albums)
}

func showDashboard(c echo.Context) error {
	c.Response().Status = http.StatusOK
	fmt.Print("Inside Dashboard")
	fmt.Fprintf(c.Response().Writer, "The request type is: %v\n", c.Request().Method)
	fmt.Fprintf(c.Response().Writer, "Welcome")
	return nil
}

func middleWare1(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Print("Inside Middleware 1")
		x := next(c)
		fmt.Print("exiting middle ware 1")
		return x
	}
}

func middleWare2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Print("Inside Middleware 2")
		return next(c)
		fmt.Print("exiting middle ware 2")
		return nil
	}
}

func middleWare3(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Print("Inside Middleware 3")
		err := next(c)
		fmt.Print("exiting middle ware 3")
		return err
	}
}

func main() {
	router := echo.New()
	router.GET("/", middleWare1(middleWare2(showDashboard)))
	custom := router.Group("/albums")
	custom.Use(middleWare3)

	custom.GET("", getAlbums)
	router.GET("/samplePath/:id", samplePath)

	router.Start(":8080")
}

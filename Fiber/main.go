package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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

func getAlbums(c *fiber.Ctx) error {
	var results []album
	id := c.Params("id")
	for _, val := range albums {
		if val.ID == id || id == "" {
			results = append(results, val)
		}
	}
	c.Status(200)
	return c.JSON(results)

}

func samplePath(c *fiber.Ctx) error {
	parameter := c.Params("id")
	query := c.Query("name")
	fmt.Fprintf(c.Response().BodyWriter(), "The paramter is %v and query is %v", parameter, query)
	return nil
}

func postAlbums(c *fiber.Ctx) error {
	var newAlbum album
	err := c.BodyParser(&newAlbum)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	albums = append(albums, newAlbum)
	c.Status(200)
	return c.JSON(albums)
}

func showDashboard(c *fiber.Ctx) error {
	fmt.Print("Inside Dashboard")
	fmt.Fprintf(c.Response().BodyWriter(), "The request type is: %v\n", c.Request().Header.Method())
	fmt.Fprintf(c.Response().BodyWriter(), "Welcome")
	return nil
}

func middleWare1(c *fiber.Ctx) error {
	fmt.Print("Inside Middleware 1")
	err := c.Next()
	fmt.Print("exiting middle ware 1")
	return err
}

func middleWare2(c *fiber.Ctx) error {
	fmt.Print("Inside Middleware 2")
	err := c.Next()
	fmt.Print("exiting middle ware 2")
	return err
}

func middleWare3(c *fiber.Ctx) error {
	fmt.Print("Inside Middleware 3")
	err := c.Next()
	fmt.Print("exiting middle ware 3")
	return err
}

func main() {
	router := fiber.New()
	router.Get("/", middleWare1, middleWare2, showDashboard)
	custom := router.Group("/albums")
	custom.Get("", getAlbums)
	router.Get("/samplePath/:id", samplePath)

	router.Listen(":8000")
}

package main

import (
	"fiber-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := 	fiber.New()

	setupRoutes(app)
	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	

	app.Get("/",func (c *fiber.Ctx) error{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"succes":true,
			"message":"you are at the endpoint",
		})
	})

	api := app.Group("/api")


	api.Get("",func (c *fiber.Ctx) error{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"succes":true,
			"message":"you are at the endpoint",
		})
	})

	routes.TodoRoute(api.Group("/todos"))
}
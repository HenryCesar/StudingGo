package main

import (
	"api-rest/routes"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	// api group
	api := app.Group("/fib")
	// give response when at /fib
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the fib endpoint 😉",
		})
	})

	//connect fibonacci routes
	routes.FibonacciRoute(api.Group("/"))
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}

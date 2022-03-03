package routes

import (
	"api-rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func FibonacciRoute(route fiber.Router) {
	route.Get("/all", controllers.GetAll)
	route.Get("/:input", controllers.GetNumber)
}

package routes

import "github.com/gofiber/fiber/v2"

func AuthRoute(route fiber.Router) {
	route.Get("/login", func(c *fiber.Ctx) error {
		return c.JSON("asdasdasdas")
	})
}
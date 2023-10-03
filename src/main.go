package main

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/db"
	"ecommerce-backend/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config := configs.LoadConfig()

	// connect to db
	_ = db.DBConn(config)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	setupRoutes(app)

	if err := app.Listen(":" + config.App.ServerPort); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the root endpoint 😛",
		})
	})

	api := app.Group("/api")
	routes.AuthRoute(api.Group("/auth"))
}

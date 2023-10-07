package main

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/db"
	"ecommerce-backend/src/logs"
	"ecommerce-backend/src/middlewares"
	"ecommerce-backend/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func main() {
	config := configs.LoadConfig()

	// connect to db
	_ = db.DBConn(config)

	// conf := session.Config{
	// 	Expiration:     24 * time.Hour,
	// 	Storage:        memory.New(),
	// 	KeyLookup:      "cookie:_gothic_session",
	// 	CookieDomain:   "",
	// 	CookiePath:     "",
	// 	CookieSecure:   false,
	// 	CookieHTTPOnly: true,
	// 	CookieSameSite: "Lax",
	// 	KeyGenerator:   utils.UUIDv4,
	// }

	// session := session.New(conf)
	// goth_fiber.SessionStore = session
	goth.UseProviders(
		google.New(configs.Cfg.Google.ClientID, configs.Cfg.Google.ClientSecret, "http://localhost:5050/api/auth/google/callback"),
	)

	app := fiber.New(fiber.Config{
		RequestMethods: fiber.DefaultMethods,
		ErrorHandler:   fiber.DefaultErrorHandler,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: cors.ConfigDefault.AllowMethods,
	}))
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))
	app.Use(middlewares.Recover())

	app.Use(func(c *fiber.Ctx) error {
		logs.SetContext(c)
		return c.Next()
	})

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

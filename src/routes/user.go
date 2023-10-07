package routes

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/db"
	"ecommerce-backend/src/handlers"
	"ecommerce-backend/src/middlewares"
	"ecommerce-backend/src/repositories"
	"ecommerce-backend/src/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router) {
	userRepo := repositories.NewUserRepository(db.GetCollection(configs.Cfg, db.DB, "users"))
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	//

	route.Get("/:id", middlewares.Authenticate(), userHandler.GetUserProfile)
}

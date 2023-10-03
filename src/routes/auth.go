package routes

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/db"
	"ecommerce-backend/src/models"
	"ecommerce-backend/src/repositories"
	"fmt"

	"github.com/gofiber/fiber/v2"
	gf "github.com/shareed2k/goth_fiber"
)

func AuthRoute(route fiber.Router) {
	userRepo := repositories.NewUserRepository(db.GetCollection(configs.Cfg, db.DB, "users"))
	route.Get("/login", func(c *fiber.Ctx) error {
		user, _ := userRepo.FindByEmail("go22@go.com")
		return c.JSON(user)
	})

	route.Get("/:provider", func(c *fiber.Ctx) error {
		if gothUser, err := gf.CompleteUserAuth(c); err == nil {
			c.JSON(gothUser)
		} else {
			gf.BeginAuthHandler(c)
		}
		return nil
	})

	route.Get("/:provider/callback", func(c *fiber.Ctx) error {
		user, err := gf.CompleteUserAuth(c)
		if err != nil {
			return err
		}
		fmt.Println(user.Provider)
		newUser := models.NewUser(models.User{
			Provider:  models.GoogleProvider,
			Email:     user.Email,
			Firstname: user.FirstName,
			Lastname:  user.LastName,
			Avatar:    user.AvatarURL,
			GoogleID:  user.UserID,
		})
		userRepo.Create(newUser)
		c.JSON(user)
		return nil
	})
}

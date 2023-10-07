package middlewares

import (
	"ecommerce-backend/src/configs"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Authenticate() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(configs.Cfg.Jwt.Secret)},
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":        fiber.StatusUnauthorized,
				"message":     fiber.ErrUnauthorized.Message,
				"description": err.Error(),
			})
		},
	})
}

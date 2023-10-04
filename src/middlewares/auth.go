package middlewares

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/errs"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Authenticate() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(configs.Cfg.Jwt.Secret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return errs.FiberError(c, fiber.ErrUnauthorized)
		},
	})
}

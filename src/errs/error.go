package errs

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func FiberError(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*fiber.Error); ok {
		return c.Status(appErr.Code).JSON(appErr)
	}
	log.Println(err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    fiber.StatusInternalServerError,
		"message": err.Error(),
	})
}

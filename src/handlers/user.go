package handlers

import (
	"ecommerce-backend/src/services"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) GetUserProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.ErrBadRequest
	}
	user, err := h.userService.GetProfile(id)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(user)
}

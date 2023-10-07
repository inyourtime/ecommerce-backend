package handlers

import (
	"ecommerce-backend/src/models"
	"ecommerce-backend/src/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	userService services.UserService
}

func NewAuthHandler(userService services.UserService) authHandler {
	return authHandler{userService: userService}
}

func (h authHandler) Register(c *fiber.Ctx) error {
	req := models.User{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrUnprocessableEntity
	}

	if err := validator.New().Struct(req); err != nil {
		return FiberError(c, fiber.ErrBadRequest)
	}
	req.Provider = models.LocalProvider
	user, err := h.userService.CreateUser(req)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(user)
}

func (h authHandler) Login(c *fiber.Ctx) error {
	req := models.LoginUserDto{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrUnprocessableEntity
	}
	if err := validator.New().Struct(req); err != nil {
		return FiberError(c, fiber.ErrBadRequest)
	}

	result, err := h.userService.LoginUser(req)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(result)
}

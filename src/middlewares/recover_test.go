package middlewares_test

import (
	"ecommerce-backend/src/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	t.Run("No panic", func(t *testing.T) {
		app := fiber.New()
		app.Use(middlewares.Recover())
		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendStatus(200)
		})
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		res, err := app.Test(req)
		assert.Equal(t, nil, err)
		assert.Equal(t, 200, res.StatusCode)
	})

	t.Run("Panic occurs", func(t *testing.T) {
		app := fiber.New()
		app.Use(middlewares.Recover())
		app.Get("/test", func(c *fiber.Ctx) error {
			panic("test")
		})
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		res, err := app.Test(req)
		assert.Equal(t, nil, err)
		assert.Equal(t, 500, res.StatusCode)
	})
}

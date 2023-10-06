package middlewares_test

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	configs.InitConfigMock()
}

func setupApp() *fiber.App {
	return fiber.New()
}

func TestAuth(t *testing.T) {
	// mock
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjkyODAzMjA4LCJuYW1lIjoiSm9obiBEb2UifQ.SNrr_DxyjESSkMQNkI4qODo-csjBazIgj2PkZlGz90s"
	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJlYW1AZ21haWwuY29tIiwiZXhwIjoxNjk2NTYwNTUzLCJpZCI6IjY1MWQzNzIzYjBjNmY5OTViNTQzMDViYyIsInJvbGUiOiJ1c2VyIn0.Tm3yeEJoeVWTn3vDUXPKEl7rau8IwdKKXNMDY994wNI"
	otherSecretToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJlYW1AZ21haWwuY29tIiwiZXhwIjoyNTYwNDcyMDI0LCJpZCI6IjY1MWQzNzIzYjBjNmY5OTViNTQzMDViYyIsInJvbGUiOiJ1c2VyIn0.dH6iyl-wNuuM7Uzq-YoGVP9YqkYdoxUPAuNvUBQR3LI"
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJlYW1AZ21haWwuY29tIiwiZXhwIjoyNTYwNDcyMjY3LCJpZCI6IjY1MWQzNzIzYjBjNmY5OTViNTQzMDViYyIsInJvbGUiOiJ1c2VyIn0.On4ybj-YG3a4yajeAbLpPMG5hEh-M59Ap2FrTsk_KeA"

	type testCase struct {
		description  string
		auth         bool
		token        *string
		expectedCode int
		expectedErr  error
	}

	cases := []testCase{
		{description: "No auth middleware", auth: false, token: nil, expectedCode: 200, expectedErr: nil},
		{description: "Have auth: response 401 (no token)", auth: true, token: nil, expectedCode: 401, expectedErr: nil},
		{description: "Have auth: reponse 401 (token invalid signature)", auth: true, token: &invalidToken, expectedCode: 401, expectedErr: nil},
		{description: "Have auth: reponse 401 (token expired)", auth: true, token: &expiredToken, expectedCode: 401, expectedErr: nil},
		{description: "Have auth: reponse 401 (incorrect secret)", auth: true, token: &otherSecretToken, expectedCode: 401, expectedErr: nil},
		{description: "Have auth: reponse 200", auth: true, token: &validToken, expectedCode: 200, expectedErr: nil},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			app := setupApp()
			if c.auth {
				app.Use(middlewares.Authenticate())
			}
			app.Get("/test", func(c *fiber.Ctx) error {
				return c.SendStatus(200)
			})

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if c.token != nil {
				req.Header.Add("Authorization", "Bearer "+*c.token)
			}
			res, err := app.Test(req)
			assert.Equal(t, c.expectedErr, err)
			assert.Equal(t, c.expectedCode, res.StatusCode)
		})
	}
}

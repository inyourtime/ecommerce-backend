package utils_test

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/utils"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func init() {
	configs.InitConfigMock()
}

func TestSignToken(t *testing.T) {
	claimsMock := jwt.MapClaims{
		"id":      "1234",
		"email":   "mock@email.com",
		"role":    "admin",
		"yesorno": true,
		"num":     888,
	}
	secret := configs.Cfg.Jwt.Secret
	t.Run("Signed token success", func(t *testing.T) {
		token, err := utils.Token(claimsMock, secret)
		assert.NotEmpty(t, token.AccessToken)
		assert.NotEmpty(t, token.RefreshToken)
		assert.Nil(t, err)
	})

	t.Run("Decode success", func(t *testing.T) {
		token, _ := utils.Token(claimsMock, secret)
		decoded, err := utils.Decode(token.AccessToken, secret)
		assert.EqualValues(t, claimsMock["id"], decoded["id"])
		assert.EqualValues(t, claimsMock["email"], decoded["email"])
		assert.EqualValues(t, claimsMock["role"], decoded["role"])
		assert.EqualValues(t, claimsMock["yesorno"], decoded["yesorno"])
		assert.EqualValues(t, claimsMock["num"], decoded["num"])
		assert.Nil(t, err)
	})

	t.Run("Token expired", func(t *testing.T) {
		expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1vY2tAZW1haWwuY29tIiwiZXhwIjoxNjk2NTE0NzE5LCJpZCI6IjEyMzQiLCJyb2xlIjoiYWRtaW4ifQ._1hwS1dT-BNC8BXwfFuBHCKuUjX56mAfNgOKXfvDyp4"
		_, err := utils.Decode(expiredToken, secret)
		assert.Error(t, err)
		assert.ErrorIs(t, err, jwt.ErrTokenExpired)
	})

	t.Run("Token invalid signature", func(t *testing.T) {
		invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjkyODAzMjA4LCJuYW1lIjoiSm9obiBEb2UifQ.SNrr_DxyjESSkMQNkI4qODo-csjBazIgj2PkZlGz90s"
		_, err := utils.Decode(invalidToken, secret)
		assert.Error(t, err)
		assert.ErrorIs(t, err, jwt.ErrTokenSignatureInvalid)
	})
}

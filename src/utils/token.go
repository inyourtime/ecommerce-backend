package utils

import (
	"ecommerce-backend/src/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Token(claims jwt.MapClaims, secret string) (models.Token, error) {
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := at.SignedString([]byte(secret))

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := rt.SignedString([]byte(secret))

	return models.Token{AccessToken: accessToken, RefreshToken: refreshToken}, err
}

func Decode(token string, secret string) (jwt.MapClaims, error) {
	c := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return c, err
}

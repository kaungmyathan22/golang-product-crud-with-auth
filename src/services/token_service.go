package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
)

type TokenService struct {
}

func (tokenService *TokenService) SignAccessToken(userId string) (string, error) {
	secretKey := []byte(config.AppConfigInstance.ACCESS_TOKEN_SECRET)
	expTime := time.Now().Add(1 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": expTime.Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (tokenService *TokenService) SignRefreshToken() {}

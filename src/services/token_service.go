package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common/interfaces"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
)

type TokenService struct {
}

func (tokenService *TokenService) SignAccessToken(userId string) (string, error) {
	secretKey := []byte(config.AppConfigInstance.ACCESS_TOKEN_SECRET)
	expTime := tokenService.GetAccessTokenExpiration()
	claims := interfaces.JwtCustomClaims{
		Sub: userId,
		Iat: time.Now().Unix(),
		Exp: expTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (tokenService *TokenService) GetAccessTokenExpiration() time.Time {
	return time.Now().Add(1 * time.Hour)
}

func (tokenService *TokenService) SignRefreshToken() {}

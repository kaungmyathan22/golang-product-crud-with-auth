package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common/interfaces"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
)

type TokenService struct {
}

func signToken(expTime time.Time, userId string, secretKey []byte) (string, error) {
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

func (tokenService *TokenService) SignAccessToken(userId string) (string, error) {
	return signToken(tokenService.GetAccessTokenExpiration(), userId, []byte(config.AppConfigInstance.ACCESS_TOKEN_SECRET))
}

func (tokenService *TokenService) GetAccessTokenExpiration() time.Time {
	return time.Now().Add(1 * time.Hour)
}

func (tokenService *TokenService) GetRefreshTokenExpiration() time.Time {
	return time.Now().Add(7 * 24 * time.Hour)
}

func (tokenService *TokenService) SignRefreshToken(userId string) (string, error) {
	return signToken(tokenService.GetRefreshTokenExpiration(), userId, []byte(config.AppConfigInstance.REFRESH_TOKEN_SECRET))
}

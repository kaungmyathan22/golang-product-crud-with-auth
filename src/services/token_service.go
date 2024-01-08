package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common/interfaces"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/logger"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
	"golang.org/x/crypto/scrypt"
)

type TokenService struct {
	Repository *repositories.TokenRepository
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

func (tokenService *TokenService) GetPasswordResetTokenExpiration() time.Time {
	return time.Now().Add(5 * time.Minute)
}

func (tokenService *TokenService) SignPasswordResetToken(userId string, code string) (string, error) {
	expTime := tokenService.GetPasswordResetTokenExpiration()
	claims := interfaces.ResetJwtClaims{
		Code: code,
		JwtCustomClaims: interfaces.JwtCustomClaims{
			Sub: userId,
			Iat: time.Now().Unix(),
			Exp: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfigInstance.PASSWORD_RESET_TOKEN_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (tokenService *TokenService) SignRefreshToken(userId string) (string, error) {

	token, err := signToken(tokenService.GetRefreshTokenExpiration(), userId, []byte(config.AppConfigInstance.REFRESH_TOKEN_SECRET))
	if err != nil {
		return "", err
	}
	// hash token before saving to database
	hashedToken, err := common.EncryptToken(token, config.AppConfigInstance.REFRESH_TOKEN_ENCRYPT_KEY)
	if err != nil {
		return "", err
	}
	payload := dto.CreateRefreshTokenDTO{
		UserID:         userId,
		ExpirationTime: tokenService.GetRefreshTokenExpiration(),
		TokenHash:      string(hashedToken),
	}
	if err = tokenService.Repository.CreateNewToken(&payload); err != nil {
		return "", err
	}
	return token, nil
}

func (tokenService *TokenService) VerifyRefreshToken(token string, userID string) error {
	tokenModel, err := tokenService.Repository.GetRefreshTokenByUserID(userID)
	if err != nil {
		return err
	}
	decryptedToken, err := common.DecryptToken(tokenModel.TokenHash, config.AppConfigInstance.REFRESH_TOKEN_ENCRYPT_KEY)
	if err != nil {
		return err
	}
	if decryptedToken != token {
		logger.Debug("token doesn't match.")
		return fmt.Errorf("invalid refresh token")
	}
	return nil
}

func (tokenService *TokenService) VerifyToken(tokenString, label string, secretKey []byte) (*interfaces.JwtCustomClaims, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("%s token is required", label)
	}
	claims := &interfaces.JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid %s token", label)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid %s token", label)
	}

	expires := claims.Exp
	if time.Now().Unix() > expires {
		return nil, fmt.Errorf("%s token expired", label)
	}
	return claims, nil
}

func GenerateAESKey(password, salt []byte) ([]byte, error) {
	return scrypt.Key(password, salt, 32768, 8, 1, 32)
}

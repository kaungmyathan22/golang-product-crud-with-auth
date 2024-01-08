package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common/interfaces"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthenticationService struct {
	AuthRepository *repositories.AuthenticationRepository
}

func (service *AuthenticationService) CreateNewPasswordResetLink(payload *dto.SavePasswordResetDTO) (string, error) {
	code := common.GenerateRandomNumber()
	encryptedToken, err := common.EncryptToken(strconv.Itoa(code), config.AppConfigInstance.PASSWORD_RESET_TOKEN_ENCRYPT_KEY)
	payload.Code = strconv.Itoa(code)
	if err != nil {
		return "", err
	}
	err = service.AuthRepository.GenerateResetPasswordLink(*payload)
	if err != nil {
		return "", err
	}
	return encryptedToken, nil
}

func (service *AuthenticationService) VerifyPasswordResetToken(tokenString string) (string, error) {
	claims := &interfaces.ResetJwtClaims{}
	secretKey := []byte(config.AppConfigInstance.PASSWORD_RESET_TOKEN_SECRET)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return "", errors.New("error verifying password reset token")
	}

	if !token.Valid {
		return "", errors.New("invalid password reset token")
	}

	expires := claims.Exp
	if time.Now().Unix() > expires {
		return "", errors.New("expired password reset token")
	}
	decryptedCode, err := common.DecryptToken(claims.Code, config.AppConfigInstance.PASSWORD_RESET_TOKEN_ENCRYPT_KEY)
	if err != nil {
		return "", err
	}
	tokenModel, err := service.AuthRepository.GetPasswordResetLink(bson.M{"userID": claims.Sub, "code": decryptedCode})
	if err != nil {
		return "", err
	}
	if time.Now().Unix() > tokenModel.ExpirationTime.Unix() {
		return "", errors.New("invalid password reset token")
	}
	return claims.Sub, nil
}

func (service *AuthenticationService) DeleteResetToken(userId string) error {
	return service.AuthRepository.DeleteResetToken(userId)
}

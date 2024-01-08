package services

import (
	"strconv"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
)

type AuthenticationService struct {
	AuthRepository *repositories.AuthenticationRepository
}

func (service *AuthenticationService) CreateNewPasswordResetLink(payload *dto.SavePasswordResetDTO) (string, error) {
	code := common.GenerateRandomNumber()
	encryptedToken, err := common.EncryptToken(strconv.Itoa(code), config.AppConfigInstance.PASSWORD_RESET_TOKEN_ENCRYPT_KEY)
	payload.Code = encryptedToken
	if err != nil {
		return "", err
	}
	err = service.AuthRepository.GenerateResetPasswordLink(*payload)
	if err != nil {
		return "", err
	}
	return encryptedToken, nil
}

func (service *AuthenticationService) VerifyPasswordResetToken(token string) {
	// verify token
	common.DecryptToken(token, config.AppConfigInstance.PASSWORD_RESET_TOKEN_ENCRYPT_KEY)
	// set new password
}

package services

import (
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
)

type AuthenticationService struct {
	AuthRepository *repositories.AuthenticationRepository
}

func (service *AuthenticationService) CreateNewPasswordResetLink(payload *dto.SavePasswordResetDTO) (string, error) {
	encryptedToken, err := common.EncryptToken(payload.Token, config.AppConfigInstance.PASSWORD_RESET_TOKEN_ENCRYPT_KEY)
	payload.Token = encryptedToken
	if err != nil {
		return "", err
	}
	err = service.AuthRepository.GenerateResetPasswordLink(*payload)
	if err != nil {
		return "", err
	}
	return encryptedToken, nil
}

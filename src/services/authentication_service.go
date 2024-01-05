package services

import (
	"strconv"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
)

type AuthenticationService struct {
	AuthRepository *repositories.AuthenticationRepository
}

func (service *AuthenticationService) CreateNewPasswordResetCode(user *dto.UserDTO) (*int, error) {
	code := common.GenerateRandomNumber()
	err := service.AuthRepository.GenerateResetPasswordCode(dto.CreatePasswordResetDTO{
		UserID:         user.ID.Hex(),
		Code:           strconv.Itoa(code),
		ExpirationTime: service.AuthRepository.GetPasswordResetCodeExpirationTime(),
	})
	if err != nil {
		return nil, err
	}
	return &code, nil
}

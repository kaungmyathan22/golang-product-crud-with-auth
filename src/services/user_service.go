package services

import (
	"time"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func (svc *UserService) createUser(payload *dto.CreateUserDTO) (*models.UserModel, error) {
	//TODO: decrypt password
	hashed_password := payload.Password
	user := &models.UserModel{
		Username:   payload.Username,
		Password:   hashed_password,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		IsDisabled: false,
	}
	result, err := svc.Repository.CreateUser(user)
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, err
}

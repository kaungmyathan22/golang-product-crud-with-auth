package services

import (
	"fmt"
	"time"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func (svc *UserService) CreateUser(payload *dto.CreateUserDTO) (*models.UserModel, error) {
	//TODO: decrypt password
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost) //payload.Password
	if err != nil {
		return nil, err
	}

	user := dto.UserDTO{
		Username:   payload.Username,
		Password:   string(hashed_password),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		IsDisabled: false,
	}
	result, err := svc.Repository.CreateUser(&user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("user with the name %s already exists", user.Username)
		}
		return nil, err
	}
	fmt.Println(result.InsertedID)
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user.ToModel(), err
}

func (svc *UserService) GetUserByUsername(username string) (*dto.UserDTO, error) {
	return svc.Repository.GetUserByUsername(username)
}

func (svc *UserService) GetUserByUserId(userId string) (*dto.UserDTO, error) {
	return svc.Repository.GetUserByUserId(userId)
}

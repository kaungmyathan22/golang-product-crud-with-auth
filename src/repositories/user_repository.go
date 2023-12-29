package repositories

import (
	"context"
	"fmt"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	UserCollection *mongo.Collection
}

var ctx = context.TODO()

func (repo *UserRepository) CreateUser(payload *dto.UserDTO) (*mongo.InsertOneResult, error) {
	result, err := repo.UserCollection.InsertOne(ctx, payload)
	return result, err
}

func (repo *UserRepository) GetUserByUsername(username string) (*dto.UserDTO, error) {
	var user dto.UserDTO
	err := repo.UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	fmt.Println(user.ID)
	if err != nil {
		return nil, err
	}
	fmt.Println(user.ID)
	return &user, nil
}

func (repo *UserRepository) GetUserByUserId(userId string) (*dto.UserDTO, error) {
	var user dto.UserDTO
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	err = repo.UserCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

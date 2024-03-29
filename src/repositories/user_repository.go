package repositories

import (
	"errors"
	"fmt"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	UserCollection *mongo.Collection
}

func (repo *UserRepository) CreateUser(payload *models.UserModel) (*models.UserModel, error) {
	result, err := repo.UserCollection.InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}
	payload.ID = result.InsertedID.(primitive.ObjectID)
	return payload, err
}

func (repo *UserRepository) GetUserByUsername(username string) (*dto.UserDTO, error) {
	var user dto.UserDTO
	err := repo.UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
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

func (repo *UserRepository) GetUserByEmail(email string) (*dto.UserDTO, error) {
	var user dto.UserDTO
	err := repo.UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("User not found with given email address.")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) DeleteUserByUserId(userId primitive.ObjectID) error {
	result := repo.UserCollection.FindOneAndDelete(ctx, bson.M{"_id": userId})
	return result.Err()
}

func (repo *UserRepository) UpdateUser(userId primitive.ObjectID, payload bson.M) error {
	_, err := repo.UserCollection.UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$set": payload})
	if err != nil {
		return err
	}
	return nil
}

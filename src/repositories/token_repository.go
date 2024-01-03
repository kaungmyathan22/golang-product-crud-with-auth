package repositories

import (
	"errors"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct {
	TokenCollection *mongo.Collection
}

func (repository *TokenRepository) CreateNewToken(payload *dto.CreateRefreshTokenDTO) error {
	userId, err := primitive.ObjectIDFromHex(payload.UserID)
	if err != nil {
		return err
	}
	existingToken := repository.TokenCollection.FindOne(ctx, bson.M{"userId": userId})
	if err != nil {
		if errors.Is(existingToken.Err(), mongo.ErrNoDocuments) {
			// create new token
			_, err := repository.TokenCollection.InsertOne(ctx, payload)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	// update existing token
	var tokenModel models.RefreshTokenModel
	if err := existingToken.Decode(&tokenModel); err != nil {
		return err
	}
	_, err = repository.TokenCollection.UpdateOne(ctx, bson.M{"_id": tokenModel.ID}, payload)
	if err != nil {
		return err
	}
	return nil
}

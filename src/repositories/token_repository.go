package repositories

import (
	"errors"
	"fmt"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct {
	TokenCollection *mongo.Collection
}

func (repository *TokenRepository) GetRefreshTokenByUserID(userID string) (*models.RefreshTokenModel, error) {
	var tokenModel models.RefreshTokenModel
	if err := repository.TokenCollection.FindOne(ctx, bson.M{"userID": userID}).Decode(&tokenModel); err != nil {
		return nil, fmt.Errorf("invalid refresh token")
	}
	return &tokenModel, nil
}

func (repository *TokenRepository) CreateNewToken(payload *dto.CreateRefreshTokenDTO) error {
	var tokenModel models.RefreshTokenModel
	if err := repository.TokenCollection.FindOne(ctx, bson.M{"userID": payload.UserID}).Decode(&tokenModel); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// create new token
			_, err := repository.TokenCollection.InsertOne(ctx, payload)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	_, err := repository.TokenCollection.UpdateOne(ctx, bson.M{"_id": tokenModel.ID}, bson.M{"$set": payload})
	if err != nil {
		return err
	}
	return nil
}

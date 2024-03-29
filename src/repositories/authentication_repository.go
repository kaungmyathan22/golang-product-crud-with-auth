package repositories

import (
	"errors"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthenticationRepository struct {
	PasswordResetCollecion *mongo.Collection
}

func (repository *AuthenticationRepository) GenerateResetPasswordLink(payload dto.SavePasswordResetDTO) error {
	var tokenModel models.PasswordResetTokenModel
	if err := repository.PasswordResetCollecion.FindOne(ctx, bson.M{"userID": payload.UserID}).Decode(&tokenModel); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			_, err := repository.PasswordResetCollecion.InsertOne(ctx, payload)
			if err != nil {
				return err
			}
			return nil
		}
		return nil
	}
	_, err := repository.PasswordResetCollecion.UpdateOne(ctx, bson.M{"_id": tokenModel.ID}, bson.M{"$set": bson.M{
		"code":           payload.Code,
		"expirationTime": payload.ExpirationTime,
	}})
	if err != nil {
		return err
	}
	return nil
}

func (repository *AuthenticationRepository) GetPasswordResetLink(filter bson.M) (*models.PasswordResetTokenModel, error) {
	var tokenModel *models.PasswordResetTokenModel
	err := repository.PasswordResetCollecion.FindOne(ctx, filter).Decode(&tokenModel)
	if err != nil {
		return nil, err
	}
	return tokenModel, nil
}

func (repo *AuthenticationRepository) DeleteResetToken(userId string) error {
	result := repo.PasswordResetCollecion.FindOneAndDelete(ctx, bson.M{"userID": userId})
	return result.Err()
}

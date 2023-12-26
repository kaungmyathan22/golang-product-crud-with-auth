package repositories

import (
	"context"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	UserCollection *mongo.Collection
}

var ctx = context.TODO()

func (repo *UserRepository) CreateUser(payload *models.UserModel) (*mongo.InsertOneResult, error) {
	result, err := repo.UserCollection.InsertOne(ctx, payload)
	return result, err
}

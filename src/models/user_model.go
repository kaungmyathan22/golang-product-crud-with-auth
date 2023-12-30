package models

import (
	"time"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID         primitive.ObjectID `bson:"_id"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	Username   string             `bson:"username"`
	Password   string             `bson:"password"`
	IsDisabled bool               `bson:"isDisabled"`
}

func (model *UserModel) ToUserDTO() *dto.UserDTO {
	return &dto.UserDTO{
		ID:         model.ID,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
		Username:   model.Username,
		Password:   model.Password,
		IsDisabled: model.IsDisabled,
	}
}

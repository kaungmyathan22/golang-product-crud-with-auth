package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDTO struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	Username   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"-"`
	IsDisabled bool               `bson:"isDisabled" json:"isDisabled"`
}

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3" error:"username field is required and must be a minimum of 3."`
	Password string `json:"password" validate:"required,min=6" error:"password field is required and must be a minimum of 6."`
}

type UpdateUserDTO struct {
	Username string `json:"username"`
}

type UpdatePasswordDTO struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}

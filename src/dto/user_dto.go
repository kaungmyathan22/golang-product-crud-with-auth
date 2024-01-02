package dto

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserDTO struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
	Username        string             `bson:"username" json:"username"`
	Password        string             `bson:"password" json:"-"`
	Email           string             `bson:"email" json:"email"`
	IsEmailVerified bool               `bson:"isEmailVerified" json:"isEmailVerified"`
	IsDisabled      bool               `bson:"isDisabled" json:"isDisabled"`
}

func (user *UserDTO) IsPasswordMach(password string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("invalid username / password")
	}
	return nil
}

type CreateUserDTO struct {
	Email    string `json:"email" validate:"required,email" error:"email field is required and must be a valid email address."`
	Username string `json:"username" validate:"required,min=3" error:"username field is required and must be a minimum of 3."`
	Password string `json:"password" validate:"required,min=6" error:"password field is required and must be a minimum of 6."`
}

type UpdateUserDTO struct {
	Username string `json:"username"`
}

type ChangePasswordDTO struct {
	NewPassword string `json:"newPassword" validate:"required,min=6"`
	OldPassword string `json:"oldPassword" validate:"required,min=6"`
}

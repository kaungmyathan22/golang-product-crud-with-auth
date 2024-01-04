package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginDTO struct {
	Username string `json:"username" validate:"required,min=3" error:"username field is required and must be a minimum of 3."`
	Password string `json:"password" validate:"required,min=6" error:"password field is required and must be a minimum of 6."`
}

type CreatePasswordResetDTO struct {
	UserId         primitive.ObjectID
	Code           int
	ExpirationTime time.Time
}

type ForgotPasswordDTO struct {
	Email string `json:"email" validate:"required,email" error:"email field is required and must be a valid email address."`
}
